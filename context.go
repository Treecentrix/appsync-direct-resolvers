package resolvers

import "encoding/json"

type info struct {
	FieldName string `json:"fieldName"`
}

type Identity struct {
	// common
	SourceIp []string `json:"sourceIp"`
	Username string   `json:"username"`

	// AWS_IAM authorization
	AccountId                   string `json:"accountId"`
	CognitoIdentityPoolId       string `json:"cognitoIdentityPoolId"`
	CognitoIdentityId           string `json:"cognitoIdentityId"`
	UserArn                     string `json:"userArn"`
	CognitoIdentityAuthType     string `json:"cognitoIdentityAuthType"`
	CognitoIdentityAuthProvider string `json:"cognitoIdentityAuthProvider"`

	// AMAZON_COGNITO_USER_POOLS authorization
	Sub                 string          `json:"sub"`
	Issuer              string          `json:"issuer"`
	Claims              json.RawMessage `json:"claims"`
	DefaultAuthStrategy string          `json:"defaultAuthStrategy"`
}

type request struct {
	Headers json.RawMessage `json:"headers"`
}

type context struct {
	Arguments json.RawMessage `json:"arguments"`
	Info      info            `json:"info"`
	Identity  *Identity       `json:"Identity"`
	Request   request         `json:"request"`
}

func (ctx context) resolver() string {
	return ctx.Info.FieldName
}

func (ctx context) payload() json.RawMessage {
	return ctx.Arguments
}

func (ctx context) headers() json.RawMessage {
	return ctx.Request.Headers
}

func (ctx context) identity() *Identity {
	return ctx.Identity
}
