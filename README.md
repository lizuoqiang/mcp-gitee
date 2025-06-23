# Gitee MCP Server

Gitee MCP Server is a Model Context Protocol (MCP) server implementation for Gitee. It provides a set of tools for interacting with Gitee's API, allowing AI assistants to manage repositories, issues, pull requests, and more.

[![Install MCP Server](https://cursor.com/deeplink/mcp-install-dark.svg)](https://cursor.com/install-mcp?name=gitee&config=eyJjb21tYW5kIjoibnB4IC15IEBnaXRlZS9tY3AtZ2l0ZWVAbGF0ZXN0IiwiZW52Ijp7IkdJVEVFX0FDQ0VTU19UT0tFTiI6Ijx5b3VyIHBlcnNvbmFsIGFjY2VzcyB0b2tlbj4ifX0%3D)

## Available Tools

The server provides various tools for interacting with Gitee:

| Tool                                | Category | Description |
|-------------------------------------|----------|-------------|
| **list_user_repos**                 | Repository | List user authorized repositories |
| **get_file_content**                | Repository | Get the content of a file in a repository |
| **create_user_repo**                | Repository | Create a user repository |
| **create_org_repo**                 | Repository | Create an organization repository |
| **create_enter_repo**               | Repository | Create an enterprise repository |
| **fork_repository**                 | Repository | Fork a repository |
| **create_release**                  | Repository | Create a release for a repository |
| **list_releases**                   | Repository | List repository releases |
| **search_open_source_repositories** | Repository | Search open source repositories on Gitee |
| **create_branch**                   | Repository | Create a branch |
| **list_repo_pulls**                 | Pull Request | List pull requests in a repository |
| **merge_pull**                      | Pull Request | Merge a pull request |
| **create_pull**                     | Pull Request | Create a pull request |
| **update_pull**                     | Pull Request | Update a pull request |
| **get_pull_detail**                 | Pull Request | Get details of a pull request |
| **comment_pull**                    | Pull Request | Comment on a pull request |
| **list_pull_comments**              | Pull Request | List all comments for a pull request |
| **create_issue**                    | Issue | Create an issue |
| **update_issue**                    | Issue | Update an issue |
| **get_repo_issue_detail**           | Issue | Get details of a repository issue |
| **list_repo_issues**                | Issue | List repository issues |
| **comment_issue**                   | Issue | Comment on an issue |
| **list_issue_comments**             | Issue | List comments on an issue |
| **get_user_info**                   | User | Get current authenticated user information |
| **search_users**                    | User | Search for users |
| **list_user_notifications**         | Notification | List user notifications |


## Features

- Interact with Gitee repositories, issues, pull requests, and notifications
- Configurable API base URL to support different Gitee instances
- Command-line flags for easy configuration
- Supports both personal, organization, and enterprise operations
- Dynamic toolset enable/disable

## Installation

### Prerequisites

- Go 1.23.0 or higher
- Gitee account with an access token, [Go to get](https://gitee.com/profile/personal_access_tokens)

### Building from Source

1. Clone the repository:
   ```bash
   git clone https://github.com/lizuoqiang/mcp-gitee.git
   cd mcp-gitee
   ```

2. Build the project:
   ```bash
   make build
   ```
   Move ./bin/mcp-gitee PATH env

### Use go install
   ```bash
   go install gitee.com/oschina/mcp-gitee@latest
   ```

## Usage

Check mcp-gitee version:

```bash
mcp-gitee --version
```