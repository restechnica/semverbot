package config

type Env struct {
	Files   []string          `yaml:"files,omitempty"`
	Scripts []EnvScript       `yaml:"scripts,omitempty"`
	Vars    map[string]string `yaml:"vars,omitempty"`
}

func NewEnv() Env {
	return Env{
		Files:   []string{},
		Scripts: []EnvScript{},
		Vars:    map[string]string{},
	}
}

type EnvScript struct {
	Bin  string `yaml:"bin,omitempty"`
	Path string `yaml:"path,omitempty"`
}

type Git struct {
	Config    map[string]string `yaml:"config,omitempty"`
	Unshallow bool              `yaml:"unshallow,omitempty"`
}

func NewGit() Git {
	return Git{
		Config:    map[string]string{},
		Unshallow: true,
	}
}

type Root struct {
	Env    Env    `yaml:"env,omitempty"`
	Git    Git    `yaml:"git,omitempty"`
	Semver Semver `yaml:"semver,omitempty"`
}

func NewRoot() (root Root) {
	return Root{
		Env:    NewEnv(),
		Git:    NewGit(),
		Semver: NewSemver(),
	}
}

type Semver struct {
	Bin      string            `yaml:"bin,omitempty"`
	Strategy string            `yaml:"strategy,omitempty"`
	Matches  map[string]string `yaml:"matches,omitempty"`
	Path     string            `yaml:"path,omitempty"`
}

func NewSemver() Semver {
	return Semver{Strategy: "auto", Matches: map[string]string{}}
}
