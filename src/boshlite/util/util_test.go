package util

import (
  "boshlite/configuration"
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestRouteCmd(t *testing.T) {
  config := configuration.Default()
  cmd := RouteCmd(config)
  assert.Equal(t, cmd, "route delete -net 10.244.0.0/19 192.168.50.4 > /dev/null 2>&1;route add -net 10.244.0.0/19 192.168.50.4")
}

func TestIsVersionNewer(t *testing.T) {
  assert.Equal(t, IsVersionNewer("5.0", "4.2.18"), true)
  assert.Equal(t, IsVersionNewer("4.3", "4.2.18"), true)
  assert.Equal(t, IsVersionNewer("4.2.18.1", "4.2.18"), true)
  assert.Equal(t, IsVersionNewer("4.2.18", "4.2.18"), true)
  assert.Equal(t, IsVersionNewer("4.2.8", "4.2.18"), false)
  assert.Equal(t, IsVersionNewer("4.2.1", "4.2.18"), false)
  assert.Equal(t, IsVersionNewer("4.1", "4.2.18"), false)
  assert.Equal(t, IsVersionNewer("4.0", "4.2.18"), false)
  assert.Equal(t, IsVersionNewer("4", "4.2.18"), false)
  assert.Equal(t, IsVersionNewer("3", "4.2.18"), false)
  assert.Equal(t, IsVersionNewer("3.9", "4.2.18"), false)
  assert.Equal(t, IsVersionNewer("4.9", "4.20.18"), false)
  assert.Equal(t, IsVersionNewer("9.9", "19.20.18"), false)
  assert.Equal(t, IsVersionNewer("4.2.pre.8", "4.2.pre.18"), false)
  assert.Equal(t, IsVersionNewer("4.2.alpha.98", "4.2.beta.18"), false)
}
