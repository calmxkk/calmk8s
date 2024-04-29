// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DisksDao is the data access object for table ck8s_disks.
type DisksDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns DisksColumns // columns contains all the column names of Table for convenient usage.
}

// DisksColumns defines and stores column names for table ck8s_disks.
type DisksColumns struct {
	Id          string // 主键
	Timestamp   string // 时间戳
	Device      string // 设备名称
	DiskPercent string // 磁盘使用百分比
	DiskTotal   string // 磁盘总量
	DiskUsed    string // 磁盘已使用量
	DiskRead    string // 磁盘读取量
	DiskWrite   string // 磁盘写入量
	CreatedAt   string // 创建时间
	UpdatedAt   string // 修改时间
}

// disksColumns holds the columns for table ck8s_disks.
var disksColumns = DisksColumns{
	Id:          "id",
	Timestamp:   "timestamp",
	Device:      "device",
	DiskPercent: "disk_percent",
	DiskTotal:   "disk_total",
	DiskUsed:    "disk_used",
	DiskRead:    "disk_read",
	DiskWrite:   "disk_write",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewDisksDao creates and returns a new DAO object for table data access.
func NewDisksDao() *DisksDao {
	return &DisksDao{
		group:   "default",
		table:   "ck8s_disks",
		columns: disksColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DisksDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DisksDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DisksDao) Columns() DisksColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DisksDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DisksDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DisksDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
