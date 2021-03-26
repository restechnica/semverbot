package config

type Loader interface {
	Load(path string) (config Root, err error)
	Overload(path string, config Root) (Root, error)
}
