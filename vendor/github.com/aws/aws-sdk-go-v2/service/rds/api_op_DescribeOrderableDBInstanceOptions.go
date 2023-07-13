// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of orderable DB instance options for the specified DB engine, DB
// engine version, and DB instance class.
func (c *Client) DescribeOrderableDBInstanceOptions(ctx context.Context, params *DescribeOrderableDBInstanceOptionsInput, optFns ...func(*Options)) (*DescribeOrderableDBInstanceOptionsOutput, error) {
	if params == nil {
		params = &DescribeOrderableDBInstanceOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeOrderableDBInstanceOptions", params, optFns, c.addOperationDescribeOrderableDBInstanceOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeOrderableDBInstanceOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeOrderableDBInstanceOptionsInput struct {

	// The name of the engine to retrieve DB instance options for. Valid Values:
	//
	// *
	// aurora (for MySQL 5.6-compatible Aurora)
	//
	// * aurora-mysql (for MySQL
	// 5.7-compatible and MySQL 8.0-compatible Aurora)
	//
	// * aurora-postgresql
	//
	// *
	// mariadb
	//
	// * mysql
	//
	// * oracle-ee
	//
	// * oracle-ee-cdb
	//
	// * oracle-se2
	//
	// *
	// oracle-se2-cdb
	//
	// * postgres
	//
	// * sqlserver-ee
	//
	// * sqlserver-se
	//
	// * sqlserver-ex
	//
	// *
	// sqlserver-web
	//
	// This member is required.
	Engine *string

	// The Availability Zone group associated with a Local Zone. Specify this parameter
	// to retrieve available offerings for the Local Zones in the group. Omit this
	// parameter to show the available offerings in the specified Amazon Web Services
	// Region. This setting doesn't apply to RDS Custom.
	AvailabilityZoneGroup *string

	// The DB instance class filter value. Specify this parameter to show only the
	// available offerings matching the specified DB instance class.
	DBInstanceClass *string

	// The engine version filter value. Specify this parameter to show only the
	// available offerings matching the specified engine version.
	EngineVersion *string

	// This parameter isn't currently supported.
	Filters []types.Filter

	// The license model filter value. Specify this parameter to show only the
	// available offerings matching the specified license model. RDS Custom supports
	// only the BYOL licensing model.
	LicenseModel *string

	// An optional pagination token provided by a previous
	// DescribeOrderableDBInstanceOptions request. If this parameter is specified, the
	// response includes only records beyond the marker, up to the value specified by
	// MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that you can retrieve the remaining results.
	// Default: 100 Constraints: Minimum 20, maximum 10000.
	MaxRecords *int32

	// A value that indicates whether to show only VPC or non-VPC offerings. RDS Custom
	// supports only VPC offerings. RDS Custom supports only VPC offerings. If you
	// describe non-VPC offerings for RDS Custom, the output shows VPC offerings.
	Vpc *bool

	noSmithyDocumentSerde
}

// Contains the result of a successful invocation of the
// DescribeOrderableDBInstanceOptions action.
type DescribeOrderableDBInstanceOptionsOutput struct {

	// An optional pagination token provided by a previous OrderableDBInstanceOptions
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string

	// An OrderableDBInstanceOption structure containing information about orderable
	// options for the DB instance.
	OrderableDBInstanceOptions []types.OrderableDBInstanceOption

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeOrderableDBInstanceOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeOrderableDBInstanceOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeOrderableDBInstanceOptions{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeOrderableDBInstanceOptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeOrderableDBInstanceOptions(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// DescribeOrderableDBInstanceOptionsAPIClient is a client that implements the
// DescribeOrderableDBInstanceOptions operation.
type DescribeOrderableDBInstanceOptionsAPIClient interface {
	DescribeOrderableDBInstanceOptions(context.Context, *DescribeOrderableDBInstanceOptionsInput, ...func(*Options)) (*DescribeOrderableDBInstanceOptionsOutput, error)
}

var _ DescribeOrderableDBInstanceOptionsAPIClient = (*Client)(nil)

// DescribeOrderableDBInstanceOptionsPaginatorOptions is the paginator options for
// DescribeOrderableDBInstanceOptions
type DescribeOrderableDBInstanceOptionsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that you can retrieve the remaining results.
	// Default: 100 Constraints: Minimum 20, maximum 10000.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeOrderableDBInstanceOptionsPaginator is a paginator for
// DescribeOrderableDBInstanceOptions
type DescribeOrderableDBInstanceOptionsPaginator struct {
	options   DescribeOrderableDBInstanceOptionsPaginatorOptions
	client    DescribeOrderableDBInstanceOptionsAPIClient
	params    *DescribeOrderableDBInstanceOptionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeOrderableDBInstanceOptionsPaginator returns a new
// DescribeOrderableDBInstanceOptionsPaginator
func NewDescribeOrderableDBInstanceOptionsPaginator(client DescribeOrderableDBInstanceOptionsAPIClient, params *DescribeOrderableDBInstanceOptionsInput, optFns ...func(*DescribeOrderableDBInstanceOptionsPaginatorOptions)) *DescribeOrderableDBInstanceOptionsPaginator {
	if params == nil {
		params = &DescribeOrderableDBInstanceOptionsInput{}
	}

	options := DescribeOrderableDBInstanceOptionsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeOrderableDBInstanceOptionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeOrderableDBInstanceOptionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeOrderableDBInstanceOptions page.
func (p *DescribeOrderableDBInstanceOptionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeOrderableDBInstanceOptionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeOrderableDBInstanceOptions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeOrderableDBInstanceOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeOrderableDBInstanceOptions",
	}
}
