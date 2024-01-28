package rds

import (
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/rds/rdsiface"
)

type MockRDSAPI struct {
	rdsiface.RDSAPI
	DescribeDBInstancesOutput *rds.DescribeDBInstancesOutput
	Error                     error
}

func (m MockRDSAPI) DescribeDBInstances(*rds.DescribeDBInstancesInput) (*rds.DescribeDBInstancesOutput, error) {
	return m.DescribeDBInstancesOutput, m.Error
}
