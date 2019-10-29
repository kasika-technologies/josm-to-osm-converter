package entities

type BoundingBox struct {
	MinLongitude float64 `xml:"minlon,attr"`
	MinLatitude  float64 `xml:"minlat,attr"`
	MaxLongitude float64 `xml:"maxlon,attr"`
	MaxLatitude  float64 `xml:"maxlat,attr"`
}
