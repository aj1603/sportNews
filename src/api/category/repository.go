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

func get_by_category_id_(id int) res.Category {
	var category res.Category

	db.DB.QueryRow(
		context.Background(),
		`SELECT
			c.id,
			c.name,
			c.img_url,
			(
				SELECT
					jsonb_agg(jsonb_build_object(
						'id', sc.id,
						'name', sc.name,
						'img_url', sc.img_url,
						'categories_id', sc.categories_id,
						'news', (
							SELECT
								jsonb_agg(jsonb_build_object(
									'id', n.id,
									'title', n.title,
									'context', n.context,
									'type', n.type,
									'subcategories_id', n.subcategories_id,
									'categories_id', n.categories_id,
									'news_images', (
										SELECT
											jsonb_agg(jsonb_build_object(
												'id', ni.id,
												'img_url', ni.img_url,
												'news_id', ni.news_id
											))
										FROM news_img ni
										WHERE ni.news_id = n.id
									)
								))
							FROM news n
							WHERE n.subcategories_id = sc.id
						)
					))
				FROM subcategories sc
				WHERE sc.categories_id = c.id
			) AS news_category
		FROM categories c
		WHERE c.id = $1
		GROUP BY c.id`,
		id,
	).Scan(&category.ID, &category.NAME, &category.IMG_URL, &category.NEWS)
	return category
}
