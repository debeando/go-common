package rds

import (
	"github.com/debeando/go-common/retry"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
)

type RDS struct {
	Client     *rds.RDS `json:"-"`          // AWS RDS connection.
	Instance   string   `json:"instance"`   // New instance (replica).
	Class      string   `json:"class"`      // New instance class.
	Region     string   `json:"region"`     // AWS region account.
	Primary    string   `json:"primary"`    // Source instance (primary).
	Status     string   `json:"status"`     // New instance status.
	Endpoint   string   `json:"endpoint"`   // New instance endpoint.
	Protection bool     `json:"protection"` // New instance is set to deletion protection.
	Port       uint16   `json:"port"`       // New instance port of endpoint.
	Zone       string   `json:"zone"`       // New instance Availability Zone.
}

func (r *RDS) Init() (err error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(r.Region),
	})

	r.Client = rds.New(sess)

	return
}

func (r *RDS) List() Instances {
	instances := Instances{}
	instance := Instance{}
	result, _ := r.Client.DescribeDBInstances(nil)

	for _, d := range result.DBInstances {
		instances = append(instances, *instance.New(d))
	}

	return instances
}

func (r *RDS) Describe(identifier string) (Instance, error) {
	instance := Instance{}
	input := &rds.DescribeDBInstancesInput{
		DBInstanceIdentifier: aws.String(identifier),
	}

	result, err := r.Client.DescribeDBInstances(input)
	if err != nil {
		return Instance{}, err
	}

	return *instance.New(result.DBInstances[0]), nil
}

func (r *RDS) Logs(identifier, filename string) (Logs, error) {
	logs := Logs{}
	input := &rds.DescribeDBLogFilesInput{
		DBInstanceIdentifier: aws.String(identifier),
		FilenameContains:     aws.String(filename),
	}

	result, err := r.Client.DescribeDBLogFiles(input)
	if err != nil {
		return Logs{}, err
	}

	return *logs.New(result), nil
}

func (r *RDS) PollLogs(identifier, filename string) (string, error) {
	params := &rds.DownloadDBLogFilePortionInput{
		DBInstanceIdentifier: aws.String(identifier),
		LogFileName:          aws.String(filename),
	}

	result, err := r.Client.DownloadDBLogFilePortion(params)
	if err != nil {
		return "", err
	}

	if result.LogFileData != nil {
		return aws.StringValue(result.LogFileData), nil
	}

	return "", nil
}

func (r *RDS) ParametersGroup() (ParametersGroup, error) {
	pg := ParametersGroup{}

	params := &rds.DescribeDBParameterGroupsInput{}

	result, err := r.Client.DescribeDBParameterGroups(params)
	if err != nil {
		return ParametersGroup{}, err
	}

	return *pg.New(result), nil
}

func (r *RDS) Parameters(name string) (Parameters, error) {
	p := Parameters{}
	pageNum := 0

	params := &rds.DescribeDBParametersInput{
		DBParameterGroupName: aws.String(name),
	}

	err := r.Client.DescribeDBParametersPages(
		params,
		func(page *rds.DescribeDBParametersOutput, lastPage bool) bool {
			pageNum++
			// fmt.Println(page)
			p.New(page)

			return pageNum <= 3
		})
	if err != nil {
		return Parameters{}, err
	}

	// return *p.New(result), nil
	return p, nil
}

func (r *RDS) Create() (err error) {
	input := &rds.CreateDBInstanceReadReplicaInput{
		AutoMinorVersionUpgrade:         aws.Bool(false),
		AvailabilityZone:                aws.String(r.Zone),
		DBInstanceClass:                 aws.String(r.Class),
		DBInstanceIdentifier:            aws.String(r.Instance),
		DeletionProtection:              aws.Bool(false),
		EnableCustomerOwnedIp:           aws.Bool(false),
		EnableIAMDatabaseAuthentication: aws.Bool(false),
		EnablePerformanceInsights:       aws.Bool(false),
		MultiAZ:                         aws.Bool(false),
		PubliclyAccessible:              aws.Bool(false),
		SourceDBInstanceIdentifier:      aws.String(r.Primary),
	}

	_, err = r.Client.CreateDBInstanceReadReplica(input)
	return
}

func (r *RDS) Delete() (err error) {
	input := &rds.DeleteDBInstanceInput{
		DBInstanceIdentifier: aws.String(r.Instance),
		SkipFinalSnapshot:    aws.Bool(true),
	}

	_, err = r.Client.DeleteDBInstance(input)
	return
}

func (r *RDS) DescribeOld() {
	r.Endpoint = ""
	r.Status = ""

	input := &rds.DescribeDBInstancesInput{
		DBInstanceIdentifier: aws.String(r.Instance),
	}

	result, err := r.Client.DescribeDBInstances(input)
	if err != nil {
		return
	}

	r.Status = *result.DBInstances[0].DBInstanceStatus
	r.Protection = *result.DBInstances[0].DeletionProtection

	if result.DBInstances[0].ReadReplicaSourceDBInstanceIdentifier != nil {
		r.Primary = *result.DBInstances[0].ReadReplicaSourceDBInstanceIdentifier
	}

	if result.DBInstances[0].Endpoint != nil {
		r.Endpoint = *result.DBInstances[0].Endpoint.Address
		r.Port = uint16(*result.DBInstances[0].Endpoint.Port)
	}
}

func (r *RDS) Exist() bool {
	r.DescribeOld()
	return len(r.Endpoint) > 0 && len(r.Status) > 0
}

func (r *RDS) IsAvailable() bool {
	r.DescribeOld()
	return len(r.Endpoint) > 0 && r.Status == "available"
}

func (r *RDS) IsPrimary() bool {
	return r.IsAvailable() && len(r.Primary) == 0
}

func (r *RDS) IsReplica() bool {
	return r.IsAvailable() && len(r.Primary) > 0
}

func (r *RDS) IsNotAvailable() bool {
	r.DescribeOld()
	return len(r.Endpoint) == 0 && len(r.Status) == 0
}

func (r *RDS) WaitUntilAvailable() error {
	return retry.Do(30, 60, func() bool {
		return r.IsAvailable()
	})
}

func (r *RDS) WaitUntilUnavailable() error {
	return retry.Do(30, 60, func() bool {
		return r.IsNotAvailable()
	})
}
