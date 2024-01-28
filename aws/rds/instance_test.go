package rds_test

import (
	"testing"

	"github.com/debeando/go-common/aws/rds"

	"github.com/stretchr/testify/assert"
)

func TestInstanceExist(t *testing.T) {
	i := rds.Instance{}

	i.Endpoint = "node01.abc.x-y-1.rds.amazonaws.com"
	i.Status = "creating"
	assert.True(t, i.Exist())
}

func TestInstanceIsAvailable(t *testing.T) {
	i := rds.Instance{}

	i.Endpoint = "node01.abc.x-y-1.rds.amazonaws.com"
	i.Status = "available"
	assert.True(t, i.IsAvailable())

	i.Status = "modifying"
	assert.False(t, i.IsAvailable())
}

func TestInstanceIsNotAvailable(t *testing.T) {
	i := rds.Instance{}

	i.Endpoint = ""
	i.Status = ""
	assert.True(t, i.IsNotAvailable())

	i.Endpoint = "node01.abc.x-y-1.rds.amazonaws.com"
	i.Status = "available"
	assert.False(t, i.IsNotAvailable())
}

func TestInstanceIsPrimary(t *testing.T) {
	i := rds.Instance{}

	i.Endpoint = "node01.abc.x-y-1.rds.amazonaws.com"
	i.Status = "available"
	assert.True(t, i.IsPrimary())

	i.ReplicaIdentifier = "mysql02"
	assert.False(t, i.IsPrimary())
}

func TestInstanceIsReplica(t *testing.T) {
	i := rds.Instance{}

	i.Endpoint = "node01.abc.x-y-1.rds.amazonaws.com"
	i.Status = "available"
	i.ReplicaIdentifier = "mysql02"
	assert.True(t, i.IsReplica())

	i.ReplicaIdentifier = ""
	assert.False(t, i.IsReplica())
}
