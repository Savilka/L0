package main

import (
	"github.com/nats-io/stan.go"
	"io/ioutil"
	"log"
)

// script for publish data in channel
func main() {
	sc, err := stan.Connect("test-cluster", "stan", stan.NatsURL("nats://0.0.0.0:4222"))
	if err != nil {
		log.Println(err)
		return
	}

	fileNames := []string{
		"script/model.json",
		"script/model2.json",
		"script/badJson1.json",
		"script/badJson2.json",
		"script/bad.txt",
	}

	for _, name := range fileNames {
		file, err := ioutil.ReadFile(name)
		if err != nil {
			log.Println(err)
			return
		}
		err = sc.Publish("service", file)
		if err != nil {
			log.Println(err)
			return
		}
	}

	err = sc.Close()
	if err != nil {
		log.Println(err)
		return
	}
}
