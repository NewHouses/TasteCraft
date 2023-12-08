package models

import (
	"context"
	vo "tastecraft/internal/value-objects"
	"time"
)

func (m *DBModel) InsertProperties(foodID int, propertie *vo.Properties) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
			INSERT INTO properties
	 			(food_id, kilocalories, proteins, carbohydrates, fat, fiber, calcium, iron, zinc, vitamin_a, vitamin_b, vitamin_b1, vitamin_b2, vitamin_b3, vitamin_b6, vitamin_b12, vitamin_c, vitamin_d, vitamin_e, vitamin_k)	
			VALUES 
				(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
		`

	_, err := m.DB.ExecContext(ctx, stmt,
		foodID,
		propertie.Kilocalories,
		propertie.Proteins,
		propertie.Carbohydrates,
		propertie.Fat,
		propertie.Fiber,
		propertie.Calcium,
		propertie.Iron,
		propertie.Zinc,
		propertie.VitaminA,
		propertie.VitaminB,
		propertie.VitaminB1,
		propertie.VitaminB2,
		propertie.VitaminB3,
		propertie.VitaminB6,
		propertie.VitaminB12,
		propertie.VitaminC,
		propertie.VitaminD,
		propertie.VitaminE,
		propertie.VitaminK)

	return err
}
