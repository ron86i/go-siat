# Guía de Contribución

Agradecemos el interés en contribuir a go-siat. Este documento describe el proceso y las mejores prácticas.

[English Version](../../.github/CONTRIBUTING.md)

## 📋 Requisitos Previos

- Go 1.25 o superior
- Git
- Familiaridad con SIAT (Sistema Integrado de Administración Tributaria)
- Lectura de la [Guía de Arquitectura](../../docs/es/arquitectura.md)

## 🚀 Comenzar

### 1. Fork y Clonar

```bash
# Fork el repositorio en GitHub
git clone https://github.com/tu-usuario/go-siat.git
cd go-siat
```

### 2. Crear una Rama

```bash
# Cree una rama para su feature o fix
git checkout -b feature/mi-feature
# o
git checkout -b fix/numero-issue
```

### 3. Configurar el Ambiente

```bash
# Descargar dependencias
go mod tidy

# Verificar que todo compila
go build ./...

# Ejecutar los tests
go test ./...
```

## 💻 Desarrollo

### Estándares de Código

#### Comentarios y Documentación

- **Godoc**: Todo tipo, función o método público debe tener comentario en formato godoc
- **Ejemplos**: Incluir ejemplos en tests para documentación

```go
// SolicitudCuis obtiene el Código Único de Inicio de Sistemas necesario para operar ante el SIAT.
//
// Parámetros:
//   - ctx: Contexto para cancelación e timeouts
//   - config: Configuración con token de autenticación
//   - req: Solicitud CUIS construida con builder
//
// Retorna EnvelopeResponse con la respuesta SOAP o error.
func (s *SiatCodigosService) SolicitudCuis(ctx context.Context, config Config, req models.Cuis) (*soap.EnvelopeResponse[codigos.CuisResponse], error) {
```

#### Nombres

- Variables: `camelCase` (ej: `cufsResponse`, `errorMessage`)
- Constantes: `UPPER_SNAKE_CASE` (ej: `DEFAULT_TIMEOUT`, `MAX_RETRIES`)
- Funciones Públicas: `PascalCase` (ej: `SolicitudCuis`)
- Métodos Privados: `camelCase` (ej: `buildSoapEnvelope`)

#### Errores

- Siempre usar `fmt.Errorf` o `errors.New` con contexto
- Evitar retornar `nil` para errores sin usar `error` como tipo

```go
// Bien
if err != nil {
    return nil, fmt.Errorf("falló compilar SOAP: %w", err)
}

// Evitar
if err != nil {
    return nil, err  // Pierde contexto
}
```

#### Type-Safety

- Preferir tipos específicos sobre `interface{}`
- Usar builders cuando sea posible

```go
// Bien
type SolicitudCuis struct {
    CodigoAmbiente int64
    NIT            int64
}

// Evitar
type SolicitudCuis map[string]interface{}
```

### Estructura de Tests

Los tests deben seguir el patrón de tabla (table-driven tests):

```go
func TestSolicitudCuis(t *testing.T) {
    tests := []struct {
        name      string
        req       models.Cuis
        wantErr   bool
        setupMock func()
    }{
        {
            name: "solicitud válida",
            req: models.Codigos().NewCuisBuilder().
                WithCodigoAmbiente(1).
                Build(),
            wantErr: false,
        },
        {
            name: "ambiente inválido",
            req: models.Codigos().NewCuisBuilder().
                WithCodigoAmbiente(99).
                Build(),
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation
        })
    }
}
```

### Commits

- **Mensajes claros**: Describe qué y por qué, no cómo
- **Scope**: `[feature|fix|docs|test|refactor]: descripción`

```bash
# Bien
git commit -m "fix: validar NIT antes de enviar CUFD" \
           -m "El SIAT rechaza CUFD con inconsistencias en el NIT" \
           -m "Closes #123"

# Evitar
git commit -m "arreglo de bug"
```

## 🧪 Testing

### Cobertura Mínima

- **Código nuevo**: 80% de cobertura
- **Cambios a código existente**: No disminuir la cobertura actual

```bash
# Ejecutar tests con cobertura
go test -cover ./...

# Generar reporte HTML de cobertura
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Tipos de Tests

| Tipo | Ubicación | Propósito |
|------|-----------|-----------|
| **Unit** | `*_test.go` en el mismo paquete | Verificar lógica en aislamiento |
| **Integration** | `*_integration_test.go` | Verificar comportamiento con SIAT real (solo en CI) |
| **Regression** | `testdata/` | Validar contra respuestas conocidas |

### Variables de Entorno para Tests

```bash
# Tests de integración requieren estos valores
export SIAT_API_TOKEN="tu_token"
export SIAT_NIT="tu_nit"
export SIAT_BASE_URL="https://pilotosiatservicios.impuestos.gob.bo/v2"

