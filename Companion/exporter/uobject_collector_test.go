package exporter_test

import (
    "github.com/AP-Hunt/FicsitRemoteMonitoringCompanion/Companion/exporter"
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
)

var _ = Describe("UObjectCollector", func() {
    var url string
    var sessionName = "default"
    var collector *exporter.UObjectCollector

    BeforeEach(func() {
        FRMServer.Reset()
        url = FRMServer.server.URL
        collector = exporter.NewUObjectCollector("/getUObjectCount")

        FRMServer.ReturnsUObjectData(exporter.UObjectDetails{
            UObjectCount:    1234,
            UObjectCapacity: 2000,
        })
    })

    AfterEach(func() {
        collector = nil
    })

    Describe("UObject metrics collection", func() {
        It("sets uobject_count with url/session labels", func() {
            collector.Collect(url, sessionName)

            val, err := gaugeValue(exporter.UObjectCount, url, sessionName)
            Expect(err).ToNot(HaveOccurred())
            Expect(val).To(Equal(float64(1234)))
        })

        It("sets uobject_capacity with url/session labels", func() {
            collector.Collect(url, sessionName)

            val, err := gaugeValue(exporter.UObjectCapacity, url, sessionName)
            Expect(err).ToNot(HaveOccurred())
            Expect(val).To(Equal(float64(2000)))
        })
    })
})