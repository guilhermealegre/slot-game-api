package database

import (
	"github.com/guilhermealegre/go-clean-arch-core-lib/database/table"
)

// schemas
const (
	SchemaUser = "user"
	SchemaAuth = "auth"
	SchemaSlot = "slot"
)

var (
	//User Schema Tables
	UserTableUser   = table.New(SchemaUser, "user")
	UserTableWallet = table.New(SchemaUser, "wallet")

	// Auth Schema Tables
	AuthTableAuth = table.New(SchemaAuth, "auth")

	// Slot Game Schema Tables
	SlotTableSpin           = table.New(SchemaSlot, "spin")
	SlotTableSpinResultType = table.New(SchemaSlot, "spin_result_type")
)
