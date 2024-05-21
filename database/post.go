package database

import (
	"errors"
	"log"

	"github.com/zerodot618/zerokk-go-blog/models"
)

func GetPostPage(page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from blog_post limit ?, ?", page, pageSize)
	if err != nil {
		log.Println("GetPostPage 查询出错:", err)
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryID,
			&post.UserID,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			log.Println("GetPostPage 取值出错:", err)
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func GetPostBySlug(slug string, page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from blog_post where slug = ? limit ?, ?", slug, page, pageSize)
	if err != nil {
		log.Println("GetPostPage 查询出错:", err)
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryID,
			&post.UserID,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			log.Println("GetPostPage 取值出错:", err)
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func GetAllPostCount() (int, error) {
	row := DB.QueryRow("select count(1) from blog_post")
	if err := row.Err(); err != nil {
		log.Println("GetAllPostCount 查询出错:", err)
		return 0, err
	}
	var count int
	err := row.Scan(&count)
	if err != nil {
		log.Println("GetAllPostCount 取值出错:", err)
		return 0, err
	}

	return count, nil
}

func GetAllPostCountBySlug(slug string) (int, error) {
	row := DB.QueryRow("select count(1) from blog_post where slug = ?", slug)
	if err := row.Err(); err != nil {
		log.Println("GetAllPostCountBySlug 查询出错:", err)
		return 0, err
	}
	var count int
	err := row.Scan(&count)
	if err != nil {
		log.Println("GetAllPostCount 取值出错:", err)
		return 0, err
	}

	return count, nil
}

func GetAllPostCountByCategoryID(cateID int) (int, error) {
	row := DB.QueryRow("select count(1) from blog_post where category_id = ?", cateID)
	if err := row.Err(); err != nil {
		log.Println("GetAllPostCount 查询出错:", err)
		return 0, err
	}
	var count int
	err := row.Scan(&count)
	if err != nil {
		log.Println("GetAllPostCount 取值出错:", err)
		return 0, err
	}

	return count, nil
}

func GetPostAll() ([]models.Post, error) {
	rows, err := DB.Query("select * from blog_post")
	if err != nil {
		log.Println("GetPostAll 查询出错:", err)
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryID,
			&post.UserID,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			log.Println("GetPostPage 取值出错:", err)
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func GetPostsByCategoryID(page, pageSize, categoryId int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from blog_post where category_id = ? limit ?, ?", categoryId, page, pageSize)
	if err != nil {
		log.Println("GetPostPage 查询出错:", err)
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryID,
			&post.UserID,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			log.Println("GetPostPage 取值出错:", err)
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func GetPostsByID(pId int) (*models.Post, error) {
	row := DB.QueryRow("select * from blog_post where pid = ? limit 1", pId)
	if err := row.Err(); err != nil {
		log.Println("GetPostsByID 查询出错:", err)
		return nil, err
	}
	var post models.Post

	err := row.Scan(
		&post.Pid,
		&post.Title,
		&post.Content,
		&post.Markdown,
		&post.CategoryID,
		&post.UserID,
		&post.ViewCount,
		&post.Type,
		&post.Slug,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		log.Println("GetPostPage 取值出错:", err)
		return nil, err
	}

	return &post, nil
}

func SavePost(post *models.Post) error {
	ret, err := DB.Exec("insert into blog_post(title, content, markdown, category_id, user_id, view_count, type, slug, created_at, updated_at) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryID,
		post.UserID,
		post.ViewCount,
		post.Type,
		post.Slug,
		post.CreatedAt,
		post.UpdatedAt,
	)
	if err != nil {
		return err
	}
	pid, err := ret.LastInsertId()
	if err != nil {
		return err
	}
	post.Pid = int(pid)
	return nil
}

func UpdatePost(post *models.Post) error {
	ret, err := DB.Exec("update blog_post set title = ?, content = ?, markdown = ?, category_id = ?, user_id = ?, view_count = ?, type = ?, slug = ?, created_at = ?, updated_at = ? where pid = ?",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryID,
		post.UserID,
		post.ViewCount,
		post.Type,
		post.Slug,
		post.CreatedAt,
		post.UpdatedAt,
		post.Pid,
	)
	if err != nil {
		return err
	}

	affect, err := ret.RowsAffected()
	if err != nil {
		return err
	}
	if affect == 0 {
		return errors.New("更新失败")
	}

	return nil
}

func SearchPost(condition string) ([]models.Post, error) {
	rows, err := DB.Query("select * from blog_post where title like ?", "%"+condition+"%")
	if err != nil {
		log.Println("SearchPost 查询出错:", err)
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryID,
			&post.UserID,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			log.Println("GetPostPage 取值出错:", err)
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}
