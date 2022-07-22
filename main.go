package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	_ "github.com/benthosdev/benthos/v4/public/components/all"
	"github.com/benthosdev/benthos/v4/public/service"
)

func main() {
	panicOnErr := func(err error) {
		if err != nil {
			panic(err)
		}
	}

	content, err := ioutil.ReadFile("pipeline.yaml")
	if err != nil {
		log.Fatal(err)
	}

	builder := service.NewStreamBuilder()
	err = builder.SetYAML(string(content))
	panicOnErr(err)

	stream, err := builder.Build()
	panicOnErr(err)

	ctx, _ := context.WithCancel(context.Background())
	go stream.Run(ctx)
	fmt.Println("waiting")
	stream.StopWithin(10 * time.Second)
	fmt.Println("stopping")
	// cancel()
	// fmt.Println("canceled")
	fmt.Println("waiting ...")
	time.Sleep(15 * time.Second)
	fmt.Println("done")
}
