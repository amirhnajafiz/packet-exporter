# K8S Sidecar

In Kubernetes there is a feature enabling users to define
“sidecar containers” in specifications.
The new feature is intended to help define behavior for helper
containers in multi-container pods that might assist in
configuration, networking, log and metrics collection, and so on.

In this project we are going to set up a controller in order to add
prometheus metrics exporter, elk logs exporter, and envoy proxy containers
in to user deployed pods.

