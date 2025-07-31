package tests

import (
	"fmt"
	"log"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/br8km/gofox"
	"github.com/br8km/gofox/utils"
)


func InitConfig() gofox.Config {

	c := gofox.NewConfig(
		true,
		EXE_PATH,
		"placeholder",
		gofox.PROFILE_SEPERATOR,
	)
	c.DirProfiles = filepath.Join(
		c.DirRoot, 
		TestFolder,
		TestProfilesFolder,
	)
	return c

}




// test config
func TestConfig(t *testing.T) {
	// Passed @ 20250731

	var c1, c2 gofox.Config
	var ok bool
	var err error

	// default
	c1 = InitConfig()
	// check pass
	ok, err = c1.Check()
	assert.True(t, ok)
	assert.Nil(t, err)

	dir_profiles_temp := filepath.Join(
		c1.DirRoot, TestFolder, TestProfilesFolderTemp,
	)
	c2 = InitConfig()
	c2.DirProfiles = dir_profiles_temp
	// remove if exists already
	ok, err = utils.DirRemove(dir_profiles_temp)
	assert.True(t, ok)
	assert.Nil(t, err)
	// check fail
	ok, err = c2.Check()
	assert.False(t, ok)
	assert.NotNil(t, err)

	// create profiles folder
	ok, err = utils.DirCreate(dir_profiles_temp)
	assert.True(t, ok)
	assert.Nil(t, err)
	// check pass
	ok, err = c2.Check()
	assert.True(t, ok)
	assert.Nil(t, err)

	// clean up temp folder
	ok, err = utils.DirRemove(dir_profiles_temp)
	assert.True(t, ok)
	assert.Nil(t, err)


}

// test asset
// prepare asset files before testing
func TestAsset(t *testing.T) {
	// Passed @ 20250731

	var c gofox.Config
	var assets []gofox.Asset

	var ok bool
	var err error

	c = InitConfig()

	c.DirProfiles = filepath.Join(
		c.DirRoot, TestFolder, TestProfilesFolder,
	)
	dir_profile_demo := filepath.Join(c.DirProfiles, TestDemoProfileName)

	ok, err = c.Check()
	assert.True(t, ok)
	assert.Nil(t, err)

	assets = gofox.DefaultAssets(
		filepath.Dir(c.ExePath),
		dir_profile_demo,
	)

	for _, asset := range assets {
		ok, err = asset.Check()
		assert.True(t, ok)
		assert.Nil(t, err)
	}
}

// test fingerprints
// func TestFingerPrints(t *testing.T) {}

func init_manager() gofox.FPManager {
	// Passed @ 20250731

	var c gofox.Config
	var assets []gofox.Asset
	var fps []gofox.FingerPrint
	var m gofox.FPManager

	c = InitConfig()

	c.DirProfiles = filepath.Join(
		c.DirRoot, TestFolder, TestProfilesFolder,
	)
	dir_profile_demo := filepath.Join(c.DirProfiles, TestDemoProfileName)

	assets = gofox.DefaultAssets(
		filepath.Dir(c.ExePath),
		dir_profile_demo,
	)

	fps = []gofox.FingerPrint{}

	m = gofox.FPManager{
		Config: c, 
		Assets:  assets, 
		FingerPrints: fps,
	}

	_, err := m.Check()
	if err != nil {
		log.Fatalf("init manager failed: %s", err)
	}
	return m
}

func create_fp(m *gofox.FPManager) gofox.FingerPrint {
	var fp gofox.FingerPrint
	var err error

	// create fp
	group := "group"
	id := "id"
	lang := "en-US"
	proxy_url := "http://usr:pwd@1.2.3.4:8080"
	os_family := "win"
	br_family := "Firefox"
	device_type := "desktop"

	fp, err = m.CreateFP(
		group,
		id,
		lang,
		proxy_url,
		os_family,
		br_family,
		device_type,
	)
	if err != nil {
		log.Fatalf("create fp failed: %s", err)
	}
	return fp

}

// test management
func TestManager(t *testing.T) {
	// Passed @ 20250731

	var ok bool 
	var err error

	group := "group"
	id := "id"
	lang := "en-US"
	proxy_url := "http://usr:pwd@1.2.3.4:8080"
	os_family := "win"
	br_family := "Firefox"
	device_type := "desktop"

	m := init_manager()
	fp := create_fp(&m)

	assert.Equal(t, fp.Group, group)
	assert.Equal(t, fp.Id, id)
	// assert.Equal(t, fp.Person.Male, male)
	// assert.Equal(t, fp.Person.Age, age)
	assert.Equal(t, fp.Device.Language, lang)
	assert.Equal(t, fp.Location.Proxy.ToUrl(), proxy_url)
	assert.Equal(t, fp.Device.OS.Family, os_family)
	assert.Equal(t, fp.Device.Browser.Family, br_family)
	assert.Equal(t, fp.Device.Type, device_type)

	// test UpdateFP/GetFP
	m.UpdateFP(&fp, true)
	fp_got, err := m.GetFP(fp.Id)
	assert.Equal(t, fp_got.Id, fp.Id)
	assert.Nil(t, err)

	// test FillProfile
	ok, err = m.FillProfile(&fp)
	fmt.Printf("FillProfile ok:\n%#v\n", ok)
	fmt.Printf("FillProfile err:\n%v\n", err)
	assert.True(t, ok)
	assert.Nil(t, err)

	// test SaveFPs
	ff := filepath.Join(m.Config.DirRoot, TestFolder, TestFPsFileName)
	ok, err = m.SaveFPs(ff, true)
	assert.True(t, ok)
	assert.Nil(t, err)
	assert.FileExists(t, ff)

	// test LoadFPs
	fps_got, err := m.LoadFPs(ff)
	assert.Nil(t, err)
	fmt.Printf("fps_got: \n%#v\n", fps_got)
	assert.Equal(t, len(fps_got), 1)
	assert.Equal(t, fps_got[0].Id, m.FingerPrints[0].Id)

	// test RemoveProfile
	ok, err = m.RemoveProfile(&fp)
	assert.True(t, ok)
	assert.Nil(t, err)

}

func TestStartProfile(t *testing.T) {
	// Passed @ 20250731

	var ok bool 
	var err error

	m := init_manager()
	fp := create_fp(&m)

	// test UpdateFP/GetFP
	m.UpdateFP(&fp, true)
	fp_got, err := m.GetFP(fp.Id)
	assert.Equal(t, fp_got.Id, fp.Id)
	assert.Nil(t, err)

	// test FillProfile
	ok, err = m.FillProfile(&fp)
	fmt.Printf("FillProfile ok:\n%#v\n", ok)
	fmt.Printf("FillProfile err:\n%v\n", err)
	assert.True(t, ok)
	assert.Nil(t, err)

	// test StartProfile
	ok, err = m.StartProfile(&fp)
	assert.True(t, ok)
	assert.Nil(t, err)

	// test StartProfileById
	ok, err = m.StartProfileByID(fp.Id)
	assert.True(t, ok)
	assert.Nil(t, err)

	// cleanup 
	ok, err = utils.DirRemove(fp.DirProfile)
	assert.True(t, ok)
	assert.Nil(t, err)

}
