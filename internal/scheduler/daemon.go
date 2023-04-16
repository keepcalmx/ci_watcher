package scheduler

import (
	"context"
	"fmt"
	"net"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

// Daemon 守护进程
type Daemon struct {
	conn *net.Conn
}

func NewDaemon(url string) *Daemon {
	conn, _, _, err := ws.DefaultDialer.Dial(
		context.TODO(),
		"ws://localhost:9527/client",
	)
	if err != nil {
		// handle error
	}

	return &Daemon{
		conn: &conn,
	}
}

// Read 接收DataHub的消息
func (d *Daemon) Read() {
	msg, _, err := wsutil.ReadServerData(*d.conn)
	if err != nil {
		// handle error
	}

	// handle msg
	fmt.Print(msg)
}

// Write 发送消息给DataHub
func (d *Daemon) Write(msg []byte) {
	for {
		err := wsutil.WriteClientMessage(*d.conn, ws.OpText, msg)
		if err != nil {
			// handle error
		}
	}
}

func (d *Daemon) Start() {

}
