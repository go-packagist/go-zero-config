package go_zero_config

import (
	"github.com/go-packagist/framework/foundation"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig(t *testing.T) {
	app := foundation.NewApplication()

	app.Register(NewProvider(app.Container, ".test/config.yaml"))

	assert.Equal(t, "test.config", app.MustMake("zero.config").(*Config).Name)
	assert.Equal(t, "127.0.0.1:8080", app.MustMake("zero.config").(*Config).ListenOn)
}
