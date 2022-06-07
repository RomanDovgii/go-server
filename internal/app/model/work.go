package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Work struct {
	ID            int      `json:id`
	CreatorId     int      `json:"creator_id,omitempty"`
	Name          string   `json:"name,omitempty"`
	Description   string   `json:"description,omitempty"`
	DocumentLinks []string `json:"links"`
}

func (w *Work) Validate() error {
	return validation.ValidateStruct(
		w,
		validation.Field(&w.CreatorId, validation.Required, is.Int),
		validation.Field(&w.Name, validation.Required, validation.Length(5, 100)),
	)
}
