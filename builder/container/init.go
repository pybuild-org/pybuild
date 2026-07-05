package container

var MetaConfig = struct {
	Tag string `prop:"tag"`
}{}

var Targets = []*struct {
	Base struct {
		Image string `prop:"image"`
	} `prop:"base"`

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
