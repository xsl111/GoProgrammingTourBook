package server

import (
	"GoProgrammingTourBook/chatroom/logic"
	"net/http"
)

func RegisterHandle() {

	//广播消息处理
	go logic.Broadcaster.Start()

	http.HandleFunc("/", homeHandleFunc)
	http.HandleFunc("/user_list", userListHandleFunc)
	http.HandleFunc("/ws", WebSocketHandleFunc)
}
