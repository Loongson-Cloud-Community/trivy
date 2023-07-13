// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns a database user name and temporary password with temporary authorization
// to log in to an Amazon Redshift database. The database user is mapped 1:1 to the
// source Identity and Access Management (IAM) identity. For more information about
// IAM identities, see IAM Identities (users, user groups, and roles)
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id.html) in the Amazon Web
// Services Identity and Access Management User Guide. The Identity and Access
// Management (IAM) identity that runs this operation must have an IAM policy
// attached that allows access to all necessary actions and resources. For more
// information about permissions, see Using identity-based policies (IAM policies)
// (https://docs.aws.amazon.com/redshift/latest/mgmt/redshift-iam-access-control-identity-based.html)
// in the Amazon Redshift Cluster Management Guide.
func (c *Client) GetClusterCredentialsWithIAM(ctx context.Context, params *GetClusterCredentialsWithIAMInput, optFns ...func(*Options)) (*GetClusterCredentialsWithIAMOutput, error) {
	if params == nil {
		params = &GetClusterCredentialsWithIAMInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetClusterCredentialsWithIAM", params, optFns, c.addOperationGetClusterCredentialsWithIAMMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetClusterCredentialsWithIAMOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetClusterCredentialsWithIAMInput struct {

	// The unique identifier of the cluster that contains the database for which you
	// are requesting credentials.
	//
	// This member is required.
	ClusterIdentifier *string

	// The name of the database for which you are requesting credentials. If the
	// database name is specified, the IAM policy must allow access to the resource
	// dbname for the specified database name. If the database name is not specified,
	// access to all databases is allowed.
	DbName *string

	// The number of seconds until the returned temporary password expires. Range:
	// 900-3600. Default: 900.
	DurationSeconds *int32

	noSmithyDocumentSerde
}

type GetClusterCredentialsWithIAMOutput struct {

	// A temporary password that you provide when you connect to a database.
	DbPassword *string

	// A database user name that you provide when you connect to a database. The
	// database user is mapped 1:1 to the source IAM identity.
	DbUser *string

	// The time (UTC) when the temporary password expires. After this timestamp, a log
	// in with the temporary password fails.
	Expiration *time.Time

	// Reserved for future use.
	NextRefreshTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetClusterCredentialsWithIAMMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetClusterCredentialsWithIAM{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetClusterCredentialsWithIAM{}, middleware.After)
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
	if err = addOpGetClusterCredentialsWithIAMValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetClusterCredentialsWithIAM(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetClusterCredentialsWithIAM(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "GetClusterCredentialsWithIAM",
	}
}
