package version

var (
	Version, GitCommit string
)

//BuildVersion buils version
func BuildVersion() string {
	if len(Version) == 0 {
		return "dev"
	}
	return Version
}
