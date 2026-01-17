package models

import "time"

type GitHubWebhook struct {
	Action     string      `json:"action"`
	Discussion *Discussion `json:"discussion,omitempty"`
	Comment    *Comment    `json:"comment,omitempty"`
	Repository Repository  `json:"repository"`
	Sender     User        `json:"sender"`
	Hook       *Hook       `json:"hook,omitempty"`
	Zen        string      `json:"zen,omitempty"`
}

type Discussion struct {
	ID                int                `json:"id"`
	NodeID            string             `json:"node_id"`
	Number            int                `json:"number"`
	Title             string             `json:"title"`
	Body              string             `json:"body"`
	User              User               `json:"user"`
	State             string             `json:"state"`
	StateReason       *string            `json:"state_reason"`
	Locked            bool               `json:"locked"`
	Comments          int                `json:"comments"`
	CreatedAt         time.Time          `json:"created_at"`
	UpdatedAt         time.Time          `json:"updated_at"`
	AuthorAssociation string             `json:"author_association"`
	ActiveLockReason  *string            `json:"active_lock_reason"`
	Category          DiscussionCategory `json:"category"`
	AnswerHtmlURL     *string            `json:"answer_html_url"`
	AnswerChosenAt    *string            `json:"answer_chosen_at"`
	AnswerChosenBy    *User              `json:"answer_chosen_by"`
	HTMLURL           string             `json:"html_url"`
	TimelineURL       string             `json:"timeline_url"`
	RepositoryURL     string             `json:"repository_url"`
	Reactions         Reactions          `json:"reactions"`
	Labels            []interface{}      `json:"labels"`
}

type DiscussionCategory struct {
	ID           int       `json:"id"`
	NodeID       string    `json:"node_id"`
	RepositoryID int       `json:"repository_id"`
	Emoji        string    `json:"emoji"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Slug         string    `json:"slug"`
	IsAnswerable bool      `json:"is_answerable"`
}

type Comment struct {
	ID                int       `json:"id"`
	NodeID            string    `json:"node_id"`
	HTMLURL           string    `json:"html_url"`
	User              User      `json:"user"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	AuthorAssociation string    `json:"author_association"`
	Body              string    `json:"body"`
	Reactions         Reactions `json:"reactions"`
}

type Repository struct {
	ID                       int       `json:"id"`
	NodeID                   string    `json:"node_id"`
	Name                     string    `json:"name"`
	FullName                 string    `json:"full_name"`
	Private                  bool      `json:"private"`
	Owner                    User      `json:"owner"`
	HTMLURL                  string    `json:"html_url"`
	Description              string    `json:"description"`
	Fork                     bool      `json:"fork"`
	URL                      string    `json:"url"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
	PushedAt                 time.Time `json:"pushed_at"`
	GitURL                   string    `json:"git_url"`
	SSHURL                   string    `json:"ssh_url"`
	CloneURL                 string    `json:"clone_url"`
	SVNURL                   string    `json:"svn_url"`
	Homepage                 string    `json:"homepage"`
	Size                     int       `json:"size"`
	StargazersCount          int       `json:"stargazers_count"`
	WatchersCount            int       `json:"watchers_count"`
	Language                 string    `json:"language"`
	HasIssues                bool      `json:"has_issues"`
	HasProjects              bool      `json:"has_projects"`
	HasDownloads             bool      `json:"has_downloads"`
	HasWiki                  bool      `json:"has_wiki"`
	HasPages                 bool      `json:"has_pages"`
	HasDiscussions           bool      `json:"has_discussions"`
	ForksCount               int       `json:"forks_count"`
	Archived                 bool      `json:"archived"`
	Disabled                 bool      `json:"disabled"`
	OpenIssuesCount          int       `json:"open_issues_count"`
	License                  *string   `json:"license"`
	AllowForking             bool      `json:"allow_forking"`
	IsTemplate               bool      `json:"is_template"`
	WebCommitSignoffRequired bool      `json:"web_commit_signoff_required"`
	Topics                   []string  `json:"topics"`
	Visibility               string    `json:"visibility"`
	Forks                    int       `json:"forks"`
	OpenIssues               int       `json:"open_issues"`
	Watchers                 int       `json:"watchers"`
	DefaultBranch            string    `json:"default_branch"`
	MirrorURL                *string   `json:"mirror_url"`
}

type User struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	UserViewType      string `json:"user_view_type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type Reactions struct {
	URL        string `json:"url"`
	TotalCount int    `json:"total_count"`
	PlusOne    int    `json:"+1"`
	MinusOne   int    `json:"-1"`
	Laugh      int    `json:"laugh"`
	Hooray     int    `json:"hooray"`
	Confused   int    `json:"confused"`
	Heart      int    `json:"heart"`
	Rocket     int    `json:"rocket"`
	Eyes       int    `json:"eyes"`
}

type Hook struct {
	Type          string       `json:"type"`
	ID            int          `json:"id"`
	Name          string       `json:"name"`
	Active        bool         `json:"active"`
	Events        []string     `json:"events"`
	Config        HookConfig   `json:"config"`
	UpdatedAt     time.Time    `json:"updated_at"`
	CreatedAt     time.Time    `json:"created_at"`
	URL           string       `json:"url"`
	TestURL       string       `json:"test_url"`
	PingURL       string       `json:"ping_url"`
	DeliveriesURL string       `json:"deliveries_url"`
	LastResponse  LastResponse `json:"last_response"`
}

type HookConfig struct {
	ContentType string `json:"content_type"`
	InsecureSSL string `json:"insecure_ssl"`
	URL         string `json:"url"`
}

type LastResponse struct {
	Code    *int    `json:"code"`
	Status  string  `json:"status"`
	Message *string `json:"message"`
}
