package acttype

import "encoding/json"

type ActionType uint8

const (
	Http ActionType = iota
	Docker
)

func (a ActionType) String() string {
	switch a {
	case Docker:
		return "docker"
	}
	return "http"
}

func ParseActionType(name string) ActionType {
	switch name {
	case "docker":
		return Docker
	}

	return Http
}

func (a ActionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.String())
}

func (a *ActionType) UnmarshalJSON(data []byte) (err error) {
	var actionType string
	if err := json.Unmarshal(data, &actionType); err != nil {
		return nil
	}
	*a = ParseActionType(actionType)
	return nil
}
