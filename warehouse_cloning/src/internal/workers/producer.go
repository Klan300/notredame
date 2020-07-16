package workers

import (
	"fmt"
	"net/http"
	"time"

	"ava.fund/alpha/Post-Covid/warehouse_cloning/src/internal/utils"
)

func Producer(securities []Security) chan *Request {
    utils.Debug("[producer.go] Begin")
    utils.Debug("[producer.go] Create a request channel")

    requests := make(chan *Request)
    go func() {
        defer utils.Debug("[producer.go] End")
        for _, security := range securities {
            for _, document := range utils.Config.Documents {  
                switch document {

                case "profile":
                    endpoint := fmt.Sprintf(
                        utils.Endpoints[document],
                        utils.Config.Source.Host, 
                        security.Symbol, 
                        utils.Config.Source.Token)

                    httpreq, _ := http.NewRequest("GET", endpoint, nil)

                    requests <- &Request {
                        Document: document,
                        Exchange: security.Exchange,
                        Symbol  : security.Symbol,
                        HttpReq : httpreq,
                    }

                case "financials":
                    for _, statement := range utils.Statements {
                        for _, frequency := range utils.Frequency {

                            endpoint := fmt.Sprintf(
                                utils.Endpoints[document],
                                utils.Config.Source.Host,
                                security.Symbol,
                                statement,
                                frequency,
                                utils.Config.Source.Token)
 
                            httpreq, _ := http.NewRequest("GET", endpoint, nil)

                            requests <- &Request{
                                Document : document,
                                Exchange : security.Exchange,
                                Symbol   : security.Symbol,
                                Statement: statement,
                                Frequency: frequency,
                                HttpReq  : httpreq,
                            }
                        }

                    }
                case "candle":
                    currentTime := time.Now()
                    tenyearago := currentTime.AddDate(-10, 0, 0)
                    
                    endpoint := fmt.Sprintf(
                        utils.Endpoints[document],
                        utils.Config.Source.Host, 
                        security.Symbol, 
                        tenyearago.Unix(),
                        currentTime.Unix(),
                        utils.Config.Source.Token)
                    
                    httpreq, _ := http.NewRequest("GET", endpoint, nil)

                    requests <- &Request{
                        Document : document,
                        Exchange : security.Exchange,
                        Symbol   : security.Symbol,
                        HttpReq  : httpreq,
                    }

                }
            }
        }
        utils.Debug("[producer.go] Close the request channel")
        close(requests)
    }()
    return requests
}
