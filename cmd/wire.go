//go:build wireinject
// +build wireinject

package main

import (
	"crud-go-server/internal/app"

	"github.com/google/wire"
)

func initGoApp() (app.App, error) {
	wire.Build(app.GraphSet)
	return nil, nil
}
