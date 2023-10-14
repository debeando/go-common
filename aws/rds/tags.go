package rds

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
)

type Tags []Tag

type Tag struct {
	Key   string
	Value string
}

func (t *Tags) New(input []*rds.Tag) *Tags {
	for _, tag := range input {
		*t = append(*t, Tag{
			Key:   aws.StringValue(tag.Key),
			Value: aws.StringValue(tag.Value),
		})
	}

	return t
}
