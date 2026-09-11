# Prompt para Mejorar el Codigo Base

Copia y pega el siguiente contenido completo en un asistente de IA (Claude, ChatGPT, etc.)
para obtener un ZIP con el proyecto arrancable. Si el adjunto es una carcasa (docs/placeholders),
el asistente debe materializar la estructura del stack del briefing, sin resolver las fases del reto.

---

```
## Briefing del reto (autoridad)
Este bloque manda sobre los archivos adjuntos. El stack y el rol salen de AQUÍ, no de un topic genérico ni de markdown placeholder.

### Contexto técnico original
Diseñar microservicios en Go con gRPC, Docker y Kubernetes

### Reto
- Tema: Microservicios en Go
- Seniority: junior-l2
- Tipo: practical
- Título: Diseño y Despliegue de Microservicios en Go
- Tiempo estimado: 8 horas

### Fases (trabajo del HUMANO — PROHIBIDO completarlas)
No implementes estos entregables. Dejalos como hueco pedagógico. El asistente solo materializa el proyecto arrancable para que el participante pueda trabajar.
- Fase 1: Definición del Microservicio — objetivo: Definir las operaciones y las reglas de validación del microservicio de procesamiento de pagos. — entregable (NO resolver): Documento de diseño que describe las operaciones, reglas de validación y umbrales de rendimiento.
- Fase 2: Implementación del Microservicio — objetivo: Implementar el microservicio en Go utilizando gRPC para la comunicación entre servicios. — entregable (NO resolver): Código fuente del microservicio implementado en Go con gRPC.
- Fase 3: Despliegue del Microservicio — objetivo: Desplegar el microservicio en un entorno de Kubernetes utilizando Docker. — entregable (NO resolver): Configuración de Kubernetes y Docker para el despliegue del microservicio.

Eres un asistente experto en análisis, corrección y generación de archivos de cualquier tipo:
código fuente, documentación, hojas de cálculo, documentos Word, configuraciones, entre otros.
Voy a enviarte una cadena de texto que contiene uno o más archivos. Cada archivo está delimitado por un marcador con el siguiente formato:
// === ARCHIVO: ruta/del/archivo.extension ===
o también puede aparecer como:
## === ARCHIVO: ruta/del/archivo.extension ===
Lo que sigue al marcador puede ser:

El contenido real del archivo (código, texto, YAML, etc.)
Una descripción en lenguaje natural de lo que debe contener el archivo


TU TAREA
PASO 0 — ¿Esto es un proyecto o una carcasa?
Antes de extraer archivos, leé el Briefing (si está) y diagnosticá el adjunto.

Es CARCASA si ocurre CUALQUIERA de estas:
- No hay manifiesto de dependencias del stack del briefing (manifest.json de VTEX IO / package.json / pom.xml / build.gradle / requirements.txt / go.mod / *.tf / *.csproj, según corresponda)
- Hay un "binario" que en realidad es un comentario ("no puede ser mostrado como texto plano", placeholder .fig/.docx vacío)
- Los markdowns ya completan entregables de fases posteriores ("se implementó fade-in", lista de áreas ya resuelta)

Si es CARCASA:
- MATERIALIZÁ un proyecto que arranca en el stack del briefing (VTEX IO Store Framework, Angular, Terraform, pytest, Nest, etc.). Incluí manifiesto, punto de entrada y capa de interfaz reales.
- NO copies los markdowns de "solución" como si fueran el producto. Son ruido de generación.
- NO resuelvas las fases del briefing (están marcadas PROHIBIDO). Dejá el hueco pedagógico: el flujo existe, las microinteracciones/calidad/infra que el reto pide NO están hechas.
- Después seguí al PASO 5 (ZIP).

Si es un proyecto REAL (manifiesto + código que compila o arranca):
- Seguí PASO 1 en adelante. 🔴 compilación sí. 🟡 pedagógico no.

PASO 1 — Detección y extracción
Identifica todos los archivos presentes en la cadena. Para cada archivo extrae:

Su ruta completa (ej: src/main/java/com/pragma/Service.java)
Su contenido o descripción

PASO 2 — Clasificación por tipo
Clasifica cada archivo en una de estas categorías:
A) Código fuente (Java, Python, TypeScript, JavaScript, Kotlin, etc.)
B) Configuración / documentación (YAML, properties, Markdown, JSON, txt, etc.)
C) Excel (.xlsx, .xls, .csv)
D) Word (.docx, .doc)
E) Otro tipo de archivo binario o especial
PASO 3 — Clasificación de errores en código fuente

Objetivo prioritario: que el proyecto compile. No corrijas flujo de negocio ni lógica funcional.

Antes de modificar cualquier archivo de código fuente, clasifica cada problema encontrado en una de estas dos categorías:
🔴 ERROR DE COMPILACIÓN — corregir siempre
Son errores que impiden que el proyecto arranque, sin valor pedagógico:

Import faltante o incorrecto
Clase, método o variable referenciada que no existe en ningún archivo del proyecto
Error de sintaxis
Anotación con atributos inválidos
Dependencia ausente en pom.xml, package.json, etc.
Archivo referenciado que no existe y debe ser creado con implementación mínima

→ CORREGIR estos errores.
🟡 PROBLEMA FUNCIONAL O DE CALIDAD — preservar siempre
Son problemas que no impiden compilar. Pueden ser intencionales para el aprendizaje:

Clave secreta hardcodeada ("secret", "password123")
API deprecada que funciona pero tiene reemplazo moderno
Lógica de negocio incorrecta o incompleta
Código redundante o de baja legibilidad
Falta de validaciones en flujo de negocio
Patrones de diseño incorrectos pero funcionales
Concurrencia no segura
Configuración funcional pero no óptima

→ PRESERVAR tal cual. No corregir, no mejorar, no comentar.
PASO 4 — Procesamiento según tipo de archivo
Tipo A — Código fuente
Aplica únicamente las correcciones clasificadas como 🔴 ERROR DE COMPILACIÓN.
No alteres ningún elemento clasificado como 🟡 PROBLEMA FUNCIONAL O DE CALIDAD.
Si falta un archivo referenciado, créalo con la implementación mínima necesaria para compilar.
Tipo B — Configuración / documentación
Extrae el contenido tal cual, sin modificaciones salvo errores evidentes de sintaxis
(ej: YAML mal indentado).
Tipo C — Excel (.xlsx)
Si viene con contenido real, genera el archivo respetando ese contenido.
Si viene con descripción en lenguaje natural, genera un archivo Excel funcional con:

Fila de encabezados en negrita con color de fondo distintivo
Columnas con ancho ajustado al contenido
Tipos de dato correctos por columna
Validaciones si la descripción lo indica
Hojas nombradas descriptivamente si hay más de una
Filas de ejemplo si no hay datos reales

Tipo D — Word (.docx)
Si viene con contenido real, genera el archivo respetando ese contenido.
Si viene con descripción en lenguaje natural, genera un documento Word funcional con:

Estilos de título (Título 1, Título 2) para jerarquía de secciones
Fuente legible (Calibri o equivalente), tamaño 11-12pt para cuerpo
Márgenes estándar
Tabla de contenido si tiene múltiples secciones
Tablas con encabezados en negrita si aplica

Tipo E — Otro
Genera el archivo con el contenido o estructura más apropiada según la descripción.
PASO 5 — Exportación en ZIP
Empaqueta todos los archivos en un único archivo ZIP descargable respetando exactamente
la estructura de rutas indicada por los marcadores.
El ZIP debe incluir:

Archivos de código con únicamente los errores de compilación corregidos
Archivos de configuración y documentación sin cambios
Archivos nuevos creados para resolver dependencias de compilación faltantes
Archivos Excel y Word generados desde descripción

IMPORTANTE: El ZIP debe estar listo para descargar al finalizar. No preguntes si el usuario
quiere generarlo. Simplemente genera el archivo y proporciona el enlace de descarga; No debes desplegar en el chat el resumen de lo que arreglaste al Zip, solo entregalo.

REGLAS IMPORTANTES

No omitas ningún archivo aunque no tenga errores ni modificaciones
Respeta los nombres y rutas exactas indicadas por los marcadores
Si un archivo no tiene marcador claro, infiere el nombre desde su contenido
Si la cadena contiene solo documentación, placeholders o binarios fake, NO la reproduzcas:
aplicá PASO 0 (materializar el proyecto del briefing). Reproducir la carcasa es un fallo.
No agregues texto después del enlace de descarga del ZIP
No preguntes si el usuario quiere el ZIP: simplemente generalo siempre
Si detectas que falta un archivo de configuración necesario para compilar
(pom.xml, package.json, requirements.txt, build.gradle, etc.), créalo e inclúyelo
inferiendo su contenido desde los imports y frameworks detectados en el código
Nunca corrijas problemas 🟡 aunque parezcan obvios o fáciles de mejorar.
El participante que recibirá este proyecto los debe encontrar y resolver él mismo.


INPUT
Aquí está la cadena con los archivos:

// === ARCHIVO: src/payment/payment.go ===
package payment

import (
	"context"
	"log"
	"time"

	pb "fintech/payment/api"
	auditpb "fintech/audit/api"
	"google.golang.org/grpc"
)

type server struct {}

func (s *server) ProcessPayment(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	log.Println("Processing payment", req)
	// Validación de la solicitud de pago
	if req.Amount <= 0 {
		return nil, &pb.PaymentError{Message: "Invalid amount"}
	}
	// Simular procesamiento de pago
	time.Sleep(100 * time.Millisecond)
	// Comunicación con el servicio de auditoría
	conn, err := grpc.DialContext(ctx, "audit:50051", grpc.WithInsecure())
	if err!= nil {
		return nil, err
	}
	defer conn.Close()
	auditClient := auditpb.NewAuditServiceClient(conn)
	_, err = auditClient.LogPayment(ctx, &auditpb.LogPaymentRequest{PaymentId: req.Id})
	if err!= nil {
		return nil, err
	}
	return &pb.PaymentResponse{Status: "SUCCESS"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err!= nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPaymentServiceServer(s, &server{})
	log.Println("Server started at ":50051")
	if err := s.Serve(lis); err!= nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// === ARCHIVO: src/audit/audit.go ===
package audit

import (
	"context"
	"log"

	pb "fintech/audit/api"
	"google.golang.org/grpc"
)

type server struct {}

func (s *server) LogPayment(ctx context.Context, req *pb.LogPaymentRequest) (*pb.LogPaymentResponse, error) {
	log.Println("Logging payment", req)
	// Simular registro de auditoría
	time.Sleep(100 * time.Millisecond)
	return &pb.LogPaymentResponse{Status: "LOGGED"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err!= nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuditServiceServer(s, &server{})
	log.Println("Server started at ":50052")
	if err := s.Serve(lis); err!= nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// === ARCHIVO: config/config.yaml ===
payment:
  port: 50051
audit:
  port: 50052

// === ARCHIVO: Dockerfile ===
FROM golang:1.22 AS builder
WORKDIR /app
COPY..
RUN go build -o payment./src/payment

FROM alpine:latest
COPY --from=builder /app/payment /usr/local/bin/payment
EXPOSE 50051
CMD ["payment"]

// === ARCHIVO: docker-compose.yml ===
version: '3'
services:
  payment:
    build:.
    ports:
      - "50051:50051"
  audit:
    image: fintech/audit
    ports:
      - "50052:50052"

// === ARCHIVO: scripts/healthcheck.sh ===
#!/bin/sh
set -e

PAYMENT_PORT=${PAYMENT_PORT:-50051}
AUDIT_PORT=${AUDIT_PORT:-50052}

check_port() {
  local port=$1
  for i in $(seq 1 5); do
    if nc -z localhost $port; then
      return 0
    fi
    sleep 1
  done
  return 1
}

if! check_port $PAYMENT_PORT; then
  echo "Payment service is not healthy"
  exit 1
fi

if! check_port $AUDIT_PORT; then
  echo "Audit service is not healthy"
  exit 1
fi

echo "Services are healthy"

// === ARCHIVO: tests/payment_test.go ===
package payment

import (
	"context"
	"testing"

	pb "fintech/payment/api"
	"google.golang.org/grpc"
)

func TestProcessPayment(t *testing.T) {
	conn, err := grpc.DialContext(context.Background(), "localhost:50051", grpc.WithInsecure())
	if err!= nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewPaymentServiceClient(conn)
	req := &pb.PaymentRequest{Id: "1", Amount: 100}
	res, err := client.ProcessPayment(context.Background(), req)
	if err!= nil {
		t.Fatalf("ProcessPayment failed: %v", err)
	}
	if res.Status!= "SUCCESS" {
		t.Fatalf("expected SUCCESS, got %v", res.Status)
	}
}
```
