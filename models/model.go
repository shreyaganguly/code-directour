package models

type Model interface {
	BucketName() string
	ID() string
	Value() interface{}
}
