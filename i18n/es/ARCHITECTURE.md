# Arquitectura de go-siat

Este documento describe la arquitectura, estructura de carpetas y patrones de diseño del SDK go-siat.

[English Version](../../ARCHITECTURE.md)

## 🏗️ Visión General

go-siat implementa una **arquitectura modular inspirada en Hexagonal Architecture (Puertos y Adaptadores)** con una clara separación de responsabilidades:

```
┌─────────────────────────────────────────────────────────────┐
│                   Código del Usuario                         │
│          (Utiliza el SDK a través de SiatServices)          │
└─────────────────────┬───────────────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────────────┐
│              SiatServices (Punto de Entrada)                │
│   - Operaciones()                                           │
│   - Sincronizacion()                                        │
│   - Codigos()                                               │
│   - CompraVenta()                                           │
│   - Computarizada()                                         │
│   - Electronica()                                           │
└─────────────────────┬───────────────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────────────┐
│              Puertos (Interfaces)                            │
│   - SiatCodigosService                                      │
│   - SiatSincronizacionService                               │
│   - SiatOperacionesPort                                     │
│   - SiatCompraVentaService                                  │
│   - SiatComputarizadaService                                │
│   - SiatElectronicaService                                  │
└─────────────────────┬───────────────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────────────┐
│            Adaptadores (Implementaciones)                    │
│   - SiatCodigosService (HTTP/SOAP)                          │
│   - SiatSincronizacionService (HTTP/SOAP)                   │
│   - SiatOperacionesService (HTTP/SOAP)                      │
│   - ... (otros servicios)                                   │
└─────────────────────┬───────────────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────────────┐
│              Núcleo de Dominio                              │
│   - Modelos SOAP                                            │
│   - Estructuras de Respuesta                                │
│   - Tipos de Datos del SIAT                                 │
└─────────────────────┬───────────────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────────────┐
│             Utilidades (Capa de Utilidad)                   │
│   - XML Signing (XMLDSig)                                   │
│   - Compresión y Hash                                       │
│   - Codificación Base64                                     │
│   - Generación de CUF                                       │
└─────────────────────────────────────────────────────────────┘
```

## 📂 Estructura de Carpetas

```
go-siat/
├── README.md                      # Documentación principal
├── ARCHITECTURE.md               # Este archivo
├── LICENSE                       # MIT License
├── go.mod                        # Módulo Go
│
├── pkg/                          # Código público del SDK
│   ├── config/                   # Configuración
│   ├── models/                   # Modelos de datos del usuario
│   │   ├── codigos.go           # Builders para códigos
│   │   ├── electronica.go       # Builders para facturas electrónicas
│   │   ├── computarizada.go     # Builders para facturas computarizadas
│   │   ├── compra_venta.go      # Builders para compra-venta
│   │   ├── operaciones.go       # Builders para operaciones
│   │   ├── sincronizacion.go    # Builders para sincronización
│   │   └── invoices/             # 35 sectores específicos
│   │       ├── compra_venta.go
│   │       ├── hoteles.go
│   │       ├── mineria.go
│   │       ├── ... (otros sectores)
│   │
│   └── utils/                    # Utilidades públicas
│       ├── crypto.go             # Funciones criptográficas
│       ├── cuf.go                # Generación de CUF
│       ├── encoding.go           # Codificación/Decodificación
│       ├── parse.go              # Parseo de estructuras
│       └── signXML.go            # Firma digital XMLDSig
│
├── internal/                     # Código privado del SDK
│   │
│   ├── adapter/                  # Implementaciones concretas (Adaptadores)
│   │   └── service/              # Servicios HTTP/SOAP
│   │       ├── siat_codigos_service.go
│   │       ├── siat_electronica_service.go
│   │       ├── siat_operaciones_service.go
│   │       ├── siat_sincronizacion_service.go
│   │       ├── siat_compra_venta_service.go
│   │       └── siat_computarizada_service.go
│   │
│   └── core/                     # Lógica de dominio y puertos
│       ├── domain/               # Modelos del dominio
│       │   ├── datatype/         # Tipos de datos especializados
│       │   │   ├── soap/         # Estructuras SOAP
│       │   │   ├── nilable.go    # Tipos anulables seguros
│       │   │   └── ... (tipos)
│       │   │
│       │   └── siat/             # Modelos específicos del SIAT
│       │       ├── codigos/      # Respuestas de códigos
│       │       ├── electronica/  # Respuestas electrónicas
│       │       └── ... (otros)
│       │
│       └── port/                 # Puertos (Interfaces)
│           ├── config.go         # Configuración compartida
│           ├── siat_codigos_port.go
│           ├── siat_electronica_port.go
│           ├── siat_operaciones_port.go
│           ├── siat_sincronizacion_port.go
│           ├── siat_compra_venta_port.go
│           └── siat_computarizada_port.go
│
├── siat.go                       # Punto de entrada (SiatServices)
├── constants.go                  # Constantes globales
├── config.go                     # Exposición pública de Config
└── codigos_errores.go            # Catálogo de códigos de error
```

