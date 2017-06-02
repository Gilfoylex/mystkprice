# mystkprice
模拟行情模拟生成器
1.编译golang文件夹下的messenger.go源文件：go build messenger.go

2.从百度云下载自己的镜像（基于Ubuntu 16.04 安装了nodejs和golang运行环境）：
http://pan.baidu.com/s/1sl8VxZf

3.将自己的镜像“mydocker.tar”载入docker：
docker load --input mydocker.tar

4.根据当前目录的Dockerfile构建容器：
docker build -t mess:latest .

5.用交互模式启动mess 容器，并转发8080端口：
docker run -it -p 8080:8080 mess

6.先启动go语言的web服务器：
cd /home/golang  => ./messenger&  (使用后台启动)

7.启动随机生成股票信息的js程序：
node /home/nodejs/stocks.js

8.打开浏览器查看信息：
http://localhost:8080/mystk0.html

![Image text](https://raw.githubusercontent.com/Gilfoylex/mystkprice/master/image.jpg)
