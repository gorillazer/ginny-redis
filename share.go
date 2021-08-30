package redis

import (
	"log"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Provider
var Provider = wire.NewSet(NewConfig, New)

// NewConfig
func NewConfig(v *viper.Viper) (*Config, error) {
	var err error
	o := new(Config)
	if err = v.UnmarshalKey("redis", o); err != nil {
		return nil, errors.Wrap(err, "unmarshal app option error")
	}
	return o, err
}

func New(config *Config, logger *zap.Logger) *Manager {
	mgr, err := NewManager(config, logger)
	if err != nil {
		log.Fatalf("redis manager error: %s", err.Error())
	}
	return mgr
}
