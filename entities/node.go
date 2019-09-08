package entities

import "time"

type Node struct {
	Id        int64      `xml:"id,attr"`
	Longitude float64    `xml:"lon,attr"`
	Latitude  float64    `xml:"lat,attr"`
	Version   int64      `xml:"version,attr"`
	Timestamp time.Time  `xml:"timestamp,attr"`
	Changeset int64      `xml:"changeset,attr"`
	Uid       int64      `xml:"uid,attr"`
	User      string     `xml:"user,attr"`
	Tags      []*NodeTag `xml:"tag"`
}
