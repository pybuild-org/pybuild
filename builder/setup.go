package builder

import "path/filepath"

func SetupBuilder() {
	CleanDir(TempDir, false)
	CleanDir(BuilderConfig.Output, false)
}

func SetupPython() {
	url := MkPyDownloadUrl(
		BuilderConfig.Version,
		BuilderConfig.Release,
		PythonConfig.Arch,
		PythonConfig.OS,
	)

	func() {
		s := GetDownloadStream(url)
		defer s.Close()
		Decompress(s, TempDir)
	}()

	Python = filepath.Join(TempDir, "python", MkPyBinPath(
		PythonConfig.OS,
		BuilderConfig.Version,
	))

	RunCommand(Python, "-m", "ensurepip")
}
