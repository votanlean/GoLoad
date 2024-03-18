package configs

import (
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	NewConfig,
	wire.FieldsOf(new(Config), "GRPC"),
	wire.FieldsOf(new(Config), "HTTP"),
)