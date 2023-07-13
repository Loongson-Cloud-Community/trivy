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

// Changes the audit policy state of a database activity stream to either locked
// (default) or unlocked. A locked policy is read-only, whereas an unlocked policy
// is read/write. If your activity stream is started and locked, you can unlock it,
// customize your audit policy, and then lock your activity stream. Restarting the
// activity stream isn't required. For more information, see  Modifying a database
// activity stream
// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/DBActivityStreams.Modifying.html)
// in the Amazon RDS User Guide. This operation is supported for RDS for Oracle
// only.
func (c *Client) ModifyActivityStream(ctx context.Context, params *ModifyActivityStreamInput, optFns ...func(*Options)) (*ModifyActivityStreamOutput, error) {
	if params == nil {
		params = &ModifyActivityStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyActivityStream", params, optFns, c.addOperationModifyActivityStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyActivityStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyActivityStreamInput struct {

	// The audit policy state. When a policy is unlocked, it is read/write. When it is
	// locked, it is read-only. You can edit your audit policy only when the activity
	// stream is unlocked or stopped.
	AuditPolicyState types.AuditPolicyState

	// The Amazon Resource Name (ARN) of the RDS for Oracle DB instance, for example,
	// arn:aws:rds:us-east-1:12345667890:instance:my-orcl-db.
	ResourceArn *string

	noSmithyDocumentSerde
}

type ModifyActivityStreamOutput struct {

	// Indicates whether engine-native audit fields are included in the database
	// activity stream.
	EngineNativeAuditFieldsIncluded *bool

	// The name of the Amazon Kinesis data stream to be used for the database activity
	// stream.
	KinesisStreamName *string

	// The Amazon Web Services KMS key identifier for encryption of messages in the
	// database activity stream.
	KmsKeyId *string

	// The mode of the database activity stream.
	Mode types.ActivityStreamMode

	// The status of the modification to the policy state of the database activity
	// stream.
	PolicyStatus types.ActivityStreamPolicyStatus

	// The status of the modification to the database activity stream.
	Status types.ActivityStreamStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyActivityStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyActivityStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyActivityStream{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyActivityStream(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyActivityStream(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ModifyActivityStream",
	}
}
