package chameleon

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/goccy/go-json"
)

// Settings load from local file
func Load(dir_profile string) (Settings, error) {
	var cs Settings
	fp := filepath.Join(dir_profile, "chameleon_settings.json")
	b, err := os.ReadFile(fp)
	if err != nil {
		return cs, fmt.Errorf("failed to read file: `%s` @ %v", fp, err)
	}
	err = json.NewDecoder(bytes.NewBuffer(b)).Decode(&cs)
	if err != nil {
		return cs, fmt.Errorf("failed to decode file: `%s` @ %v", fp, err)
	}
	// fmt.Printf("%#v", s) // debug
	return cs, nil
}

// Settings save to local file
func (cs *Settings) Save(
	dir_profile string,
	indent bool,
) (bool, error) {
	var b []byte
	var err error

	if indent {
		b, err = json.MarshalIndent(&cs, "", "\t")
	} else {
		b, err = json.Marshal(&cs)
	}

	if err != nil {
		return false, err
	}

	f := filepath.Join(dir_profile, "chameleon_settings.json")
	err = os.WriteFile(f, b, os.ModePerm)
	if err != nil {
		return false, err
	}
	return true, nil

}

// Settings delete local file
func (cs *Settings) Delete(fp string) error {
	return os.Remove(fp)
}

func (cs *Settings) SetEnable(on bool) *Settings {
	cs.Config.Enabled = on
	return cs
}

func (cs *Settings) GetEnable() bool {
	return cs.Config.Enabled
}

func (cs *Settings) SetTheme(name string) *Settings {
	if !slices.Contains(Theme_Codes[:], name) {
		log.Fatalf("Chameleon Settings Theme: `%s` not allowed", name)
	}
	cs.Config.Theme = name
	return cs
}

func (cs *Settings) GetTheme() string {
	return cs.Config.Theme
}

func (cs *Settings) SetLanguage(enabled bool, code string) *Settings {
	if !slices.Contains(Lang_Codes[:], code) {
		log.Fatalf("Chameleon Settings Language: `%s` not allowed", code)
	}
	cs.Headers.SpoofAcceptLang.Value = code
	cs.Headers.SpoofAcceptLang.Enabled = enabled
	return cs
}

func (cs *Settings) GetLanguage() (bool, string) {
	b := cs.Headers.SpoofAcceptLang.Enabled
	c := cs.Headers.SpoofAcceptLang.Value
	return b, c
}

func (cs *Settings) SetTimezone(code string) *Settings {
	codes := []string{}
	for _, tz := range CHAMELEON_TimeZones {
		codes = append(codes, tz.Zone)
	}
	if !slices.Contains(codes[:], code) {
		log.Fatalf("Chameleon Settings TimeZone: `%s` not allowed", code)
	}
	cs.Options.TimeZone = code
	return cs
}

func (cs *Settings) GetTimezone() string {
	return cs.Options.TimeZone
}

// set Chameleon_Settings.Profile
func (cs *Settings) SetProfile(
	os_code string, 
	br_code string,
) *Settings {
	if _, ok := CHAMELEON_OS[os_code]; !ok {
		log.Fatalf("chameleon Profile error os_code: %s", os_code)
	}

	if _, ok := CHAMELEON_BROWSER[br_code]; !ok {
		log.Fatalf("chameleon Profile error br_code: %s", br_code)
	} 
	
	cs.Profile.Selected = fmt.Sprintf("%s-%s", os_code, br_code)
	return cs
}

// set profile
func (cs *Settings) GetProfile() (string, string) {
	s := strings.Split(cs.Profile.Selected, "-")
	if len(s) != 2{
		log.Fatalf("chameleon get Profile error: split")
	}
	o, b := s[0], s[1]
	if o == "" || b == "" {
		log.Fatalf("chameleon get Profile error: empty")
	}
	return o, b
}

func (cs *Settings) SetScreenSize(size string) *Settings {
	// check if allowed...?
	cs.Options.ScreenSize = size
	return cs
}

func (cs *Settings) GetScreenSize() string {
	return cs.Options.ScreenSize
}
