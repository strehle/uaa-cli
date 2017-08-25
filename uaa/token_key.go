package uaa

import (
	"encoding/json"
)

type JWK struct {
	Kty string
	E string
	Use string
	Kid string
	Alg string
	Value string
	N string
}

func TokenKey(context UaaContext) (JWK, error) {
	body, err := AuthenticatedGetter{}.GetBytes(context, "token_key", "")
	if err != nil {
		return JWK{}, err
	}

	key := JWK{}
	err = json.Unmarshal(body,&key)
	if err != nil {
		return JWK{}, parseError("/token_key", body)
	}

	return key, nil
}