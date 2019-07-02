package elasticSearch

import (
	"bytes"
	"context"
	"go-server/utility"
	"io/ioutil"
	"net/http"
	"time"

	loghttp "github.com/motemen/go-loghttp"

	// "github.com/motemen/go-nuts/roundtime"
	"gopkg.in/olivere/elastic.v5"
	elasticapi "gopkg.in/olivere/elastic.v5"
)

var client *elastic.Client
var err error

func NewElasticClient(ctx context.Context, sniff bool, responseSize int) (*elasticapi.Client, error) {
	if client != nil {
		return client, nil
	} else {
		config := utility.GetConfig()
		url := config.ElasticSearchUrl
		// fmt.Println("check1212")
		var httpClient = &http.Client{
			Transport: &loghttp.Transport{
				LogRequest: func(req *http.Request) {
					var bodyBuffer []byte
					if req.Body != nil {
						bodyBuffer, _ = ioutil.ReadAll(req.Body) // after this operation body will equal 0
						// Restore the io.ReadCloser to request
						req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBuffer))
					}
					// fmt.Println("--------- Elasticsearch ---------")
					// fmt.Println("Request URL : ", req.URL)
					// fmt.Println("Request Method : ", req.Method)
					// fmt.Println("Request Body : ", string(bodyBuffer))
				},
				LogResponse: func(resp *http.Response) {
					ctx := resp.Request.Context()
					if _, ok := ctx.Value(loghttp.ContextKeyRequestStart).(time.Time); ok {
						// fmt.Println("Response Status : ", resp.StatusCode)
						// fmt.Println("Response Duration : ", roundtime.Duration(time.Now().Sub(start), 2))
					} else {
						// fmt.Println("Response Status : ", resp.StatusCode)
					}
					// fmt.Println("--------------------------------")
				},
			},
		}
		// fmt.Println("client1111", client)
		client, err = elasticapi.NewClient(elasticapi.SetURL(url), elasticapi.SetSniff(sniff), elasticapi.SetHttpClient(httpClient))
		// fmt.Println("reflect type of", reflect.TypeOf(client))
		if err != nil {
			return nil, err
		}

		err = Ping(ctx, url)
		if err != nil {
			return nil, err
		}

		return client, nil
	}
}

// func ping(ctx context.Context, client *elasticapi.Client, url string) error {

// 	// Ping the Elasticsearch server to get HttpStatus, version number
// 	if client != nil {
// 		info, code, err := client.Ping(url).Do(ctx)
// 		if err != nil {
// 			return err
// 		}

// 		fmt.Printf("Elasticsearch returned with code %d and version %s \n", code, info.Version.Number)
// 		return nil
// 	}
// 	return errors.New("elastic client is nil")
// }
