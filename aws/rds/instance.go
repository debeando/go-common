package rds

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
)

type Instance struct {
	AvailabilityZone    string
	Class               string
	Created             *time.Time
	DeletionProtection  bool
	Endpoint            string
	Engine              string
	Identifier          string
	MultiAZ             bool
	ParameterGroup      ParameterGroup
	PerformanceInsights bool
	Port                int64
	Public              bool
	Status              string
	Tags                Tags
	Username            string
	Version             string
}

type Instances []Instance
type Tags []Tag

type ParameterGroup struct {
	Name   string
	Status string
}

type Tag struct {
	Key   string
	Value string
}

func (i *Instance) New(instance *rds.DBInstance) *Instance {
	tags := Tags{}

	*i = Instance{
		AvailabilityZone:    aws.StringValue(instance.AvailabilityZone),
		Class:               aws.StringValue(instance.DBInstanceClass),
		Created:             instance.InstanceCreateTime,
		DeletionProtection:  aws.BoolValue(instance.DeletionProtection),
		Endpoint:            aws.StringValue(instance.Endpoint.Address),
		Engine:              aws.StringValue(instance.Engine),
		Identifier:          aws.StringValue(instance.DBInstanceIdentifier),
		MultiAZ:             aws.BoolValue(instance.MultiAZ),
		PerformanceInsights: aws.BoolValue(instance.PerformanceInsightsEnabled),
		Port:                aws.Int64Value(instance.Endpoint.Port),
		Public:              aws.BoolValue(instance.PubliclyAccessible),
		Status:              aws.StringValue(instance.DBInstanceStatus),
		Username:            aws.StringValue(instance.MasterUsername),
		Version:             aws.StringValue(instance.EngineVersion),
		ParameterGroup: ParameterGroup{
			Name:   aws.StringValue(instance.DBParameterGroups[0].DBParameterGroupName),
			Status: aws.StringValue(instance.DBParameterGroups[0].ParameterApplyStatus),
		},
		Tags: *tags.New(instance.TagList),
	}

	return i
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