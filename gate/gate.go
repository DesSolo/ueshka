package gate

import "ueshka/ueshka"

// Sender ...
type Sender interface {
	Send(string) error
	RenderMessage(*ueshka.Operation) string
}
