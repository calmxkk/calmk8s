// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MemoriesDao is the data access object for table ck8s_memories.
type MemoriesDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns MemoriesColumns // columns contains all the column names of Table for convenient usage.
}

// MemoriesColumns defines and stores column names for table ck8s_memories.
type MemoriesColumns struct {
	Id         string // 主键
	Timestamp  string // 时间戳
	MemPercent string // 内存使用百分比
	MemTotal   string // 内存总量
	MemUsed    string // 内存已使用量
	CreatedAt  string // 创建时间
	UpdatedAt  string // 修改时间
}

// memoriesColumns holds the columns for table ck8s_memories.
var memoriesColumns = MemoriesColumns{
	Id:         "id",
	Timestamp:  "timestamp",
	MemPercent: "mem_percent",
	MemTotal:   "mem_total",
	MemUsed:    "mem_used",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewMemoriesDao creates and returns a new DAO object for table data access.
func NewMemoriesDao() *MemoriesDao {
	return &MemoriesDao{
		group:   "default",
		table:   "ck8s_memories",
		columns: memoriesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MemoriesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MemoriesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MemoriesDao) Columns() MemoriesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MemoriesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MemoriesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MemoriesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
