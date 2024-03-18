//go:build wireinject
// +build wireinject

//
//go:generate go run github.com/google/wire/cmd/wire

package wiring

import (
	"github.com/google/wire"
	"github.com/votanlean/GoLoad/internal/app"
	"github.com/votanlean/GoLoad/internal/configs"
	"github.com/votanlean/GoLoad/internal/handler"
)

var WireSet = wire.NewSet(
	app.WireSet,
	configs.WireSet,
	handler.WireSet,
)

func InitializeServer(configFilePath configs.ConfigFilePath) (*app.Server, func(), error) {
	wire.Build(WireSet)
	return nil, nil, nil
}