# Ejecutar tests de integración
go test -tags=integration ./...
```

## 📦 Tipos de Contribución

### 1. Nuevas Funcionalidades

- Crear issue primero describiendo el cambio
- Esperar aprobación antes de comenzar
- Incluir tests y documentación
- Actualizar ejemplos si aplica

### 2. Bug Fixes

- Incluir test que reproduca el bug
- Arreglarlo
- Verificar que el test pase
- Documentar en el commit

### 3. Mejoras en Documentación

- Actualizar README.md o archivos en el directorio `docs/`
- Mejorar comentarios godoc
- Agregar ejemplos
- No necesita aprobación previa

### 4. Mejoras en Performance

- Incluir benchmarks que demuestren la mejora
- No sacrificar legibilidad
- Documentar la decisión de diseño

## 🔄 Proceso de Pull Request

### Paso 1: Preparar el PR

```bash
# Verificar que todo está ok
go fmt ./...
go vet ./...
go test ./...

# Push a tu rama
git push origin feature/mi-feature
```

### Paso 2: Crear el Pull Request

En GitHub:

1. Título: `[tipo] descripción breve`
2. Descripción: Explicar el cambio, por qué, y qué pruebas se incluyeron
3. Referencia: `Closes #issueNumber` si aplica

**Template de descripción:**

```markdown
## Descripción
Breve descripción del cambio.

## Tipo de Cambio
- [ ] Bug fix
- [ ] Nueva funcionalidad
- [ ] Breaking change
- [ ] Mejora de documentación

## Testing
- [ ] Tests unitarios agregados/modificados
- [ ] Tests de integración (si aplica)
- [ ] Cobertura >= 80%

## Checklist
- [ ] Código sigue los estándares del proyecto
- [ ] Todos los tests pasan
- [ ] Documentación actualizada
- [ ] Sin problemas de linting
```

### Paso 3: Revisión

- Esperar revisión de los maintainers
- Responder comentarios en el PR
- Hacer cambios solicitados
- Re-request review

### Paso 4: Merge

Una vez aprobado, un maintainer hará el merge.

## 📋 Checklist de Calidad

Antes de abrir un PR:

- [ ] `go fmt ./...` ha sido ejecutado
- [ ] `go vet ./...` sin errores
- [ ] `go test ./...` pasan todos los tests
- [ ] Cobertura de tests >= 80% para código nuevo
- [ ] Todos los tipos/funciones públicas tienen documentación godoc
- [ ] Ejemplos en comentarios si es complejo
- [ ] Sin cambios no relacionados (keep PRs focused)
- [ ] Commits con mensajes claros
- [ ] README.md actualizado si hay cambios en la API

## 🐛 Reportar Bugs

1. Verificar que no exista un issue similar
2. Crear issue con:
   - Título descriptivo
   - Descripción detallada
   - Código que reproduce el problema
   - Versión de Go y go-siat
   - Versión del SIAT (producción/pruebas)

**Ejemplo:**

```markdown
## Descripción
CUFD rechazado con error "Código de Control inválido"

## Pasos para Reproducir
1. Llamar a SolicitudCufd con...
2. Observar el error...

## Salida Esperada
Código CUFD válido

## Información del Sistema
- Go: 1.25
- go-siat: v1.2.0
- SIAT: Producción
```

## 🎓 Convenciones del Proyecto

### Estructura de Paquetes

- `pkg/`: API pública
- `internal/`: Implementación privada
- Mantener paquetes enfocados en una responsabilidad

### Dependencias

- Mantener `go.mod` limpio
- Justificar nuevas dependencias externas
- Preferir stdlib cuando sea posible

### Versionado Semántico

- MAJOR: Breaking changes
- MINOR: Nuevas funcionalidades (backward compatible)
- PATCH: Bug fixes

## 💬 Comunicación

- **Issues**: Para reportar bugs o proponer features
- **Discussions**: Para preguntas y conversaciones generales
- **Email**: Contactar a los maintainers si es necesario

## 📖 Recursos

- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Guía de Arquitectura](../../docs/es/arquitectura.md) - Diseño del proyecto
- [SIAT Documentation](https://www.impuestos.gob.bo/) - Documentación oficial SIAT

## ✨ Gracias

Gracias por contribuir a go-siat. Tus aportes hacen que la facturación electrónica en Bolivia sea más accesible para todos los desarrolladores.
