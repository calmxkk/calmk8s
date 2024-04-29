// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserDao is the data access object for table ck8s_user.
type UserDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns UserColumns // columns contains all the column names of Table for convenient usage.
}

// UserColumns defines and stores column names for table ck8s_user.
type UserColumns struct {
	Id                 string // 主键
	Role               string // 权限等级
	RealName           string // 真实姓名
	Username           string // 帐号
	PasswordHash       string // 密码
	Salt               string // 密码盐
	PasswordResetToken string // 密码重置令牌
	Avatar             string // 头像
	Sex                string // 性别
	Qq                 string // qq
	Email              string // 邮箱
	Mobile             string // 手机号码
	Birthday           string // 生日
	Address            string // 联系地址
	LastActiveAt       string // 最后活跃时间
	Remark             string // 备注
	Status             string // 状态
	CreatedAt          string // 创建时间
	UpdatedAt          string // 修改时间
}

// userColumns holds the columns for table ck8s_user.
var userColumns = UserColumns{
	Id:                 "id",
	Role:               "role",
	RealName:           "real_name",
	Username:           "username",
	PasswordHash:       "password_hash",
	Salt:               "salt",
	PasswordResetToken: "password_reset_token",
	Avatar:             "avatar",
	Sex:                "sex",
	Qq:                 "qq",
	Email:              "email",
	Mobile:             "mobile",
	Birthday:           "birthday",
	Address:            "address",
	LastActiveAt:       "last_active_at",
	Remark:             "remark",
	Status:             "status",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
}

// NewUserDao creates and returns a new DAO object for table data access.
func NewUserDao() *UserDao {
	return &UserDao{
		group:   "default",
		table:   "ck8s_user",
		columns: userColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserDao) Columns() UserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