## 🔄 Patrones de Diseño

### 1. **Puertos y Adaptadores (Hexagonal Architecture)**

- **Puertos**: Son interfaces que definen contratos (ubicadas en `internal/core/port/`)
- **Adaptadores**: Son implementaciones concretas de esos puertos (ubicadas en `internal/adapter/service/`)

**Beneficio**: Fácil testing y cambio de implementaciones sin afectar el código cliente.

### 2. **Builder Pattern**

Utilizamos el patrón Builder para la construcción intuitiva de solicitudes complejas, proporcionando una sintaxis fluida y verificación de tipos en tiempo de compilación.

**Beneficio**: Sintaxis fluida, type-safe, y validaciones en tiempo de compilación.

### 3. **Dependency Injection**

Los servicios reciben dependencias (como `http.Client`) en su inicialización, en lugar de crear instancias globales.

**Beneficio**: Testeable, flexible, sin estado global.

### 4. **Type Aliasing**

Se utiliza para exponer tipos internos de forma pública sin duplicar definiciones.

**Beneficio**: Abstracción clara del paquete `internal/`.

## 🎯 Flujo de Datos

Toda operación sigue un flujo predecible:

1. El usuario construye una solicitud usando builders
2. Llama al método correspondiente en SiatServices 
3. SiatServices delega a la interfaz (Puerto)
4. El adaptador maneja la comunicación SOAP/HTTP
5. Se retorna la respuesta parseada y actualizada

## 🔌 Responsabilidades por Capa

| Capa | Responsabilidad |
|------|-----------------|
| **SiatServices** | Orquestar y exponer servicios al usuario |
| **Puertos** | Definir contratos y abstracciones |
| **Adaptadores** | Implementar lógica de comunicación SOAP/HTTP |
| **Dominio** | Modelar estructuras de datos SIAT |
| **Utilidades** | Proporcionar funciones transversales |

## 🧪 Arquitectura para Testing

Debido a la arquitectura de puertos y adaptadores:

1. **Tests Unitarios**: Mock fácil de interfaces (puertos)
2. **Tests de Integración**: Utilizar el cliente real con ambiente de pruebas
3. **Tests de Regresión**: Validar modelos contra respuestas conocidas

## 🔐 Seguridad

- **TLS 1.2+**: Configurado por defecto en `New()`
- **XMLDSig**: Integrado en utilidades para firma de facturas
- **Tokens**: Manejados securely en `Config`

## 📊 Decisiones Arquitectónicas

| Decisión | Razón |
|----------|-------|
| Struct methods sobre funciones globales | Namespace y organización |
| Interfaces pequeñas (un servicio por interfaz) | Flexibilidad y cleanup |
| Type-safe builders | Evitar errores en compilación |
| Context en todas las operaciones | Cancellations y timeouts |
| http.Client inyectable | Testing y configuración personalizada |

## 🚀 Cómo Agregar un Nuevo Servicio

1. Crear la interfaz en `internal/core/port/siat_nuevo_service.go`
2. Implementarly en `internal/adapter/service/siat_nuevo_service.go`
3. Crear modelos en `pkg/models/nuevo_modulo.go`
4. Exponer en `SiatServices` mediante getter
5. Actualizar documentación y ejemplos (revisar [CONTRIBUTING.md](../../CONTRIBUTING.md))

## 📚 Referencias

- [Hexagonal Architecture (Alistair Cockburn)](https://alistair.cockburn.us/hexagonal-architecture/)
- [Clean Architecture (Robert C. Martin)](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
