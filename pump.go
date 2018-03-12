package go_pumps

type Pump struct {
	handler Handler
	pipe    *chan interface{}
	stop    *chan struct{}
}

type Feed struct {
	pipe *chan interface{}
	stop *chan struct{}
}

type Handler func(obj *interface{})

func NewPump(stop *chan struct{}, pipe *chan interface{}, handler Handler) *Pump {
	return &Pump{handler, pipe, stop}
}

func NewFeed(stop *chan struct{}, pipe *chan interface{}) *Feed {
	return &Feed{pipe, stop}
}

func (f *Feed) Input(obj interface{}) bool {
	select {
	case <-*f.stop:
		return false
	default:
	}

	select {
	case <-*f.stop:
		return false
	case *f.pipe <- obj:

	}

	return true
}

func (p *Pump) Run() (*Feed, error) {
	go func() {
		for {
			select {
			case <-*p.stop:
				return
			case obj := <-*p.pipe:
				p.handler(&obj)
			}
		}
	}()

	select {
	case <-*p.stop:
		return nil, NewPumpError("Pump already stopped")
	default:
		return NewFeed(p.stop, p.pipe), nil
	}
}

func (p *Pump) Stopped() bool {
	select {
	case <-*p.stop:
		return true
	default:
	}

	select {
	case <-*p.stop:
		return true
	default:
		return false
	}
}

func NewPumpError(description string) *PumpError {
	return &PumpError{description: description}
}

type PumpError struct {
	description string
}

func (e *PumpError) Error() string {
	return e.description
}
