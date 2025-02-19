package vismanet

// Segment is a segment as represented in the body of responses from Visma.net
type Segment struct {
	SegmentID               int    `json:"segmentId"`
	SegmentDescription      string `json:"segmentDescription"`
	SegmentValue            string `json:"segmentValue"`
	SegmentValueDescription string `json:"segmentValueDescription"`
}

// RequestBodySegment is a segment as represented in the body of requests to Visma.net
type RequestBodySegment struct {
	SegmentID    int    `json:"segmentId"`
	SegmentValue string `json:"segmentValue"`
}
