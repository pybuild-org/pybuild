package job

func Register(name string, job func()) {
	jobs[name] = job
}
