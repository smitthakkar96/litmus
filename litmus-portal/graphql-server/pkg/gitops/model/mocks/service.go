// Package mocks ...
package mocks

import (
	"context"

	"github.com/litmuschaos/litmus/litmus-portal/graphql-server/graph/model"
	dbSchemaCluster "github.com/litmuschaos/litmus/litmus-portal/graphql-server/pkg/database/mongodb/cluster"
	"github.com/stretchr/testify/mock"
)

// GitOpsService is an autogenerated mock type for the GitOpsService type
type GitOpsService struct {
	mock.Mock
}

// GitOpsNotificationHandler provides a mock function with given fields: ctx, cluster, workflowID
func (g *GitOpsService) GitOpsNotificationHandler(ctx context.Context, cluster *dbSchemaCluster.Cluster, workflowID string) (string, error) {
	args := g.Called(ctx, cluster, workflowID)
	return args.String(0), args.Error(1)
}

// EnableGitOpsHandler provides a mock function with given fields: ctx, config
func (g *GitOpsService) EnableGitOpsHandler(ctx context.Context, config model.GitConfig) (bool, error) {
	args := g.Called(ctx, config)
	return args.Bool(0), args.Error(1)
}

// DisableGitOpsHandler provides a mock function with given fields: ctx, projectID
func (g *GitOpsService) DisableGitOpsHandler(ctx context.Context, projectID string) (bool, error) {
	args := g.Called(ctx, projectID)
	return args.Bool(0), args.Error(1)
}

// UpdateGitOpsDetailsHandler provides a mock function with given fields: ctx, config
func (g *GitOpsService) UpdateGitOpsDetailsHandler(ctx context.Context, config model.GitConfig) (bool, error) {
	args := g.Called(ctx, config)
	return args.Bool(0), args.Error(1)
}

// GetGitOpsDetails provides a mock function with given fields: ctx, projectID
func (g *GitOpsService) GetGitOpsDetails(ctx context.Context, projectID string) (*model.GitConfigResponse, error) {
	args := g.Called(ctx, projectID)
	return args.Get(0).(*model.GitConfigResponse), args.Error(1)
}

// UpsertWorkflowToGit provides a mock function with given fields: ctx, workflow
func (g *GitOpsService) UpsertWorkflowToGit(ctx context.Context, workflow *model.ChaosWorkFlowRequest) error {
	args := g.Called(ctx, workflow)
	return args.Error(0)
}

// DeleteWorkflowFromGit provides a mock function with given fields: ctx, workflow
func (g *GitOpsService) DeleteWorkflowFromGit(ctx context.Context, workflow *model.ChaosWorkFlowRequest) error {
	args := g.Called(ctx, workflow)
	return args.Error(0)
}

// GitOpsSyncHandler provides a mock function with given fields: singleRun
func (g *GitOpsService) GitOpsSyncHandler(singleRun bool) {
	g.Called(singleRun)
}
