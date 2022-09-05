package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
)

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {

    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;

    var print = function(message) {
        var d = document.createElement("div");
        d.textContent = message;
        output.appendChild(d);
        output.scroll(0, output.scrollHeight);
    };

    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };

    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };
    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };

});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output" style="max-height: 70vh;overflow-y: scroll;"></div>
</td></tr></table>
</body>
</html>
`))

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
		logrus.Error("Error ws: " + reason.Error())
	},
}

func (h *Handler) connect(c *gin.Context) {
	err := homeTemplate.Execute(c.Writer, "ws://"+c.Request.Host+"/ws/echo")
	if err != nil {
		logrus.Fatalf("Connect error: %s", err.Error())
	}

	logrus.Infof("connect from: %s", c.Request.RemoteAddr)
}

func (h *Handler) echo(c *gin.Context) {
	writer := c.Writer
	request := c.Request

	con, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		logrus.Fatalf("Upgrade error: %s", err.Error())
	}

	defer func(con *websocket.Conn) {
		err := con.Close()
		if err != nil {
			logrus.Warnf("Connect close with error: %s", err.Error())
		}
	}(con)

	for {
		mt, message, err := con.ReadMessage()
		if err != nil {
			logrus.Warnf("ReadMessage error: %s", err.Error())
		}

		if message != nil {
			logrus.Infof("Recieve message: %s", message)

			err = con.WriteMessage(mt, message)
			if err != nil {
				logrus.Warnf("WriteMessage error: %s", err.Error())
			}
		}
	}
}
