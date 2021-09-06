package models

import "github.com/bandros/framework"

func GetProductList() ([]map[string]interface{}, error) {
	db := framework.Database{}
	defer db.Close()

	db.Select("*").From("product")
	return db.Result()
}
