package main

import (
    "log"
    "net/http"

    "golang.org/x/net/webdav"
)

func main() {
    dir := "."
    port := "8080"

    handler := &webdav.Handler{
        FileSystem: webdav.Dir(dir),
        LockSystem: webdav.NewMemLS(),
        Logger: func(request *http.Request, err error) {
            if err != nil {
                log.Printf("[%s]: %s, ERROR: %s\n",
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

    log.Fatal(http.ListenAndServe(":" + port, nil))
}
