// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Stops a running task. Any tags associated with the task will be deleted. When
// StopTask is called on a task, the equivalent of docker stop is issued to the
// containers running in the task. This results in a SIGTERM value and a default
// 30-second timeout, after which the SIGKILL value is sent and the containers are
// forcibly stopped. If the container handles the SIGTERM value gracefully and
// exits within 30 seconds from receiving it, no SIGKILL value is sent. The default
// 30-second timeout can be configured on the Amazon ECS container agent with the
// ECS_CONTAINER_STOP_TIMEOUT variable. For more information, see Amazon ECS
// Container Agent Configuration
// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-config.html)
// in the Amazon Elastic Container Service Developer Guide.
func (c *Client) StopTask(ctx context.Context, params *StopTaskInput, optFns ...func(*Options)) (*StopTaskOutput, error) {
	if params == nil {
		params = &StopTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopTask", params, optFns, c.addOperationStopTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopTaskInput struct {

	// The task ID or full Amazon Resource Name (ARN) of the task to stop.
	//
	// This member is required.
	Task *string

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts the
	// task to stop. If you do not specify a cluster, the default cluster is assumed.
	Cluster *string

	// An optional message specified when a task is stopped. For example, if you're
	// using a custom scheduler, you can use this parameter to specify the reason for
	// stopping the task here, and the message appears in subsequent DescribeTasks API
	// operations on this task. Up to 255 characters are allowed in this message.
	Reason *string

	noSmithyDocumentSerde
}

type StopTaskOutput struct {

	// The task that was stopped.
	Task *types.Task

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStopTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopTask{}, middleware.After)
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
	if err = addOpStopTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStopTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "StopTask",
	}
}
