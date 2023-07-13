// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of versions of the package, along with their creation time and
// commit message.
func (c *Client) GetPackageVersionHistory(ctx context.Context, params *GetPackageVersionHistoryInput, optFns ...func(*Options)) (*GetPackageVersionHistoryOutput, error) {
	if params == nil {
		params = &GetPackageVersionHistoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPackageVersionHistory", params, optFns, c.addOperationGetPackageVersionHistoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPackageVersionHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for request parameters to GetPackageVersionHistory operation.
type GetPackageVersionHistoryInput struct {

	// Returns an audit history of versions of the package.
	//
	// This member is required.
	PackageID *string

	// Limits results to a maximum number of versions.
	MaxResults int32

	// Used for pagination. Only necessary if a previous API call includes a non-null
	// NextToken value. If provided, returns results for the next page.
	NextToken *string

	noSmithyDocumentSerde
}

// Container for response returned by GetPackageVersionHistory operation.
type GetPackageVersionHistoryOutput struct {
	NextToken *string

	PackageID *string

	// List of PackageVersionHistory objects.
	PackageVersionHistoryList []types.PackageVersionHistory

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetPackageVersionHistoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetPackageVersionHistory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetPackageVersionHistory{}, middleware.After)
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
	if err = addOpGetPackageVersionHistoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPackageVersionHistory(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

// GetPackageVersionHistoryAPIClient is a client that implements the
// GetPackageVersionHistory operation.
type GetPackageVersionHistoryAPIClient interface {
	GetPackageVersionHistory(context.Context, *GetPackageVersionHistoryInput, ...func(*Options)) (*GetPackageVersionHistoryOutput, error)
}

var _ GetPackageVersionHistoryAPIClient = (*Client)(nil)

// GetPackageVersionHistoryPaginatorOptions is the paginator options for
// GetPackageVersionHistory
type GetPackageVersionHistoryPaginatorOptions struct {
	// Limits results to a maximum number of versions.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetPackageVersionHistoryPaginator is a paginator for GetPackageVersionHistory
type GetPackageVersionHistoryPaginator struct {
	options   GetPackageVersionHistoryPaginatorOptions
	client    GetPackageVersionHistoryAPIClient
	params    *GetPackageVersionHistoryInput
	nextToken *string
	firstPage bool
}

// NewGetPackageVersionHistoryPaginator returns a new
// GetPackageVersionHistoryPaginator
func NewGetPackageVersionHistoryPaginator(client GetPackageVersionHistoryAPIClient, params *GetPackageVersionHistoryInput, optFns ...func(*GetPackageVersionHistoryPaginatorOptions)) *GetPackageVersionHistoryPaginator {
	if params == nil {
		params = &GetPackageVersionHistoryInput{}
	}

	options := GetPackageVersionHistoryPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetPackageVersionHistoryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetPackageVersionHistoryPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetPackageVersionHistory page.
func (p *GetPackageVersionHistoryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetPackageVersionHistoryOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.GetPackageVersionHistory(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opGetPackageVersionHistory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "GetPackageVersionHistory",
	}
}
