// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package dynamodbstreams

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/service"
	"github.com/aws/aws-sdk-go/aws/service/serviceinfo"
	"github.com/aws/aws-sdk-go/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go/private/signer/v4"
)

// This is the Amazon DynamoDB Streams API Reference. This guide describes the
// low-level API actions for accessing streams and processing stream records.
// For information about application development with DynamoDB Streams, see
// the Amazon DynamoDB Developer Guide (http://docs.aws.amazon.com/amazondynamodb/latest/developerguide//Streams.html).
//
// Note that this document is intended for use with the following DynamoDB
// documentation:
//
//    Amazon DynamoDB Developer Guide (http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/)
//
//    Amazon DynamoDB API Reference (http://docs.aws.amazon.com/amazondynamodb/latest/APIReference/)
//
//   The following are short descriptions of each low-level DynamoDB Streams
// API action, organized by function.
//
//  DescribeStream - Returns detailed information about a particular stream.
//
//  GetRecords - Retrieves the stream records from within a shard.
//
//   GetShardIterator - Returns information on how to retrieve the streams
// record from a shard with a given shard ID.
//
//   ListStreams - Returns a list of all the streams associated with the current
// AWS account and endpoint.
type DynamoDBStreams struct {
	*service.Service
}

// Used for custom service initialization logic
var initService func(*service.Service)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// New returns a new DynamoDBStreams client.
func New(config *aws.Config) *DynamoDBStreams {
	service := &service.Service{
		ServiceInfo: serviceinfo.ServiceInfo{
			Config:       defaults.DefaultConfig.Merge(config),
			ServiceName:  "streams.dynamodb",
			SigningName:  "dynamodb",
			APIVersion:   "2012-08-10",
			JSONVersion:  "1.0",
			TargetPrefix: "DynamoDBStreams_20120810",
		},
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &DynamoDBStreams{service}
}

// newRequest creates a new request for a DynamoDBStreams operation and runs any
// custom request initialization.
func (c *DynamoDBStreams) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
