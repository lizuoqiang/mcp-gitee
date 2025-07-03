package repository

import (
	"context"
	"fmt"

	"github.com/lizuoqiang/mcp-gitee/operations/types"
	"github.com/lizuoqiang/mcp-gitee/utils"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	CreateTagToolName = "create_tag"
)

var CreateTagTool = mcp.NewTool(
	CreateTagToolName,
	mcp.WithDescription("Create a tag"),
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
		"tag_name",
		mcp.Description("New tag name"),
		mcp.Required(),
	),
	mcp.WithString(
		"tag_message",
		mcp.Description("tag description"),
		mcp.DefaultString("发版"),
	),
)

func CreateTagHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args, _ := utils.ConvertArgumentsToMap(request.Params.Arguments)
	if checkResult, err := utils.CheckRequired(args, "owner", "repo", "tag_name", "refs"); err != nil {
		return checkResult, err
	}
	owner := args["owner"].(string)
	repo := args["repo"].(string)

	apiUrl := fmt.Sprintf("/repos/%s/%s/tags", owner, repo)
	giteeClient := utils.NewGiteeClient("POST", apiUrl, utils.WithContext(ctx), utils.WithPayload(request.Params.Arguments))

	data := &types.TagItem{}
	return giteeClient.HandleMCPResult(data)
}
