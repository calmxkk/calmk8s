// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NetsDao is the data access object for table ck8s_nets.
type NetsDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns NetsColumns // columns contains all the column names of Table for convenient usage.
}

// NetsColumns defines and stores column names for table ck8s_nets.
type NetsColumns struct {
	Id        string // 主键
	Timestamp string // 时间戳
	Ethernet  string // 以太网接口
	NetRecv   string // 网络接收量
	NetSend   string // 网络发送量
	CreatedAt string // 创建时间
	UpdatedAt string // 修改时间
}

// netsColumns holds the columns for table ck8s_nets.
var netsColumns = NetsColumns{
	Id:        "id",
	Timestamp: "timestamp",
	Ethernet:  "ethernet",
	NetRecv:   "net_recv",
	NetSend:   "net_send",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewNetsDao creates and returns a new DAO object for table data access.
func NewNetsDao() *NetsDao {
	return &NetsDao{
		group:   "default",
		table:   "ck8s_nets",
		columns: netsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NetsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NetsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NetsDao) Columns() NetsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NetsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NetsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NetsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
