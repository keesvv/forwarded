package cli

import (
	"log"
	"os"
)

func PrintUsage() {
	log.Fatalf("%s <host> <lport>:<rport>", os.Args[0])
}
