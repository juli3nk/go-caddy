package plugins

type API struct {
	LatestCaddy LatestCaddy `json:"latest_caddy"`
	Plugins     Plugins     `json:"plugins"`
}

type LatestCaddy struct {
	Version    string
	Timestamp  string
	ReleasedBy string
}

type Plugins []Plugin

type Plugin struct {
	ID             string
	OwnerAccountID string
	Name           string
	Type           Type
	Description    string
	Website        string
	Support        string
	Docs           string
	Examples       []Example
	ImportPath     string
	SourceRepo     string
	ReleaseBranch  string
	Releases       []Release
	Published      string
	Updated        string
	DownloadCount  int
	Hidden         bool
	OwnerName      string
	LastUpdate     string
	Required       bool
}

type Type struct {
	ID            string
	Name          string
	CategoryTitle string
	Description   string
	Package       string
	Function      string
}

type Example struct {
	Title       string
	Code        string
	Explanation string
}

type Release struct {
	Version         string
	CaddyVersion    string
	Timestamp       string
	TestedPlatforms string
}
