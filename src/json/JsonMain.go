package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Server struct {
	ServerName string
	ServerIP   string
}
 
type Serverslice struct {
	Servers []Server
}

func main() {
	http.HandleFunc("/test", handler)
	http.ListenAndServe("localhost:8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	
	var s Serverslice
	s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}

	fmt.Println(string(b))
	fmt.Fprint(w, string(b))

	//fmt.Println(string(b))
}
