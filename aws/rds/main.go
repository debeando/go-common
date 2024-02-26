package rds

import (
	"errors"

	"github.com/debeando/go-common/env"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/rds/rdsiface"
)

type RDS struct {
	Client     rdsiface.RDSAPI
	Identifier string
	Instance   Instance
	Region     string
	CheckMode  bool
}

func (r *RDS) Init() error {
	if !env.Exist("AWS_REGION") {
		return errors.New("You must specify a region. Define value of environment variable AWS_REGION.")
	}

	sess := session.Must(session.NewSession())

	r.Region = aws.StringValue(sess.Config.Region)
	r.Client = rds.New(sess)

	return nil
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

func (r *RDS) Load() error {
	instance := Instance{}

	input := &rds.DescribeDBInstancesInput{
		DBInstanceIdentifier: aws.String(r.Identifier),
	}

	result, err := r.Client.DescribeDBInstances(input)
	if err != nil {
		return err
	}

	r.Instance = *instance.New(result.DBInstances[0])

	return nil
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

	return *logs.Decode(result), nil
}

func (r *RDS) PollLogs(identifier, filename string) (string, error) {
	params := &rds.DownloadDBLogFilePortionInput{
		DBInstanceIdentifier: aws.String(identifier),
		LogFileName:          aws.String(filename),
		Marker:               aws.String("0"),
	}

	var body string

	err := r.Client.DownloadDBLogFilePortionPages(
		params,
		func(page *rds.DownloadDBLogFilePortionOutput, lastPage bool) bool {
			body = body + aws.StringValue(page.LogFileData)
			return !lastPage
		})
	if err != nil {
		return "", err
	}

	return body, nil
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
			p.New(page)

			return pageNum <= 3
		})
	if err != nil {
		return Parameters{}, err
	}

	return p, nil
}

func (r *RDS) CreateReplica(instance Instance) error {
	if r.CheckMode {
		return nil
	}

	input := &rds.CreateDBInstanceReadReplicaInput{
		AutoMinorVersionUpgrade:         aws.Bool(false),
		AvailabilityZone:                aws.String(instance.AvailabilityZone),
		DBInstanceClass:                 aws.String(instance.Class),
		DBInstanceIdentifier:            aws.String(r.Identifier),
		DeletionProtection:              aws.Bool(false),
		EnableCustomerOwnedIp:           aws.Bool(false),
		EnableIAMDatabaseAuthentication: aws.Bool(false),
		EnablePerformanceInsights:       aws.Bool(true),
		MultiAZ:                         aws.Bool(false),
		PubliclyAccessible:              aws.Bool(false),
		SourceDBInstanceIdentifier:      aws.String(instance.Identifier),
	}

	_, err := r.Client.CreateDBInstanceReadReplica(input)
	return err
}

func (r *RDS) Delete() error {
	if r.CheckMode {
		return nil
	}

	input := &rds.DeleteDBInstanceInput{
		DBInstanceIdentifier: aws.String(r.Identifier),
		SkipFinalSnapshot:    aws.Bool(true),
	}

	_, err := r.Client.DeleteDBInstance(input)
	return err
}
