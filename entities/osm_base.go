package entities

import "encoding/xml"

type OsmBase struct {
	XMLName     xml.Name `xml:"osm"`
	Generator   string   `xml:"generator,attr"`
	Copyright   string   `xml:"copyright,attr"`
	Attribution string   `xml:"attribution,attr"`
	License     string   `xml:"license,attr"`
}
