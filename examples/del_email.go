package main

import (
	"fmt"
	"github.com/unixzen/go-yandex-pdd"
	"log"
)

func main() {
	client := yandex_pdd_api_go.AuthClient("YOUR-YANDEX-API-KEY")
	delEmail, err := client.DelEmail("YOUR-DOMAIN", "test", "113000004")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", delEmail)
}
