package kaca

type disptcher struct {
	//Registered conenctions.
	connections map[*connection]bool
	broadcast   chan []byte
	sub         chan string
	pub         chan string
	register    chan *connection
	unregister  chan *connection
}

func NewDispatcher() *disptcher {
	return &disptcher{
		broadcast:  make(chan []byte),
		sub:        make(chan string),
		pub:        make(chan string),
		register:   make(chan *connection),
		unregister: make(chan *connection),
	}
}

func (d *disptcher) run() {
	for {

	}
}
