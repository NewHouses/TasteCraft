package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// DBModel is the type for database connection values
type DBModel struct {
	DB *sql.DB
}

// Models is the wrapper for all models
type Models struct {
	DB DBModel
}

// NewModels returns a model type with database connection pool
func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{
			DB: db,
		},
	}
}

func (m *DBModel) CleanDB() error {
	stmt := `
	TRUNCATE steps;
	`
	_, err := m.DB.Exec(stmt)
	if err != nil {
		return err
	}

	stmt = `
	TRUNCATE properties;
	`
	_, err = m.DB.Exec(stmt)
	if err != nil {
		return err
	}

	stmt = `
	TRUNCATE ingredients;
	`
	_, err = m.DB.Exec(stmt)
	if err != nil {
		return err
	}

	stmt = `
	DELETE FROM foods WHERE food_id != 0;
	`
	_, err = m.DB.Exec(stmt)
	if err != nil {
		return err
	}

	stmt = `
	DELETE FROM recipes WHERE recipe_id != 0;
	`
	_, err = m.DB.Exec(stmt)
	if err != nil {
		return err
	}

	return err
}

func (m *DBModel) LastID(idColumn, tableName string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var lastID *int

	query := fmt.Sprintf(`
	SELECT 
		MAX(%s)
	FROM 
		%s;
	`, idColumn, tableName)

	row := m.DB.QueryRowContext(ctx, query)
	err := row.Scan(
		&lastID,
	)

	if lastID == nil {
		return 0, err
	}
	return *lastID, err
}
