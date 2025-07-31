package firefox

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/br8km/GoFox/utils"
)

type Config struct {

	Debug bool `toml:"Debug"`

	// Root dir of this project, auto resolved, Don't change this
	DirRoot string `toml:"DirRoot"`

	// profiles folder
	DirProfiles string `toml:"DirProfiles"`

	// Firefox Exeutable, eg `D:\Portable\FirefoxPortable\App\Firefox`
	ExePath string `toml:"ExePath"`

	// profile name = prefix + seperator + id
	Seperator string `toml:"Seperator"`

}

// create new Config
func NewConfig(
	debug bool,
	exepath string,
	dir_profiles string,
	seperator string,
) Config {

	if exepath == "" {
		log.Fatalf("parameter exepath empty")
	}

	if dir_profiles == "" {
		log.Fatalf("parameter dir_profiles empty")
	}

	if seperator == "" {
		log.Fatalf("parameter seperator empty")
	}

	var _, b, _, _ = runtime.Caller(0)
	root := filepath.Join(filepath.Dir(b), "../")

	return Config{
		Debug:     debug,
		DirRoot: 	root,
		Seperator:  seperator,
		ExePath:   exepath,
		DirProfiles: dir_profiles,
	}

}


func (c *Config) Check() (bool, error) {

	if c.Seperator == "" {
		return false, fmt.Errorf("config.Seperator CANNOT be empty")
	}

	dirs := []string{c.DirRoot, c.ToDirAssets(), c.DirProfiles}
	for _, d := range dirs {
		if !utils.IsDirExists(d){
			return false, fmt.Errorf("dir Not exists: `%s`", d)
		}
	}

	fs := []string{c.ExePath}
	for _, f := range fs {
		if !utils.IsFileExists(f) {
			return false, fmt.Errorf("file Not exists: `%s`", f)
		}

	}

	return true, nil
}

func (c *Config) ToDirAssets() string {
	return filepath.Join(
		c.DirRoot,
		"assets",
	)
}