// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"calmk8s/internal/dao/internal"
)

// internalNetsDao is internal type for wrapping internal DAO implements.
type internalNetsDao = *internal.NetsDao

// netsDao is the data access object for table ck8s_nets.
// You can define custom methods on it to extend its functionality as you wish.
type netsDao struct {
	internalNetsDao
}

var (
	// Nets is globally public accessible object for table ck8s_nets operations.
	Nets = netsDao{
		internal.NewNetsDao(),
	}
)

// Fill with you ideas below.
