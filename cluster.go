package main

import (
	"context"
	"fmt"

	"k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// clusterConfigs loads configs of k8s cluster
func clusterConfigs() (*rest.Config, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get cluster configs: %w", err)
	}

	return config, nil
}

// clientSet opens a new client
func clientSet(config *rest.Config) (*kubernetes.Clientset, error) {
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create client set: %w", err)
	}

	return client, nil
}

// get list of pods based on the deployment
func getPodsOfDeployment(client *kubernetes.Clientset, namespace string, deployment string) ([]v1.Pod, error) {
	options := metaV1.ListOptions{
		LabelSelector: fmt.Sprintf("app=%s", deployment),
	}

	pods, err := client.CoreV1().Pods(namespace).List(context.Background(), options)
	if err != nil {
		return nil, fmt.Errorf("failed to get pods: %w", err)
	}

	return pods.Items, nil
}
