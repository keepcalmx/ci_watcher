package websocket

import (
	"encoding/json"
	"net"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/kataras/iris/v12"
	guid "github.com/satori/go.uuid"
)

type SchedulerClient struct {
	uuid string
	Conn *net.Conn

	requests chan *Request
	responds chan *Response
}

func (c *SchedulerClient) Read() {
	defer func() {
		GetHub().Unregister("")
		(*c.Conn).Close()
	}()

	reader := wsutil.NewReader(*c.Conn, ws.StateServerSide)
	decoder := json.NewDecoder(reader)

	for {

		header, err := reader.NextFrame()
		if err != nil {
			panic(err)
		}
		if header.OpCode == ws.OpClose {
			panic(err)
		}

		var req Request
		if err := decoder.Decode(&req); err != nil {
			panic(err)
		}
		// 请求交给Hub处理
		c.requests <- &req
	}
}

func (c *SchedulerClient) Write() {
	defer func() {
		GetHub().Unregister("")
		(*c.Conn).Close()
	}()

	writer := wsutil.NewWriter(*c.Conn, ws.StateServerSide, ws.OpText)
	encoder := json.NewEncoder(writer)

	for {
		resp := <-c.responds
		if err := encoder.Encode(&resp); err != nil {
			panic(err)
		}
		if err := writer.Flush(); err != nil {
			panic(err)
		}
	}
}

func HandleNewSchedulerConn(c iris.Context) {
	// 处理来自调度器的新连接
	// 创建一个新的调度器客户端
	// 将其添加到调度器客户端管理器中
	w, r := c.ResponseWriter(), c.Request()
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		panic(err)
	}

	client := &SchedulerClient{
		uuid:     guid.NewV1().String(),
		Conn:     &conn,
		requests: nil,
		responds: make(chan *Response),
	}

	GetHub().Register(client)

	go client.Read()
	go client.Write()
}
