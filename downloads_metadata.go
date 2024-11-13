package vangogh_local_data

type DownloadMetadata struct {
	Id            string         `json:"id"`
	Slug          string         `json:"slug"`
	Title         string         `json:"title"`
	DownloadLinks []DownloadLink `json:"download-links,omitempty"`
}

type DownloadLink struct {
	ManualUrl      string `json:"manual-url"`
	LocalFilename  string `json:"local-filename"`
	Md5            string `json:"md5"`
	OS             string `json:"os"`
	Type           string `json:"type"`
	LanguageCode   string `json:"language-code"`
	Version        string `json:"version"`
	EstimatedBytes int    `json:"estimated-bytes"`
}
