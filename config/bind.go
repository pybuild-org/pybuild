package config

func Bind(name string, target any) {
	cfg[name] = target
}
