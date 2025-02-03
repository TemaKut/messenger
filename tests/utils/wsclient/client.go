package wsclient

import (
	"context"
	"fmt"

	requestv1 "github.com/TemaKut/messenger/pkg/proto/client/gen/request"
	"golang.org/x/net/websocket"
)

type Client struct {
	conn *websocket.Conn
}

func NewClient() (*Client, error) {
	conn, err := websocket.Dial("ws://localhost:8001/ws", "", "http://localhost:8334")
	if err != nil {
		return nil, fmt.Errorf("error dial ws client. %w", err)
	}

	return &Client{conn: conn}, nil
}

func (c *Client) Request(ctx context.Context, req *requestv1.Request) {
	// websocket.Message.Send(c.conn, )
}

func (c *Client) Close() {
	c.conn.Close()
}
