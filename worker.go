package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"

	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

type Worker struct {
	CS *kubernetes.Clientset
}

func (w Worker) Do(ctx context.Context, pod v1.Pod) {
	PodLogsConnection := w.CS.CoreV1().Pods(pod.Namespace).GetLogs(pod.Name, &v1.PodLogOptions{
		Follow:    true,
		TailLines: &[]int64{int64(10)}[0],
	})

	LogStream, _ := PodLogsConnection.Stream(context.Background())
	defer func(LogStream io.ReadCloser) {
		err := LogStream.Close()
		if err != nil {
			log.Println(err)
		}
	}(LogStream)

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
