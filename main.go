package main

import (
    "github.com/sangyun-han/strongswan-go/vici"
    "encoding/json"
    "fmt"
)

type TunnelStats struct {
    TunnelId string `json:"tunnel_id"`
    SaId     string `json:"sa_id"`
    Version  string `json:"version"`
    Status   string `json:"status"`
    BytesIn  string `json:"bytes_in"`
    BytesOut string `json:"bytes_out"`
}

func main() {
    executeTunnelMonitor()
}

// passive mode monitor
func executeTunnelMonitor() {
    client, err := vici.NewViciClientFromDefaultSocket()
    if err != nil {
        panic(err)
    }
    defer client.Close()

    // To be implemented
    sasList, err := client.ListSas("", "")
    var statsList []TunnelStats

    for _, sa := range sasList {
        for k, v := range sa {
            stats := &TunnelStats{
                TunnelId: k,
                SaId: v.Uniqueid,
                Version: v.Version,
                Status: v.State,
                BytesIn: v.Child_sas[k].Bytes_in,
                BytesOut: v.Child_sas[k].Bytes_out,
            }
            statsList = append(statsList, *stats)
        }
    }
    output, _ := json.Marshal(statsList)
    fmt.Println(string(output))
}

/*
- IPsec tunnel monitor as a daemon mode
- IPsec tunnel monitor as a passive mode
 */