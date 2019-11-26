package main

import (
	"fmt"
	"github.com/unixzen/go-yandex-pdd"
	"log"
)

func main() {
	client := yandex_pdd_api_go.AuthClient("YOUR-YANDEX-API-KEY")
	countersEmail, err := client.GetEmailCounters("YOUR-DOMAIN", "test", "1130000041")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", countersEmail)
}
