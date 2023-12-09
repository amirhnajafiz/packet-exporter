package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

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
			// create get logs request
			podLogOpts := v1.PodLogOptions{}
			req := cs.CoreV1().Pods(pod.Namespace).GetLogs(pod.Name, &podLogOpts)

			podLogs, e := req.Stream(ctx) // opening pod logs request
			if e != nil {
				log.Println(e)

				continue
			}

			// create a buffer to read logs
			buf := new(bytes.Buffer)

			_, err = io.Copy(buf, podLogs)
			if err != nil {
				log.Println(err)

				continue
			}

			// split logs by the ending delimiter
			logs := strings.Split(buf.String(), "\n")
			for _, tmp := range logs {
				if len(tmp) < 2 {
					continue
				}

				// encode logs
				fmt.Println(EncodeLog(tmp))
			}
		}
	}

	// TODO: create logs option
	// TODO: get deployment name
	// TODO: get deployment pods
	// TODO: create a worker for each pod
	// TODO: connect to NATS
	// TODO: monitor each pod for incoming logs
	// TODO: unmarshal to an struct
	// TODO: publish over a topic
}
