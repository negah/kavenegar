kavenegar.go [![Build Status](https://travis-ci.org/negah/kavenegar.svg?branch=master)](https://travis-ci.org/negah/kavenegar)
============

[Kavenegar](http://kavenegar.com/) SMS web service client in golang.

Install
-------
```shell
go get github.com/negah/kavenegar
```

Example
-------
```go
package main

import (
  "github.com/negah/kavenegar"
  "log"
  "fmt"
)

func main(){
  // Setup new Client
  // kavenegar.SMS(api_key, sender)
  // sender is your dedicate phone_number in kavenagar. if you dont have one, left it blank but pass ""
  sms := kavenegar.NewSMS("api_key","")

  // Send sms
  status, err := sms.Send("09121231231", "سلام من به تو یار قدیمی")

  if err != nil {
    log.Printf("Response status code: %d", status)
    log.Fatal(err)
  } else {
    fmt.Printf("SMS sent successfully.")
  }
}

```

License
-------
MIT
