package api

type addThingReq struct {
	token       string
	ExternalID  string `json:"external_id"`
	ExternalKey string `json:"external_key"`
}

func (req addThingReq) validate() error {
	if req.ExternalID == "" || req.ExternalKey == "" {
		return errUnauthorized
	}
	return nil
}
