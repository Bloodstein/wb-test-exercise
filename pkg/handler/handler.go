package handler

import (
	"errors"
	"fmt"

	"github.com/Bloodstein/wb-test-exercise/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	// ginprometheus "github.com/zsais/go-gin-prometheus"
)

const (
	success_result = "ok"
	error_result   = "error"
)

type Handler struct {
	services *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		services: s,
	}
}

func initPrometheus() {
	var (
		counter = prometheus.NewCounter(
			prometheus.CounterOpts{
				Namespace: "golang",
				Name:      "my_counter",
				Help:      "This is my counter",
			})

		gauge = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "golang",
				Name:      "my_gauge",
				Help:      "This is my gauge",
			})

		histogram = prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Namespace: "golang",
				Name:      "my_histogram",
				Help:      "This is my histogram",
			})

		summary = prometheus.NewSummary(
			prometheus.SummaryOpts{
				Namespace: "golang",
				Name:      "my_summary",
				Help:      "This is my summary",
			})
	)

	prometheus.MustRegister(counter)
	prometheus.MustRegister(gauge)
	prometheus.MustRegister(histogram)
	prometheus.MustRegister(summary)
}

func (h *Handler) Routes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	initPrometheus()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			rel := v1.Group("/telegram-to-office-relations")
			{
				rel.GET("/items", h.GetAll)
				rel.GET("/items/:id", h.GetOne)
				rel.POST("/create/", h.Create)
				rel.POST("/delete/:id", h.Delete)
				rel.POST("/update/:id", h.Update)
			}
		}
		api.GET("/metrics", gin.WrapH(promhttp.Handler()))
	}

	return router
}

func getRowId(ctx *gin.Context) (string, error) {
	param := ctx.Param("id")

	if len(param) == 0 {
		return "", errors.New(fmt.Sprintf("You've missed the required param \"ID\": %s", param))
	}

	return param, nil
}
