package models

import (
	"context"
	e "tastecraft/internal/entities"
	"time"
)

func (m *DBModel) InsertIngredients(recipeID int, ingredients []*e.Ingredient) error {
	for _, ingredient := range ingredients {
		err := m.InsertIngredient(recipeID, ingredient)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *DBModel) InsertIngredient(recipeID int, ingredient *e.Ingredient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	food, err := m.Food(ingredient.Food.Name)
	if err != nil {
		m.InsertFood(&ingredient.Food)
	} else {
		ingredient.Food.ID = food.ID
	}

	stmt := `
		INSERT INTO ingredients
			(recipe_id, food_id, quantity, measurement)
		VALUES 
			(?, ?, ?, ?);
		`

	_, err = m.DB.ExecContext(ctx, stmt,
		recipeID,
		ingredient.Food.ID,
		ingredient.Quantity,
		ingredient.Measurement)

	return err
}
