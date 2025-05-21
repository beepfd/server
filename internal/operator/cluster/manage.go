package cluster

import (
	k8scluster "github.com/beepfd/server/internal/store/cluster"
	"github.com/beepfd/server/models"
	"github.com/beepfd/server/pkg/utils"
	"github.com/pkg/errors"
)

type Operator struct {
	QueryParma   *utils.Query
	Cluster      *models.Cluster
	ClusterStore *k8scluster.Store
	User         string
}

func NewOperator() *Operator {
	return &Operator{
		ClusterStore: &k8scluster.Store{},
	}
}

func (o *Operator) WithCluster(c *models.Cluster) *Operator {
	o.Cluster = c
	return o
}

func (o *Operator) WithQueryParma(q *utils.Query) *Operator {
	o.QueryParma = q
	return o
}

func (o *Operator) checkCluster() (err error) {
	basic := o.Cluster.ClusterBasic
	// 必填校验
	if err = basic.Validate(); err != nil {
		err = errors.Wrap(err, "真实集群校验失败")
		return
	}

	return
}
