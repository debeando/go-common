package rds

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
)

type ParametersGroup []ParameterGroup

type ParameterGroup struct {
	Name        string
	Status      string
	Family      string
	Description string
}

func (pg *ParametersGroup) New(input *rds.DescribeDBParameterGroupsOutput) *ParametersGroup {
	for _, parameterGroup := range input.DBParameterGroups {
		*pg = append(*pg, ParameterGroup{
			Name:        aws.StringValue(parameterGroup.DBParameterGroupName),
			Family:      aws.StringValue(parameterGroup.DBParameterGroupFamily),
			Description: aws.StringValue(parameterGroup.Description),
		})
	}

	return pg
}
