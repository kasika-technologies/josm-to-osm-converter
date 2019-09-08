package entities

type RelationTag struct {
	RelationId int64  `xml:"-"`
	Key        string `xml:"k,attr"`
	Value      string `xml:"v,attr"`
}
