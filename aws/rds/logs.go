package rds

import (
	"sort"

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

// Len is part of sort.Interface.
func (l Logs) Len() int {
	return len(l)
}

// Less is part of sort.Interface.
// We use count as the value to sort by
func (l Logs) Less(i, j int) bool {
	return l[i].Size < l[j].Size
}

// Swap is part of sort.Interface.
func (l Logs) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l Logs) SortBySize() {
	sort.Sort(l)
}
