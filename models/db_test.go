package models

import "testing"

func TestConnectDB(t *testing.T) {
	if err := ConnectDB(); err != nil {
		t.Fatal(err)
	}
}
