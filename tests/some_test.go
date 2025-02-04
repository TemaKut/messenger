package tests_test

import (
	"context"
	"fmt"
	"testing"

	requestv1 "github.com/TemaKut/messenger/pkg/proto/client/gen/request"
	"github.com/TemaKut/messenger/tests/utils/wsclient"
)

func TestSome(t *testing.T) {
	client, err := wsclient.NewClient()
	if err != nil {
		t.Fatalf("error build client. %s", err)
	}

	defer client.Close()

	fmt.Println(client.Request(context.Background(), &requestv1.Request{}))
}
