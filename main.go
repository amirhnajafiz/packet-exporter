package main

import (
	"context"
	"fmt"
	"os"

	"k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

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

func main() {
	// get cluster configs
	config, err := ClusterConfigs()
	if err != nil {
		panic(err)
	}

	// create an instance of client set
	cs, err := ClientSet(config)
	if err != nil {
		panic(err)
	}

	// get namespace & deployment
	namespace := os.Getenv("NAMESPACE")
	deploymentName := os.Getenv("DEPLOYMENT")

	ctx := context.Background()

	// main loop
	for {
		// list pods
		pods, er := getPodsOfDeployment(cs, namespace, deploymentName)
		if er != nil {
			panic(er)
		}

		// iterate pods
		for _, pod := range pods {
			go Worker(ctx, cs, pod)
		}
	}

	// TODO: connect to NATS
	// TODO: publish over a topic
}
