package vangogh_local_data

const (
	defaultRootDir = "/var/lib/vangogh"
)

var DefaultDirs = map[string]string{
	"backups":     defaultRootDir + "/backups",
	"downloads":   defaultRootDir + "/downloads",
	"images":      defaultRootDir + "/images",
	"input":       defaultRootDir + "/input",
	"items":       defaultRootDir + "/items",
	"logs":        "/var/log/vangogh",
	"metadata":    defaultRootDir + "/metadata",
	"output":      defaultRootDir + "/output",
	"recycle_bin": defaultRootDir + "/recycle_bin",
	"videos":      defaultRootDir + "/videos",
}
