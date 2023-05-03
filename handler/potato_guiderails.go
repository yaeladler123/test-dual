package handler

import (
	"github.com/calculi-corp/api/example/go/potato"
	handler "github.com/calculi-corp/grpc-handler"
	healthchecks "github.com/calculi-corp/grpc-handler/pb"
	"github.com/calculi-corp/log"
	"time"
)

// This file contains required implementations but very specific to GuideRails.
// That's why this code is not located in the `potato.go` file.

func (ph *PotatoHandler) initialize() error {
	log.Debugf("Initialize ldapMap...")
	beginTs := time.Now()

	interval := float64(time.Since(beginTs)) / float64(time.Millisecond)
	ph.metrics.NewGauge("inittime", "ms", true).Set(interval)
	log.Debugf("%s initialize() took: %f ms", handlerName, interval)
	return nil
}

// Description is used to register the service in the service registry automatically
func (ph *PotatoHandler) Description() *handler.ServiceDesc {
	return &handler.ServiceDesc{Name: potato.PotatoService_ServiceDesc.ServiceName, ProtoDesc: potato.PotatoService_ServiceDesc}
}

// MetricMap returns the metrics for this service
func (ph *PotatoHandler) MetricMap() *handler.Map {
	return ph.metrics
}

// Healthy Return health status
func (ph *PotatoHandler) Healthy() error {
	err := ph.initialize()
	log.CheckErrorf(err, "could not initialize %s handler", handlerName)
	return nil
}

// Dependencies returns a list of services which are dependents of entitlement-service
func (ph *PotatoHandler) Dependencies() []string {
	return dependencies
}

// HealthDependency checks the service's own health and all of its dependencies healths based on the received depth
func (ph *PotatoHandler) HealthDependency(depth int32, task string) []*healthchecks.ServiceHealthResponse {
	return handler.HealthCheck.HealthDependency(ph, depth, task, ph.Dependencies())
}
