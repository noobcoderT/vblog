/**
 * vblog: A multi user blog platform, written by vue and go
 * Author: tangzjxb@gmail.com
 * Date: 2020-02-28
**/
package main

import (
    "fmt"
    "net/http"
    "context"
    "vblog/common/logging"
    "os"
    "os/signal"
    "time"
    "vblog/router"
    "vblog/common/setting"
)

func main() {
    router := router.InitRouter()

    s := &http.Server{
        Addr:           fmt.Sprintf("%s:%d", setting.HTTPHost, setting.HTTPPort),
        Handler:        router,
        ReadTimeout:    setting.ReadTimeout,
        WriteTimeout:   setting.WriteTimeout,
        MaxHeaderBytes: 1 << 20,
    }

    go func() {
        if err := s.ListenAndServe(); err != nil {
            logging.Info("Listen: %s\n", err)
        }
    }()

    quit := make(chan os.Signal)
    signal.Notify(quit, os.Interrupt)
    <- quit

    logging.Info("Shutdown Server ...")

    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()
    if err := s.Shutdown(ctx); err != nil {
        logging.Fatal("Server Shutdown:", err)
    }

    logging.Info("Server exiting")
}
