package builder

type Project struct {
	Meta     *ProjectMeta
	HTMLBody string
}

type ProjectMeta struct {
	Title       string   `json:"title"`
	Slug        string   `json:"slug"`
	Date        string   `json:"date"`
	Description string   `json:"description"`
	CoverImage  string   `json:"coverImage"`
	Read        string   `json:"read"`
	Sponsors    []string `json:"sponsors"`
}

type StaticPage struct {
	Slug    string `json:"slug"`
	Path    string `json:"path"`
	LastMod string `json:"lastMod"`
}

type SitemapEntry struct {
	Loc        string
	LastMod    string
	ChangeFreq string
	Priority   string
}
