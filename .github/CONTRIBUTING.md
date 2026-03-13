# Guía de Contribución

¡Gracias por tu interés en colaborar en **go-siat**! Toda ayuda es bienvenida para mejorar este proyecto y facilitar la integración con el SIAT.

Al participar en este proyecto, te comprometes a seguir nuestro [Código de Conducta](./CODE_OF_CONDUCT.md).

## Cómo empezar

1. **Reportar Bugs**: Si encuentras un error, por favor abre un *Issue* en GitHub detallando los pasos para reproducirlo y el comportamiento esperado.
2. **Sugerir Mejoras**: Si tienes una idea para una nueva funcionalidad, abre un *Issue* para discutirla antes de empezar a trabajar en ella.
3. **Pull Requests (PRs)**: 
    - Crea una rama (*branch*) descriptiva para tus cambios.
    - Asegúrate de que el código siga las convenciones estándar de Go (`go fmt`).
    - Incluye pruebas unitarias si estás agregando lógica.
    - Abre el PR hacia la rama principal y espera la revisión.

## Configuración del Entorno

1. **Requisitos**: Go 1.21 o superior.
2. **Clonar el repositorio**:
   ```bash
   git clone https://github.com/ron86i/go-siat.git
   cd go-siat
   ```
3. **Dependencias**:
   ```bash
   go mod download
   ```
4. **Variables de Entorno**: Copia el archivo `.env.example` (si existe) a `.env` y configura tus credenciales de prueba del SIAT.

## Estructura del Proyecto

El SDK sigue una arquitectura de puertos y adaptadores para desacoplar la lógica de negocio de las implementaciones técnicas:

- `pkg/`: Lógica pública y puntos de entrada para el usuario.
  - `models/`: Constructores (Builders) e interfaces opacas para las solicitudes.
  - `config/`: Tipos de configuración para el cliente.
  - `utils/`: Utilidades generales (CUF, formatos de fecha).
- `internal/`: Lógica interna protegida.
  - `core/domain/`: Estructuras de datos internas y mapeos XML/SOAP.
  - `core/port/`: Definición de interfaces de los servicios (puertos).
  - `adapter/service/`: Implementaciones concretas de los servicios SIAT (adaptadores).
- `siat.go`: Orquestador principal y punto de entrada unificado.

## Flujo de Trabajo (Git)

- **Ramas**: Usa nombres descriptivos como `feat/nueva-funcionalidad` o `fix/error-especifico`.
- **Commits**: Se recomienda seguir los [Conventional Commits](https://www.conventionalcommits.org/es/v1.0.0/):
  - `feat`: Nueva funcionalidad.
  - `chore`: Cambios en la configuración o dependencias.
  - `refactor`: Refactorización del código.
  - `perf`: Cambios que mejoran el rendimiento.
  - `fix`: Corrección de un error.
  - `docs`: Cambios en la documentación.
  - `test`: Añadir o corregir pruebas.

## Estilo de Código y Calidad

- **Formato**: Ejecuta `go fmt ./...` antes de enviar tus cambios.
- **Idioma**: Mantén la coherencia con la estructura ya establecida. Usa nombres descriptivos en español para la lógica de negocio técnica (ej. `SolicitudCufd`) y comentarios en español.
- **Pruebas**:
  ```bash
  go test ./...
  ```
  Asegúrate de que todas las pruebas existentes pasen y añade nuevas pruebas para cubrir la funcionalidad agregada.

---

¡Esperamos tus contribuciones!
