// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package codedeploy provides the client and types for making API
// requests to AWS CodeDeploy.
//
// AWS CodeDeploy is a deployment service that automates application deployments
// to Amazon EC2 instances or on-premises instances running in your own facility.
//
// You can deploy a nearly unlimited variety of application content, such as
// code, web and configuration files, executables, packages, scripts, multimedia
// files, and so on. AWS CodeDeploy can deploy application content stored in
// Amazon S3 buckets, GitHub repositories, or Bitbucket repositories. You do
// not need to make changes to your existing code before you can use AWS CodeDeploy.
//
// AWS CodeDeploy makes it easier for you to rapidly release new features, helps
// you avoid downtime during application deployment, and handles the complexity
// of updating your applications, without many of the risks associated with
// error-prone manual deployments.
//
// AWS CodeDeploy Components
//
// Use the information in this guide to help you work with the following AWS
// CodeDeploy components:
//
//    * Application: A name that uniquely identifies the application you want
//    to deploy. AWS CodeDeploy uses this name, which functions as a container,
//    to ensure the correct combination of revision, deployment configuration,
//    and deployment group are referenced during a deployment.
//
//    * Deployment group: A set of individual instances. A deployment group
//    contains individually tagged instances, Amazon EC2 instances in Auto Scaling
//    groups, or both.
//
//    * Deployment configuration: A set of deployment rules and deployment success
//    and failure conditions used by AWS CodeDeploy during a deployment.
//
//    * Deployment: The process, and the components involved in the process,
//    of installing content on one or more instances.
//
//    * Application revisions: An archive file containing source content—source
//    code, web pages, executable files, and deployment scripts—along with an
//    application specification file (AppSpec file). Revisions are stored in
//    Amazon S3 buckets or GitHub repositories. For Amazon S3, a revision is
//    uniquely identified by its Amazon S3 object key and its ETag, version,
//    or both. For GitHub, a revision is uniquely identified by its commit ID.
//
// This guide also contains information to help you get details about the instances
// in your deployments and to make on-premises instances available for AWS CodeDeploy
// deployments.
//
// AWS CodeDeploy Information Resources
//
//    * AWS CodeDeploy User Guide (http://docs.aws.amazon.com/codedeploy/latest/userguide)
//
//    * AWS CodeDeploy API Reference Guide (http://docs.aws.amazon.com/codedeploy/latest/APIReference/)
//
//    * AWS CLI Reference for AWS CodeDeploy (http://docs.aws.amazon.com/cli/latest/reference/deploy/index.html)
//
//    * AWS CodeDeploy Developer Forum (https://forums.aws.amazon.com/forum.jspa?forumID=179)
//
// See https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06 for more information on this service.
//
// See codedeploy package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/codedeploy/
//
// Using the Client
//
// To use the client for AWS CodeDeploy you will first need
// to create a new instance of it.
//
// When creating a client for an AWS service you'll first need to have a Session
// already created. The Session provides configuration that can be shared
// between multiple service clients. Additional configuration can be applied to
// the Session and service's client when they are constructed. The aws package's
// Config type contains several fields such as Region for the AWS Region the
// client should make API requests too. The optional Config value can be provided
// as the variadic argument for Sessions and client creation.
//
// Once the service's client is created you can use it to make API requests the
// AWS service. These clients are safe to use concurrently.
//
//   // Create a session to share configuration, and load external configuration.
//   sess := session.Must(session.NewSession())
//
//   // Create the service's client with the session.
//   svc := codedeploy.New(sess)
//
// See the SDK's documentation for more information on how to use service clients.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws package's Config type for more information on configuration options.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS CodeDeploy client CodeDeploy for more
// information on creating the service's client.
// https://docs.aws.amazon.com/sdk-for-go/api/service/codedeploy/#New
//
// Once the client is created you can make an API request to the service.
// Each API method takes a input parameter, and returns the service response
// and an error.
//
// The API method will document which error codes the service can be returned
// by the operation if the service models the API operation's errors. These
// errors will also be available as const strings prefixed with "ErrCode".
//
//   result, err := svc.AddTagsToOnPremisesInstances(params)
//   if err != nil {
//       // Cast err to awserr.Error to handle specific error codes.
//       aerr, ok := err.(awserr.Error)
//       if ok && aerr.Code() == <error code to check for> {
//           // Specific error code handling
//       }
//       return err
//   }
//
//   fmt.Println("AddTagsToOnPremisesInstances result:")
//   fmt.Println(result)
//
// Using the Client with Context
//
// The service's client also provides methods to make API requests with a Context
// value. This allows you to control the timeout, and cancellation of pending
// requests. These methods also take request Option as variadic parameter to apply
// additional configuration to the API request.
//
//   ctx := context.Background()
//
//   result, err := svc.AddTagsToOnPremisesInstancesWithContext(ctx, params)
//
// See the request package documentation for more information on using Context pattern
// with the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/request/
package codedeploy
