// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HostsDao is the data access object for table ck8s_hosts.
type HostsDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns HostsColumns // columns contains all the column names of Table for convenient usage.
}

// HostsColumns defines and stores column names for table ck8s_hosts.
type HostsColumns struct {
	Id              string // 主键
	Timestamp       string // 时间戳
	Uptime          string // 运行时间
	Hostname        string // 主机名
	Os              string // 操作系统
	Platform        string // 平台
	PlatformVersion string // 平台版本
	KernelVersion   string // 内核版本
	KernelArch      string // 内核架构
	CreatedAt       string // 创建时间
	UpdatedAt       string // 修改时间
}

// hostsColumns holds the columns for table ck8s_hosts.
var hostsColumns = HostsColumns{
	Id:              "id",
	Timestamp:       "timestamp",
	Uptime:          "uptime",
	Hostname:        "hostname",
	Os:              "os",
	Platform:        "platform",
	PlatformVersion: "platform_version",
	KernelVersion:   "kernel_version",
	KernelArch:      "kernel_arch",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}

// NewHostsDao creates and returns a new DAO object for table data access.
func NewHostsDao() *HostsDao {
	return &HostsDao{
		group:   "default",
		table:   "ck8s_hosts",
		columns: hostsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HostsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HostsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HostsDao) Columns() HostsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HostsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HostsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HostsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
