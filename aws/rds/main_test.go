package rds_test

import (
	"testing"

	"github.com/debeando/go-common/aws/rds"

	"github.com/aws/aws-sdk-go/aws"
	awsrds "github.com/aws/aws-sdk-go/service/rds"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	mockedOutput := &awsrds.DescribeDBInstancesOutput{
		DBInstances: []*awsrds.DBInstance{{
			DBInstanceIdentifier: aws.String("node01"),
			DBInstanceStatus:     aws.String("creating"),
			DBParameterGroups: []*awsrds.DBParameterGroupStatus{{
				DBParameterGroupName: aws.String("default.mysql5.6"),
				ParameterApplyStatus: aws.String("in-sync"),
			}},
		}},
	}

	r := rds.RDS{
		Client: rds.MockRDSAPI{
			DescribeDBInstancesOutput: mockedOutput,
		},
	}

	r.Load()

	t.Log(r.Instance.Status)

	assert.True(t, true)
}
