package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc/credentials"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

var (
	serviceName  = os.Getenv("SERVICE_NAME")
	collectorURL = os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	insecure     = os.Getenv("INSECURE_MODE")
)

func main() {
	// Validações de variáveis de ambiente
	if serviceName == "" {
		log.Fatal("SERVICE_NAME não está definido")
	}
	if collectorURL == "" {
		log.Fatal("OTEL_EXPORTER_OTLP_ENDPOINT não está definido")
	}

	// Inicializa o tracer
	cleanup := initTracer()
	defer cleanup(context.Background()) // ou um contexto mais adequado

	r := gin.Default()
	r.Use(otelgin.Middleware(serviceName))
	// Outras rotas e lógica...
}

func initTracer() func(context.Context) error {
	// Define as opções de segurança para o exportador
	secureOption := otlptracegrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
	if len(insecure) > 0 {
		secureOption = otlptracegrpc.WithInsecure()
	}

	// Usa um contexto com timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Cria o exportador com o contexto
	exporter, err := otlptrace.New(
		ctx,
		otlptracegrpc.NewClient(
			secureOption,
			otlptracegrpc.WithEndpoint(collectorURL),
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	// Define os recursos com o contexto
	resources, err := resource.New(
		ctx,
		resource.WithAttributes(
			attribute.String("service.name", serviceName),
			attribute.String("library.language", "go"),
		),
	)
	if err != nil {
		log.Printf("Could not set resources: %v", err)
	}

	// Configura o tracer provider
	otel.SetTracerProvider(
		sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithBatcher(exporter),
			sdktrace.WithResource(resources),
		),
	)
	return exporter.Shutdown
}
