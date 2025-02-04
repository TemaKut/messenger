package wsclient

import (
	"context"
	"fmt"

	requestv1 "github.com/TemaKut/messenger/pkg/proto/client/gen/request"
	reponsev1 "github.com/TemaKut/messenger/pkg/proto/client/gen/response"
	"golang.org/x/net/websocket"
	"google.golang.org/protobuf/proto"
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

// Синхронный метод. В нём упускается сравнение id запроса и ответа
func (c *Client) Request(ctx context.Context, req *requestv1.Request) (*reponsev1.Response, error) {
	b, err := proto.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshal request. %w", err)
	}

	if err := websocket.Message.Send(c.conn, b); err != nil {
		return nil, fmt.Errorf("error send request. %w", err)
	}

	var data []byte

	if err := websocket.Message.Receive(c.conn, &data); err != nil {
		return nil, fmt.Errorf("error receive response. %w", err)
	}

	var resp reponsev1.Response

	if err := proto.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("error unmarshal response. %w", err)
	}

	return &resp, nil
}

func (c *Client) Close() {
	c.conn.Close()
}
