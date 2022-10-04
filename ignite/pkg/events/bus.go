package events

import (
	"fmt"

	"github.com/ignite/cli/ignite/pkg/cliui/colors"
)

// DefaultBufferSize defines the default maximum number
// of events that the bus can cache before they are handled.
const DefaultBufferSize = 50

type (
	// Bus defines a bus to send and receive events.
	Bus struct {
		evChan chan Event
	}

	// BusOption configures the Bus.
	BusOption func(*Bus)
)

// WithBufferSize assigns the size of the buffer to use for buffering events.
func WithBufferSize(size int) BusOption {
	return func(bus *Bus) {
		bus.evChan = make(chan Event, size)
	}
}

// NewBus creates a new event bus.
func NewBus(options ...BusOption) Bus {
	bus := Bus{
		evChan: make(chan Event, DefaultBufferSize),
	}

	for _, apply := range options {
		apply(&bus)
	}

	return bus
}

// Send sends a new event to bus.
// This method will block if the event bus buffer is full.
func (b Bus) Send(message string, options ...Option) {
	if b.evChan == nil {
		return
	}

	b.evChan <- New(message, options...)
}

// Sendf sends a new event with a formatted message to bus.
func (b Bus) Sendf(format string, a ...any) {
	b.Send(fmt.Sprintf(format, a...))
}

// SendInfo sends an info event to the bus.
func (b Bus) SendInfo(message string, options ...Option) {
	b.Send(colors.Info(message), options...)
}

// SendInfo sends an error event to the bus.
func (b Bus) SendError(err error, options ...Option) {
	b.Send(colors.Error(err.Error()), options...)
}

// SendView sends a new event for a view to the bus.
// Views are types that implement the `fmt.Stringer` interface
// which allow events with complex message formats.
func (b Bus) SendView(s fmt.Stringer, options ...Option) {
	b.Send(s.String(), options...)
}

// Events returns a read only channel to read the events.
func (b *Bus) Events() <-chan Event {
	return b.evChan
}

// Shutdown shutdowns event bus.
func (b Bus) Shutdown() {
	if b.evChan == nil {
		return
	}

	close(b.evChan)
}
