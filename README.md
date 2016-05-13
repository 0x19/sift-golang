# sift-golang
Unofficial Sift Science API (Golang client)


## Usage

Lorem

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
