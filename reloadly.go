package reloadly

type Client struct {
	clientID string
	clientSecret string
	sandBox bool
}

func NewClient(clientId, clientSecret string, sandbox bool)(*Client, error){
	if clientId == ""{
		return nil, &Error{Message: "INVALID_CREDENTIALS"}
	}
	return &Client{
		clientID:     clientId,
		clientSecret: clientSecret,
		sandBox:      sandbox,
	}, nil
}
