package models

import (
	"context"
	e "tastecraft/internal/entities"
	"time"
)

func (m *DBModel) Food(name string) (e.Food, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var recipe e.Food

	row := m.DB.QueryRowContext(ctx, `
		SELECT 
			food_id, food_name
		FROM 
			foods
		WHERE
		food_name = ?`, name)
	err := row.Scan(
		&recipe.ID,
		&recipe.Name,
	)

	return recipe, err
}

func (m *DBModel) InsertFood(food *e.Food) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if food.ID == 0 {
		lastFoodID, err := m.LastFoodID()
		if err != nil {
			return err
		}

		food.ID = lastFoodID + 1
	}

	stmt := `
		INSERT INTO foods
		(food_id, food_name)
		VALUES (?, ?)`

	_, err := m.DB.ExecContext(ctx, stmt, food.ID, food.Name)
	if err != nil {
		return err
	}

	err = m.InsertProperties(food.ID, &food.Properties)

	return err
}

func (m *DBModel) LastFoodID() (int, error) {
	return m.LastID("food_id", "foods")
}
