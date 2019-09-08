package entities

type RelationMember struct {
	RelationId int64  `xml:"-"`
	MemberType string `xml:"type,attr"`
	MemberId   int64  `xml:"ref,attr"`
	MemberRole string `xml:"role,attr"`
	SequenceId int64  `xml:"-"`
}
