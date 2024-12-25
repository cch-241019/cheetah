package version

/*
* @author: Chen Chiheng
* @date: 2024/12/25 20:40:02
* @description:
**/

var (
	gitVersion   string = "v0.0.0-master+$Format:%H$"
	gitCommit    string = "$Format:%H$"
	gitTreeState string = ""
	buildDate    string = "1970-01-01T00:00:00Z"
)

type Info struct {
	Major        string
	Minor        string
	GitVersion   string
	GitCommit    string
	GitTreeState string
	BuildDate    string
	GoVersion    string
	Compiler     string
	Platform     string
}

func (i Info) String() string {
	return i.GitVersion
}
