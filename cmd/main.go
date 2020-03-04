package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-redis/redis"
	"github.com/mronfim/invoicify/pkg/billing"
	"github.com/mronfim/invoicify/pkg/http/rest"
	redisdb "github.com/mronfim/invoicify/pkg/storage/redis"
)

func main() {
	// set up storage
	storageType := flag.String("storage", "redis", "storage type [redis]")

	var invoiceRepo billing.InvoiceRepository

	switch *storageType {
	case "redis":
		invoiceRepo = redisdb.NewRedisInvoiceRepository(redisConnection("localhost:6379"))
	default:
		panic("Unknown storage type")
	}

	invoiceService := billing.NewInvoiceService(invoiceRepo)
	invoiceHandler := billing.NewInvoiceHandler(invoiceService)

	_ = rest.Handler(invoiceHandler)

	errs := make(chan error, 2)
	go func() {
		fmt.Println("Listening on port :3001")
		errs <- http.ListenAndServe(":3001", nil)
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("terminated %s", <-errs)
}

func redisConnection(url string) *redis.Client {
	fmt.Println("Connecting to Redis DB")
	client := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "",
		DB:       0,
	})
	err := client.Ping().Err()

	if err != nil {
		panic(err)
	}
	return client
}
