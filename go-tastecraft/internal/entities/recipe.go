package entity

import (
	vo "tastecraft/internal/value-objects"
)

type Recipe struct {
	ID                int                  `json:"id"`
	Name              string               `json:"name"`
	DishType          vo.DishType          `json:"dish-type"`
	Ingredients       []*Ingredient        `json:"ingredients"`
	Steps             []string             `json:"steps"`
	NutritionalValues vo.NutritionalValues `json:"nutritional_values"`
}

type Ingredient struct {
	Food        Food   `json:"food"`
	Quantity    int    `json:"quantity"`
	Measurement string `json:"measurement"`
}
