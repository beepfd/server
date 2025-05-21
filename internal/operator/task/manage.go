package task

import (
	"github.com/beepfd/server/internal/store/component"
	"github.com/beepfd/server/internal/store/task"
	"github.com/beepfd/server/models"
	"github.com/beepfd/server/pkg/utils"
	"github.com/pkg/errors"
)

type Operator struct {
	QueryParma     *utils.Query
	Task           *models.Task
	TaskStore      *task.Store
	ComponentStore *component.Store
	User           string
}

func NewOperator() *Operator {
	return &Operator{
		TaskStore:      &task.Store{},
		ComponentStore: &component.Store{},
	}
}

func (o *Operator) WithTask(t *models.Task) *Operator {
	o.Task = t
	return o
}

func (o *Operator) WithQueryParma(q *utils.Query) *Operator {
	o.QueryParma = q
	return o
}

func (o *Operator) checkTask() (err error) {
	if err = o.Task.Validate(); err != nil {
		err = errors.Wrap(err, "任务校验失败")
		return
	}

	return
}
