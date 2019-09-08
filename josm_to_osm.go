package josm_to_osm_converter

import (
	"codes.musubu.co.jp/musubu/josm"
	"codes.musubu.co.jp/musubu/josm-to-osm-converter/entities"
	"io"
	"time"
)

func Convert(reader io.Reader) (*entities.OsmRoot, error) {
	osmRoot := &entities.OsmRoot{
		OsmBase: entities.OsmBase{
			Generator:   "josm-to-osm-converter",
			Copyright:   "OpenStreetMap and contributors",
			Attribution: "http://www.openstreetmap.org/copyright",
			License:     "http://opendatacommons.org/licenses/odbl/1-0/",
		},
		Bounds:    nil,
		Nodes:     nil,
		Ways:      nil,
		Relations: nil,
	}

	now := time.Now().UTC()
	uid := int64(1)
	user := "foo"

	josmRoot, err := josm.Decode(reader)
	if err != nil {
		return osmRoot, err
	}

	bounds := &entities.BoundingBox{
		MinLongitude: josmRoot.Bounds.Minlon,
		MinLatitude:  josmRoot.Bounds.Minlat,
		MaxLongitude: josmRoot.Bounds.Maxlon,
		MaxLatitude:  josmRoot.Bounds.Maxlat,
	}

	osmRoot.Bounds = bounds

	for _, josmNode := range josmRoot.Nodes {
		id := josmNode.ID
		if josmNode.ID < 0 {
			id = -1 * josmNode.ID
		}

		node := &entities.Node{
			Id:        id,
			Longitude: josmNode.Lon,
			Latitude:  josmNode.Lat,
			Version:   1,
			Timestamp: now,
			Changeset: 1,
			Uid:       uid,
			User:      user,
			Tags:      nil,
		}

		for _, josmTag := range josmNode.Tags {
			tag := &entities.NodeTag{
				NodeId: id,
				Key:    josmTag.Key,
				Value:  josmTag.Value,
			}

			node.Tags = append(node.Tags, tag)
		}

		osmRoot.Nodes = append(osmRoot.Nodes, node)
	}

	for _, josmWay := range josmRoot.Ways {
		id := josmWay.ID
		if josmWay.ID < 0 {
			id = -1 * josmWay.ID
		}

		way := &entities.Way{
			Id:        id,
			Version:   1,
			Timestamp: now,
			Changeset: 1,
			Uid:       uid,
			User:      user,
			Nodes:     nil,
			Tags:      nil,
		}

		for i, josmNode := range josmWay.Nds {
			nodeId := josmNode.ID
			if nodeId < 0 {
				nodeId = -1 * josmNode.ID
			}

			node := &entities.WayNode{
				WayId:      id,
				NodeId:     nodeId,
				SequenceId: int64(i + 1),
			}

			way.Nodes = append(way.Nodes, node)
		}

		for _, josmTag := range josmWay.Tags {
			tag := &entities.WayTag{
				WayId: id,
				Key:   josmTag.Key,
				Value: josmTag.Value,
			}

			way.Tags = append(way.Tags, tag)
		}

		osmRoot.Ways = append(osmRoot.Ways, way)
	}

	for _, josmRelation := range josmRoot.Relations {
		id := josmRelation.ID
		if josmRelation.ID < 0 {
			id = -1 * josmRelation.ID
		}

		relation := &entities.Relation{
			Id:        id,
			Version:   1,
			Timestamp: now,
			Changeset: 1,
			Uid:       uid,
			User:      user,
			Members:   nil,
			Tags:      nil,
		}

		for i, josmMember := range josmRelation.Members {
			memberId := josmMember.Ref
			if josmMember.Ref < 0 {
				memberId = -1 * josmMember.Ref
			}

			member := &entities.RelationMember{
				RelationId: id,
				MemberType: josmMember.Type,
				MemberId:   memberId,
				MemberRole: josmMember.Role,
				SequenceId: int64(i + 1),
			}

			relation.Members = append(relation.Members, member)
		}

		for _, josmTag := range josmRelation.Tags {
			tag := &entities.RelationTag{
				RelationId: id,
				Key:        josmTag.Key,
				Value:      josmTag.Value,
			}

			relation.Tags = append(relation.Tags, tag)
		}

		osmRoot.Relations = append(osmRoot.Relations, relation)
	}

	return osmRoot, nil
}
