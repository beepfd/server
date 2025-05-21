package cluster

import "github.com/beepfd/server/models"

type Cluster interface {
	Create(cluster *models.Cluster) (err error)
	Get(id int) (cluster models.Cluster, err error)
	List(pageSize, pageNum int) (total int64, clusters []*models.Cluster, err error)
	Count() (total int64, err error)
	Update(cluster *models.Cluster) (err error)
	Delete(cluster *models.Cluster) (err error)
}

type Store struct {
}
