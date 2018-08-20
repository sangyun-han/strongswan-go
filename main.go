package main

import (
    "github.com/sangyun-han/strongswan-go/vici"
    "fmt"
    "encoding/json"
    "github.com/sangyun-han/strongswan-go/monitoring"
)

func main() {
    monitoring.ExecuteTunnelMonitor()

    vips := []string{"1", "2", "3"}

    testSa := vici.IkeSa{
        Uniqueid: "876",
        Version: "1",
        State: "ESTABLISHED",
        Local_host: "Local_host",
        Local_id: "Local_id",
        Remote_host: "Remote_host",
        Remote_id: "Remote_id",
        Initiator: "yes",
        Initiator_spi: "3fc123661bb52609",
        Responder_spi: "Responder_spi",
        Encr_alg: "Encr_alg",
        Encr_keysize: "Encr_keysize",
        Integ_alg: "Integ_alg",
        Integ_keysize: "Integ_keysize",
        Prf_alg: "Prf_alg",
        Dh_group: "Dh_group",
        Established: "Established",
        Rekey_time: "Rekey_time",
        Reauth_time: "Reauth_time",
        Remote_vips: vips,
    }

    fmt.Println(testSa)
    fmt.Println()
    fmt.Println()
    resultSa, _ := json.Marshal(testSa)
    fmt.Println(string(resultSa))


}
