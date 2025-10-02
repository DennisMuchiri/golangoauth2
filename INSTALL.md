### Install

```bash
git clone github.com/DennisMuchiri/ke-soundstream-oauth2.git
cd ./ke-soundstream-oauth2
go install
```

### Add dependency

To use the lib on a another project:

```bash
go env -w GOPRIVATE="github.com/DennisMuchiri/*"
cd /path/to/another-project-dir
go get github.com/DennisMuchiri/ke-soundstream-oauth2
go mod tidy
```

### Usage

#### Sending a message to a queue

main.go

```go
package main

import (
	"log"

	"github.com/DennisMuchiri/ke-soundstream-oauth2/events/kafka"
)

func main() {
	conf := kafka.Conf{
		GroupId:                  "kesoundstream-integrations-synqussd",
		ConsumerBootstrapServers: []string{"localhost:9092"},
		ProducerBootstrapServers: []string{"localhost:9092"},
	}
	sender := kafka.NewSender(&conf)
	topic := "Put-Store-v1"
	m := []byte("My JSON store data")
	if err := sender.Send(topic, m); err != nil {
		log.Println(err)
	}
}
```
