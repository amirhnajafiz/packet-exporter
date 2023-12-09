package main

import (
	"context"
	"fmt"
	"os"
	"sync"

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
	config, err := clusterConfigs()
	if err != nil {
		panic(err)
	}

	// create an instance of client set
	cs, err := clientSet(config)
	if err != nil {
		panic(err)
	}

	// get namespace & deployment
	namespace := os.Getenv("NAMESPACE")
	deploymentName := os.Getenv("DEPLOYMENT")

	ctx := context.Background()
	wg := sync.WaitGroup{}
	workerInstance := worker{CS: cs}

	// list pods
	pods, er := getPodsOfDeployment(cs, namespace, deploymentName)
	if er != nil {
		panic(er)
	}

	// iterate pods
	for _, pod := range pods {
		wg.Add(1)

		// start a new go routine
		go func(p v1.Pod) {
			workerInstance.Do(ctx, p)
			wg.Done()
		}(pod)
	}

	wg.Wait()

	// TODO: connect to NATS
	// TODO: publish over a topic
}
