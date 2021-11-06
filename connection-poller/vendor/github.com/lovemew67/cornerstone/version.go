package cornerstone

import (
	"runtime"
)

var (
	appName    = "unknown"
	appVersion = "unknown"
	gitBranch  = "master"
	gitCommit  = "$Format:%H$"
	buildDate  = "1970-01-01T00:00:00Z"
	goVer      = runtime.Version()
	goOS       = runtime.GOOS
	goArch     = runtime.GOARCH
)

type Version struct {
	AppName    string `json:"app_name"`
	AppVersion string `json:"app_version"`
	GitBranch  string `json:"git_branch"`
	GitCommit  string `json:"git_commit"`
	BuildDate  string `json:"build_date"`
	GoOS       string `json:"go_os"`
	GoArch     string `json:"go_arch"`
	GoVer      string `json:"go_ver"`
}

func GetVersion() (result Version) {
	result = Version{
		appName,
		appVersion,
		gitBranch,
		gitCommit,
		buildDate,
		goOS,
		goArch,
		goVer,
	}
	return
}

func GetAppName() (result string) {
	result = appName
	return
}
