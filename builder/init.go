package builder

var (
	TempDir string = ".pybuild"
	Python  string
)

var BuilderConfig = struct {
	Source  string `prop:"source"`
	Output  string `prop:"output"`
	Version string `prop:"version"`
	Release string `prop:"release"`
}{}

var PythonConfig = struct {
	Arch string `prop:"arch"`
	OS   string `prop:"os"`
}{}
