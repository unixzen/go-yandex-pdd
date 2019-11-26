package main

import (
	"fmt"
	"github.com/unixzen/go-yandex-pdd"
	"log"
)

func main() {
	client := yandex_pdd_api_go.AuthClient("YOUR-YANDEX-API-KEY")
	result, err := client.DelDomain("domain.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", result)
}
