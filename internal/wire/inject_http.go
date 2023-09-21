package wire

import (
	"net/http"

	httpPort "github.com/divyam234/bolt-ui/ports/http"
	"github.com/google/wire"
)

//lint:ignore U1000 because
var httpSet = wire.NewSet(
	httpPort.NewServer,
	httpPort.NewHandler,
	httpPort.NewTokenAuthProvider,
	wire.Bind(new(http.Handler), new(*httpPort.Handler)),
	wire.Bind(new(httpPort.AuthProvider), new(*httpPort.TokenAuthProvider)),
)
