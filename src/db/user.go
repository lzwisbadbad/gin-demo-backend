package db

const USER_TABLE_NAME = "user"

type UserRoleType int32

const (
	ADMIN UserRoleType = iota + 1
	USER
)

const (
	ADMIN_STR = "管理员"

	USER_STR = "普通用户"
)

var UserRoleTypeName = map[UserRoleType]string{
	ADMIN: ADMIN_STR,
	USER:  USER_STR,
}

var UserRoleTypeValue = map[string]UserRoleType{
	ADMIN_STR: ADMIN,
	USER_STR:  USER,
}

type User struct {
	GeneralField
	UserName     string `gorm:"uniqueIndex"`
	UserRole     UserRoleType
	UserPwd      string
	UserNickName string
	UserPhoneNum string
	UserEmail    string
}

func (u *User) TableName() string {
	return USER_TABLE_NAME
}

func init() {
	user := new(User)
	TableSlice = append(TableSlice, &user)
}
