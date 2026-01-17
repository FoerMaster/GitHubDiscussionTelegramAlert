package githubEnums

type GitHubEvent string
type GitHubAction string

const (
	PING               GitHubEvent = "ping"
	DISCUSSION         GitHubEvent = "discussion"
	DISCUSSION_COMMENT GitHubEvent = "discussion_comment"

	CREATED GitHubAction = "created"
	EDITED  GitHubAction = "edited"
	DELETED GitHubAction = "deleted"
)
