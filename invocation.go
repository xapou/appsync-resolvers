package resolvers

import "encoding/json"

// Identity contains the identity context information (cognito userpool or IAM)
type Identity struct {
	AccountID             string            `json:"accountId"`
	CognitoIdentityPoolID string            `json:"cognitoIdentityPoolId"`
	CognitoIdentityID     string            `json:"cognitoIdentityId"`
	SourceIP              []string          `json:"sourceIp"`
	Username              string            `json:"username"`
	UserArn               string            `json:"userArn"`
	Sub                   string            `json:"sub"`
	Issuer                string            `json:"issuer"`
	Claims                map[string]string `json:"claims"`
	DefaultAuthStrategy   string            `json:"defaultAuthStrategy"`
}

type context struct {
	Arguments json.RawMessage `json:"arguments"`
	Source    json.RawMessage `json:"source"`
	Identity  Identity        `json:"identity"`
}

type invocation struct {
	Resolve string  `json:"resolve"`
	Context context `json:"context"`
}

func (in invocation) isRoot() bool {
	return in.Context.Source == nil || string(in.Context.Source) == "null"
}

func (in invocation) payload() json.RawMessage {
	if in.isRoot() {
		return in.Context.Arguments
	}

	return in.Context.Source
}
