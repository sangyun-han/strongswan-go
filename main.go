package main

import (
    "github.com/sangyun-han/strongswan-go/vici"
    "os"
    "fmt"
    "encoding/json"
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
    if len(os.Args) < 2 {
        statsList := executeTunnelMonitor()
        if statsList == nil {
            fmt.Println("null")
        }
        output, _ := json.Marshal(statsList)
        fmt.Println(string(output))
    } else {
        runMonitoringDaemon()
    }
}

// passive mode monitor
func executeTunnelMonitor() []TunnelStats {
    client, err := vici.NewViciClientFromDefaultSocket()
    if err != nil {
        panic(err)
    }
    defer client.Close()

    sasList, err := client.ListSas("", "")
    var statsList []TunnelStats

    if len(sasList) == 0 {
        return nil
    }

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

    return statsList
}

// daemon mode monitor
func runMonitoringDaemon() {
    client, err := vici.NewViciClientFromDefaultSocket()
    if err != nil {
        panic(err)
    }
    defer client.Close()

    go func() {
        sasList, err := client.ListSas("", "")
        if err != nil {
            fmt.Println(err)
        }
        if len(sasList) == 0 {
            fmt.Println("null")
        }

        var statsList []TunnelStats

        // TODO add waitgroup(lazy withdraw)
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
    }()
}