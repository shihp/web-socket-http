package main

import (
	"golang.org/x/net/websocket"
	"fmt"
	"net/http"
	"log"
	"strings"
	"io/ioutil"
	"github.com/buger/jsonparser"
	"encoding/json"
)

var conns = make(map[string]*websocket.Conn)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Echo(ws *websocket.Conn) {
	println("连接头: %s", ws)
	var err error
	for {
		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}
		params := strings.Split(reply, "?")
		if len(params) != 2 {
			websocket.Message.Send(ws, "参数异常")
			ws.Close()
		}

		userId := strings.Split(params[1], "=")[1]

		switch params[0] {

		case "/login":
			conns[userId] = ws
			break
		default:
			websocket.Message.Send(ws, "不存在的方法")
			ws.Close()
		}

		println(conns, userId)

		fmt.Println("Received back from client: " + reply)

		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

func Login(ws *websocket.Conn) {
	var err error
	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Login Can't receive")
			break
		}

		fmt.Println("Login Received back from client: " + reply)

		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

func Callback(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("read all err %s", err)
	}

	fmt.Println(string(body))

	callId, err := jsonparser.GetString(body, "content", "[0]", "callId")
	if err != nil {
		log.Fatalf("jsonparser GetString %s", err)
	}
	fmt.Println(callId)

	rs := Response{
		Code:    0,
		Message: "success",
	}
	rsJson, err2 := json.Marshal(rs)
	if err2 != nil {
		log.Fatalf("format %s", err2)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(rsJson)

	//var err error
	//msg := "ok"
	//if err = websocket.Message.Send(conns["9873"], msg); err != nil {
	//	fmt.Println("call back")
	//}
}
func Upload(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(w, "helloworld")
}

func main() {
	http.Handle("/", websocket.Handler(Echo))
	http.Handle("/login", websocket.Handler(Login))
	http.HandleFunc("/Upload", Upload)
	http.HandleFunc("/Callback", Callback)
	if err := http.ListenAndServe(":8989", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
