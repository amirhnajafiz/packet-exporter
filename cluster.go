package main

import (
	"fmt"

	"k8s.io/client-go/rest"
)

func ClusterConfigs() (*rest.Config, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get cluster configs: %w", err)
	}

	return config, nil
}
