package entities

type WayNode struct {
	WayId      int64 `xml:"-"`
	NodeId     int64 `xml:"ref,attr"`
	SequenceId int64 `xml:"-"`
}
