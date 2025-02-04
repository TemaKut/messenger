package websocket

import (
	"context"
	"fmt"
	"sync"

	requestv1 "github.com/TemaKut/messenger/pkg/proto/client/gen/request"
	responsev1 "github.com/TemaKut/messenger/pkg/proto/client/gen/response"
	authv1 "github.com/TemaKut/messenger/pkg/proto/service/gen/auth"
	"github.com/google/uuid"
	ws "golang.org/x/net/websocket"
	"google.golang.org/protobuf/proto"
)

type ConnectedSession struct {
	conn *ws.Conn
	sess *authv1.Session
}

type SessionManager struct {
	sessions map[string]*ConnectedSession

	mu sync.Mutex
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessions: make(map[string]*ConnectedSession),
	}
}

func (s *SessionManager) RegisterSessionFromConnection(conn *ws.Conn) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	// TODO возможно придётся регать сессии в БД

	sessId := uuid.NewString()

	s.sessions[sessId] = &ConnectedSession{
		conn: conn,
		sess: &authv1.Session{
			Id:        sessId,
			Ip:        conn.Request().RemoteAddr, // TODO сделать real ip
			UserAgent: conn.Request().UserAgent(),
			IsActive:  true,
		},
	}

	return sessId
}

func (s *SessionManager) HandleSessionRequests(ctx context.Context, id string) error {
	cs, err := s.session(id)
	if err != nil {
		return fmt.Errorf("error fetch session. %w", err)
	}

	// ctx = interconnect.SessionContext(ctx, cs.sess)

	for {
		if !cs.sess.GetIsActive() {
			break
		}

		var data []byte

		if err := ws.Message.Receive(cs.conn, &data); err != nil {
			return fmt.Errorf("error receive message from session id %s", id)
		}

		var req requestv1.Request

		if err := proto.Unmarshal(data, &req); err != nil {
			return fmt.Errorf("error unmarshal request. %w", err)
		}

		resp := &responsev1.Response{
			Id: "req.GetId()",
		}

		b, err := proto.Marshal(resp)
		if err != nil {
			return fmt.Errorf("error marshal response. %w", err)
		}

		if err := ws.Message.Send(cs.conn, b); err != nil {
			return fmt.Errorf("error send message to session id %s", id)
		}
	}

	return nil
}

func (s *SessionManager) session(id string) (*ConnectedSession, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	cs, ok := s.sessions[id]
	if !ok {
		return nil, fmt.Errorf("error session %s not found", id)
	}

	return cs, nil
}
