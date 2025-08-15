package exporter

import "github.com/prometheus/client_golang/prometheus"

// No extra labels; RegisterNewGaugeVec adds url and session_name automatically.
var (
    UObjectCount = RegisterNewGaugeVec(prometheus.GaugeOpts{
        Name: "uobjects_loaded",
        Help: "Number of UObjects currently loaded",
    }, []string{})

    UObjectCapacity = RegisterNewGaugeVec(prometheus.GaugeOpts{
        Name: "uobject_capacity",
        Help: "UObject capacity limit",
    }, []string{})
)