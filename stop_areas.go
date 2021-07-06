package osmtransit

import "github.com/paulmach/osm"

func isStopArea(n *osm.Node) bool {
	if v := n.Tags.Find("public_transport"); v == "stop_area" {
		return true
	}

	return false
}

type StopArea struct {
	Id   string
	Name string
	Code string
	Lat  float64
	Long float64
}

func (s StopArea) Type() TransitDataType {
	return StopAreaType
}

func StopAreaFromNode(n *osm.Node) StopArea {
	return StopArea{
		Id:   n.ID.FeatureID().String(),
		Name: n.Tags.Find("name"),
		Code: n.Tags.Find("ref"),
		Lat:  n.Lat,
		Long: n.Lon,
	}
}
