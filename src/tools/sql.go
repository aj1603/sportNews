package tools

import "os"

var Faculty string

func Init_queries() {
	Faculty = get_query("queries/faculty/get_faculty.sql")
}

func get_query(path string) string {
	query, _ := os.ReadFile(path)
	return string(query)
}
