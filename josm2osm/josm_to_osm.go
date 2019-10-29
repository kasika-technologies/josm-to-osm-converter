package josm2osm

import (
	"fmt"
	"github.com/kasika-technologies/josm"
	"github.com/kasika-technologies/josm-to-osm-converter/entities"
	quadTile "github.com/kasika-technologies/quad_tile"
	"io"
	"strings"
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

func ConvertToSql(root *entities.OsmRoot) (string, error) {
	var query string
	var err error

	var queries []string

	now := time.Now().UTC().Format(time.RFC3339)

	for _, node := range root.Nodes {
		email := fmt.Sprintf("%d@example.com", node.Uid)
		passCrypt := "sample"

		changesetId := 1

		lat := int64(node.Latitude * 10000000)
		lon := int64(node.Longitude * 10000000)
		visible := "true"
		tile := quadTile.TileForPoint(node.Longitude, node.Latitude)

		qUser := fmt.Sprintf("insert into users (id, email, pass_crypt, creation_time, display_name, description) values (%d, '%s', '%s', '%s', '%s', '%s') on conflict on constraint users_pkey do update set display_name='%s';", node.Uid, email, passCrypt, now, node.User, "", node.User)
		queries = append(queries, qUser)

		qChangeset := fmt.Sprintf("insert into changesets (id, user_id, created_at, closed_at) values (%d, %d, '%s', '%s') on conflict on constraint changesets_pkey do update set user_id=%d;", changesetId, node.Uid, now, now, node.Uid)
		queries = append(queries, qChangeset)

		qNode := fmt.Sprintf("insert into current_nodes (id, latitude, longitude, changeset_id, visible, timestamp, tile, version) values (%d, %d, %d, %d, %s, '%s', %d, %d);", node.Id, lat, lon, node.Changeset, visible, node.Timestamp.Format(time.RFC3339), tile, node.Version)
		queries = append(queries, qNode)

		for _, tag := range node.Tags {
			qNodeTag := fmt.Sprintf("insert into current_node_tags (node_id, k, v) values (%d, '%s', '%s');", tag.NodeId, tag.Key, tag.Value)
			queries = append(queries, qNodeTag)
		}
	}

	for _, way := range root.Ways {
		email := fmt.Sprintf("%d@example.com", way.Uid)
		passCrypt := "sample"

		changesetId := 1

		qUser := fmt.Sprintf("insert into users (id, email, pass_crypt, creation_time, display_name, description) values (%d, '%s', '%s', '%s', '%s', '%s') on conflict on constraint users_pkey do update set display_name='%s';", way.Uid, email, passCrypt, now, way.User, "", way.User)
		queries = append(queries, qUser)

		qChangeset := fmt.Sprintf("insert into changesets (id, user_id, created_at, closed_at) values (%d, %d, '%s', '%s') on conflict on constraint changesets_pkey do update set user_id=%d;", changesetId, way.Uid, now, now, way.Uid)
		queries = append(queries, qChangeset)

		qWay := fmt.Sprintf("insert into current_ways (id, changeset_id, timestamp, visible, version) values (%d, %d, '%s', true, %d);", way.Id, changesetId, way.Timestamp.Format(time.RFC3339), way.Version)
		queries = append(queries, qWay)

		for _, tag := range way.Tags {
			qWayTag := fmt.Sprintf("insert into current_way_tags (way_id, k, v) values (%d, '%s', '%s');", tag.WayId, tag.Key, tag.Value)
			queries = append(queries, qWayTag)
		}

		for i, node := range way.Nodes {
			qWayNode := fmt.Sprintf("insert into current_way_nodes (way_id, node_id, sequence_id) values (%d, %d, %d);", node.WayId, node.NodeId, int64(i+1))
			queries = append(queries, qWayNode)
		}
	}

	for _, relation := range root.Relations {
		email := fmt.Sprintf("%d@example.com", relation.Uid)
		passCrypt := "sample"

		changesetId := 1

		qUser := fmt.Sprintf("insert into users (id, email, pass_crypt, creation_time, display_name, description) values (%d, '%s', '%s', '%s', '%s', '%s') on conflict on constraint users_pkey do update set display_name='%s';", relation.Uid, email, passCrypt, now, relation.User, "", relation.User)
		queries = append(queries, qUser)

		qChangeset := fmt.Sprintf("insert into changesets (id, user_id, created_at, closed_at) values (%d, %d, '%s', '%s') on conflict on constraint changesets_pkey do update set user_id=%d;", changesetId, relation.Uid, now, now, relation.Uid)
		queries = append(queries, qChangeset)

		qRelation := fmt.Sprintf("insert into current_relations (id, changeset_id, timestamp, visible, version) values (%d, %d, '%s', true, %d);", relation.Id, changesetId, relation.Timestamp.Format(time.RFC3339), relation.Version)
		queries = append(queries, qRelation)

		for _, tag := range relation.Tags {
			qRelationTag := fmt.Sprintf("insert into current_relation_tags (relation_id, k, v) values (%d, '%s', '%s');", tag.RelationId, tag.Key, tag.Value)
			queries = append(queries, qRelationTag)
		}

		for i, member := range relation.Members {
			qMember := fmt.Sprintf("insert into current_relation_members (relation_id, member_type, member_id, member_role, sequence_id) values (%d, '%s', %d, '%s', %d);", member.RelationId, strings.Title(member.MemberType), member.MemberId, member.MemberRole, int64(i+1))
			queries = append(queries, qMember)
		}
	}

	if err != nil {
		return query, err
	}

	query = strings.Join(queries, "\n")

	return query, nil
}
