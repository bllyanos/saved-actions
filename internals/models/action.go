package models

import (
	"time"

	"github.com/bllyanos/saved-actions/internals/models/subtypes/actrunstat"
	"github.com/bllyanos/saved-actions/internals/models/subtypes/acttype"
)

type Action struct {
	ID uint64 `json:"id" gorm:"primaryKey"`

	Name       string             `json:"name"`
	ActionType acttype.ActionType `json:"actionType"`

	CreatedAt time.Time `json:"createdAt" gorm:"index"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ActionParameter struct {
	ID uint64 `json:"id" gorm:"primaryKey"`

	Name     string `json:"name" gorm:"type:varchar(128); uniqueIndex:act_param_uniq_idx"`
	Required bool   `json:"required"`
	Default  string `json:"default" gorm:"type:text"`

	CreatedAt time.Time `json:"createdAt" gorm:"index"`
	UpdatedAt time.Time `json:"updatedAt"`

	// belongs to Action
	ActionID uint64 `json:"actionId" gorm:"uniqueIndex:act_param_uniq_idx"`
	Action   Action `json:"-"`
}

type HttpActionDetail struct {
	ID uint64 `json:"id" gorm:"primaryKey"`

	Url       string `json:"url" gorm:"type:text"`
	Method    string `json:"method" gorm:"type:text"`
	Body      string `json:"body" gorm:"type:text"`
	Headers   string `json:"headers" gorm:"type:text"`
	TimeoutMs uint64 `json:"timeoutMs"`

	CreatedAt time.Time `json:"createdAt" gorm:"index"`
	UpdatedAt time.Time `json:"updatedAt"`

	// belongs to Action
	ActionID uint64 `json:"actionId"`
	Action   Action `json:"-"`
}

type ActionRun struct {
	ID uint64 `json:"id" gorm:"primaryKey"`

	CreatedAt time.Time `json:"createdAt" gorm:"index"`

	// belongs to Action
	ActionID uint64 `json:"actionId"`
	Action   Action `json:"-"`
}

type ActionRunParameter struct {
	ID uint64 `json:"id" gorm:"primaryKey"`

	Name  string `json:"name" gorm:"type:varchar(128)"`
	Value string `json:"value" gorm:"type:text"`

	// belongs to Action
	ActionID uint64 `json:"actionId"`
	Action   Action `json:"-"`
}

type ActionRunEvent struct {
	ID uint64 `json:"id" gorm:"primaryKey"`

	Status  actrunstat.ActionRunStatus `json:"status" gorm:"uniqueIndex:act_run_uniq_idx"`
	Message string                     `json:"message" gorm:"type:text"`

	CreatedAt time.Time `json:"createdAt" gorm:"index"`

	// belongs to Action
	ActionID uint64 `json:"actionId" gorm:"uniqueIndex:act_run_uniq_idx"`
	Action   Action `json:"-"`

	// belongs to ActionRun
	ActionRunID uint64    `json:"actionRunId" gorm:"uniqueIndex:act_run_uniq_idx"`
	ActionRun   ActionRun `json:"-"`
}
