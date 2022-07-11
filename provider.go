package redis

import (
	"github.com/google/wire"
)

// Provider
var Provider = wire.NewSet(NewConfig, NewRedis)
