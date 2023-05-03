package internal

import (
	"fmt"
	server "github.com/calculi-corp/grpc-server"
	"github.com/calculi-corp/hippo-service/handler"
	"github.com/calculi-corp/log"
	"strings"
)

const (
	// ServiceName is the name of the entitlement service
	ServiceName = "PotatoService"
	ServiceDesc = "Potato Example Service"
)

type Configuration struct {
	LogLevel              string
	PotatoServiceEndpoint string
}

func Start(c *Configuration) error {
	configureLogging(c.LogLevel)

	s, err := server.NewServer()
	if log.CheckErrorf(err, "Unable to instantiate s") {
		return fmt.Errorf("error instanciating grpc server: %w", err)
	}
	defer s.Stop()

	h := handler.NewPotatoHandler()
	err = s.AddHandler(h)
	if log.CheckErrorf(err, "Failure adding handler") {
		return fmt.Errorf("error adding handler: %w", err)
	}

	s.Start()
	s.WaitForExit()

	return nil
}

// configureLogging set up the Logs configuration depending on the Configuration
func configureLogging(level string) {
	log.Initialize(ServiceName)
	log.SetLoggingLevel(level)
	log.Info(fmt.Sprintf("Logs set to %s", strings.ToUpper(level)))
}
