package funcjob

var jobs map[string]func()

func init() {
	jobs = make(map[string]func())
}
