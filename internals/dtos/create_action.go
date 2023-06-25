package dtos

import (
	"time"

	"github.com/bllyanos/saved-actions/internals/models"
	"github.com/bllyanos/saved-actions/internals/models/subtypes/acttype"
)

type CreateActionParameter struct {
	Name     string `json:"key"`
	Required bool   `json:"required"`
	Default  string `json:"default"`
}

type CreateHttpAction struct {
	Name       string                  `json:"name"`
	Url        string                  `json:"url"`
	Method     string                  `json:"method"`
	Body       string                  `json:"body"`
	Headers    string                  `json:"headers"`
	TimeoutMs  uint64                  `json:"timeoutMs"`
	Parameters []CreateActionParameter `json:"parameters"`
}

func (d CreateHttpAction) GetAction() models.Action {
	return models.Action{
		Name:       d.Name,
		ActionType: acttype.Http,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

func (d CreateHttpAction) GetHttpActionDetail(actionId uint64) models.HttpActionDetail {
	return models.HttpActionDetail{
		Url:       d.Url,
		Body:      d.Body,
		Method:    d.Method,
		Headers:   d.Headers,
		TimeoutMs: d.TimeoutMs,
		ActionID:  actionId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (d CreateHttpAction) GetActionParameter(actionId uint64) []models.ActionParameter {
	newActionParameters := []models.ActionParameter{}
	for _, param := range d.Parameters {
		newActionParameters = append(newActionParameters,
			models.ActionParameter{
				ActionID:  actionId,
				Name:      param.Name,
				Required:  param.Required,
				Default:   param.Default,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		)
	}
	return newActionParameters
}
