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
			array_agg(jsonb_build_object(
				'id', p.id,
				'name', p.name,
				'description', p.description,
				'price', p.price::float,
				'product_sku', p.product_sku,
				'quantity', p.quantity,
				'categories_id', p.categories_id,
				'discounts_id', p.discounts_id,
				'brands_id', p.brands_id,
				'images', pi.images
			)) AS products
		FROM categories c
		LEFT JOIN products p ON p.categories_id = c.id
		LEFT JOIN (
			SELECT
				products_id,
				array_agg(jsonb_build_object(
					'id', id,
					'img_url', img_url,
					'products_id', products_id
				)) AS images
			FROM products_images
			GROUP BY products_id
		) pi ON pi.products_id = p.id
		WHERE c.id = $1
		GROUP BY c.id, c.name`,
		id,
	).Scan(&category.ID, &category.NAME, &category.IMG_URL, &category.NEWS)
	return category
}
