// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package gluedatabrewiface provides an interface to enable mocking the AWS Glue DataBrew service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package gluedatabrewiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/gluedatabrew"
)

// GlueDataBrewAPI provides an interface to enable mocking the
// gluedatabrew.GlueDataBrew service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Glue DataBrew.
//    func myFunc(svc gluedatabrewiface.GlueDataBrewAPI) bool {
//        // Make svc.BatchDeleteRecipeVersion request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := gluedatabrew.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockGlueDataBrewClient struct {
//        gluedatabrewiface.GlueDataBrewAPI
//    }
//    func (m *mockGlueDataBrewClient) BatchDeleteRecipeVersion(input *gluedatabrew.BatchDeleteRecipeVersionInput) (*gluedatabrew.BatchDeleteRecipeVersionOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockGlueDataBrewClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type GlueDataBrewAPI interface {
	BatchDeleteRecipeVersion(*gluedatabrew.BatchDeleteRecipeVersionInput) (*gluedatabrew.BatchDeleteRecipeVersionOutput, error)
	BatchDeleteRecipeVersionWithContext(aws.Context, *gluedatabrew.BatchDeleteRecipeVersionInput, ...request.Option) (*gluedatabrew.BatchDeleteRecipeVersionOutput, error)
	BatchDeleteRecipeVersionRequest(*gluedatabrew.BatchDeleteRecipeVersionInput) (*request.Request, *gluedatabrew.BatchDeleteRecipeVersionOutput)

	CreateDataset(*gluedatabrew.CreateDatasetInput) (*gluedatabrew.CreateDatasetOutput, error)
	CreateDatasetWithContext(aws.Context, *gluedatabrew.CreateDatasetInput, ...request.Option) (*gluedatabrew.CreateDatasetOutput, error)
	CreateDatasetRequest(*gluedatabrew.CreateDatasetInput) (*request.Request, *gluedatabrew.CreateDatasetOutput)

	CreateProfileJob(*gluedatabrew.CreateProfileJobInput) (*gluedatabrew.CreateProfileJobOutput, error)
	CreateProfileJobWithContext(aws.Context, *gluedatabrew.CreateProfileJobInput, ...request.Option) (*gluedatabrew.CreateProfileJobOutput, error)
	CreateProfileJobRequest(*gluedatabrew.CreateProfileJobInput) (*request.Request, *gluedatabrew.CreateProfileJobOutput)

	CreateProject(*gluedatabrew.CreateProjectInput) (*gluedatabrew.CreateProjectOutput, error)
	CreateProjectWithContext(aws.Context, *gluedatabrew.CreateProjectInput, ...request.Option) (*gluedatabrew.CreateProjectOutput, error)
	CreateProjectRequest(*gluedatabrew.CreateProjectInput) (*request.Request, *gluedatabrew.CreateProjectOutput)

	CreateRecipe(*gluedatabrew.CreateRecipeInput) (*gluedatabrew.CreateRecipeOutput, error)
	CreateRecipeWithContext(aws.Context, *gluedatabrew.CreateRecipeInput, ...request.Option) (*gluedatabrew.CreateRecipeOutput, error)
	CreateRecipeRequest(*gluedatabrew.CreateRecipeInput) (*request.Request, *gluedatabrew.CreateRecipeOutput)

	CreateRecipeJob(*gluedatabrew.CreateRecipeJobInput) (*gluedatabrew.CreateRecipeJobOutput, error)
	CreateRecipeJobWithContext(aws.Context, *gluedatabrew.CreateRecipeJobInput, ...request.Option) (*gluedatabrew.CreateRecipeJobOutput, error)
	CreateRecipeJobRequest(*gluedatabrew.CreateRecipeJobInput) (*request.Request, *gluedatabrew.CreateRecipeJobOutput)

	CreateSchedule(*gluedatabrew.CreateScheduleInput) (*gluedatabrew.CreateScheduleOutput, error)
	CreateScheduleWithContext(aws.Context, *gluedatabrew.CreateScheduleInput, ...request.Option) (*gluedatabrew.CreateScheduleOutput, error)
	CreateScheduleRequest(*gluedatabrew.CreateScheduleInput) (*request.Request, *gluedatabrew.CreateScheduleOutput)

	DeleteDataset(*gluedatabrew.DeleteDatasetInput) (*gluedatabrew.DeleteDatasetOutput, error)
	DeleteDatasetWithContext(aws.Context, *gluedatabrew.DeleteDatasetInput, ...request.Option) (*gluedatabrew.DeleteDatasetOutput, error)
	DeleteDatasetRequest(*gluedatabrew.DeleteDatasetInput) (*request.Request, *gluedatabrew.DeleteDatasetOutput)

	DeleteJob(*gluedatabrew.DeleteJobInput) (*gluedatabrew.DeleteJobOutput, error)
	DeleteJobWithContext(aws.Context, *gluedatabrew.DeleteJobInput, ...request.Option) (*gluedatabrew.DeleteJobOutput, error)
	DeleteJobRequest(*gluedatabrew.DeleteJobInput) (*request.Request, *gluedatabrew.DeleteJobOutput)

	DeleteProject(*gluedatabrew.DeleteProjectInput) (*gluedatabrew.DeleteProjectOutput, error)
	DeleteProjectWithContext(aws.Context, *gluedatabrew.DeleteProjectInput, ...request.Option) (*gluedatabrew.DeleteProjectOutput, error)
	DeleteProjectRequest(*gluedatabrew.DeleteProjectInput) (*request.Request, *gluedatabrew.DeleteProjectOutput)

	DeleteRecipeVersion(*gluedatabrew.DeleteRecipeVersionInput) (*gluedatabrew.DeleteRecipeVersionOutput, error)
	DeleteRecipeVersionWithContext(aws.Context, *gluedatabrew.DeleteRecipeVersionInput, ...request.Option) (*gluedatabrew.DeleteRecipeVersionOutput, error)
	DeleteRecipeVersionRequest(*gluedatabrew.DeleteRecipeVersionInput) (*request.Request, *gluedatabrew.DeleteRecipeVersionOutput)

	DeleteSchedule(*gluedatabrew.DeleteScheduleInput) (*gluedatabrew.DeleteScheduleOutput, error)
	DeleteScheduleWithContext(aws.Context, *gluedatabrew.DeleteScheduleInput, ...request.Option) (*gluedatabrew.DeleteScheduleOutput, error)
	DeleteScheduleRequest(*gluedatabrew.DeleteScheduleInput) (*request.Request, *gluedatabrew.DeleteScheduleOutput)

	DescribeDataset(*gluedatabrew.DescribeDatasetInput) (*gluedatabrew.DescribeDatasetOutput, error)
	DescribeDatasetWithContext(aws.Context, *gluedatabrew.DescribeDatasetInput, ...request.Option) (*gluedatabrew.DescribeDatasetOutput, error)
	DescribeDatasetRequest(*gluedatabrew.DescribeDatasetInput) (*request.Request, *gluedatabrew.DescribeDatasetOutput)

	DescribeJob(*gluedatabrew.DescribeJobInput) (*gluedatabrew.DescribeJobOutput, error)
	DescribeJobWithContext(aws.Context, *gluedatabrew.DescribeJobInput, ...request.Option) (*gluedatabrew.DescribeJobOutput, error)
	DescribeJobRequest(*gluedatabrew.DescribeJobInput) (*request.Request, *gluedatabrew.DescribeJobOutput)

	DescribeProject(*gluedatabrew.DescribeProjectInput) (*gluedatabrew.DescribeProjectOutput, error)
	DescribeProjectWithContext(aws.Context, *gluedatabrew.DescribeProjectInput, ...request.Option) (*gluedatabrew.DescribeProjectOutput, error)
	DescribeProjectRequest(*gluedatabrew.DescribeProjectInput) (*request.Request, *gluedatabrew.DescribeProjectOutput)

	DescribeRecipe(*gluedatabrew.DescribeRecipeInput) (*gluedatabrew.DescribeRecipeOutput, error)
	DescribeRecipeWithContext(aws.Context, *gluedatabrew.DescribeRecipeInput, ...request.Option) (*gluedatabrew.DescribeRecipeOutput, error)
	DescribeRecipeRequest(*gluedatabrew.DescribeRecipeInput) (*request.Request, *gluedatabrew.DescribeRecipeOutput)

	DescribeSchedule(*gluedatabrew.DescribeScheduleInput) (*gluedatabrew.DescribeScheduleOutput, error)
	DescribeScheduleWithContext(aws.Context, *gluedatabrew.DescribeScheduleInput, ...request.Option) (*gluedatabrew.DescribeScheduleOutput, error)
	DescribeScheduleRequest(*gluedatabrew.DescribeScheduleInput) (*request.Request, *gluedatabrew.DescribeScheduleOutput)

	ListDatasets(*gluedatabrew.ListDatasetsInput) (*gluedatabrew.ListDatasetsOutput, error)
	ListDatasetsWithContext(aws.Context, *gluedatabrew.ListDatasetsInput, ...request.Option) (*gluedatabrew.ListDatasetsOutput, error)
	ListDatasetsRequest(*gluedatabrew.ListDatasetsInput) (*request.Request, *gluedatabrew.ListDatasetsOutput)

	ListDatasetsPages(*gluedatabrew.ListDatasetsInput, func(*gluedatabrew.ListDatasetsOutput, bool) bool) error
	ListDatasetsPagesWithContext(aws.Context, *gluedatabrew.ListDatasetsInput, func(*gluedatabrew.ListDatasetsOutput, bool) bool, ...request.Option) error

	ListJobRuns(*gluedatabrew.ListJobRunsInput) (*gluedatabrew.ListJobRunsOutput, error)
	ListJobRunsWithContext(aws.Context, *gluedatabrew.ListJobRunsInput, ...request.Option) (*gluedatabrew.ListJobRunsOutput, error)
	ListJobRunsRequest(*gluedatabrew.ListJobRunsInput) (*request.Request, *gluedatabrew.ListJobRunsOutput)

	ListJobRunsPages(*gluedatabrew.ListJobRunsInput, func(*gluedatabrew.ListJobRunsOutput, bool) bool) error
	ListJobRunsPagesWithContext(aws.Context, *gluedatabrew.ListJobRunsInput, func(*gluedatabrew.ListJobRunsOutput, bool) bool, ...request.Option) error

	ListJobs(*gluedatabrew.ListJobsInput) (*gluedatabrew.ListJobsOutput, error)
	ListJobsWithContext(aws.Context, *gluedatabrew.ListJobsInput, ...request.Option) (*gluedatabrew.ListJobsOutput, error)
	ListJobsRequest(*gluedatabrew.ListJobsInput) (*request.Request, *gluedatabrew.ListJobsOutput)

	ListJobsPages(*gluedatabrew.ListJobsInput, func(*gluedatabrew.ListJobsOutput, bool) bool) error
	ListJobsPagesWithContext(aws.Context, *gluedatabrew.ListJobsInput, func(*gluedatabrew.ListJobsOutput, bool) bool, ...request.Option) error

	ListProjects(*gluedatabrew.ListProjectsInput) (*gluedatabrew.ListProjectsOutput, error)
	ListProjectsWithContext(aws.Context, *gluedatabrew.ListProjectsInput, ...request.Option) (*gluedatabrew.ListProjectsOutput, error)
	ListProjectsRequest(*gluedatabrew.ListProjectsInput) (*request.Request, *gluedatabrew.ListProjectsOutput)

	ListProjectsPages(*gluedatabrew.ListProjectsInput, func(*gluedatabrew.ListProjectsOutput, bool) bool) error
	ListProjectsPagesWithContext(aws.Context, *gluedatabrew.ListProjectsInput, func(*gluedatabrew.ListProjectsOutput, bool) bool, ...request.Option) error

	ListRecipeVersions(*gluedatabrew.ListRecipeVersionsInput) (*gluedatabrew.ListRecipeVersionsOutput, error)
	ListRecipeVersionsWithContext(aws.Context, *gluedatabrew.ListRecipeVersionsInput, ...request.Option) (*gluedatabrew.ListRecipeVersionsOutput, error)
	ListRecipeVersionsRequest(*gluedatabrew.ListRecipeVersionsInput) (*request.Request, *gluedatabrew.ListRecipeVersionsOutput)

	ListRecipeVersionsPages(*gluedatabrew.ListRecipeVersionsInput, func(*gluedatabrew.ListRecipeVersionsOutput, bool) bool) error
	ListRecipeVersionsPagesWithContext(aws.Context, *gluedatabrew.ListRecipeVersionsInput, func(*gluedatabrew.ListRecipeVersionsOutput, bool) bool, ...request.Option) error

	ListRecipes(*gluedatabrew.ListRecipesInput) (*gluedatabrew.ListRecipesOutput, error)
	ListRecipesWithContext(aws.Context, *gluedatabrew.ListRecipesInput, ...request.Option) (*gluedatabrew.ListRecipesOutput, error)
	ListRecipesRequest(*gluedatabrew.ListRecipesInput) (*request.Request, *gluedatabrew.ListRecipesOutput)

	ListRecipesPages(*gluedatabrew.ListRecipesInput, func(*gluedatabrew.ListRecipesOutput, bool) bool) error
	ListRecipesPagesWithContext(aws.Context, *gluedatabrew.ListRecipesInput, func(*gluedatabrew.ListRecipesOutput, bool) bool, ...request.Option) error

	ListSchedules(*gluedatabrew.ListSchedulesInput) (*gluedatabrew.ListSchedulesOutput, error)
	ListSchedulesWithContext(aws.Context, *gluedatabrew.ListSchedulesInput, ...request.Option) (*gluedatabrew.ListSchedulesOutput, error)
	ListSchedulesRequest(*gluedatabrew.ListSchedulesInput) (*request.Request, *gluedatabrew.ListSchedulesOutput)

	ListSchedulesPages(*gluedatabrew.ListSchedulesInput, func(*gluedatabrew.ListSchedulesOutput, bool) bool) error
	ListSchedulesPagesWithContext(aws.Context, *gluedatabrew.ListSchedulesInput, func(*gluedatabrew.ListSchedulesOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*gluedatabrew.ListTagsForResourceInput) (*gluedatabrew.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *gluedatabrew.ListTagsForResourceInput, ...request.Option) (*gluedatabrew.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*gluedatabrew.ListTagsForResourceInput) (*request.Request, *gluedatabrew.ListTagsForResourceOutput)

	PublishRecipe(*gluedatabrew.PublishRecipeInput) (*gluedatabrew.PublishRecipeOutput, error)
	PublishRecipeWithContext(aws.Context, *gluedatabrew.PublishRecipeInput, ...request.Option) (*gluedatabrew.PublishRecipeOutput, error)
	PublishRecipeRequest(*gluedatabrew.PublishRecipeInput) (*request.Request, *gluedatabrew.PublishRecipeOutput)

	SendProjectSessionAction(*gluedatabrew.SendProjectSessionActionInput) (*gluedatabrew.SendProjectSessionActionOutput, error)
	SendProjectSessionActionWithContext(aws.Context, *gluedatabrew.SendProjectSessionActionInput, ...request.Option) (*gluedatabrew.SendProjectSessionActionOutput, error)
	SendProjectSessionActionRequest(*gluedatabrew.SendProjectSessionActionInput) (*request.Request, *gluedatabrew.SendProjectSessionActionOutput)

	StartJobRun(*gluedatabrew.StartJobRunInput) (*gluedatabrew.StartJobRunOutput, error)
	StartJobRunWithContext(aws.Context, *gluedatabrew.StartJobRunInput, ...request.Option) (*gluedatabrew.StartJobRunOutput, error)
	StartJobRunRequest(*gluedatabrew.StartJobRunInput) (*request.Request, *gluedatabrew.StartJobRunOutput)

	StartProjectSession(*gluedatabrew.StartProjectSessionInput) (*gluedatabrew.StartProjectSessionOutput, error)
	StartProjectSessionWithContext(aws.Context, *gluedatabrew.StartProjectSessionInput, ...request.Option) (*gluedatabrew.StartProjectSessionOutput, error)
	StartProjectSessionRequest(*gluedatabrew.StartProjectSessionInput) (*request.Request, *gluedatabrew.StartProjectSessionOutput)

	StopJobRun(*gluedatabrew.StopJobRunInput) (*gluedatabrew.StopJobRunOutput, error)
	StopJobRunWithContext(aws.Context, *gluedatabrew.StopJobRunInput, ...request.Option) (*gluedatabrew.StopJobRunOutput, error)
	StopJobRunRequest(*gluedatabrew.StopJobRunInput) (*request.Request, *gluedatabrew.StopJobRunOutput)

	TagResource(*gluedatabrew.TagResourceInput) (*gluedatabrew.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *gluedatabrew.TagResourceInput, ...request.Option) (*gluedatabrew.TagResourceOutput, error)
	TagResourceRequest(*gluedatabrew.TagResourceInput) (*request.Request, *gluedatabrew.TagResourceOutput)

	UntagResource(*gluedatabrew.UntagResourceInput) (*gluedatabrew.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *gluedatabrew.UntagResourceInput, ...request.Option) (*gluedatabrew.UntagResourceOutput, error)
	UntagResourceRequest(*gluedatabrew.UntagResourceInput) (*request.Request, *gluedatabrew.UntagResourceOutput)

	UpdateDataset(*gluedatabrew.UpdateDatasetInput) (*gluedatabrew.UpdateDatasetOutput, error)
	UpdateDatasetWithContext(aws.Context, *gluedatabrew.UpdateDatasetInput, ...request.Option) (*gluedatabrew.UpdateDatasetOutput, error)
	UpdateDatasetRequest(*gluedatabrew.UpdateDatasetInput) (*request.Request, *gluedatabrew.UpdateDatasetOutput)

	UpdateProfileJob(*gluedatabrew.UpdateProfileJobInput) (*gluedatabrew.UpdateProfileJobOutput, error)
	UpdateProfileJobWithContext(aws.Context, *gluedatabrew.UpdateProfileJobInput, ...request.Option) (*gluedatabrew.UpdateProfileJobOutput, error)
	UpdateProfileJobRequest(*gluedatabrew.UpdateProfileJobInput) (*request.Request, *gluedatabrew.UpdateProfileJobOutput)

	UpdateProject(*gluedatabrew.UpdateProjectInput) (*gluedatabrew.UpdateProjectOutput, error)
	UpdateProjectWithContext(aws.Context, *gluedatabrew.UpdateProjectInput, ...request.Option) (*gluedatabrew.UpdateProjectOutput, error)
	UpdateProjectRequest(*gluedatabrew.UpdateProjectInput) (*request.Request, *gluedatabrew.UpdateProjectOutput)

	UpdateRecipe(*gluedatabrew.UpdateRecipeInput) (*gluedatabrew.UpdateRecipeOutput, error)
	UpdateRecipeWithContext(aws.Context, *gluedatabrew.UpdateRecipeInput, ...request.Option) (*gluedatabrew.UpdateRecipeOutput, error)
	UpdateRecipeRequest(*gluedatabrew.UpdateRecipeInput) (*request.Request, *gluedatabrew.UpdateRecipeOutput)

	UpdateRecipeJob(*gluedatabrew.UpdateRecipeJobInput) (*gluedatabrew.UpdateRecipeJobOutput, error)
	UpdateRecipeJobWithContext(aws.Context, *gluedatabrew.UpdateRecipeJobInput, ...request.Option) (*gluedatabrew.UpdateRecipeJobOutput, error)
	UpdateRecipeJobRequest(*gluedatabrew.UpdateRecipeJobInput) (*request.Request, *gluedatabrew.UpdateRecipeJobOutput)

	UpdateSchedule(*gluedatabrew.UpdateScheduleInput) (*gluedatabrew.UpdateScheduleOutput, error)
	UpdateScheduleWithContext(aws.Context, *gluedatabrew.UpdateScheduleInput, ...request.Option) (*gluedatabrew.UpdateScheduleOutput, error)
	UpdateScheduleRequest(*gluedatabrew.UpdateScheduleInput) (*request.Request, *gluedatabrew.UpdateScheduleOutput)
}

var _ GlueDataBrewAPI = (*gluedatabrew.GlueDataBrew)(nil)
