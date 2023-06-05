package domain

import (
	"encoding/json"
	"github.com/pkg/errors"
	"navy/internal/common/nerrors"
	"navy/internal/common/validate"
)

type Port struct {
	ID          string    `redis:"id" json:"id" validate:"required"`
	Name        string    `redis:"name" json:"name" validate:"required"`
	Code        string    `json:"code"`
	City        string    `json:"city" validate:"required"`
	Country     string    `json:"country" validate:"required"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float64 `json:"coordinates"`
	Province    string    `json:"province"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
}

// Validate validates a Port
func (p Port) Validate() error {
	if err := validate.Validate(p); err != nil {
		return errors.Wrapf(nerrors.ErrInvalidInput, err.Error())
	}

	return nil
}

func (p Port) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}
