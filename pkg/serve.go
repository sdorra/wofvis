//go:generate go-bindata-assetfs -prefix "../client" -pkg "pkg" ../client/build/...
package pkg

import (
	"encoding/json"
	"net/http"

	"fmt"

	"github.com/urfave/cli"
)

type WebServer struct {
	createNodes NodeFactory
}

func (ws *WebServer) jsonNodesHandler(w http.ResponseWriter, r *http.Request) {
	nodes, err := ws.createNodes()
	if err != nil {
		http.Error(w, "failed to read nodes", 500)
		return
	}

	data, err := json.Marshal(nodes)
	if err != nil {
		http.Error(w, "failed to marshal nodes", 500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func Serve(ctx *cli.Context) error {
	ws := &WebServer{
		createNodes: CreateNodeFactory(ctx),
	}

	http.Handle("/", http.FileServer(assetFS()))
	http.HandleFunc("/nodes.json", ws.jsonNodesHandler)
	addr := ctx.String("addr")
	fmt.Println("webserver is listening at", addr)
	return http.ListenAndServe(addr, nil)
}
