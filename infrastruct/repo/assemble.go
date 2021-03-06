package repo

import (
	"github.com/liaojuntao/infrastruct"
	"time"
)

const defaultTimeFormat = "2006-01-02 15:04:05"

// toSqliteModel dto与model转换
func toSqliteModel(userDto *infrastruct.User) *user {
	if userDto == nil {
		return nil
	}
	return &user{
		UserId:      userDto.UserId,
		UserName:    userDto.UserName,
		BirthOfDate: userDto.BirthOfDate,
		Address:     userDto.Address,
		Description: userDto.Description,
	}
}

// toUser model与dto转换
func toUser(model *user) *infrastruct.User {
	if model == nil {
		return nil
	}
	user := &infrastruct.User{
		UserId:      model.UserId,
		UserName:    model.UserName,
		BirthOfDate: model.BirthOfDate,
		Address:     model.Address,
		Description: model.Description,
	}
	if model.CreateAt != nil {
		user.CreateAt = model.CreateAt.Format(defaultTimeFormat)
	}
	return user
}

type user struct {
	UserId      int `gorm:"primary_key"`
	UserName    string
	BirthOfDate string
	Address     string
	Description string
	CreateAt    *time.Time
}
