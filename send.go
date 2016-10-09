package kavenegar

import (
  "net/http"
  "net/url"
  "encoding/json"
  "io/ioutil"
  "log"
  "errors"
)

const (
  api_prefix = "https://api.kavenegar.com/v1/"
  api_postfix = "/sms/send.json"
)

type SMS struct {
  api_key string
  sender string
}


func NewSMS(key string, sender string) SMS{
  return SMS{ api_key: key, sender: sender}
}

func (s *SMS)Send(receptor string, message string)(int , error) {
  // api base url that contains api key
  api_url := api_prefix + s.api_key + api_postfix

  // SMS parameters
  params := url.Values{ "receptor" : {receptor}, "message" : {message}, "sender": {s.sender}}

  // Send sms
  response, err := http.PostForm(api_url, params)

  // Check if an error occured with sms
  if err != nil {
    log.Print(err)
  }

  defer response.Body.Close()

  // Read body of response
  body, err := ioutil.ReadAll(response.Body)

  status := struct {
          Return struct  {
            Status int `json: "status"`
          } `json:"return"`
        }{}

  if err := json.Unmarshal(body, &status); err != nil {
		log.Print(err)
	}

  if status.Return.Status != 200 {
    return status.Return.Status, errors.New("an error occured")
  }

  return status.Return.Status, nil
}
