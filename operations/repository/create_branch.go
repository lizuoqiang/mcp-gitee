package repository

import (
	"context"
	"fmt"

	"github.com/lizuoqiang/mcp-gitee/operations/types"
	"github.com/lizuoqiang/mcp-gitee/utils"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	CreateBranchToolName = "create_branch"
)

var CreateBranchTool = mcp.NewTool(
	CreateBranchToolName,
	mcp.WithDescription("Create a branch"),
	mcp.WithString(
		"owner",
		mcp.Description("The space address to which the repository belongs (the address path of the enterprise, organization or individual)"),
		mcp.Required(),
	),
	mcp.WithString(
		"repo",
		mcp.Description("The path of the repository"),
		mcp.Required(),
	),
	mcp.WithString(
		"refs",
		mcp.Description("Source branch name"),
		mcp.DefaultString("master"),
	),
	mcp.WithString(
		"branch_name",
		mcp.Description("New branch name"),
		mcp.Required(),
	),
)

func CreateBranchHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args, _ := utils.ConvertArgumentsToMap(request.Params.Arguments)
	if checkResult, err := utils.CheckRequired(args, "owner", "repo", "branch_name"); err != nil {
		return checkResult, err
	}
	owner := args["owner"].(string)
	repo := args["repo"].(string)

	apiUrl := fmt.Sprintf("/repos/%s/%s/branches", owner, repo)
	giteeClient := utils.NewGiteeClient("POST", apiUrl, utils.WithContext(ctx), utils.WithPayload(request.Params.Arguments))

	data := &types.BranchItem{}
	return giteeClient.HandleMCPResult(data)
}
