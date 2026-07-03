package funcjob

func Run(name string) {
	job, ok := jobs[name]
	if !ok {
		return
	}

	job()
}
