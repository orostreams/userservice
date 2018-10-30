package models

//Interface an interface that must be satisified by all models
type Interface interface {
	ToMap() map[string]interface{}
}
