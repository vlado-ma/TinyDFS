package messaging

import (
	"os"
	"runtime"
	"testing"

	"github.com/google/uuid"
)

var queueConnParams = ConnParams{
	"localhost", "3333", "tcp",
}

func init() {
	runtime.LockOSThread()
}

func TestMain(m *testing.M) {
	// setup
	var mainNode = NewNode(queueConnParams, queueConnParams, true)
	go mainNode.Run()
	go func() {
		retCode := m.Run()
		//cleanup
		mainNode.CloseConn()
		os.Exit(retCode)
	}()
}

func TestConnectingToQueue(t *testing.T) {
	var node = NewNode(queueConnParams, queueConnParams, false)
	err := node.Run()
	if err != nil {
		t.Fail()
	}
}

func TestSendingToQueue(t *testing.T) {
	var node = NewNode(queueConnParams, queueConnParams, false)
	err := node.Run()
	if err != nil {
		t.Fail()
	}
	var message = Message{Key: uuid.New(), Topic: "Test", Payload: []byte("Hello world!")}
	node.SendMessage(message)
}

func TestCloseNode(t *testing.T) {
	var node = NewNode(queueConnParams, queueConnParams, false)
	err := node.Run()
	if err != nil {
		t.Fail()
	}
	closeerr := node.CloseConn()
	if closeerr != nil {
		t.Fail()
	}
}
