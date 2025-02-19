package vismanet

type Segment struct {
	SegmentID               int    `json:"segmentId"`
	SegmentDescription      string `json:"segmentDescription,omitempty"`
	SegmentValue            string `json:"segmentValue"`
	SegmentValueDescription string `json:"segmentValueDescription,omitempty"`
}
