[![License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://github.com/0x19/sift-golang/tree/master/LICENSE)
[![Build Status](https://travis-ci.org/0x19/sift-golang.svg?branch=master)](https://travis-ci.org/0x19/sift-golang)
[![Go 1.3 Ready](https://img.shields.io/badge/Go%201.3-Ready-green.svg?style=flat)]()
[![Go 1.4 Ready](https://img.shields.io/badge/Go%201.4-Ready-green.svg?style=flat)]()
[![Go 1.5 Ready](https://img.shields.io/badge/Go%201.5-Ready-green.svg?style=flat)]()
[![Go 1.6 Ready](https://img.shields.io/badge/Go%201.6-Ready-green.svg?style=flat)]()

# sift-golang
Unofficial Sift Science API (Golang client)

## Installation

Installation for this package should be straight forward.

```sh
go get github.com/0x19/sift-golang
```

Once it's installed you can import it as

```go
import (
  sift "github.com/0x19/sift-golang"
)
```

## Usage

Bellow you can find few examples on how to use [sift-golang].

### Track an event

Here's an example that sends a $transaction event to sift.

```go
package main

import (
  "log"
  sift "github.com/0x19/sift-golang"
)

sift := sift.New("your_api_key")

// Name of the event. Can be pre-defined such as $transaction and custom
// such as "my_custom_event"
eventName := "$transaction"

data := map[string]interface{}{
  "$user_id":          "someone@someone.com",
  "$transaction_id":   "1233456",
  "$currency_code":    "USD",
  "$amount":           15230000,
  "$time":             1327604222,
  "trip_time":         930,
  "distance_traveled": 5.26,
  "$order_id":         "ORDER-123124124",
}

extras := map[string]interface{}{
  "return_score":  true,
  "return_action": true,
},

record, err := sift.Track(eventName, data, record)
if err != nil {
  panic(err)
}

log.Printf("Got tracking record: %v", record)
```

## Label a user as good or bad

```go
package main

import (
  "log"
  sift "github.com/0x19/sift-golang"
)

sift := sift.New("your_api_key")

data := map[string]interface{}{
  "$is_bad":       true,
  "$reasons":      []string{"$chargeback", "$fraud"},
  "$description":  "Some description about this user..."
}

record, err := sift.Label("some-user-id", data)
if err != nil {
  panic(err)
}

log.Printf("Got label record: %v", record)
```

## Remove label from user

```go
package main

import (
  "log"
  sift "github.com/0x19/sift-golang"
)

sift := sift.New("your_api_key")

record, err := sift.UnLabel("some-user-id")
if err != nil {
  panic(err)
}

log.Printf("Got un label record: %v", record)
```

## Get user's score

```go
package main

import (
  "log"
  sift "github.com/0x19/sift-golang"
)

sift := sift.New("your_api_key")

record, err := sift.Score("some-user-id")
if err != nil {
  panic(err)
}

log.Printf("Got user (score: %v) record: %v", record.Score, record)
```

## Contributions

Please make sure to read [Contribution Guide] if you are interested into code contributions :)

## License

(The MIT License)

Copyright (c) 2016 Nevio Vesic, http://www.neviovesic.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

[Contribution Guide]: <https://github.com/0x19/sift-golang/blob/master/CONTRIBUTING.md>
[sift-golang]: <https://github.com/0x19/sift-golang>
