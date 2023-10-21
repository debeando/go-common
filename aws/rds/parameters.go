package rds

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
)

type Parameters []Parameter

type Parameter struct {
	AllowedValues string
	ApplyMethod   string
	ApplyType     string
	DataType      string
	Description   string
	IsModifiable  bool
	Name          string
	Source        string
	Value         string
}

func (p *Parameters) New(input *rds.DescribeDBParametersOutput) *Parameters {
	for _, parameter := range input.Parameters {
		*p = append(*p, Parameter{
			AllowedValues: aws.StringValue(parameter.AllowedValues),
			ApplyMethod:   aws.StringValue(parameter.ApplyMethod),
			ApplyType:     aws.StringValue(parameter.ApplyType),
			DataType:      aws.StringValue(parameter.DataType),
			Description:   aws.StringValue(parameter.Description),
			IsModifiable:  aws.BoolValue(parameter.IsModifiable),
			Name:          aws.StringValue(parameter.ParameterName),
			Source:        aws.StringValue(parameter.Source),
			Value:         aws.StringValue(parameter.ParameterValue),
		})
	}

	return p
}
