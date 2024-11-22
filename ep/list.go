package ep

import (
	"log"

	"github.com/Peter-Bird/ws"
)

func List() {
	services := ws.ListServices()
	log.Printf("Services: %v", services)
}
