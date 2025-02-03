package tests_test

import (
	"testing"

	"github.com/TemaKut/messenger/tests/utils/wsclient"
)

func TestSome(t *testing.T) {
	client, err := wsclient.NewClient()
	if err != nil {
		t.Fatalf("error build client. %s", err)
	}

	defer client.Close()
}
