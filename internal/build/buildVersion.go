package build

var (
	// Do not change, these variables are replaced by ldflags -X at build time
	gitVersion = "development"
	gitCommit  = "commit"
	date       = "1970-01-01T00:00:00Z"
)

type Version struct {
	Version   string
	Commit    string
	BuildDate string
}

func Get() *Version {
	return &Version{
		Version:   gitVersion,
		Commit:    gitCommit,
		BuildDate: date,
	}
}

func (v *Version) IsDevelopment() bool {
	return v.Version == "development"
}
