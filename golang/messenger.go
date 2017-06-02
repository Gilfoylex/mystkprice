//golang程序，从nodejs程序获取行情数据并通过websocket推送至网页

package main

import (
    "golang.org/x/net/websocket"
    "fmt"
    "log"
    "net/http"
    "net"  
    "os"
    "container/list"
    "sync"
)

//行情数据缓存队列
var stkprice = list.New()

//channel 和 lock
var c chan int
var l sync.Mutex

//websocket连接容器
var users map[*websocket.Conn]string

//-----------------------------连接行情生成器 begin--------------------
func Net() {  
    //建立socket，监听端口  
    netListen, err := net.Listen("tcp", "localhost:1377")  
    CheckError(err)  
    defer netListen.Close()

    log.Println("Waiting for clients")  
    for {  
        conn, err := netListen.Accept()  
        if err != nil {  
            continue  
        }  
        log.Println(conn.RemoteAddr().String(), " tcp connect success")  
        handleConnection(conn)  
    }
	c <- 2;
} 

//处理连接  
func handleConnection(conn net.Conn) {  
    buffer := make([]byte, 2048)  
    for { 
        n, err := conn.Read(buffer)  
        if err != nil {
            log.Println(conn.RemoteAddr().String(), " connection error: ", err)  
            return  
        }  
        log.Println(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))
		
		l.Lock()
		stkprice.PushBack(string(buffer[:n]))
		l.Unlock()
    }
}

func CheckError(err error) {  
    if err != nil {  
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())  
        os.Exit(1)  
    }  
}
  
//-----------------------------连接行情生成器 end--------------------

//-----------------------------创建webserver及websockserver begin--------------------

func echoHandler(ws *websocket.Conn) {
	
	defer ws.Close()
	
	if _, ok := users[ws]; !ok {  
            users[ws] = "匿名"  
        }
	
	//var err error
	var str string

	for {
		/*if err = websocket.Message.Receive(ws, &str); err != nil {
			break
		} else {
			fmt.Println("从客户端收到：", str)
		}*/
		
		l.Lock()
		i := stkprice.Front()
		if i != nil{
			stkprice.Remove(i)
			str = i.Value.(string)
			l.Unlock()
		}else{
			l.Unlock()
			continue
		}

		//str = "hello, I'm server."
		
		for key, _ := range users {  
                errMarshl := websocket.Message.Send(key, str)  
                if errMarshl != nil {  
                    //移除出错的链接  
                    delete(users, key)  
                    fmt.Println("发送出错...")  
                    break  
                }  
            }

		/*if err = websocket.Message.Send(ws, str); err != nil {
			break
		} else {
			fmt.Println("向客户端发送：", str)
		}*/
	}
}

func Web() {

	users = make(map[*websocket.Conn]string)
	
    http.Handle("/echo", websocket.Handler(echoHandler))
    http.Handle("/", http.FileServer(http.Dir(".")))

    err := http.ListenAndServe(":8080", nil)

    if err != nil {
        panic("ListenAndServe: " + err.Error())
    }
	c <- 1;
}

//-----------------------------创建webserver及websockserver end--------------------

func main() {
	c = make(chan int)
	go Web()
	go Net()
	<- c
	<- c
}
