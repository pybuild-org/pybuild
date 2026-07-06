package docker

var MetaConfig = struct {
	Tag string `prop:"tag"`
}{}

var Targets = []*struct {
	Image struct {
		Base string `prop:"base"`
		Arch string `prop:"arch"`
	} `prop:"image"`

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
