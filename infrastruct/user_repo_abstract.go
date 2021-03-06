package infrastruct

// UserRepo 存储服务的CRUD接口模型（通过接口实现依赖倒置，解耦）
type UserRepo interface {
	Create(user *User) error
	Update(user *User) error
	DeleteById(id int) error
	GetByUserName(userName string) (*User, error)
	GetByUserId(userId int) (*User, error)
	NotExistByName(userName string) (bool, error)
}
