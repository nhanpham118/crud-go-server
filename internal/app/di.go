package app

import (
	"crud-go-server/internal/pkg/api"
	"crud-go-server/internal/pkg/db"
	"crud-go-server/internal/pkg/repo"
	"crud-go-server/internal/setting"

	"github.com/google/wire"
)

var deps = wire.NewSet(
	api.GraphSet,
	db.GraphSet,
	repo.GraphSet,
	setting.GraphSet,
)

var GraphSet = wire.NewSet(
	deps,
	NewApp,
	NewHttpServer,
)
