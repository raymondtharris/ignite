package igniteconnect

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var connection *websocket.Conn

//ConnectToMagna function sets up a socket for magna to return data to the user.
func ConnectToMagna() {
	var res *http.Response
	var err error
	req, _ := http.NewRequest("GET", "http://localhost:8080", nil)
	conn, _ := net.Dial("tcp", "http://localhost:8080")
	connection, res, err = websocket.NewClient(conn, req.URL, req.Header, 1024, 1024)
	_ = res
	_ = err
	defer connection.Close()
	for {
		messageType, p, err1 := connection.ReadMessage()
		if err1 != nil {
			return
		}
		fmt.Println(string(p))
		err1 = connection.WriteMessage(messageType, p)
		if err1 != nil {
			return
		}
	}
}

//FormatQuery fixes any problems with the query for Magna
func FormatQuery(queryInput string) {
	//Cursory look at the query to strip out any text that might cause trouble with magna's implementation.
}

func sendQuery(formatedQuery string) {
	req, _ := http.NewRequest("POST", "http://localhost:8080/query", nil)
	// change body to formatedQuery to be sent to magna
	handleRequestResponse(req)
}

func handleRequestResponse(req *http.Request) {

}
