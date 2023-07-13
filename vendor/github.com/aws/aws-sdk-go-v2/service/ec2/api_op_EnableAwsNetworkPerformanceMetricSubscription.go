// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables Infrastructure Performance subscriptions.
func (c *Client) EnableAwsNetworkPerformanceMetricSubscription(ctx context.Context, params *EnableAwsNetworkPerformanceMetricSubscriptionInput, optFns ...func(*Options)) (*EnableAwsNetworkPerformanceMetricSubscriptionOutput, error) {
	if params == nil {
		params = &EnableAwsNetworkPerformanceMetricSubscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableAwsNetworkPerformanceMetricSubscription", params, optFns, c.addOperationEnableAwsNetworkPerformanceMetricSubscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableAwsNetworkPerformanceMetricSubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableAwsNetworkPerformanceMetricSubscriptionInput struct {

	// The target Region or Availability Zone that the metric subscription is enabled
	// for. For example, eu-west-1 .
	Destination *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The metric used for the enabled subscription.
	Metric types.MetricType

	// The source Region or Availability Zone that the metric subscription is enabled
	// for. For example, us-east-1 .
	Source *string

	// The statistic used for the enabled subscription.
	Statistic types.StatisticType

	noSmithyDocumentSerde
}

type EnableAwsNetworkPerformanceMetricSubscriptionOutput struct {

	// Indicates whether the subscribe action was successful.
	Output *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEnableAwsNetworkPerformanceMetricSubscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpEnableAwsNetworkPerformanceMetricSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpEnableAwsNetworkPerformanceMetricSubscription{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnableAwsNetworkPerformanceMetricSubscription(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opEnableAwsNetworkPerformanceMetricSubscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "EnableAwsNetworkPerformanceMetricSubscription",
	}
}
