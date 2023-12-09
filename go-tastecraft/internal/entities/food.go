package entity

import (
	vo "tastecraft/internal/value-objects"
)

type Food struct {
	ID                int                  `json:"id"`
	Name              string               `json:"name"`
	NutritionalValues vo.NutritionalValues `json:"nutrional_values"`
}
