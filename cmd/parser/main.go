package main

import (
	"BeTest-AlexanderBergasov/internal/service"
	"BeTest-AlexanderBergasov/internal/service/parser"
	"flag"
	"log"
)

var dataFilesPath = flag.String("path", "/", "Data files path")

func main() {
	log.Println("Starting application")
	flag.Parse()
	*dataFilesPath = "/home/alejandro/go/src/BeTest-AlexanderBergasov"

	app := service.NewOrchestra()
	app.RegisterDecoder(".csv", parser.NewCSVDecoder())
	app.RegisterDecoder(".prn", parser.NewPRNDecoder())

	log.Println("parsing in dir: ", *dataFilesPath)
	if err := app.Run(*dataFilesPath); err != nil {
		log.Fatalf("failed to parse source data: %v", err)
	}
}
