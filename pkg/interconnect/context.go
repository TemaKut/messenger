package interconnect

import (
	"context"
	"fmt"

	authv1 "github.com/TemaKut/messenger/pkg/proto/service/gen/auth"
)

func SessionContext(parent context.Context, sess *authv1.Session) context.Context {
	return context.WithValue(parent, KeySession, sess)
}

func Session(ctx context.Context) (*authv1.Session, error) {
	sess := ctx.Value(KeySession)
	if sess == nil {
		return nil, fmt.Errorf("error session not found in context")
	}

	if sess, ok := sess.(*authv1.Session); ok {
		return sess, nil
	}

	return nil, fmt.Errorf("error wrong type of session %T", sess)
}
