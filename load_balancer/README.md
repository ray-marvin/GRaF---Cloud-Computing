# Load Balancer

Distribute functions to worker nodes. There are 5 different scheduling algorithms in the framework like in the below:

- GRaF: Proposing algorithm that maximizes locality while minimizing load imbalance

### Configuring worker nodes

We recommended you to configure two or more worker nodes to use the load balancer correctly.
After [configuring worker nodes](../worker_front), edit `nodes.config.json` before starting the load balancer.
Note that the http protocol and port number should be included in the file.

### How to run

```bash
go run *.go
```

Load balancer will send a similar response with worker node, but additional information from load balancing will be appended to the response like below.

```json
{
	"Result": {
		"statusCode": 200,
		"body": "Hello World2999663"
	},
	"ExecutionTime": 929,
	"InternalExecutionTime": 98,
	"Meta": {
		"ImageBuilt": false,
		"ContainerName": "hf_w1__9572_27887",
		"ImageName": "hf_w1"
	},
	"LoadBalancingInfo": {
		"WorkerNodeId": 0,
		"WorkerNodeUrl": "http://localhost:8222",
		"Algorithm": "ours",
		"AlgorithmLatency": 0.015869140625
	}
}
```