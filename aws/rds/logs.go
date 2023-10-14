package rds

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
)

type Logs []Log

type Log struct {
	LastWritten int64
	FileName    string
	Size        int64 // In bytes, file with 75396 bytes is empty.
}

func (l *Logs) New(input *rds.DescribeDBLogFilesOutput) *Logs {
	for _, log := range input.DescribeDBLogFiles {
		*l = append(*l, Log{
			LastWritten: aws.Int64Value(log.LastWritten),
			FileName:    aws.StringValue(log.LogFileName),
			Size:        aws.Int64Value(log.Size),
		})
	}

	return l
}
