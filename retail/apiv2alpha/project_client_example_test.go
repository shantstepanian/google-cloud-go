// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package retail_test

import (
	"context"

	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	retail "cloud.google.com/go/retail/apiv2alpha"
	retailpb "cloud.google.com/go/retail/apiv2alpha/retailpb"
	"google.golang.org/api/iterator"
)

func ExampleNewProjectClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := retail.NewProjectClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleNewProjectRESTClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := retail.NewProjectRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleProjectClient_AcceptTerms() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := retail.NewProjectClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.AcceptTermsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/retail/apiv2alpha/retailpb#AcceptTermsRequest.
	}
	resp, err := c.AcceptTerms(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleProjectClient_EnrollSolution() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := retail.NewProjectClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.EnrollSolutionRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/retail/apiv2alpha/retailpb#EnrollSolutionRequest.
	}
	op, err := c.EnrollSolution(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleProjectClient_GetAlertConfig() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := retail.NewProjectClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.GetAlertConfigRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/retail/apiv2alpha/retailpb#GetAlertConfigRequest.
	}
	resp, err := c.GetAlertConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleProjectClient_GetLoggingConfig() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := retail.NewProjectClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.GetLoggingConfigRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/retail/apiv2alpha/retailpb#GetLoggingConfigRequest.
	}
	resp, err := c.GetLoggingConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleProjectClient_GetProject() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := retail.NewProjectClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.GetProjectRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/retail/apiv2alpha/retailpb#GetProjectRequest.
	}
	resp, err := c.GetProject(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleProjectClient_ListEnrolledSolutions() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := retail.NewProjectClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.ListEnrolledSolutionsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/retail/apiv2alpha/retailpb#ListEnrolledSolutionsRequest.
	}
	resp, err := c.ListEnrolledSolutions(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleProjectClient_UpdateAlertConfig() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := retail.NewProjectClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.UpdateAlertConfigRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/retail/apiv2alpha/retailpb#UpdateAlertConfigRequest.
	}
	resp, err := c.UpdateAlertConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleProjectClient_UpdateLoggingConfig() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := retail.NewProjectClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.UpdateLoggingConfigRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/retail/apiv2alpha/retailpb#UpdateLoggingConfigRequest.
	}
	resp, err := c.UpdateLoggingConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleProjectClient_GetOperation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := retail.NewProjectClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.GetOperationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/longrunning/autogen/longrunningpb#GetOperationRequest.
	}
	resp, err := c.GetOperation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleProjectClient_ListOperations() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := retail.NewProjectClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.ListOperationsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/longrunning/autogen/longrunningpb#ListOperationsRequest.
	}
	it := c.ListOperations(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp

		// If you need to access the underlying RPC response,
		// you can do so by casting the `Response` as below.
		// Otherwise, remove this line. Only populated after
		// first call to Next(). Not safe for concurrent access.
		_ = it.Response.(*longrunningpb.ListOperationsResponse)
	}
}
