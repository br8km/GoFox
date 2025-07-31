package tests

// test extension Chameleon Settings manipulation

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	gofox "github.com/br8km/GoFox"
	"github.com/br8km/GoFox/chameleon"
)


func TestChameleon(t *testing.T) {
	// Passed @ 20250520

	var c gofox.Config
	var cs chameleon.Settings

	var ok bool
	var err error
	var debug bool = true

	assert.Equal(t, 1, 1)
	c = gofox.NewConfig(debug, EXE_PATH, TestProfilesFolder, gofox.PROFILE_SEPERATOR)
	c.DirProfiles = filepath.Join(
		c.DirRoot,
		TestFolder,
		TestProfilesFolder,
	)
	fmt.Println(c.DirProfiles)
	dir_profile := filepath.Join(
		c.DirProfiles, 
		TestDemoProfileName,
	)

	// test load
	cs, err = chameleon.Load(dir_profile)
	assert.Nil(t, err)

	// test Enable
	cs.SetEnable(true)
	assert.True(t, cs.GetEnable())

	// test Theme
	theme := "dark"
	cs.SetTheme(theme)
	assert.Equal(t, cs.GetTheme(), theme)

	// test Language
	on := true
	lang := "en-US"
	cs.SetLanguage(true, lang)
	on, lang = cs.GetLanguage() 
	assert.True(t, on)
	assert.Equal(t, lang, "en-US")

	// test TimeZone
	tz := "US/Central"
	cs.SetTimezone(tz)
	assert.Equal(t, tz, cs.GetTimezone())

	// test Profile
	o, b := "win4" ,"gcr"
	cs.SetProfile(o, b)
	oo, bb := cs.GetProfile()
	assert.Equal(t, o, oo)
	assert.Equal(t, b, bb)

	// test save
	ok, err = cs.Save(dir_profile, true)
	assert.True(t, ok)
	assert.Nil(t, err)

}