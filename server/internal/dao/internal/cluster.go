// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ClusterDao is the data access object for table ck8s_cluster.
type ClusterDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns ClusterColumns // columns contains all the column names of Table for convenient usage.
}

// ClusterColumns defines and stores column names for table ck8s_cluster.
type ClusterColumns struct {
	Id        string // ä¸»é”®
	Name      string // é›†ç¾¤åç§°
	Configstr string // é…ç½®æ–‡ä»¶
	CreatedBy string // åˆ›å»ºç”¨æˆ·
	CreatedAt string // åˆ›å»ºæ—¶é—´
	UpdatedAt string // ä¿®æ”¹æ—¶é—´
}

// clusterColumns holds the columns for table ck8s_cluster.
var clusterColumns = ClusterColumns{
	Id:        "id",
	Name:      "name",
	Configstr: "configstr",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewClusterDao creates and returns a new DAO object for table data access.
func NewClusterDao() *ClusterDao {
	return &ClusterDao{
		group:   "default",
		table:   "ck8s_cluster",
		columns: clusterColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ClusterDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ClusterDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ClusterDao) Columns() ClusterColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ClusterDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ClusterDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ClusterDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
