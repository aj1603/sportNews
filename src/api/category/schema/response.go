package schema

type Category struct {
	ID      int             `json:"id"`
	NAME    string          `json:"name"`
	IMG_URL string          `json:"img_url"`
	NEWS    []News_category `json:"news_category"`
}

type News_category struct {
	ID             int     `json:"id"`
	TITLE          string  `json:"title"`
	CONTEXT        string  `json:"context"`
	TYPE           int     `json:"type"`
	CATEGORY_ID    int     `json:"categories_id"`
	SUBCATEGORY_ID int     `json:"subcategories_id"`
	Images         []Image `json:"images"`
}

type Image struct {
	ID      int    `json:"id"`
	IMG_URL string `json:"img_url"`
	NEWS_ID int    `json:"news_id"`
}
