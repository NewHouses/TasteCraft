package models

import (
	"context"
	vo "tastecraft/internal/value-objects"
	"time"
)

func (m *DBModel) InsertNutritionalValues(foodID int, nutrionalValues *vo.NutritionalValues) error {
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
		nutrionalValues.Kilocalories,
		nutrionalValues.Proteins,
		nutrionalValues.Carbohydrates,
		nutrionalValues.Fat,
		nutrionalValues.Fiber,
		nutrionalValues.Calcium,
		nutrionalValues.Iron,
		nutrionalValues.Zinc,
		nutrionalValues.VitaminA,
		nutrionalValues.VitaminB,
		nutrionalValues.VitaminB1,
		nutrionalValues.VitaminB2,
		nutrionalValues.VitaminB3,
		nutrionalValues.VitaminB6,
		nutrionalValues.VitaminB12,
		nutrionalValues.VitaminC,
		nutrionalValues.VitaminD,
		nutrionalValues.VitaminE,
		nutrionalValues.VitaminK)

	return err
}
