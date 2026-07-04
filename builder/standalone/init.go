package standalone

var Targets = []*struct {
	Python struct {
		Arch string `prop:"arch"`
		OS   string `prop:"os"`
	} `prop:"python"`

	Pip struct {
		Platform  string   `prop:"platform"`
		Downloads []string `prop:"download"`
	} `prop:"pip"`

	Launcher struct {
		Run string `prop:"run"`
	} `prop:"launcher"`
}{}
