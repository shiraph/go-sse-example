package event_server

import (
	"time"

	"github.com/r3labs/sse/v2"
)

type EventServer struct {
	Server     *sse.Server
	StreamName string
	Interval   time.Duration
}

func NewEventServer(server *sse.Server, streamName string, interval time.Duration) *EventServer {
	return &EventServer{
		Server:     server,
		StreamName: streamName,
		Interval:   interval,
	}
}

func (e *EventServer) Init(numberOfMessages *int64) {
	e.createStream()
	go func() {
		i := 0
		for {
			i++
			e.publishEvent(time.Now().String())
			if numberOfMessages != nil && int64(i) == *numberOfMessages {
				e.removeStream()
				break
			}
			time.Sleep(e.Interval)
		}
	}()
}

func (e *EventServer) createStream() {
	e.Server.CreateStream(e.StreamName)
}

func (e *EventServer) removeStream() {
	e.Server.RemoveStream(e.StreamName)
}

func (e *EventServer) publishEvent(content string) {
	e.Server.Publish(e.StreamName, &sse.Event{
		Data: []byte(content),
	})
}
