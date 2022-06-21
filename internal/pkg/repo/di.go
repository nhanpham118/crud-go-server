package repo

import "github.com/google/wire"

var GraphSet = wire.NewSet(
	NewStudentRepo,
	NewModuleRepo,
	NewMarkRepo,
)
