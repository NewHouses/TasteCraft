package models

import (
	"context"
	"fmt"
	e "tastecraft/internal/entities"
	valueobjects "tastecraft/internal/value-objects"
	"time"
)

func (m *DBModel) Recipe(id int) (e.Recipe, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	recipe := e.Recipe{
		ID: id,
	}

	var dishType int
	row := m.DB.QueryRowContext(ctx, `
		SELECT 
			recipe_name, dish_type
		FROM 
			recipes
		WHERE
			recipe_id = ?`, id)
	err := row.Scan(
		&recipe.Name,
		&dishType,
	)
	recipe.DishType = valueobjects.DishType(dishType)

	recipe.Ingredients, err = m.RecipeIngredients(id)
	if err != nil {
		return recipe, err
	}

	recipe.Steps, err = m.RecipeSteps(id)
	if err != nil {
		return recipe, err
	}

	return recipe, err
}

func (m *DBModel) RecipeIngredients(recipeID int) ([]*e.Ingredient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var ingredients []*e.Ingredient

	query := `
		SELECT 
			i.food_id,
			f.food_name,
			nv.kilocalories,
			nv.proteins,
			nv.carbohydrates,
			nv.fat,
			nv.fiber,
			nv.calcium,
			nv.iron,
			nv.zinc,
			nv.vitamin_a,
			nv.vitamin_b,
			nv.vitamin_b1,
			nv.vitamin_b2,
			nv.vitamin_b3,
			nv.vitamin_b6,
			nv.vitamin_b12,
			nv.vitamin_c,
			nv.vitamin_d,
			nv.vitamin_e,
			nv.vitamin_k,
			i.quantity,
			i.measurement
		FROM
			ingredients i 
		INNER JOIN
			foods f ON i.food_id = f.food_id
		INNER JOIN
			nutritional_values nv ON nv.food_id = f.food_id
		WHERE
			i.recipe_id = ?;
	`

	rows, err := m.DB.QueryContext(ctx, query, recipeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var i e.Ingredient
		err = rows.Scan(
			&i.Food.ID,
			&i.Food.Name,
			&i.Food.NutritionalValues.Kilocalories,
			&i.Food.NutritionalValues.Proteins,
			&i.Food.NutritionalValues.Carbohydrates,
			&i.Food.NutritionalValues.Fat,
			&i.Food.NutritionalValues.Fiber,
			&i.Food.NutritionalValues.Calcium,
			&i.Food.NutritionalValues.Iron,
			&i.Food.NutritionalValues.Zinc,
			&i.Food.NutritionalValues.VitaminA,
			&i.Food.NutritionalValues.VitaminB,
			&i.Food.NutritionalValues.VitaminB1,
			&i.Food.NutritionalValues.VitaminB2,
			&i.Food.NutritionalValues.VitaminB3,
			&i.Food.NutritionalValues.VitaminB6,
			&i.Food.NutritionalValues.VitaminB12,
			&i.Food.NutritionalValues.VitaminC,
			&i.Food.NutritionalValues.VitaminD,
			&i.Food.NutritionalValues.VitaminE,
			&i.Food.NutritionalValues.VitaminK,
			&i.Quantity,
			&i.Measurement,
		)
		if err != nil {
			return nil, err
		}

		ingredients = append(ingredients, &i)
	}

	return ingredients, err
}

func (m *DBModel) RecipeSteps(recipeID int) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		SELECT 
			step
		FROM
			steps
		WHERE
			recipe_id = ?;
	`

	rows, err := m.DB.QueryContext(ctx, query, recipeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var steps []string

	for rows.Next() {
		var s string
		err = rows.Scan(
			&s,
		)
		if err != nil {
			return nil, err
		}
		steps = append(steps, s)
	}

	return steps, err
}

func (m *DBModel) InsertRecipe(recipe *e.Recipe) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if recipe.ID == 0 {
		lastRecipeID, err := m.LastRecipeID()
		if err != nil {
			return err
		}

		recipe.ID = lastRecipeID + 1
	}

	stmt := fmt.Sprintf(`
		INSERT INTO recipes
			(recipe_id, recipe_name, dish_type)
			VALUES (%d, "%s", %d);
	`, recipe.ID, recipe.Name, recipe.DishType.EnumIndex())

	_, err := m.DB.ExecContext(ctx, stmt)
	if err != nil {
		return err
	}

	err = m.InsertIngredients(recipe.ID, recipe.Ingredients)
	if err != nil {
		return err
	}

	err = m.InsertSteps(recipe.ID, recipe.Steps)
	if err != nil {
		return err
	}

	return err
}

func (m *DBModel) LastRecipeID() (int, error) {
	return m.LastID("recipe_id", "recipes")
}
