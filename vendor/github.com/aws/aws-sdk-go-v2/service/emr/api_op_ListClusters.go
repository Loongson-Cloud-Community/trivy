// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides the status of all clusters visible to this Amazon Web Services
// account. Allows you to filter the list of clusters based on certain criteria;
// for example, filtering by cluster creation date and time or by status. This call
// returns a maximum of 50 clusters in unsorted order per call, but returns a
// marker to track the paging of the cluster list across multiple ListClusters
// calls.
func (c *Client) ListClusters(ctx context.Context, params *ListClustersInput, optFns ...func(*Options)) (*ListClustersOutput, error) {
	if params == nil {
		params = &ListClustersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListClusters", params, optFns, c.addOperationListClustersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListClustersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// This input determines how the ListClusters action filters the list of clusters
// that it returns.
type ListClustersInput struct {

	// The cluster state filters to apply when listing clusters. Clusters that change
	// state while this action runs may be not be returned as expected in the list of
	// clusters.
	ClusterStates []types.ClusterState

	// The creation date and time beginning value filter for listing clusters.
	CreatedAfter *time.Time

	// The creation date and time end value filter for listing clusters.
	CreatedBefore *time.Time

	// The pagination token that indicates the next set of results to retrieve.
	Marker *string

	noSmithyDocumentSerde
}

// This contains a ClusterSummaryList with the cluster details; for example, the
// cluster IDs, names, and status.
type ListClustersOutput struct {

	// The list of clusters for the account based on the given filters.
	Clusters []types.ClusterSummary

	// The pagination token that indicates the next set of results to retrieve.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListClustersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListClusters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListClusters{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListClusters(options.Region), middleware.Before); err != nil {
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

// ListClustersAPIClient is a client that implements the ListClusters operation.
type ListClustersAPIClient interface {
	ListClusters(context.Context, *ListClustersInput, ...func(*Options)) (*ListClustersOutput, error)
}

var _ ListClustersAPIClient = (*Client)(nil)

// ListClustersPaginatorOptions is the paginator options for ListClusters
type ListClustersPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListClustersPaginator is a paginator for ListClusters
type ListClustersPaginator struct {
	options   ListClustersPaginatorOptions
	client    ListClustersAPIClient
	params    *ListClustersInput
	nextToken *string
	firstPage bool
}

// NewListClustersPaginator returns a new ListClustersPaginator
func NewListClustersPaginator(client ListClustersAPIClient, params *ListClustersInput, optFns ...func(*ListClustersPaginatorOptions)) *ListClustersPaginator {
	if params == nil {
		params = &ListClustersInput{}
	}

	options := ListClustersPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListClustersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListClustersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListClusters page.
func (p *ListClustersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListClustersOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	result, err := p.client.ListClusters(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListClusters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "ListClusters",
	}
}
