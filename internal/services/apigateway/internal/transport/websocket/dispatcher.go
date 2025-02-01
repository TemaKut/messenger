package websocket

import (
	ws "golang.org/x/net/websocket"
)

type Dispatcher struct {
	conn *ws.Conn
}

func NewDispatcher(conn *ws.Conn) *Dispatcher {
	return &Dispatcher{
		conn: conn,
	}
}

func (d *Dispatcher) DispatchSingle() {

}

func (d *Dispatcher) Close() {

}
