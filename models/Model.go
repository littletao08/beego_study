package models

type IModel interface {
	AsJSON() (map[string]interface{})
}
