package main

// Stdlib
import "fmt"
import "log"
import "flag"
import "time"
import "net/http"
import "net/http/httputil"

func RequestDumpHandler(resp http.ResponseWriter, req *http.Request) {
	request_dump, err := httputil.DumpRequest(req, true)

	t := time.Now()
	if err == nil {
		reqInfo := fmt.Sprintf("\n---\n[%s]\n%s", t.Format(time.StampMilli), request_dump)
		fmt.Println(reqInfo)
		resp.Write([]byte(reqInfo))
	} else {
		msg := fmt.Sprintf("Could not handle request: %s\n", err.Error())
		fmt.Println(msg)
		resp.Write([]byte(msg))
	}
}

func main() {
	addr := flag.String("address", "", "the address to bind to. Default is 0.0.0.0")
	port := flag.String("port", "8080", "the local port to bind to")

	flag.Parse()

	http.HandleFunc("/", RequestDumpHandler)

	fmt.Printf("httpecho starting up on %s:%s...\n", *addr, *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", *addr, *port), nil))
}
