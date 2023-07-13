// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the latest status of the last upgrade or upgrade eligibility check
// that was performed on the domain.
func (c *Client) GetUpgradeStatus(ctx context.Context, params *GetUpgradeStatusInput, optFns ...func(*Options)) (*GetUpgradeStatusOutput, error) {
	if params == nil {
		params = &GetUpgradeStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetUpgradeStatus", params, optFns, c.addOperationGetUpgradeStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetUpgradeStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for request parameters to GetUpgradeStatus operation.
type GetUpgradeStatusInput struct {

	// The name of an Elasticsearch domain. Domain names are unique across the domains
	// owned by an account within an AWS region. Domain names start with a letter or
	// number and can contain the following characters: a-z (lowercase), 0-9, and -
	// (hyphen).
	//
	// This member is required.
	DomainName *string

	noSmithyDocumentSerde
}

// Container for response returned by GetUpgradeStatus operation.
type GetUpgradeStatusOutput struct {

	// One of 4 statuses that a step can go through returned as part of the
	// GetUpgradeStatusResponse object. The status can take one of the following
	// values:
	//   - In Progress
	//   - Succeeded
	//   - Succeeded with Issues
	//   - Failed
	StepStatus types.UpgradeStatus

	// A string that describes the update briefly
	UpgradeName *string

	// Represents one of 3 steps that an Upgrade or Upgrade Eligibility Check does
	// through:
	//   - PreUpgradeCheck
	//   - Snapshot
	//   - Upgrade
	UpgradeStep types.UpgradeStep

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetUpgradeStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetUpgradeStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetUpgradeStatus{}, middleware.After)
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
	if err = addOpGetUpgradeStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetUpgradeStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetUpgradeStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "GetUpgradeStatus",
	}
}
