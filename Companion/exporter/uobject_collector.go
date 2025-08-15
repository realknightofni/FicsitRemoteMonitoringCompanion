package exporter

import (
    "log"
)

type UObjectDetails struct {
    UObjectCount    float64 `json:"UObjectCount"`
    UObjectCapacity float64 `json:"UObjectCapacity"`
}

type UObjectCollector struct {
    endpoint string
}

func NewUObjectCollector(endpoint string) *UObjectCollector {
    return &UObjectCollector{
        endpoint: endpoint,
    }
}

func (c *UObjectCollector) Collect(frmAddress string, sessionName string) {
    details := UObjectDetails{}
    if err := retrieveData(frmAddress+c.endpoint, &details); err != nil {
        log.Printf("error reading UObject stats from FRM: %s\n", err)
        return
    }

    UObjectCount.WithLabelValues(frmAddress, sessionName).Set(details.UObjectCount)
    UObjectCapacity.WithLabelValues(frmAddress, sessionName).Set(details.UObjectCapacity)
}

func (c *UObjectCollector) DropCache() {}