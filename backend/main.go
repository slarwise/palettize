package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"

	"github.com/gin-contrib/cors"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var tracer = otel.Tracer("backend")

func main() {
	ctx := context.Background()
	tp, err := initTracer(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	r := gin.Default()
	r.Use(otelgin.Middleware("backend"))
	r.Use(cors.Default())

	r.POST("/convert", func(c *gin.Context) {
		file, err := c.FormFile("img")
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatus(400)
			return
		}
		inputPath := "input.png"
		if err := c.SaveUploadedFile(file, inputPath); err != nil {
			c.AbortWithStatus(500)
			return
		}
		magickCmd, err := exec.LookPath("convert")
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatus(500)
			return
		}
		colorscheme := c.Query("colorscheme")
		cmd := exec.Command(magickCmd, inputPath, "+dither", "-remap", filepath.Join("colorschemes", colorscheme+".png"), "output.png")
		if err = cmd.Run(); err != nil {
			fmt.Println(err)
			c.AbortWithStatus(500)
			return
		}
		c.File("output.png")
	})

	r.Run(":3001")
}

func initTracer(ctx context.Context) (*sdktrace.TracerProvider, error) {
	exporter, err := otlptracehttp.New(ctx, otlptracehttp.WithInsecure(), otlptracehttp.WithEndpoint("tempo:4318"))
	if err != nil {
		return nil, err
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp, nil
}
