package monitoring

import (
	"github.com/sangyun-han/strongswan-go/vici"
	"fmt"
)

func ExecuteTunnelMonitor() {
	client, err := vici.NewViciClientFromDefaultSocket()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// To be implemented
	sasList, err := client.ListSas("", "")
	fmt.Println(sasList)
	fmt.Println("End of test")
}

/*
- IPsec tunnel monitor as a daemon mode
- IPsec tunnel monitor as a passive mode
 */