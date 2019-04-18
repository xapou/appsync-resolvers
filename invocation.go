package resolvers

import "encoding/json"

type context struct {
	Arguments json.RawMessage `json:"arguments"`
	Source    json.RawMessage `json:"source"`
	Identity  json.RawMessage `json:"identity"`
}

type invocation struct {
	Resolve string  `json:"resolve"`
	Context context `json:"context"`
}
