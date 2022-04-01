package websocket

import (
	"log"
	"net/http"
	"testing"
)

func TestAll(t *testing.T) {
	str := "test websocket"

	ln, _ := Listen("localhost:5000", nil)
	http.HandleFunc("/ws", ln.(*Listener).Handler)

	go func() {
		err := http.ListenAndServe("localhost:5000", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()

	go func() {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		defer conn.Close()

		bread := make([]byte, len(str))
		_, err = conn.Read(bread)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		conn.Write(bread)
	}()

	// dial
	conn, err := Dial("ws://localhost:5000/ws")
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}
	defer conn.Close()

	// write
	bread := make([]byte, len(str))
	nwrite, err := conn.Write([]byte(str))
	if err != nil || nwrite != len(str) {
		t.Fatalf("failed to listen: %v", err)
	}

	// read
	_, err = conn.Read(bread)
	if err != nil || string(bread) != str {
		t.Fatalf("failed to listen: %v", err)
	}
}
