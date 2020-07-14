package workers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"ava.fund/alpha/Post-Covid/warehouse_cloning/src/internal/utils"
)

func consumerProcess(request *Request) ([]byte, error) {

	client := &http.Client{}

	for i := 0; i < utils.Config.Source.Attempts; i++ {

		response, err := client.Do(request.HttpReq)
		if err != nil {
			utils.Error("[consumer.go] %v", err)
		}

		defer response.Body.Close()

		if response.StatusCode == utils.HttpTooManyRequests {
			utils.Debug("[consumer.go] Too many requests. Wait for %v ", utils.Config.Source.Wait * time.Second)
			time.Sleep(utils.Config.Source.Wait * time.Second)
			continue
		}

		var data []byte
		data, err = ioutil.ReadAll(response.Body)

		if err != nil {
			utils.Error("[consumer.go] %v", err)
		}

		return data, err
	}

	err := fmt.Errorf("Reached %d attempts", utils.Config.Source.Attempts)
	return nil, err
}

func Consumer(requests chan *Request, responses chan *Response, workerGroup *sync.WaitGroup) {
	utils.Debug("[consumer.go] Begin")

	workerGroup.Add(1)
	go func() {
		defer workerGroup.Done()
		defer utils.Debug("[consumer.go] End")

		for {
			select {
			case request, more := <-requests:
				if !more {
					utils.Debug("[consumer.go] Terminate a consumer")
					return
				}

				data, err := consumerProcess(request)
				if err != nil {
					utils.Error("[consumer.go] %v", err)
				}

				responses <- &Response{
					Data:    data,
					Request: request,
				}
			}
		}
	}()
}
