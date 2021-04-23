package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	http.HandleFunc("/count", func(w http.ResponseWriter, req *http.Request) {
		err := rdb.Incr(ctx, "count").Err()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		val, err := rdb.Get(ctx, "count").Result()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "count %s", val)
	})
	fmt.Println("http server is listening on :8080")
	http.ListenAndServe(":8080", nil)
}
