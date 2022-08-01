package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/rk295/go-automox/automox"
)

func main() {
	APIKey := os.Getenv("AUTOMOX_API_KEY")
	ctx := context.Background()
	api, err := automox.New(ctx, APIKey, nil)
	if err != nil {
		log.Fatal(err)
	}

	s, err := api.Servers().List(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range s {
		fmt.Printf("%s uptime=%v\n", s.Name, s.Uptime)
	}

	fmt.Println("number of servers:", len(s))
}
