package category

import (
	"context"
	db "sport_news/database"
	res "sport_news/src/api/category/schema"
)

func get_() ([]res.Category, error) {
	var categories []res.Category

	rows, err := db.DB.Query(
		context.Background(),
		`
		SELECT * FROM categories
		`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var category res.Category

		err := rows.Scan(
			&category.ID,
			&category.NAME,
			&category.IMG_URL,
		)

		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}
