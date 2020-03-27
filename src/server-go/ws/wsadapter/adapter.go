package wsadapter

import "github.com/jeremija/peer-calls/src/server-go/ws/wsmessage"

type Client interface {
	ID() string
	WriteChannel() chan<- wsmessage.Message
	Metadata() string
}

type Adapter interface {
	Add(client Client) error
	Remove(clientID string) error
	Broadcast(msg wsmessage.Message) error
	Metadata(clientID string) (string, bool)
	Emit(clientID string, msg wsmessage.Message) error
	Clients() (map[string]string, error)
	Size() (int, error)
	Close() error
}