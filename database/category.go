package database

import (
	"log"

	"github.com/zerodot618/zerokk-go-blog/models"
)

func GetAllCategory() ([]models.Category, error) {
	rows, err := DB.Query("select * from blog_category")
	if err != nil {
		log.Println("GetAllCategory 查询出错:", err)
		return nil, err
	}
	var categorys []models.Category
	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.Cid, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			log.Println("GetAllCategory 取值出错:", err)
			return nil, err
		}
		categorys = append(categorys, category)
	}
	return categorys, nil
}

func GetCategoryNameById(cateID int) (string, error) {
	row := DB.QueryRow("select name from blog_category where cid = ?", cateID)
	if err := row.Err(); err != nil {
		log.Println("GetCategoryNameById 查询出错:", err)
		return "", err
	}
	var categoryName string
	err := row.Scan(&categoryName)
	if err != nil {
		log.Println("GetCategoryNameById 取值出错:", err)
		return "", err
	}

	return categoryName, nil
}
