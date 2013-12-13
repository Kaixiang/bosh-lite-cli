package util

import (
	"github.com/stretchr/testify/assert"
	"boshlite/configuration"
	"testing"
)

func TestRouteCmd(t *testing.T) {
  config :=configuration.Default()
  cmd := RouteCmd(config)
  assert.Equal(t, cmd, "route delete -net 10.244.0.0/22 192.168.50.4 > /dev/null 2>&1;route add -net 10.244.0.0/22 192.168.50.4")
}

