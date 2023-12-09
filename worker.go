package main

import (
	"bufio"
	"context"
	"fmt"

	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

func Worker(ctx context.Context, cs *kubernetes.Clientset, pod v1.Pod) {
	PodLogsConnection := cs.CoreV1().Pods(pod.Namespace).GetLogs(pod.Name, &v1.PodLogOptions{
		Follow:    true,
		TailLines: &[]int64{int64(10)}[0],
	})

	LogStream, _ := PodLogsConnection.Stream(context.Background())
	defer LogStream.Close()

	reader := bufio.NewScanner(LogStream)

	var line string

	for {
		select {
		case <-ctx.Done():
			break
		default:
			for reader.Scan() {
				line = reader.Text()

				fmt.Println(EncodeLog(line))
			}
		}
	}
}
