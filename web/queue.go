package web

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ISS-Dating/service-analyzer/service"
	"github.com/nsqio/go-nsq"
)

type Poller struct {
	Service service.Interface
}

type Message struct {
	Name      string
	Content   string
	Timestamp string
}

func (p *Poller) Start() {
	config := nsq.NewConfig()

	config.MaxAttempts = 10
	config.MaxInFlight = 5
	config.MaxRequeueDelay = time.Second * 900
	config.DefaultRequeueDelay = time.Second * 0

	topic := "topic"
	channel := "world"
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		log.Fatal(err)
	}

	consumer.AddHandler(p)                          //Use nsqlookupd to find nsqd instances
	consumer.ConnectToNSQLookupd("nsqlookupd:4161") // wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan // Gracefully stop the consumer.
	consumer.Stop()
}

// HandleMessage implements the Handler interface.
func (p *Poller) HandleMessage(m *nsq.Message) error { //Process the Message
	log.Println(string(m.Body))
	p.Service.UpdateWithMessage(string(m.Body))
	return nil
}
