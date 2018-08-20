package monitoring

import "github.com/sangyun-han/strongswan-go/vici"

func executeTunnelMonitor() {
	client, err := vici.NewClientConnFromDefaultSocket()
	client2 := vici.NewClientFromDefaultSocket()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// To be implemented
	client.ListSas("df", "d")
	client2.ListSas("1", "")
}

/*
- IPsec tunnel monitor as a daemon mode
- IPsec tunnel monitor as a passive mode
 */