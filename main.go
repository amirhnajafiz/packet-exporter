package main

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

	// TODO: create logs option
	// TODO: get namespace from env
	// TODO: get deployment name
	// TODO: get deployment pods
	// TODO: create a worker for each pod
	// TODO: connect to NATS
	// TODO: monitor each pod for incoming logs
	// TODO: unmarshal to an struct
	// TODO: publish over a topic
}
