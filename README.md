# civetcat
Civet cat  狸猫 远控
突破网络限制出网

开源版本实现基础功能  商业合作请联系dollarkiller#dollarkiller.com

### 实现协议
- TCP
- UDP 
- KCP
- HTTP
- HTTPS
- ICMP
- DNS
- WebSocket

### 实现功能
- 交互式shell
- 文件上传与下载

### 使用说明
``` 
./backend_ 0.0.0.0:8081 12345   // ./backedn listen_addr token
./agent_ 0.0.0.0:8081 12345     // 远控上线
```

``` 
./backend_ 0.0.0.0:8081 12345    
$ ls agent              // 获取在线机器
127.0.0.1:54920
$ use 127.0.0.1:54920   // 切换到此机器上
$ uname -r              // 执行普通shell命令
5.3.18-1-MANJARO
$ upload a.exe b.exe  (本地) (远程)  // 上传目标文件
$ download a.exe b.exe               // 下载目标文件
```
