# Arquitectura del SDK

El SDK `go-siat` está diseñado siguiendo los principios de la **Arquitectura Hexagonal (Puertos y Adaptadores)**, lo que permite un desacoplamiento total entre la lógica de negocio técnica y las implementaciones de transporte (SOAP/HTTP).

## Estructura de Capas

### 1. Dominio (Core)
Ubicación: `internal/core/domain`
Contiene las estructuras de datos puras, mapeos XML y lógica de negocio que no depende de factores externos.

### 2. Puertos (Interfaces)
Ubicación: `internal/core/ports`
Define las interfaces que deben implementar los servicios (Códigos, Sincronización, etc.). Esto permite, por ejemplo, crear mocks fácilmente para pruebas.

### 3. Adaptadores (Infraestructura)
Ubicación: `internal/adapter`
- **Services**: Implementación concreta de los puertos que se comunica con los WSDL del SIAT.
- **HTTP**: Configuración del cliente y transporte SOAP.

### 4. API Pública (Models)
Ubicación: `pkg/models`
Proporciona la interfaz "opaca" para el usuario final mediante el patrón **Builder**, ocultando la complejidad interna del esquema XML de Impuestos.

## Flujo de una Solicitud

1. El usuario inicializa el SDK mediante `siat.New()`.
2. Utiliza un **Builder** de `pkg/models/invoices` para construir su factura.
3. Llama a un método del servicio (ej. `Electronica().RecepcionFactura`).
4. El adaptador traduce el modelo público al modelo interno de dominio.
5. Se realiza la firma XML y el envío SOAP.
6. El SDK procesa la respuesta y la devuelve al usuario de forma simplificada.
