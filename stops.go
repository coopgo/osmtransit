package osmtransit

import (
	"github.com/coopgo/gtfs"
	"github.com/paulmach/osm"
)

func isStop(n *osm.Node) bool {
	if v := n.Tags.Find("public_transport"); v == "platform" || v == "stop_position" {
		return true
	}
	if v := n.Tags.Find("highway"); v == "bus_stop" {
		return true
	}
	if v := n.Tags.Find("railway"); v == "tram_stop" {
		return true
	}

	return false
}

type Stop gtfs.StopSerializable

func (s Stop) Type() TransitDataType {
	return StopType
}

func StopFromNode(n *osm.Node) Stop {
	return Stop{
		Id:   n.ID.FeatureID().String(),
		Name: n.Tags.Find("name"),
		Code: n.Tags.Find("ref"),
		Lat:  n.Lat,
		Long: n.Lon,
	}
}
