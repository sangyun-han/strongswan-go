package monitoring

import (
	"github.com/sangyun-han/strongswan-go/vici"
	"fmt"
	"encoding/json"
)

func ExecuteTunnelMonitor() {
	client, err := vici.NewViciClientFromDefaultSocket()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// To be implemented
	sasList, err := client.ListSas("", "")
	result, err := json.Marshal(sasList)
	fmt.Println(sasList)
	fmt.Println()
	fmt.Println()
	fmt.Println(string(result))
	fmt.Println("End of test")
	for key, value := range sasList {
		fmt.Println("key=", key)
		fmt.Println("value=", value)
		fmt.Println()
		fmt.Println()
		fmt.Println(value["pdc"].Child_sas["pdc"].Bytes_in)
		fmt.Println()
		fmt.Println()
		fmt.Println(value["tims"])
		fmt.Println()
		fmt.Println()
		fmt.Println(value["pdc"].State)
		fmt.Println()
		fmt.Println()
	}
	fmt.Println("End of loop")
}

/*
- IPsec tunnel monitor as a daemon mode
- IPsec tunnel monitor as a passive mode
 */