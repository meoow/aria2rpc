package main

import "encoding/json"
import "fmt"
import "flag"
import "log"
import "os"
import "net/http"
import "bytes"

var rpc = flag.String("rpc", "http://127.0.0.1:6800/jsonrpc", "Aria2 rpc server address")
var cookie = flag.String("cookie", "", "Cookies")
var dir = flag.String("dir", "", "Saved dest directory")
var out = flag.String("out", "", "Saved output file name")

func main() {
	flag.Parse()
	URIs := flag.Args()
	if len(URIs) == 0 {
		flag.PrintDefaults()
		os.Exit(0)
	}
	params := makeParamsArry(URIs, *cookie, *dir, *out)
	jsonreq, err := makeJsonStruct(params)
	if err != nil {
		log.Fatal(err)
	}
	addTask(*rpc, jsonreq)
}

func makeParamsArry(uris []string, cookie string, dir string, out string) []interface{} {
	output := make([]interface{}, 0, 2)
	output = append(output, uris)
	opts := make(map[string]interface{}, 9)
	if dir != "" {
		opts["dir"] = dir
	}
	if out != "" {
		opts["out"] = out
	}
	if cookie != "" {
		opts["header"] = []string{fmt.Sprintf("Cookie: %s", cookie)}
	}
	opts["continue"] = "true"
	opts["max-connection-per-server"] = "15"
	opts["split"] = "15"
	opts["min-split-size"] = "10M"
	output = append(output, opts)
	return output
}

func makeJsonStruct(params []interface{}) ([]byte, error) {
	output := make(map[string]interface{}, 4)
	output["jsonrpc"] = "2.0"
	output["id"] = "qwer"
	output["method"] = "aria2.addUri"
	output["params"] = params
	return json.Marshal(output)
}

func addTask(url string, json []byte) {
	body := bytes.NewReader(json)
	resp, err := http.Post(*rpc, "text/plain", body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}
