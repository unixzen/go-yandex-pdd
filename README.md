# go-yandex-pdd
Go API Yandex PDD

https://yandex.ru/dev/pdd/doc/concepts/termin-docpage/


## Features

### Domains

* Register domain
* Status domain
* Detail domain
* Remove domain

### Emails

* Add mail box
* List of email boxes
* Remove email box
* List of new emails at email box

### Email listing

* Add mailing list
* List of mailing list
* Remove mailing list
* Subscribe mailing list
* List of subscribers of mailing list
* Unsubscribe mailing list

### Deputy of administrator

* Add deputy
* List of deputies
* Delete deputy

### Manage DKIM

* Information about DKIM
* Enable DKIM
* Disable DKIM

### Dns records

* Add dns record
* List of dns records
* Edit dns records
* Remove dns records


## Install

```go
go get github.com/unixzen/go-yandex-pdd
```

## Usage

This is example for adding A dns record. More examples you can find at [directory](https://github.com/unixzen/go-yandex-pdd/tree/master/examples) 

```go
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
```

## Limitations

This implementation of API supports only for administrator (for more information see [docs](https://yandex.ru/dev/pdd/doc/concepts/termin-docpage/))


