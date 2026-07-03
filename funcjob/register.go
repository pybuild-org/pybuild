package funcjob

func Register(name string, job func()) {
	jobs[name] = job
}
