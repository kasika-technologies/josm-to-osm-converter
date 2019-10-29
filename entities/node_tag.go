package entities

type NodeTag struct {
	NodeId int64  `xml:"-"`
	Key    string `xml:"k,attr"`
	Value  string `xml:"v,attr"`
}
