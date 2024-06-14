package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

  router := http.NewServeMux()
  port := 8000

  router.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset utf-8")
    w.Write([]byte("<h1>Hello World</h1>"))
  })

  logger.Info("HTTP server is running", "port", port)
 
  srv := http.Server{
    Addr: ":8000",
    Handler: router,
  }

  srv.ListenAndServe()

}
