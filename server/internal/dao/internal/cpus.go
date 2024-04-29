// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CpusDao is the data access object for table ck8s_cpus.
type CpusDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns CpusColumns // columns contains all the column names of Table for convenient usage.
}

// CpusColumns defines and stores column names for table ck8s_cpus.
type CpusColumns struct {
	Id         string // 主键
	Timestamp  string // 时间戳
	CpuPercent string // CPU使用百分比
	CreatedAt  string // 创建时间
	UpdatedAt  string // 修改时间
}

// cpusColumns holds the columns for table ck8s_cpus.
var cpusColumns = CpusColumns{
	Id:         "id",
	Timestamp:  "timestamp",
	CpuPercent: "cpu_percent",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewCpusDao creates and returns a new DAO object for table data access.
func NewCpusDao() *CpusDao {
	return &CpusDao{
		group:   "default",
		table:   "ck8s_cpus",
		columns: cpusColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CpusDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CpusDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CpusDao) Columns() CpusColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CpusDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CpusDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CpusDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
