package database

import (
	"log"

	"github.com/zerodot618/zerokk-go-blog/models"
)

func GetUserNameById(userID int) (string, error) {
	row := DB.QueryRow("select user_name from blog_user where uid = ?", userID)
	if err := row.Err(); err != nil {
		log.Println("GetUserNameById 查询出错:", err)
		return "", err
	}
	var userName string
	err := row.Scan(&userName)
	if err != nil {
		log.Println("GetUserNameById 取值出错:", err)
		return "", err
	}

	return userName, nil
}

func GetUser(userName, passwd string) (*models.User, error) {
	var user = models.User{}
	row := DB.QueryRow("select * from blog_user where user_name = ? and passwd = ? limit 1", userName, passwd)
	if err := row.Err(); err != nil {
		log.Println("GetUser 查询出错:", err)
		return nil, err
	}
	err := row.Scan(
		&user.Uid,
		&user.UserName,
		&user.Passwd,
		&user.Avatar,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		log.Println("GetUser 取值出错:", err)
		return nil, err
	}

	return &user, nil
}
