package models

import (
	"context"
	vo "tastecraft/internal/value-objects"
	"time"
)

func (m *DBModel) InsertNutritionalValues(foodID int, nutritionalValues *vo.NutritionalValues) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
			INSERT INTO nutritional_values
	 			(food_id, kilocalories, proteins, carbohydrates, fat, fiber, calcium, iron, zinc, vitamin_a, vitamin_b, vitamin_b1, vitamin_b2, vitamin_b3, vitamin_b6, vitamin_b12, vitamin_c, vitamin_d, vitamin_e, vitamin_k)	
			VALUES 
				(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
		`

	_, err := m.DB.ExecContext(ctx, stmt,
		foodID,
		nutritionalValues.Kilocalories,
		nutritionalValues.Proteins,
		nutritionalValues.Carbohydrates,
		nutritionalValues.Fat,
		nutritionalValues.Fiber,
		nutritionalValues.Calcium,
		nutritionalValues.Iron,
		nutritionalValues.Zinc,
		nutritionalValues.VitaminA,
		nutritionalValues.VitaminB,
		nutritionalValues.VitaminB1,
		nutritionalValues.VitaminB2,
		nutritionalValues.VitaminB3,
		nutritionalValues.VitaminB6,
		nutritionalValues.VitaminB12,
		nutritionalValues.VitaminC,
		nutritionalValues.VitaminD,
		nutritionalValues.VitaminE,
		nutritionalValues.VitaminK)

	return err
}
