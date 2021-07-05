# coopgo/osmtransit

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/coopgo/gtfs.svg)](https://github.com/coopgo/osmtransit)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/coopgo/osmtransit)

Package `coopgo/osmtransit` is a Golang library to extract transit data from [OpenStreetMap](https://www.openstreetmap.org) PBF files.

The main features of this package are :

- Scan OSM files in the osm.pbf format (example : https://download.geofabrik.de/europe/france-latest.osm.pbf)
- Extract transit data (stops, routes, ...)
- Uses GTFS data structures from the [coopgo/gtfs](https://github.com/coopgo/gtfs) library to return transit data (ability to convert OSM transit data to partial GTFS files/CSV files using this library -still under development there-)

Supported transit objects (and TODO list) :

- [X] Stops
- [ ] Stop Areas
- ...

## Usage

Check examples in the [examples](examples) folder.

## Project Status

This library is still in development. 

## Contributing


We welcome any contributions following theses guidelines :
- Write simple, clear and maintainable code and avoid technical debt. 
- Leave the code cleaner than when you started.
- Refactoring existing code for better performance, better readability or better testing wins over creating a new feature.

If you want to contribute, you can fork the repository and create a pull request.

## Bug report

For reporting a bug, you can open an issue using the **Bug Report** template. Try to write a bug report that is easy to understand and explain how to reproduce the bug. 
Do not duplicate an existing issue and keep each issue specific to an individual bug.

## License

`coopgo/osmtransit` is under the Apache 2.0 license. Please refer to the [LICENSE](LICENSE) file for details.
