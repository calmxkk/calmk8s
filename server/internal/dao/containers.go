// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"calmk8s/internal/dao/internal"
)

// internalContainersDao is internal type for wrapping internal DAO implements.
type internalContainersDao = *internal.ContainersDao

// containersDao is the data access object for table ck8s_containers.
// You can define custom methods on it to extend its functionality as you wish.
type containersDao struct {
	internalContainersDao
}

var (
	// Containers is globally public accessible object for table ck8s_containers operations.
	Containers = containersDao{
		internal.NewContainersDao(),
	}
)

// Fill with you ideas below.
