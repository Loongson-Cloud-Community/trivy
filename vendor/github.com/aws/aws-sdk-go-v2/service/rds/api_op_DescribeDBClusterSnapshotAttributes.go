// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of DB cluster snapshot attribute names and values for a manual DB
// cluster snapshot. When sharing snapshots with other Amazon Web Services
// accounts, DescribeDBClusterSnapshotAttributes returns the restore attribute and
// a list of IDs for the Amazon Web Services accounts that are authorized to copy
// or restore the manual DB cluster snapshot. If all is included in the list of
// values for the restore attribute, then the manual DB cluster snapshot is public
// and can be copied or restored by all Amazon Web Services accounts. To add or
// remove access for an Amazon Web Services account to copy or restore a manual DB
// cluster snapshot, or to make the manual DB cluster snapshot public or private,
// use the ModifyDBClusterSnapshotAttribute API action.
func (c *Client) DescribeDBClusterSnapshotAttributes(ctx context.Context, params *DescribeDBClusterSnapshotAttributesInput, optFns ...func(*Options)) (*DescribeDBClusterSnapshotAttributesOutput, error) {
	if params == nil {
		params = &DescribeDBClusterSnapshotAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBClusterSnapshotAttributes", params, optFns, c.addOperationDescribeDBClusterSnapshotAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBClusterSnapshotAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDBClusterSnapshotAttributesInput struct {

	// The identifier for the DB cluster snapshot to describe the attributes for.
	//
	// This member is required.
	DBClusterSnapshotIdentifier *string

	noSmithyDocumentSerde
}

type DescribeDBClusterSnapshotAttributesOutput struct {

	// Contains the results of a successful call to the
	// DescribeDBClusterSnapshotAttributes API action. Manual DB cluster snapshot
	// attributes are used to authorize other Amazon Web Services accounts to copy or
	// restore a manual DB cluster snapshot. For more information, see the
	// ModifyDBClusterSnapshotAttribute API action.
	DBClusterSnapshotAttributesResult *types.DBClusterSnapshotAttributesResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDBClusterSnapshotAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBClusterSnapshotAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBClusterSnapshotAttributes{}, middleware.After)
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
	if err = addOpDescribeDBClusterSnapshotAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBClusterSnapshotAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDBClusterSnapshotAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBClusterSnapshotAttributes",
	}
}
