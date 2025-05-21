package service

import (
	"github.com/beepfd/server/internal/metrics"
	"github.com/beepfd/server/pkg/utils"
	"github.com/gin-gonic/gin"
)

type NodeMetrics struct {
	Metrics *metrics.NodeMetricsCollector
}

func (n *NodeMetrics) GetMetrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		metrics, err := n.Metrics.GetMetrics()
		if err != nil {
			utils.HandleError(c, err)
			return
		}

		data := &map[string]interface{}{"metrics": metrics}
		utils.HandleResult(c, data)
	}
}
