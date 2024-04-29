// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ContainersDao is the data access object for table ck8s_containers.
type ContainersDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns ContainersColumns // columns contains all the column names of Table for convenient usage.
}

// ContainersColumns defines and stores column names for table ck8s_containers.
type ContainersColumns struct {
	Id          string // 主键
	Timestamp   string // 时间戳
	ContainerId string // 容器ID
	Name        string // 名称
	Image       string // 镜像
	Ip          string // IP地址
	State       string // 状态
	Uptime      string // 运行时间
	CpuPercent  string // CPU使用率
	MemPercent  string // 内存使用率
	MemUsage    string // 内存使用量
	MemLimit    string // 内存限制
	CreatedAt   string // 创建时间
	UpdatedAt   string // 修改时间
}

// containersColumns holds the columns for table ck8s_containers.
var containersColumns = ContainersColumns{
	Id:          "id",
	Timestamp:   "timestamp",
	ContainerId: "container_id",
	Name:        "name",
	Image:       "image",
	Ip:          "ip",
	State:       "state",
	Uptime:      "uptime",
	CpuPercent:  "cpu_percent",
	MemPercent:  "mem_percent",
	MemUsage:    "mem_usage",
	MemLimit:    "mem_limit",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewContainersDao creates and returns a new DAO object for table data access.
func NewContainersDao() *ContainersDao {
	return &ContainersDao{
		group:   "default",
		table:   "ck8s_containers",
		columns: containersColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ContainersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ContainersDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ContainersDao) Columns() ContainersColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ContainersDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ContainersDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ContainersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
