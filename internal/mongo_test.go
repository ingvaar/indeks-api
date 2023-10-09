package internal_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/ingvaar/indeks-api/internal"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestMongoNew(t *testing.T) {
	ctx := context.Background()
	// Creating the container
	req := testcontainers.ContainerRequest{
		Image:        "bitnami/mongodb:5.0-debian-11",
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor:   wait.ForLog("Waiting for connections"),
	}
	// Starting the container
	mongoC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		panic(err)
	}
	// Closing when the test finishes
	defer func() {
		if err := mongoC.Terminate(ctx); err != nil {
			panic(err)
		}
	}()

	// Getting the host of the container
	host, err := mongoC.Host(ctx)
	if err != nil {
		t.Fatal(err)
	}
	// Getting the port of the container
	port, err := mongoC.MappedPort(ctx, "27017")
	if err != nil {
		t.Fatal(err)
	}

	mongoURI := fmt.Sprintf("mongodb://%s:%s", host, port.Port())

	testCases := []struct {
		name     string
		mongoURI string
		timeout  int
		gotError bool
	}{
		{"Valid timeout", mongoURI, 5, false},
		{"Invalid host", "mongodb://test:123", 5, true},
		{"Invalid port", "mongodb://" + host + ":123", 5, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			config := internal.Config{
				Port:     0,
				Timeout:  uint(tc.timeout),
				MongoURI: tc.mongoURI,
			}

			mongo, err := internal.NewMongo(context.Background(), config)
			if tc.gotError {
				assert.Error(t, err)
				assert.Nil(t, mongo)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, mongo)
			}
		})
	}
}
