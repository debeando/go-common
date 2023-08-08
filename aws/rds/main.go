package rds

import (
	"github.com/debeando/skale/retry"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
)

type Config struct {
	Client     *rds.RDS // AWS rds connection.
	Instance   string   // New instance (replica).
	Class      string   // New instance class.
	Region     string   // AWS region account.
	Primary    string   // Source instance (primary).
	Status     string   // New instance status.
	Endpoint   string   // New instance endpoint.
	Protection bool     // New instance is set to deletion protection.
	Port       int64    // New instance port of endpoint.
	Zone       string   // New instance Availability Zone.
}

func (c *Config) Init() (err error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(c.Region),
	})

	c.Client = rds.New(sess)

	return
}

func (c *Config) Create() (err error) {
	input := &rds.CreateDBInstanceReadReplicaInput{
		AutoMinorVersionUpgrade:         aws.Bool(false),
		AvailabilityZone:                aws.String(c.Zone),
		DBInstanceClass:                 aws.String(c.Class),
		DBInstanceIdentifier:            aws.String(c.Instance),
		DeletionProtection:              aws.Bool(false),
		EnableCustomerOwnedIp:           aws.Bool(false),
		EnableIAMDatabaseAuthentication: aws.Bool(false),
		EnablePerformanceInsights:       aws.Bool(false),
		MultiAZ:                         aws.Bool(false),
		PubliclyAccessible:              aws.Bool(false),
		SourceDBInstanceIdentifier:      aws.String(c.Primary),
	}

	_, err = c.Client.CreateDBInstanceReadReplica(input)
	return
}

func (c *Config) Delete() (err error) {
	input := &rds.DeleteDBInstanceInput{
		DBInstanceIdentifier: aws.String(c.Instance),
		SkipFinalSnapshot:    aws.Bool(true),
	}

	_, err = c.Client.DeleteDBInstance(input)
	return
}

func (c *Config) Describe() {
	c.Endpoint = ""
	c.Status = ""

	input := &rds.DescribeDBInstancesInput{
		DBInstanceIdentifier: aws.String(c.Instance),
	}

	result, err := c.Client.DescribeDBInstances(input)
	if err != nil {
		return
	}

	c.Status = *result.DBInstances[0].DBInstanceStatus
	c.Protection = *result.DBInstances[0].DeletionProtection

	if result.DBInstances[0].ReadReplicaSourceDBInstanceIdentifier != nil {
		c.Primary = *result.DBInstances[0].ReadReplicaSourceDBInstanceIdentifier
	}

	if result.DBInstances[0].Endpoint != nil {
		c.Endpoint = *result.DBInstances[0].Endpoint.Address
		c.Port = *result.DBInstances[0].Endpoint.Port
	}
}

func (c *Config) IsAvailable() bool {
	c.Describe()
	return len(c.Endpoint) > 0 && c.Status == "available"
}

func (c *Config) IsNotAvailable() bool {
	c.Describe()
	return len(c.Endpoint) == 0 && len(c.Status) == 0
}

func (c *Config) WaitUntilAvailable() error {
	return retry.Do(30, 60, func() bool {
		return c.IsAvailable()
	})
}

func (c *Config) WaitUntilUnavailable() error {
	return retry.Do(30, 60, func() bool {
		return c.IsNotAvailable()
	})
}
