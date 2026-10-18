package invoices_test

import (
	"encoding/json"
	"encoding/xml"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"
)

type xsdField struct {
	name, setter, kind, sample string
	nillable                   bool
}

type xsdSection struct {
	name, setter string
	count        int
	newBuilder   func() any
	fields       []xsdField
}

type xsdContract struct {
	name       string
	newBuilder func() any
	roots      map[int]string
	sections   []xsdSection
}

type xsdNode struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",any,attr"`
	Text     string     `xml:",chardata"`
	Children []xsdNode  `xml:",any"`
}

// These expectations describe the supplied XML schemas, independently of the domain
// structs. All construction goes through the SDK's public invoice builders.
func TestFacturasXSD(t *testing.T) {
	for _, contract := range xsdContracts {
		t.Run(contract.name, func(t *testing.T) {
			for modalidad, root := range contract.roots {
				t.Run(root, func(t *testing.T) {
					for _, mode := range []string{"valores", "nulos", "ceros"} {
						t.Run(mode, func(t *testing.T) {
							document := buildXSDDocument(t, contract, modalidad, mode)
							data, err := xml.Marshal(document)
							if err != nil {
								t.Fatal(err)
							}
							var node xsdNode
							if err := xml.Unmarshal(data, &node); err != nil {
								t.Fatal(err)
							}
							if node.XMLName.Local != root {
								t.Fatalf("raíz = %s; esperada %s", node.XMLName.Local, root)
							}
							foundSchema := false
							for _, attr := range node.Attrs {
								if attr.Name.Space == "http://www.w3.org/2001/XMLSchema-instance" && attr.Name.Local == "noNamespaceSchemaLocation" {
									foundSchema = attr.Value == root+".xsd"
								}
							}
							if !foundSchema {
								t.Fatalf("la ubicación XSD no corresponde a %s", root)
							}
							expectedCount := 0
							for _, section := range contract.sections {
								expectedCount += section.count
							}
							if len(node.Children) != expectedCount {
								t.Fatalf("secciones = %d; esperadas %d", len(node.Children), expectedCount)
							}
							position := 0
							for _, section := range contract.sections {
								for range section.count {
									actual := node.Children[position]
									position++
									if actual.XMLName.Local != section.name {
										t.Fatalf("sección = %s; esperada %s", actual.XMLName.Local, section.name)
									}
									if len(actual.Children) != len(section.fields) {
										t.Fatalf("%s: %d campos; esperados %d", section.name, len(actual.Children), len(section.fields))
									}
									for i, field := range section.fields {
										checkXSDField(t, actual.Children[i], field, mode)
									}
								}
							}
						})
					}
				})
			}
		})
	}
}

func buildXSDDocument(t *testing.T, contract xsdContract, modalidad int, mode string) any {
	t.Helper()
	builder := reflect.ValueOf(contract.newBuilder())
	builder.MethodByName("WithModalidad").Call([]reflect.Value{reflect.ValueOf(modalidad)})
	for _, section := range contract.sections {
		for range section.count {
			part := reflect.ValueOf(section.newBuilder())
			for _, field := range section.fields {
				method := part.MethodByName(field.setter)
				if !method.IsValid() {
					t.Fatalf("%s: falta %s", section.name, field.setter)
				}
				argType := method.Type().In(0)
				if field.kind != "json" && (argType.Kind() == reflect.Pointer) != field.nillable {
					t.Fatalf("%s: parámetro %s, nillable XSD=%t", field.setter, argType, field.nillable)
				}
				value := reflect.New(argType).Elem()
				if field.kind == "json" {
					if mode != "nulos" {
						if mode == "ceros" {
							value.Set(reflect.ValueOf(map[string]any{}))
						} else {
							setXSDValue(t, value, field)
						}
					}
					method.Call([]reflect.Value{value})
					continue
				}
				if !field.nillable || mode != "nulos" {
					target := value
					if field.nillable {
						value.Set(reflect.New(argType.Elem()))
						target = value.Elem()
					}
					if mode != "ceros" || !field.nillable {
						setXSDValue(t, target, field)
					}
				}
				method.Call([]reflect.Value{value})
			}
			setter := builder.MethodByName(section.setter)
			if !setter.IsValid() {
				t.Fatalf("falta %s", section.setter)
			}
			setter.Call(part.MethodByName("Build").Call(nil))
		}
	}
	return builder.MethodByName("Build").Call(nil)[0].Interface()
}

func TestFacturasInventarioYSectores(t *testing.T) {
	roots := make(map[string]bool)
	if len(xsdContracts) != 48 {
		t.Fatalf("implementaciones = %d; esperadas 48 para las 47 familias", len(xsdContracts))
	}
	for _, contract := range xsdContracts {
		for _, root := range contract.roots {
			if roots[root] && root != "notaFiscalElectronicaCreditoDebito" && root != "notaFiscalComputarizadaCreditoDebito" {
				t.Fatalf("raíz duplicada: %s", root)
			}
			roots[root] = true
		}
		t.Run(contract.name, func(t *testing.T) {
			header := contract.sections[0]
			value := reflect.ValueOf(header.newBuilder()).MethodByName("Build").Call(nil)[0].Interface()
			data, err := xml.Marshal(value)
			if err != nil {
				t.Fatal(err)
			}
			var node xsdNode
			if err := xml.Unmarshal(data, &node); err != nil {
				t.Fatal(err)
			}
			for _, field := range header.fields {
				if field.name != "codigoDocumentoSector" {
					continue
				}
				found := false
				for _, actual := range node.Children {
					if actual.XMLName.Local == field.name {
						checkXSDField(t, actual, field, "valores")
						found = true
					}
				}
				if !found {
					t.Fatal("falta codigoDocumentoSector")
				}
			}
		})
	}
	if len(roots) != 93 {
		t.Fatalf("raíces = %d; esperadas 93", len(roots))
	}
}

func setXSDValue(t *testing.T, value reflect.Value, field xsdField) {
	t.Helper()
	switch field.kind {
	case "json":
		var data map[string]any
		if err := json.Unmarshal([]byte(field.sample), &data); err != nil {
			t.Fatal(err)
		}
		value.Set(reflect.ValueOf(data))
	case "string", "bigInteger":
		if value.Kind() != reflect.String {
			t.Fatalf("%s: tipo %s, esperado string", field.setter, value.Type())
		}
		value.SetString(field.sample)
	case "integer":
		if value.Kind() != reflect.Int && value.Kind() != reflect.Int64 {
			t.Fatalf("%s: tipo %s, esperado entero", field.setter, value.Type())
		}
		n, err := strconv.ParseInt(field.sample, 10, 64)
		if err != nil {
			t.Fatal(err)
		}
		value.SetInt(n)
	case "decimal":
		if value.Kind() != reflect.Float64 {
			t.Fatalf("%s: tipo %s, esperado float64", field.setter, value.Type())
		}
		n, err := strconv.ParseFloat(field.sample, 64)
		if err != nil {
			t.Fatal(err)
		}
		value.SetFloat(n)
	case "dateTime":
		date, err := time.Parse(time.RFC3339, field.sample)
		if err != nil {
			t.Fatal(err)
		}
		if value.Type() != reflect.TypeFor[time.Time]() {
			t.Fatalf("%s: tipo %s, esperado time.Time", field.setter, value.Type())
		}
		value.Set(reflect.ValueOf(date))
	default:
		t.Fatalf("tipo XSD no contemplado: %s", field.kind)
	}
}

func checkXSDField(t *testing.T, actual xsdNode, expected xsdField, mode string) {
	t.Helper()
	if actual.XMLName.Local != expected.name {
		t.Fatalf("elemento = %s; esperado %s", actual.XMLName.Local, expected.name)
	}
	wantNil := expected.nillable && mode == "nulos"
	isNil := false
	for _, attr := range actual.Attrs {
		if attr.Name.Space == "http://www.w3.org/2001/XMLSchema-instance" && attr.Name.Local == "nil" && attr.Value == "true" {
			isNil = true
		}
	}
	if isNil != wantNil {
		t.Fatalf("%s: xsi:nil=%t; esperado %t", expected.name, isNil, wantNil)
	}
	want := expected.sample
	if wantNil {
		want = ""
	} else if expected.nillable && mode == "ceros" {
		switch expected.kind {
		case "dateTime":
			want = "0001-01-01T00:00:00Z"
		case "json":
			want = "{}"
		case "string", "bigInteger":
			want = ""
		default:
			want = "0"
		}
	}
	if actual.Text != want {
		t.Fatalf("%s: valor = %q; esperado %q", expected.name, actual.Text, want)
	}
}

func TestFacturasEscapeXML(t *testing.T) {
	// Arbitrary customer/product strings must remain data, including XML metacharacters.
	contract := xsdContracts[0]
	contract.sections = append([]xsdSection(nil), contract.sections...)
	for index, section := range contract.sections {
		fields := append([]xsdField(nil), section.fields...)
		for i := range fields {
			if fields[i].kind == "string" {
				fields[i].sample = `<dato> & "valor"`
			}
		}
		contract.sections[index].fields = fields
	}
	data, err := xml.Marshal(buildXSDDocument(t, contract, 1, "valores"))
	if err != nil {
		t.Fatal(err)
	}
	var node xsdNode
	if err := xml.Unmarshal(data, &node); err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(data), "<dato>") || !strings.Contains(string(data), "&lt;dato&gt;") {
		t.Fatal("el contenido no fue escapado como texto XML")
	}
	for i, section := range contract.sections {
		for j, field := range section.fields {
			checkXSDField(t, node.Children[i].Children[j], field, "valores")
		}
	}
}
