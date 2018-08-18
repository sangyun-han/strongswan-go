package monitoring

import "github.com/sangyun-han/strongswan-go/vici"

func startTunnelMonitor() {
	client, err := vici.NewClientConnFromDefaultSocket()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// To be implemented
}

/*
- IPsec tunnel monitor as a daemon mode
- IPsec tunnel monitor as a passive mode
 */