package service

import (
	"html/template"

	"github.com/zerodot618/zerokk-go-blog/config"
	"github.com/zerodot618/zerokk-go-blog/database"
	"github.com/zerodot618/zerokk-go-blog/models"
	"github.com/zerodot618/zerokk-go-blog/models/responses"
)

func GetPostsByID(pID int) (*responses.PostRes, error) {
	post, err := database.GetPostsByID(pID)
	if err != nil {
		return nil, err
	}

	categoryName, _ := database.GetCategoryNameById(post.CategoryID)
	userName, _ := database.GetUserNameById(post.UserID)
	var postMore = models.PostMore{
		Pid:          post.Pid,
		Title:        post.Title,
		Slug:         post.Slug,
		Content:      template.HTML(post.Content),
		CategoryId:   post.CategoryID,
		CategoryName: categoryName,
		UserID:       post.UserID,
		UserName:     userName,
		ViewCount:    post.ViewCount,
		Type:         post.Type,
		CreatedAt:    models.DateDay(post.CreatedAt),
		UpdatedAt:    models.DateDay(post.UpdatedAt),
	}

	res := &responses.PostRes{
		config.Cfg.Viewer,
		config.Cfg.System,
		postMore,
	}

	return res, nil
}

func SavePost(post *models.Post) error {
	return database.SavePost(post)
}

func UpdatePost(post *models.Post) error {
	return database.UpdatePost(post)
}

func FindPosts() (responses.PigeonholeRes, error) {
	categorys, err := database.GetAllCategory()
	if err != nil {
		return responses.PigeonholeRes{}, err
	}

	posts, err := database.GetPostAll()
	pigeonholeMap := make(map[string][]models.Post)
	if err != nil {
		return responses.PigeonholeRes{}, err
	}
	for _, post := range posts {
		at := post.CreatedAt
		month := at.Format("2006-01")
		pigeonholeMap[month] = append(pigeonholeMap[month], post)
	}

	return responses.PigeonholeRes{
		config.Cfg.Viewer,
		config.Cfg.System,
		categorys,
		pigeonholeMap,
	}, nil
}

func SearchPost(condition string) ([]responses.SearchRes, error) {
	posts, err := database.SearchPost(condition)
	if err != nil {
		return nil, err
	}
	var searchRes []responses.SearchRes
	for _, post := range posts {
		searchRes = append(searchRes, responses.SearchRes{
			Pid:   post.Pid,
			Title: post.Title,
		})
	}
	return searchRes, nil
}
