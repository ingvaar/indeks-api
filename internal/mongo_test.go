package internal_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/ingvaar/indeks-api/internal"
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
		expected error
	}{
		{"Valid timeout", mongoURI, 5, nil},
		{"Invalid host", "mongodb://test:123", 5, fmt.Errorf("server selection error: context deadline exceeded, current topology: { Type: Unknown, Servers: [{ Addr: test:123, Type: Unknown, Last error: dial tcp: lookup test: no such host }, ] }")},
		{"Invalid port", "mongodb://"+host+":123", 5, fmt.Errorf("server selection error: context deadline exceeded, current topology: { Type: Unknown, Servers: [{ Addr: localhost:123, Type: Unknown, Last error: dial tcp 127.0.0.1:123: connect: connection refused }, ] }")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			config := internal.Config{
				Port:     0,
				Timeout:  uint(tc.timeout),
				MongoURI: tc.mongoURI,
			}

			mongo, err := internal.NewMongo(context.Background(), config)
			if err != nil && tc.expected == nil {
				t.Fatalf("Got an error but didn't expect one: %v", err)
			}
			if err == nil && tc.expected != nil {
				t.Fatalf("Expected an error but didn't get one")
			}
			if err != nil && err.Error() != tc.expected.Error() {
				t.Fatalf("Expected error: %v, got: %v", tc.expected, err)
			}
			if mongo == nil && tc.expected == nil {
				t.Fatalf("Expected a Mongo instance, got nil")
			}
		})
	}
}