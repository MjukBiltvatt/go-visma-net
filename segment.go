package vismanet

type Segment struct {
	SegmentID               IntValue    `json:"segmentId"`
	SegmentDescription      StringValue `json:"segmentDescription,omitempty"`
	SegmentValue            StringValue `json:"segmentValue"`
	SegmentValueDescription StringValue `json:"segmentValueDescription,omitempty"`
}
