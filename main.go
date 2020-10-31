package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/webdav"
)

func main() {
	config, err := NewConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, usage())
		os.Exit(1)
	}

	handler := &webdav.Handler{
		FileSystem: webdav.Dir(config.Dir),
		LockSystem: webdav.NewMemLS(),
		Logger: func(request *http.Request, err error) {
			if err != nil {
				log.Printf("[%s]: %s, ERROR: %v\n",
					request.Method,
					request.URL,
					err)
			} else {
				log.Printf("[%s]: %s\n",
					request.Method,
					request.URL)
			}
		},
	}
	http.Handle("/", handler)

	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
