package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"fadhilla-hentino/go-kit/sample-rest/datasource"
	"fadhilla-hentino/go-kit/sample-rest/repository/user"
	"github.com/go-kit/kit/log"
)

func main() {
	ctx := context.Background()
	errChan := make(chan error)

	userRepository := user.NewUserRepo(datasource.NewRedisClient("localhost", "6379",
		"redispass")) // TODO: get from config
	userSvc := NewUserService(userRepository)

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	r := MakeHttpHandler(ctx, userSvc, logger)

	go func() {
		fmt.Println("Http Server start at port:8080")
		handler := r
		errChan <- http.ListenAndServe(":9000", handler)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	fmt.Println(<-errChan)
}
