package invoices_test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"encoding/xml"
	"math/big"
	"testing"
	"time"

	"github.com/beevik/etree"
	"github.com/ron86i/go-siat/v3/pkg/utils"
	dsig "github.com/russellhaering/goxmldsig"
)

func TestFacturasV3FirmaElectronica(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	template := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		NotBefore:    time.Now().Add(-time.Hour),
		NotAfter:     time.Now().Add(time.Hour),
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}
	der, err := x509.CreateCertificate(rand.Reader, template, template, &key.PublicKey, key)
	if err != nil {
		t.Fatal(err)
	}
	certificate, err := x509.ParseCertificate(der)
	if err != nil {
		t.Fatal(err)
	}
	signer, err := utils.NewXMLDocumentSigner(
		pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)}),
		pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der}),
	)
	if err != nil {
		t.Fatal(err)
	}
	for _, contract := range xsdV3Contracts {
		t.Run(contract.name, func(t *testing.T) {
			input, err := xml.Marshal(buildXSDV3Document(t, contract, 1, "valores"))
			if err != nil {
				t.Fatal(err)
			}
			signed, err := signer.SignXML(input)
			if err != nil {
				t.Fatal(err)
			}
			var node xsdV3Node
			if err := xml.Unmarshal(signed, &node); err != nil {
				t.Fatal(err)
			}
			position := 0
			for _, section := range contract.sections {
				for range section.count {
					if position >= len(node.Children) {
						t.Fatalf("falta %s después de firmar", section.name)
					}
					actual := node.Children[position]
					position++
					if actual.XMLName.Local != section.name || len(actual.Children) != len(section.fields) {
						t.Fatalf("se alteró %s al firmar", section.name)
					}
					for i, field := range section.fields {
						checkXSDV3Field(t, actual.Children[i], field, "valores")
					}
				}
			}
			if len(node.Children) != position+1 {
				t.Fatalf("cantidad de nodos tras la firma = %d; esperada %d", len(node.Children), position+1)
			}
			signature := node.Children[position]
			if signature.XMLName.Local != "Signature" || signature.XMLName.Space != "http://www.w3.org/2000/09/xmldsig#" {
				t.Fatal("la firma debe aparecer después de los datos fiscales")
			}
			doc := etree.NewDocument()
			if err := doc.ReadFromBytes(signed); err != nil {
				t.Fatal(err)
			}
			validation := dsig.NewDefaultValidationContext(&dsig.MemoryX509CertificateStore{Roots: []*x509.Certificate{certificate}})
			if _, err := validation.Validate(doc.Root()); err != nil {
				t.Fatalf("firma XML inválida: %v", err)
			}
		})
	}
}
