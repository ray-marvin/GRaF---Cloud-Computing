package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"

	"github.com/Prev/HotFunctions/load_balancer/scheduler"
)

var logger *log.Logger
var sched scheduler.Scheduler
var schedType string
var fakeMode = false
var nodes []*scheduler.Node

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	schedType = "ours"

	println("Load node info from `nodes.config.json`")
	nodes = initNodesFromConfig("nodes.config.json")
	fmt.Printf("%d nodes found\n", len(nodes))

	setScheduler()

	port := 8111

	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Printf("The load balancer is starting at the port :%d\n", port)

	http.Handle("/", newRequestHandler())
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic(err)
	}
}

// Init node list from the config file
func initNodesFromConfig(configFilePath string) []*scheduler.Node {
	nodeConfigFile, err := os.Open(configFilePath)
	if err != nil {
		panic(err)
	}
	defer nodeConfigFile.Close()

	var nodeConfigs []string
	byteValue, _ := ioutil.ReadAll(nodeConfigFile)
	json.Unmarshal(byteValue, &nodeConfigs)

	nodes := make([]*scheduler.Node, len(nodeConfigs))
	for i, url := range nodeConfigs {
		nodes[i] = scheduler.NewNode(i, url)
	}
	return nodes
}

func setScheduler() error {
	var err error
	// Proposing Greedy Scheduler
	var t1, t2, t3 int
	if t1, err = strconv.Atoi(os.Getenv("T1")); err != nil {
		t1 = 8
	}
	if t2, err = strconv.Atoi(os.Getenv("T2")); err != nil {
		t2 = 5
	}
	if t3, err = strconv.Atoi(os.Getenv("T3")); err != nil {
		t3 = 3
	}
	println("Using Our Scheduler")
	sched = scheduler.NewOurScheduler(&nodes, uint(t1), uint(t2), t3)
	return nil
}
