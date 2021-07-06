package osmtransit

import (
	"context"
	"os"

	"github.com/paulmach/osm"
	"github.com/paulmach/osm/osmpbf"
)

// Parser is the struct that contains all the options and data to parse OSM PBF files.
type Parser struct {
	File *os.File

	Stops     bool
	StopAreas bool
}

func NoStop(p *Parser) {
	p.Stops = false
}

func NoStopArea(p *Parser) {
	p.StopAreas = false
}

// ParserOption is a function that sets a certain config on a Parser.
//
// This is part of the self referential functions design.
// See more: https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html
type ParserOption func(*Parser)

// NewParser creates a new custom parser.
//
// Use self referential functions design to configure.
// See more: https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html
func NewParser(file string, opts ...ParserOption) (*Parser, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	parser := Parser{
		File:      f,
		Stops:     true,
		StopAreas: true,
	}

	for _, opt := range opts {
		opt(&parser)
	}

	return &parser, nil
}

// scan scans the OSM pbf file for each elements
func (p *Parser) scan(objects chan osm.Object) {
	scanner := osmpbf.New(context.Background(), p.File, 5)
	defer scanner.Close()

	for scanner.Scan() {
		o := scanner.Object()
		objects <- o
	}

	scanErr := scanner.Err()
	if scanErr != nil {
		panic(scanErr)
	}

	close(objects)
}

// extractor selects OSM objects
func (p *Parser) extractor(transit chan TransitData, objects chan osm.Object) {

	for o := range objects {
		if o.ObjectID().Type() == osm.TypeNode {
			obj := o.(*osm.Node)
			if p.Stops && isStop(obj) {
				transit <- StopFromNode(obj)
			} else if p.StopAreas && isStopArea(obj) {
				transit <- StopAreaFromNode(obj)
			}
		}
	}

	close(transit)
}

// Extract parses OSM data and extracts transit data
func (p *Parser) Extract(transit chan TransitData) {
	objects := make(chan osm.Object, 100)

	go p.scan(objects)

	go p.extractor(transit, objects)
}

// Parse returns transit data in a slice from an OSM pbf file
func (p *Parser) TransitData() []TransitData {
	result := []TransitData{}

	transit := make(chan TransitData, 100)
	p.Extract(transit)

	for t := range transit {
		result = append(result, t)
	}

	return result
}
