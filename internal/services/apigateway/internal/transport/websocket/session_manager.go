package websocket

import (
	"context"
	"fmt"
	"sync"

	authv1 "github.com/TemaKut/messenger/pkg/proto/service/gen/auth"
	"github.com/google/uuid"
	ws "golang.org/x/net/websocket"
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

		if err := ws.Message.Send(cs.conn, string(data)+"resp"); err != nil {
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
