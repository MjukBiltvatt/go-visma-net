package vismanet

// Attachment is an attachment as represented in responses from Visma.net
type Attachment struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Revision int    `json:"revision"`
}

type RequestAttachment struct {
	SendToAutoInvoice BoolValue `json:"sendToAutoInvoice"`
}

// =========================================================
// ========================== PUT ==========================
// =========================================================

// newPutAttachmentV1Request creates a new PutAttachmentV1Request
func newPutAttachmentV1Request(c *Client) PutAttachmentV1Request {
	return PutAttachmentV1Request{
		Client: c,
		Method: "PUT",
		Path:   "controller/api/v1/attachment/{{.attachment_id}}",
		Body:   JSONRequestBody{},
	}
}

// PutAttachmentV1Request represents a request to update an attachment
type PutAttachmentV1Request Request

// SetPathParams sets the path parameters of the request
func (r *PutAttachmentV1Request) SetPathParams(params PutAttachmentV1PathParams) {
	r.pathParams = params
}

// SetBody sets the body of the request
func (r *PutAttachmentV1Request) SetBody(body RequestAttachment) {
	r.Body = JSONRequestBody{body}
}

// Do performs the request and returns the response
func (r *PutAttachmentV1Request) Do() (PutAttachmentV1Response, error) {
	resp, err := r.Client.Do((*Request)(r), nil)
	return PutAttachmentV1Response{Response{resp}}, err
}

// PutAttachmentV1PathParams represents the path parameters of the PutAttachmentV1Request
type PutAttachmentV1PathParams struct {
	AttachmentID string `schema:"attachment_id"`
}

// PutAttachmentV1Response represents the response of the PutAttachmentV1Request
type PutAttachmentV1Response struct {
	Response
}
