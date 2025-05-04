package builder

type Project struct {
	Meta     *ProjectMeta
	HTMLBody string
}

type ProjectMeta struct {
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Date        string `json:"date"`
	Description string `json:"description"`
	CoverImage  string `json:"coverImage"`
}

type StaticPage struct {
	Slug string `json:"slug"`
	Path string `json:"path"`
}
