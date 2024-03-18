package handler

import (
	"github.com/google/wire"

	"github.com/votanlean/GoLoad/internal/handler/grpc"
	"github.com/votanlean/GoLoad/internal/handler/http"
)

var WireSet = wire.NewSet(
	grpc.WireSet,
	http.WireSet,
)