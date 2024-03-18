//go: build wireinject
//go:build wireinject
// +build wireinject

//
//go:generate go run github.com/google/wire/cmd/wire

package wiring

import (
	"github.com/google/wire"
	"github.com/votanlean/GoLoad/internal/app"
)

var WireSet = wire.NewSet(
	app.WireSet,
)

func InitializeServer(configFilePath string) (*app.Server, func(), error) {
	wire.Build(WireSet)
	return nil, nil, nil
}