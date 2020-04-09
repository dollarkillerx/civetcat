package socket_conn

import (
	"civetcat/pb"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
	"runtime"
	"time"
)

type AgentWebSocket struct {
	token    string
	uri      string
	conn     *websocket.Conn
	initData []byte
}

func NewAgentWebSocket(uri string, token string) *AgentWebSocket {
	result := &AgentWebSocket{}
	result.token = token
	u := url.URL{Scheme: "ws", Host: uri, Path: "/registration"}
	result.uri = u.String()
	conn, _, err := websocket.DefaultDialer.Dial(result.uri, http.Header{"Sec-Websocket-Protocol": []string{token}})
	if err != nil {
		log.Fatalln(err)
	}
	result.conn = conn
	resp := &pb.Resp{
		RespItem: &pb.Resp_Heartbeat{Heartbeat: &pb.DeHeartbeat{CPU: uint32(runtime.NumCPU())}},
	}
	marshal, err := proto.Marshal(resp)
	if err != nil {
		log.Fatalln(err)
	}
	result.initData = marshal

	result.Write(marshal)
	go result.heartbeat()
	return result
}

func (a *AgentWebSocket) heartbeat() error {
	ticker := time.NewTicker(time.Millisecond * 500)
	for {
		select {
		case <-ticker.C:
			for {
				err := a.Write(a.initData)
				if err != nil {
					a.Close()
					for {
						err = a.Reconnect()
						if err != nil {
							continue
						} else {
							break
						}
					}
					continue
				} else {
					break
				}
			}
		}
	}
}

func (a *AgentWebSocket) Reconnect() error {
	conn, _, err := websocket.DefaultDialer.Dial(a.uri, http.Header{"Sec-Websocket-Protocol": []string{a.token}})
	if err != nil {
		return err
	}
	a.conn = conn
	a.Write(a.initData)
	return nil
}

func (a *AgentWebSocket) Write(data []byte) error {
	return a.conn.WriteMessage(websocket.TextMessage, data)
}

func (a *AgentWebSocket) Read() ([]byte, error) {
	_, p, err := a.conn.ReadMessage()
	return p, err
}

func (a *AgentWebSocket) Close() {
	a.conn.Close()
}
