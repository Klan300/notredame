package main

import (
    "fmt"
    "os"
    "sync"

    "ava.fund/alpha/notredame/warehouse_cloning/src/internal/utils"
    "ava.fund/alpha/notredame/warehouse_cloning/src/internal/workers"
)

func main() {

    if len(os.Args) != 2 {
        fmt.Printf("Usage: %s <config>\n", os.Args[0])
        os.Exit(0)
    }
    utils.LoadConfig(os.Args[1])
    utils.Debug("[main.go] Begin")

    securities  := workers.RetrieveSecurities()
    workerGroup := &sync.WaitGroup{}
    requests    := workers.Producer(securities)
    responses   := make(chan *workers.Response)

    for i := 0; i < utils.Config.Source.Consumers; i++ {
        workers.Consumer(requests, responses, workerGroup)
    }
    workers.Writer(responses)
    workerGroup.Wait()

    utils.Debug("[main.go] End")
}
