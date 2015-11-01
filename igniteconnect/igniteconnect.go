package igniteconnect

import (
	"net"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var connection *websocket.Conn

func ConnectToMagna() {
	var res *http.Response
	var err error
	req, _ := http.NewRequest("GET", "http://localhost:8080", nil)
	conn, _ := net.Dial("tcp", "http://localhost:8080")
	connection, res, err = websocket.NewClient(conn, req.URL, req.Header, 1024, 1024)
	_ = res
	_ = err
}

func FormatQuery(queryInput string) {

}

func sendQuery(formatedQuery string) {

}
