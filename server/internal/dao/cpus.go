// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"calmk8s/internal/dao/internal"
)

// internalCpusDao is internal type for wrapping internal DAO implements.
type internalCpusDao = *internal.CpusDao

// cpusDao is the data access object for table ck8s_cpus.
// You can define custom methods on it to extend its functionality as you wish.
type cpusDao struct {
	internalCpusDao
}

var (
	// Cpus is globally public accessible object for table ck8s_cpus operations.
	Cpus = cpusDao{
		internal.NewCpusDao(),
	}
)

// Fill with you ideas below.
