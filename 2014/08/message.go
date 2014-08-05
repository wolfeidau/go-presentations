package main

import (
	"encoding/json"
	"net"
)

// START1 OMIT

type MessageSender struct {
	conn net.Conn
}

func (m *MessageSender) Send(msg interface{}) error {
	payload, err := json.Marshal(msg)

	if err != nil {
		return err
	}

	return m.transmitPayload(payload)
}

func (m *MessageSender) transmitPayload(payload []byte) error {
	_, err := m.conn.Write(payload)
	return err
}

// STOP1 OMIT
