package vangogh_local_data

const (
	defaultRootDir = "/var/lib/vangogh"
)

var DefaultDirs = map[string]string{
	"backups":      defaultRootDir + "/backups",
	"downloads":    defaultRootDir + "/downloads",
	"images":       defaultRootDir + "/images",
	"input_files":  defaultRootDir,
	"items":        defaultRootDir + "/items",
	"logs":         "/var/log/vangogh",
	"metadata":     defaultRootDir + "/metadata",
	"output_files": defaultRootDir,
	"recycle_bin":  defaultRootDir + "/recycle_bin",
	"videos":       defaultRootDir + "/videos",
}
