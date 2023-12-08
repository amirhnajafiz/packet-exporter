package main

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// ClusterConfigs loads configs of k8s cluster
func ClusterConfigs() (*rest.Config, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get cluster configs: %w", err)
	}

	return config, nil
}

// ClientSet opens a new client
func ClientSet(config *rest.Config) (*kubernetes.Clientset, error) {
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create client set: %w", err)
	}

	return client, nil
}
