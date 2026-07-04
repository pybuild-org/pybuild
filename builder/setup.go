package builder

var BuilderConfig = struct {
	Source  string `prop:"source"`
	Output  string `prop:"output"`
	Release string `prop:"release"`
	Version string `prop:"version"`
}{}

func Setup() {
	cleanDir(".pybuild", false)
	cleanDir(BuilderConfig.Output, false)
}
