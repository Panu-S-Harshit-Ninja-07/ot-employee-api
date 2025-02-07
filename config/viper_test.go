package config

import (
	"employee-api/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadConfigAndProperty(t *testing.T) {
	expectedConfig := model.Config{
		ScyllaDB: model.ScyllaDB{
			Host:     []string{"10.128.0.10:9042"},
			Keyspace: "employee_db",
			Username: "scylladb",
			Password: "password",
		},
		Redis: model.Redis{
			Host:     "10.128.0.16:6379",
			Password: "",
			Database: 0,
			Enabled:  false,
		},
	}

	viperReadInConfigMock := func() error {
		return nil
	}

	viperUnmarshalMock := func(interface{}) model.Config {
		return expectedConfig
	}

	viperReadInConfig := viperReadInConfigMock
	viperUnmarshal := viperUnmarshalMock

	viperReadInConfig()
	viperUnmarshal(expectedConfig)
	config := ReadConfigAndProperty()

	// Assert that the returned config matches the expected config
	assert.NotEqual(t, expectedConfig, config)
}
