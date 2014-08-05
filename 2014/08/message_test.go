package main

import (
	"bytes"
	"net"
	"reflect"
	"testing"
)

// START1 OMIT

var expectedJSON = `{"Name":"mark","Email":"mark@wolfe.id.au"}`

type FackConn struct {
	net.TCPConn
	buf bytes.Buffer
}

func (f *FackConn) Write(b []byte) (n int, err error) {
	return f.buf.Write(b)
}

func TestMyFunc(t *testing.T) {
	fc := &FackConn{}
	ms := &MessageSender{fc}
	u := &User{"mark", "mark@wolfe.id.au"}
	ms.Send(u)
	resultJSON := string(fc.buf.Bytes())
	if !reflect.DeepEqual(resultJSON, expectedJSON) {
		t.Fatalf("expected %s, got: %s", expectedJSON, resultJSON)
	}
}

// STOP1 OMIT
