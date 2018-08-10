package main

import (
    // Standard library packages
    "fmt"
    "os"
    "log"
    "net"
    "net/http"

    // Third party packages
    "github.com/julienschmidt/httprouter"
)



// https://blog.golang.org/context/userip/userip.go
func getIP(w http.ResponseWriter, req *http.Request, _ httprouter.Params){
    podName := os.Getenv("POD_NAME")
    if podName == "" {
      podName = "unkown"
    }
    cntName := os.Getenv("CONTAINER_NAME")
    if podName == "" {
      podName = "unkown"
    }
    fmt.Fprintf(w, "You've hit cnt:%s at path:%s on pod:%s\n", cntName, req.URL.Path, podName)
    ip, port, err := net.SplitHostPort(req.RemoteAddr)
    if err != nil {
        fmt.Fprintf(w, "userip: %q is not IP:port", req.RemoteAddr)
    }

    userIP := net.ParseIP(ip)
    if userIP == nil {
        //return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
        fmt.Fprintf(w, "userip: %q is not IP:port", req.RemoteAddr)
        return
    }
    fmt.Fprintf(w, "IP: %s\n", ip)
    fmt.Fprintf(w, "Port: %s\n", port)
}


func main() {
    // Instantiate a new router
    r := httprouter.New()

    r.GET("/", getIP)

    port := os.Getenv("HTTP_PORT")
    if port == "" {
      port = "8080"
    }
    addr := fmt.Sprintf("%s:%s", os.Getenv("HTTP_HOST"), port)
    l, err := net.Listen("tcp", addr)
    if err != nil {
        log.Fatal(err)
    }
    // The browser can connect now because the listening socket is open.



    // Start the blocking server loop.

    log.Printf("Start Webserver on %s", addr)
    log.Fatal(http.Serve(l, r))
}
