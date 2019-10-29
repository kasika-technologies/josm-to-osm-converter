package entities

import "time"

type Relation struct {
	Id        int64             `xml:"id,attr"`
	Version   int64             `xml:"version,attr"`
	Timestamp time.Time         `xml:"timestamp,attr"`
	Changeset int64             `xml:"changeset,attr"`
	Uid       int64             `xml:"uid,attr"`
	User      string            `xml:"user,attr"`
	Members   []*RelationMember `xml:"member"`
	Tags      []*RelationTag    `xml:"tag"`
}
