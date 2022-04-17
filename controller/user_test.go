package controller

import (
	"errors"
	"github.com/liaojuntao/infrastruct"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

const (
	existsName = "existsName"
	sameId = 111
)

// 测试user 的controller层，也需要和repo层解耦

func TestCreateUser(t *testing.T) {

	Convey("test CreateUser", t, func() {

		Convey("failed when repo system err", func() {
			err := NewUserController(&mockRepo{isMockWrong: true}).CreateUser(new(infrastruct.User))
			So(err, ShouldNotBeNil)
		})

		Convey("failed when userName is already exists", func() {
			user := new(infrastruct.User)
			user.UserName = existsName
			err := NewUserController(&mockRepo{isMockWrong: false}).CreateUser(user)
			So(err.Error(), ShouldEqual, "该用户名已存在")
		})

		Convey("success when userName is not exists", func() {
			err := NewUserController(&mockRepo{isMockWrong: false}).CreateUser(new(infrastruct.User))
			So(err, ShouldBeNil)
		})
	})
}

func TestUpdateUser(t *testing.T) {

	Convey("test UpdateUser", t, func() {

		Convey("failed when repo system err", func() {
			err := NewUserController(&mockRepo{isMockWrong: true}).UpdateUser(new(infrastruct.User))
			So(err, ShouldNotBeNil)
		})

		Convey("failed when userName is already exists", func() {
			user := new(infrastruct.User)
			user.UserId = 1
			user.UserName = existsName
			err := NewUserController(&mockRepo{isMockWrong: false}).UpdateUser(user)
			So(err.Error(), ShouldEqual, "该用户名已存在")
		})

		Convey("success when userName is not exists", func() {
			err := NewUserController(&mockRepo{isMockWrong: false}).UpdateUser(new(infrastruct.User))
			So(err, ShouldBeNil)
		})
	})
}

// mockRepo mock仓储服务
type mockRepo struct {
	isMockWrong bool
}

func (m *mockRepo) Create(user *infrastruct.User) error {
	if m.isMockWrong {
		return errors.New("wrong")
	}
	return nil
}

func (m *mockRepo) Update(user *infrastruct.User) error {
	if m.isMockWrong {
		return errors.New("wrong")
	}
	return nil
}

func (m *mockRepo) DeleteById(id int) error {
	if m.isMockWrong {
		return errors.New("wrong")
	}
	return nil
}

func (m *mockRepo) GetByUserName(userName string) (*infrastruct.User, error) {
	if m.isMockWrong {
		return nil, errors.New("wrong")
	}
	if userName == existsName {
		return &infrastruct.User{UserName: existsName,UserId: sameId}, nil
	}
	return &infrastruct.User{}, nil
}

func (m *mockRepo) NotExistByName(userName string) (bool, error) {
	if m.isMockWrong {
		return false, errors.New("wrong")
	}
	if userName == existsName {
		return false, nil
	}
	return true, nil
}
