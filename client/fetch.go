package client

import (
	"log"
  "time"
  "bytes"
  "net/http"
	"io/ioutil"
  "encoding/json"

  "github.com/fl-flow/resource-coordinator/common/error"
  httpresponse "github.com/fl-flow/resource-coordinator/http_server/http/response"
)


func fetch(method string, url string, jsonData []byte) (httpresponse.Ret, *error.Error) {
  request, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("new request to '%s' failed: %v\n", url, err)
	}
  var client = &http.Client{
		Timeout:   time.Second * 30,
	}
  response, err := client.Do(request)
  if err != nil {
		log.Fatalf("request for '%s' failed: %v\n", url, err)
	}
  defer response.Body.Close()
	var ret httpresponse.Ret
  body, _ := ioutil.ReadAll(response.Body)
  if response.StatusCode != 200 {
    log.Printf("request for '%s' status : %v\n body: %v\n", url, response.StatusCode, string(body))
    return ret, &error.Error{
      Code: 80010,
      Hits: string(body),
    }
  }
  err_ := json.Unmarshal(body, &ret)
  if err_ != nil {
    log.Fatalf("data json loads error:  %v\n", err_)
  }
  if ret.Code != 0 {
    return ret, &error.Error{
      Code: ret.Code,
    }
  }
  return ret, nil
}
