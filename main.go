package main

import (
	"flag"
	"log/slog"
	"net"
	"net/http"
	"os"
)

var headers = []string{
	"Cf-Connecting-Ip",
	"X-Forwarded-For",
	"X-Real-IP",
}

func main() {
	port := flag.String("port", "80", "listen port")
	flag.Parse()

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ip := ""
		for _, h := range headers {
			ip = r.Header.Get(h)
		}

		if ip == "" {
			ip, _, _ = net.SplitHostPort(r.RemoteAddr)
		}

		_, err := w.Write([]byte(ip))
		if err != nil {
			slog.Error("could not write response", "error", err)
		}
	})

	err := http.ListenAndServe(net.JoinHostPort("", *port), mux)
	if err != nil {
		slog.Error("could not start server", "error", err)
		os.Exit(1)
	}
}
