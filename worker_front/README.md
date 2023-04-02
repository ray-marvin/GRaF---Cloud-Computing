# Worker Front

Task runner of worker node.
Receives application request with RESTful API and executes the function with Docker.

### Base features

- Build image automatically if image not exists
	- Download the target function source code (from S3)
	- Attach environment by function's demands (python3.7, NodeJS 12, Java 8)
	- Build docker image with the function source code
- Creates containers for function execution
	- After execution, remove container and send responses

### Recommended step before running

As our framework includes automated image build & pull step on running a function.
But pulling the base image(e.g. `python`, `java`, `nodejs`) takes quite a lot time,
we recommend you to pull base images before running the worker front.

```bash
$ docker pull openjdk:8
$ docker pull python:3.7
$ docker pull node:12
``` 

### How to run

We support two arguments on starting worker front.
First argument is the port, and second argument is the option for daemon. 
Default port number is `8222`.

```bash
$ go run *.go start
$ go run *.go start 8222
$ go run *.go start 8222 &
```

After running the worker front, you can access the worker front with visiting `http://localhost:8222`.
To run the function `W1`, request `http://localhost:8222/execute?name=W1`.
Response would be like below:

```json
{
	"Result": {
		"statusCode": 200,
		"body": "Hello World2998950"
	},
	"ExecutionTime": 3373,
	"InternalExecutionTime": 98,
	"Meta": {
		"ImageBuilt": true,
		"ContainerName": "hf_w1__9437_98081",
		"ImageName": "hf_w1"
	}
}
```