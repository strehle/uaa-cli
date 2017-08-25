package uaa

import (
	"encoding/json"
	"net/http"
)

type UaaInfo struct {
	App uaaApp `json:"app"`
	Links uaaLinks `json:"links"`
	Prompts uaaPrompts `json:"prompts"`
	ZoneName string `json:"zone_name"`
	EntityId string `json:"entityID"`
	CommitId string `json:"commit_id"`
	Timestamp string `json:"timestamp"`
}

type uaaApp struct {
	Version string `json:"version"`
}

type uaaLinks struct {
	ForgotPassword string `json:"passwd"`
	Uaa string `json:"uaa"`
	Registration string `json:"register"`
	Login string `json:"login"`
}

type uaaPrompts struct {
	Username []string `json:"username"`
	Password []string `json:"password"`
}

func Info(context UaaContext, client *http.Client) (UaaInfo, error) {
	bytes, err := UnauthenticatedGetter{}.GetBytes(context, "info", "")
	if err != nil {
		return UaaInfo{}, err
	}

	info := UaaInfo{}
	err = json.Unmarshal(bytes,&info)
	if err != nil {
		return UaaInfo{}, parseError("", bytes)
	}

	return info, err
}