package osmtransit

type TransitDataType int

const (
	StopType TransitDataType = iota
	StopAreaType
)

type TransitData interface {
	Type() TransitDataType
}
