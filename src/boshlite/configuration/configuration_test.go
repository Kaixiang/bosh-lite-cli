package configuration

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaults(t *testing.T) {
	config := Default()

	assert.Equal(t, config.Target, "http://api.10.244.0.34.xip.io")
	assert.Equal(t, config.IpRange, "10.244.0.0/22")
	assert.Equal(t, config.Gateway, "192.168.50.4")
	assert.Equal(t, config.OStype, "Darwin")
	assert.Equal(t, config.Version, "1")
}
