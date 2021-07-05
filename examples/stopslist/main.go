package main

import (
	"fmt"
	"log"
	"time"

	"github.com/coopgo/osmtransit"
)

func main() {
	start := time.Now()
	parser, err := osmtransit.NewParser("/home/adelcasse/Downloads/france-latest.osm.pbf")
	//parser, err := osmtransit.NewParser("/home/adelcasse/Downloads/provence-alpes-cote-d-azur-latest.osm.pbf")
	if err != nil {
		panic(err)
	}

	transit := parser.TransitData()

	fmt.Println("Nb of transit objects in the OSM file : ", len(transit))

	elapsed := time.Since(start)
	log.Printf("Extracting transit data took %s", elapsed)
}
