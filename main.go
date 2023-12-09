package main

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/nats-io/nats.go"
	"k8s.io/api/core/v1"
)

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
	natsCluster := os.Getenv("NATS_HOST")
	topic := os.Getenv("NATS_TOPIC")

	// open nats connection
	nc, err := nats.Connect(natsCluster)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	wg := sync.WaitGroup{}

	// create a worker instance
	workerInstance := worker{
		Conn:  nc,
		CS:    cs,
		Topic: topic,
	}

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

	log.Println("logs consuming started ...")

	wg.Wait()

	log.Println("logs consuming stopped ...")
}
