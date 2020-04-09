# civetcat
Civet cat  狸猫 远控

### 实现协议
- TCP
- UDP 
- KCP
- HTTP
- HTTPS
- RPC
- Email
- ICMP
- DNS
- Socket
- WebSocket

### 实现功能
- 交互式shell
- 文件上传与下载

### 使用说明
``` 
./backend_ 0.0.0.0:8081 12345     
./agent_ 0.0.0.0:8081 12345
```

``` 
./backend_ 0.0.0.0:8081 12345    
$ ls agent // 获取在线机器
127.0.0.1:54920
$ use 127.0.0.1:54920
$ uname -r
5.3.18-1-MANJARO
$ upload a.exe b.exe  (本地) (远程)
$ download a.exe b.exe 
```