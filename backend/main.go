package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"

	"go.opentelemetry.io/otel/attribute"
	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

var db = make(map[string]string)
var tracer = otel.Tracer("backend")

func main() {
	tp, err := initTracer()
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

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/colorschemes", func(c *gin.Context) {
		var json struct {
			Name  string `json:"name" binding:"required"`
			Color string `json:"color" binding:"required"`
		}
		err := c.BindJSON(&json)
		if err == nil {
			_, span := tracer.Start(
				c.Request.Context(),
				"db",
				trace.WithAttributes(attribute.String("name", json.Name), attribute.String("operation", "set")),
			)
			defer span.End()
			db[json.Name] = json.Color
		} else {
			fmt.Println(err)
		}
	})

	r.GET("/colorschemes/:name", func(c *gin.Context) {
		name := c.Params.ByName("name")
		_, span := tracer.Start(
			c.Request.Context(),
			"db",
			trace.WithAttributes(attribute.String("name", name), attribute.String("operation", "get")),
		)
		defer span.End()
		color, ok := db[name]
		if ok {
			c.JSON(http.StatusOK, gin.H{"name": name, "color": color})
		} else {
			c.AbortWithStatus(http.StatusNotFound)
		}
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

func initTracer() (*sdktrace.TracerProvider, error) {
	exporter, err := stdout.New(stdout.WithPrettyPrint())
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
