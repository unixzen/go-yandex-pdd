package main

import (
	"fmt"
	"github.com/unixzen/go-yandex-pdd"
	"log"
)

func main() {
	client := yandex_pdd_api_go.AuthClient("YOUR-YANDEX-API-KEY")
	record, err := client.AddRecord("YOUR-DOMAIN", "A", "testttt", "60", "127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", record)
}
