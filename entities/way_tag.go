package entities

type WayTag struct {
	WayId int64  `xml:"-"`
	Key   string `xml:"k,attr"`
	Value string `xml:"v,attr"`
}
