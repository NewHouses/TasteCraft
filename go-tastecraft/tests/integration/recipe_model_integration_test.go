package integration_test

import (
	"fmt"
	"tastecraft/db/driver"
	"tastecraft/db/models"
	e "tastecraft/internal/entities"
	vo "tastecraft/internal/value-objects"
	"testing"

	"github.com/stretchr/testify/assert"
)

var DB models.DBModel

func SetUp() {
	fmt.Println("connecting DB...")
	connectionString := "test:secret@tcp(localhost:3306)/tastecraft_test?parseTime=true&tls=false"
	conn, err := driver.OpenDB(connectionString)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	DB = models.DBModel{DB: conn}
	fmt.Println("DB connected")
}

func Teardown() {
	err := DB.CleanDB()
	if err != nil {
		panic(fmt.Sprintf("Teardown failed: %s", err.Error()))
	}
	DB.DB.Close()
}

func TestGetRecipeFromDB(t *testing.T) {
	fmt.Println("starting test...")
	SetUp()
	expectedRecipe := assumeRecipeInDatabase()

	actualRecipe, err := DB.Recipe(1)
	if err != nil {
		panic(fmt.Sprintf("Test failed: %s", err.Error()))
	}

	assertRecipe(t, actualRecipe, expectedRecipe)
	Teardown()
	fmt.Println("test ended successfully")
}

func assumeRecipeInDatabase() e.Recipe {
	ingredient1 := e.Ingredient{
		Food: e.Food{
			ID:   1,
			Name: "food test 1",
			Properties: vo.Properties{
				Kilocalories:  2222,
				Proteins:      20,
				Carbohydrates: 50,
				Fat:           10,
				VitaminA:      12,
			},
		},
		Quantity:    500,
		Measurement: "ml",
	}
	ingredient2 := e.Ingredient{
		Food: e.Food{
			ID:   2,
			Name: "food test 2",
			Properties: vo.Properties{
				Kilocalories:  132,
				Proteins:      85,
				Carbohydrates: 30,
				Fat:           10,
				VitaminA:      12,
			},
		},
		Quantity:    200,
		Measurement: "grams",
	}
	recipe := e.Recipe{
		ID:          1,
		Name:        "recipe test",
		DishType:    4,
		Ingredients: []*e.Ingredient{&ingredient1, &ingredient2},
		Steps:       []string{"step 1", "step 2"},
	}
	err := DB.InsertRecipe(&recipe)
	if err != nil {
		panic(fmt.Sprintf("Test failed: %s", err.Error()))
	}

	return recipe
}

func assertRecipe(t *testing.T, actualRecipe, expectedRecipe e.Recipe) {
	assert.Equal(t, actualRecipe.ID, expectedRecipe.ID, "The two IDs should be the same.")
	assert.Equal(t, actualRecipe.DishType, expectedRecipe.DishType, "The two DishTypes should be the same.")
	assert.Equal(t, actualRecipe.Name, expectedRecipe.Name, "The two Names should be the same.")
	for i, actualIngredient := range actualRecipe.Ingredients {
		assert.Equal(t, actualIngredient.Food.ID, expectedRecipe.Ingredients[i].Food.ID, "The two Food IDs should be the same.")
		assert.Equal(t, actualIngredient.Food.Name, expectedRecipe.Ingredients[i].Food.Name, "The two Food Names should be the same.")
		assert.Equal(t, actualIngredient.Food.Properties, expectedRecipe.Ingredients[i].Food.Properties, "The two Food Properties should be the same.")
		assert.Equal(t, actualIngredient.Quantity, expectedRecipe.Ingredients[i].Quantity, "The two Ingredient Quantities should be the same.")
		assert.Equal(t, actualIngredient.Measurement, expectedRecipe.Ingredients[i].Measurement, "The two Ingredient Measurements should be the same.")
	}
	for i, actualStep := range actualRecipe.Steps {
		assert.Equal(t, actualStep, expectedRecipe.Steps[i], "The two Steps should be the same.")
	}
}
