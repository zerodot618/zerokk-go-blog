package service

import (
	"html/template"

	"github.com/zerodot618/zerokk-go-blog/config"
	"github.com/zerodot618/zerokk-go-blog/database"
	"github.com/zerodot618/zerokk-go-blog/models"
	"github.com/zerodot618/zerokk-go-blog/models/responses"
)

func GetPostsByCategoryID(page, pageSize, cateID int) (*responses.CategoryRes, error) {
	categoryName, _ := database.GetCategoryNameById(cateID)
	//页面上涉及到的所有的数据，必须有定义
	categorys, err := database.GetAllCategory()
	if err != nil {
		return nil, err
	}
	posts, err := database.GetPostsByCategoryID(page, pageSize, cateID)
	if err != nil {
		return nil, err
	}

	var postMores []models.PostMore
	for _, post := range posts {
		categoryName, _ := database.GetCategoryNameById(post.CategoryID)
		userName, _ := database.GetUserNameById(post.UserID)
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[:100]
		}

		postMore := models.PostMore{
			Pid:          post.Pid,
			Title:        post.Title,
			Slug:         post.Slug,
			Content:      template.HTML(content),
			CategoryId:   post.CategoryID,
			CategoryName: categoryName,
			UserID:       post.UserID,
			UserName:     userName,
			ViewCount:    post.ViewCount,
			Type:         post.Type,
			CreatedAt:    models.DateDay(post.CreatedAt),
			UpdatedAt:    models.DateDay(post.UpdatedAt),
		}
		postMores = append(postMores, postMore)
	}
	total, _ := database.GetAllPostCountByCategoryID(cateID)
	pagesCount := (total-1)/pageSize + 1
	var pages []int
	for i := 0; i < pagesCount; i++ {
		pages = append(pages, i+1)
	}
	var homeRes = &responses.HomeRes{
		config.Cfg.Viewer,
		categorys,
		postMores,
		total,
		page,
		pages,
		page != pagesCount,
	}

	var cateRes = &responses.CategoryRes{
		HomeRes:      *homeRes,
		CategoryName: categoryName,
	}
	return cateRes, nil
}
