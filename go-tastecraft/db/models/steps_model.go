package models

import (
	"context"
	"fmt"
	"time"
)

func (m *DBModel) InsertSteps(recipeID int, steps []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
		INSERT INTO steps
			(recipe_id, step)
		VALUES
	`

	for i, step := range steps {
		if i == len(steps)-1 {
			stmt += fmt.Sprintf(`(%d, "%s")`, recipeID, step)
		} else {
			stmt += fmt.Sprintf(`(%d, "%s"),`, recipeID, step)
		}
	}

	_, err := m.DB.ExecContext(ctx, stmt)

	return err
}
