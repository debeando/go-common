package rds

import (
	"encoding/json"
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
	ReplicaIdentifier   string
	Status              string
	Tags                Tags
	Username            string
	Version             string
}

type Instances []Instance

func (i *Instance) New(instance *rds.DBInstance) *Instance {
	tags := Tags{}

	*i = Instance{
		AvailabilityZone:    aws.StringValue(instance.AvailabilityZone),
		Class:               aws.StringValue(instance.DBInstanceClass),
		Created:             instance.InstanceCreateTime,
		DeletionProtection:  aws.BoolValue(instance.DeletionProtection),
		Engine:              aws.StringValue(instance.Engine),
		Identifier:          aws.StringValue(instance.DBInstanceIdentifier),
		MultiAZ:             aws.BoolValue(instance.MultiAZ),
		PerformanceInsights: aws.BoolValue(instance.PerformanceInsightsEnabled),
		Public:              aws.BoolValue(instance.PubliclyAccessible),
		ReplicaIdentifier:   aws.StringValue(instance.ReadReplicaSourceDBInstanceIdentifier),
		Status:              aws.StringValue(instance.DBInstanceStatus),
		Username:            aws.StringValue(instance.MasterUsername),
		Version:             aws.StringValue(instance.EngineVersion),
		ParameterGroup: ParameterGroup{
			Name:   aws.StringValue(instance.DBParameterGroups[0].DBParameterGroupName),
			Status: aws.StringValue(instance.DBParameterGroups[0].ParameterApplyStatus),
		},
		Tags: *tags.New(instance.TagList),
	}

	if instance.Endpoint != nil {
		i.Endpoint = aws.StringValue(instance.Endpoint.Address)
		i.Port = aws.Int64Value(instance.Endpoint.Port)
	}

	return i
}

func (i *Instance) JSON() (r map[string]interface{}) {
	t, _ := json.Marshal(i)
	json.Unmarshal(t, &r)

	return r
}

func (i *Instance) Exist() bool {
	return len(i.Endpoint) > 0 && len(i.Status) > 0
}

func (i *Instance) IsAvailable() bool {
	return len(i.Endpoint) > 0 && i.Status == "available"
}

func (i *Instance) IsPrimary() bool {
	return i.IsAvailable() && len(i.ReplicaIdentifier) == 0
}

func (i *Instance) IsReplica() bool {
	return i.IsAvailable() && len(i.ReplicaIdentifier) > 0
}

func (i *Instance) IsNotAvailable() bool {
	return len(i.Endpoint) == 0 && len(i.Status) == 0
}
