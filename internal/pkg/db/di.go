package db

import "github.com/google/wire"

var GraphSet = wire.NewSet(
	NewMySqlSession,
)
