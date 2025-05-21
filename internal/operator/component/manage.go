package component

import (
	"github.com/beepfd/server/internal/store/component"
	"github.com/beepfd/server/models"
	"github.com/beepfd/server/pkg/utils"
	"github.com/pkg/errors"
)

type Operator struct {
	QueryParma     *utils.Query
	Component      *models.Component
	Binary         []byte
	ComponentStore *component.Store
	User           string
}

func NewOperator() *Operator {
	return &Operator{
		ComponentStore: &component.Store{},
	}
}

func (o *Operator) WithComponent(c *models.Component) *Operator {
	o.Component = c
	return o
}

func (o *Operator) WithQueryParma(q *utils.Query) *Operator {
	o.QueryParma = q
	return o
}

func (o *Operator) checkComponent() (err error) {
	if err = o.Component.Validate(); err != nil {
		err = errors.Wrap(err, "组件校验失败")
		return
	}

	return
}
