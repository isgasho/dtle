// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationdiscoveryservice

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/private/protocol/jsonrpc"
)

// AWS Application Discovery Service helps you plan application migration projects
// by automatically identifying servers, virtual machines (VMs), software, and
// software dependencies running in your on-premises data centers. Application
// Discovery Service also collects application performance data, which can help
// you assess the outcome of your migration. The data collected by Application
// Discovery Service is securely retained in an Amazon-hosted and managed database
// in the cloud. You can export the data as a CSV or XML file into your preferred
// visualization tool or cloud-migration solution to plan your migration. For
// more information, see the Application Discovery Service FAQ (http://aws.amazon.com/application-discovery/faqs/).
//
// Application Discovery Service offers two modes of operation.
//
//    * Agentless discovery mode is recommended for environments that use VMware
//    vCenter Server. This mode doesn't require you to install an agent on each
//    host. Agentless discovery gathers server information regardless of the
//    operating systems, which minimizes the time required for initial on-premises
//    infrastructure assessment. Agentless discovery doesn't collect information
//    about software and software dependencies. It also doesn't work in non-VMware
//    environments. We recommend that you use agent-based discovery for non-VMware
//    environments and if you want to collect information about software and
//    software dependencies. You can also run agent-based and agentless discovery
//    simultaneously. Use agentless discovery to quickly complete the initial
//    infrastructure assessment and then install agents on select hosts to gather
//    information about software and software dependencies.
//
//    * Agent-based discovery mode collects a richer set of data than agentless
//    discovery by using Amazon software, the AWS Application Discovery Agent,
//    which you install on one or more hosts in your data center. The agent
//    captures infrastructure and application information, including an inventory
//    of installed software applications, system and process performance, resource
//    utilization, and network dependencies between workloads. The information
//    collected by agents is secured at rest and in transit to the Application
//    Discovery Service database in the cloud.
//
// Application Discovery Service integrates with application discovery solutions
// from AWS Partner Network (APN) partners. Third-party application discovery
// tools can query Application Discovery Service and write to the Application
// Discovery Service database using a public API. You can then import the data
// into either a visualization tool or cloud-migration solution.
//
// Application Discovery Service doesn't gather sensitive information. All data
// is handled according to the AWS Privacy Policy (http://aws.amazon.com/privacy/).
// You can operate Application Discovery Service using offline mode to inspect
// collected data before it is shared with the service.
//
// Your AWS account must be granted access to Application Discovery Service,
// a process called whitelisting. This is true for AWS partners and customers
// alike. To request access, sign up for AWS Application Discovery Service here
// (http://aws.amazon.com/application-discovery/preview/). We send you information
// about how to get started.
//
// This API reference provides descriptions, syntax, and usage examples for
// each of the actions and data types for Application Discovery Service. The
// topic for each action shows the API request parameters and the response.
// Alternatively, you can use one of the AWS SDKs to access an API that is tailored
// to the programming language or platform that you're using. For more information,
// see AWS SDKs (http://aws.amazon.com/tools/#SDKs).
//
// This guide is intended for use with the AWS Application Discovery Service
// User Guide (http://docs.aws.amazon.com/application-discovery/latest/userguide/).
// The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01
type ApplicationDiscoveryService struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "discovery" // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the ApplicationDiscoveryService client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a ApplicationDiscoveryService client from just a session.
//     svc := applicationdiscoveryservice.New(mySession)
//
//     // Create a ApplicationDiscoveryService client with additional configuration
//     svc := applicationdiscoveryservice.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *ApplicationDiscoveryService {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *ApplicationDiscoveryService {
	svc := &ApplicationDiscoveryService{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2015-11-01",
				JSONVersion:   "1.1",
				TargetPrefix:  "AWSPoseidonService_V2015_11_01",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a ApplicationDiscoveryService operation and runs any
// custom request initialization.
func (c *ApplicationDiscoveryService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}