package router

import (
	"github.com/beepfd/server/internal/cache"
	"github.com/beepfd/server/internal/metrics"
	"github.com/gin-gonic/gin"
)

func InitPrometheusMetrics(r *gin.Engine) {
	taskMetrics := metrics.NewTaskMetrics(cache.TaskRunningStore)
	r.GET("/metrics", taskMetrics.Handler())
}
