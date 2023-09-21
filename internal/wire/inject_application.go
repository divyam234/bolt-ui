package wire

import (
	"github.com/divyam234/bolt-ui/application"
	"github.com/google/wire"
)

//lint:ignore U1000 because
var appSet = wire.NewSet(
	wire.Struct(new(application.Application), "*"),
	application.NewBrowseHandler,
)
