package entities

type OsmRoot struct {
	OsmBase
	Bounds    *BoundingBox `xml:"bounds"`
	Nodes     []*Node      `xml:"node"`
	Ways      []*Way       `xml:"way"`
	Relations []*Relation  `xml:"relation"`
}
