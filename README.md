# Diseño y Despliegue de Microservicios en Go

La empresa de fintech 'FinTech Solutions' está migrando sus servicios monolíticos a una arquitectura basada en microservicios. Tu tarea es diseñar y desplegar un microservicio en Go que se comunicará con otros servicios a través de gRPC. El microservicio debe manejar solicitudes de procesamiento de pagos, validando la información del pago y comunicando el resultado a un servicio de auditoría. El sistema debe soportar un mínimo de 1 000 solicitudes por segundo con un tiempo de respuesta promedio menor a 500ms.

## Informacion General

| Campo | Valor |
|-------|-------|
| **Tema** | Microservicios en Go |
| **Nivel** | junior-l2 |
| **Tipo** | practical |
| **Tiempo estimado** | 8 horas |

## Fases del Reto

### Fase 0: Configuración del Proyecto

**Objetivo:** Obtener el proyecto base funcional enviando el Código Base a un asistente de IA, que lo analizará, corregirá errores y generará un ZIP listo para usar.

**Tiempo estimado:** 15-30 minutos

**Instrucciones:**

- Asegúrate de tener instalado para ejecutar el proyecto: Un IDE o editor de código.
- Copia todo el contenido del campo **Código Base** de este reto — incluyendo el texto de instrucciones que aparece al inicio.
- Abre un asistente de IA (Claude en claude.ai, ChatGPT o Gemini — se recomienda Claude), pega el contenido copiado en el chat y envíalo.
- El asistente analizará los archivos, corregirá errores y generará un archivo ZIP descargable. Descárgalo y extráelo en la carpeta donde quieras trabajar.
- Verifica que el proyecto arranca sin errores.

**Entregable:** El proyecto compila/arranca sin errores.

<details>
<summary>Pistas de conocimiento</summary>

- Copia el Código Base completo incluyendo el texto de instrucciones al inicio — esas instrucciones le indican al asistente exactamente qué hacer con los archivos.
- Si el asistente no genera el ZIP automáticamente al terminar el análisis, escríbele: "genera el ZIP ahora".
- Si el proyecto tiene errores al arrancar, comparte el mensaje de error con el mismo asistente para que lo corrija.

</details>

### Fase 1: Definición del Microservicio

**Objetivo:** Definir las operaciones y las reglas de validación del microservicio de procesamiento de pagos.

**Tiempo estimado:** 2 horas

**Instrucciones:**

- Identificar las operaciones que el microservicio debe soportar.
- Definir las reglas de validación para las solicitudes de pago.
- Establecer los umbrales de rendimiento y disponibilidad del microservicio.

**Entregable:** Documento de diseño que describe las operaciones, reglas de validación y umbrales de rendimiento.

<details>
<summary>Pistas de conocimiento</summary>

- Considera los diferentes tipos de pagos que el sistema debe soportar.
- Piensa en los posibles errores de validación y cómo manejarlos.

</details>

### Fase 2: Implementación del Microservicio

**Objetivo:** Implementar el microservicio en Go utilizando gRPC para la comunicación entre servicios.

**Tiempo estimado:** 4 horas

**Instrucciones:**

- Implementar las operaciones definidas en la fase anterior.
- Utilizar gRPC para la comunicación con otros servicios.
- Asegurar que el microservicio cumpla con los umbrales de rendimiento y disponibilidad establecidos.

**Entregable:** Código fuente del microservicio implementado en Go con gRPC.

<details>
<summary>Pistas de conocimiento</summary>

- Utiliza gRPC para definir los servicios y las llamadas entre microservicios.
- Asegúrate de manejar correctamente los errores y las respuestas del servicio.

</details>

### Fase 3: Despliegue del Microservicio

**Objetivo:** Desplegar el microservicio en un entorno de Kubernetes utilizando Docker.

**Tiempo estimado:** 2 horas

**Instrucciones:**

- Empaquetar el microservicio en un contenedor Docker.
- Desplegar el contenedor en un clúster de Kubernetes.
- Configurar el balanceo de carga y la escalabilidad del microservicio.

**Entregable:** Configuración de Kubernetes y Docker para el despliegue del microservicio.

<details>
<summary>Pistas de conocimiento</summary>

- Utiliza Dockerfile para empaquetar el microservicio en un contenedor.
- Configura los servicios y los despliegues en Kubernetes para asegurar el balanceo de carga y la escalabilidad.

</details>

## Dimensiones Evaluadas

- **queEs**: ¿Qué es un microservicio y por qué se utiliza en este caso?
- **paraQueSirve**: ¿Para qué sirve el microservicio de procesamiento de pagos en el contexto de FinTech Solutions?
- **comoSeUsa**: ¿Cómo se usa gRPC para la comunicación entre microservicios?
- **erroresComunes**: ¿Cuáles son los errores comunes que pueden ocurrir al implementar y desplegar un microservicio?
- **queDecisionesImplica**: ¿Qué decisiones implica el despliegue de un microservicio en Kubernetes utilizando Docker?

## Criterios de Evaluacion

- Definición clara de las operaciones y reglas de validación del microservicio.
- Implementación correcta del microservicio en Go utilizando gRPC.
- Despliegue exitoso del microservicio en Kubernetes utilizando Docker.

## Como trabajar con un asistente de IA

- **AGENTS.md** — instrucciones nativas del repo (Cursor, Codex, Copilot, Gemini, Claude Code). Abrí el proyecto y el agente las carga solo.
- **PROMPT_MEJORA.md** — el mismo prompt, para copiar y pegar en un chat (claude.ai, ChatGPT, etc.).

---

*Reto generado automaticamente por Challenge Generator - Pragma*
