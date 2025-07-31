package firefox

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/br8km/GoFox/chameleon"
	"github.com/br8km/GoFox/utils"
	"github.com/goccy/go-json"
)

// Firefox FingerPrint Manager
type FPManager struct {
	Config       Config
	Assets       []Asset
	FingerPrints []FingerPrint
	// Chameleon   chameleon.ChameleonSettings        
}


// check
func (m *FPManager) Check() (bool, error) {
	var err error

	// check config
	_, err = m.Config.Check()
	if err != nil {
		return false, nil
	}

	// check assets
	if len(m.Assets) == 0 {
		return false, fmt.Errorf("no Asset files yet")
	}

	for _, asset := range m.Assets {
		_, err := asset.Check()
		if err != nil {
			return false, err
		}
	}

	return true, nil

}


// get dir of Profile for FingerPrint
func (m *FPManager) DirOfProfile(fp *FingerPrint) string {
	folder := fmt.Sprintf("%s%s%s", fp.Group, m.Config.Seperator, fp.Id)
	return filepath.Join(
		m.Config.DirProfiles,
		folder,
	)
}

// create FingerPrint item, but no folders/files yet.
func (m *FPManager) CreateFP(
	group string,
	id string,
	lang string,
	proxy_url string,
	os_family string,
	br_family string,
	device_type string,
) (fp FingerPrint, err error) {

	if id == "" {
		id = utils.StrRndChar(8)
	}
	proxy, err := utils.FromUrl(proxy_url, true)
	if err != nil {
		return fp, err
	}
	o, err := GetRandomOS(os_family)
	if err != nil {
		return fp, err
	}

	b, err := GetRandomBrowser(br_family, device_type)
	if err != nil {
		return fp, err
	}

	s, err := GetRandomScreenSize(device_type)
	if err != nil {
		return fp, err
	}


	location := Location{
		Proxy: proxy,
		GeoAddr: GeoAddr{
			IPAddr: proxy.IPAddr,
		},
	}
	device := Device{
		Browser:    b,
		OS:         o,
		Type:       device_type,
		Language:   lang,
		ScreenSize: s,
	}

	fp = FingerPrint{
		Group:    group,
		Id:       id,
		Location: location,
		Device:   device,
	}
	fp.DirProfile = m.DirOfProfile(&fp)
	return fp, nil
}

// delete FingerPrint item, without touch folders/files yet
func (m *FPManager) DeleteFP(fp *FingerPrint) (bool, error) {
	index := utils.SliceIndex(
		len(m.FingerPrints),
		func(i int) bool { return m.FingerPrints[i] == *fp },
	)
	ret := make([]FingerPrint, len(m.FingerPrints)-1)
	ret = append(ret, m.FingerPrints[:index]...)
	ret = append(ret, m.FingerPrints[index+1:]...)
	m.FingerPrints = ret
	return true, nil
}

// update FingerPrint, or add it if not exist yet
func (m *FPManager) UpdateFP(fp *FingerPrint, add bool) {
	// TODO ?? check if this way okay
	exist := false
	for _, o := range m.FingerPrints {
		if o.Id == fp.Id {
			o = *fp
			exist = true
			break
		}
	}
	if !exist && add {
		m.FingerPrints = append(m.FingerPrints, *fp)
	}
}

// get FingerPrint by id
func (m *FPManager) GetFP(id string) (FingerPrint, error) {
	// fmt.Printf("find id = %s", id)
	// fmt.Printf("fps.len = %d", len(m.FingerPrints))
	var fp FingerPrint
	for _, item := range m.FingerPrints {
		if item.Id == id {
			return item, nil
		}
	}
	return fp, fmt.Errorf("error Profile.Id = `%s`", id)
}

// fill Profile folder from demo profile folder
func (m *FPManager) FillProfile(fp *FingerPrint) (bool, error) {
	for _, asset := range m.Assets {
		if asset.Kind == "Profile"{
			_, err := asset.Copy(fp.DirProfile)
			if err != nil {
				return false, err
			}
		}
	}
	return true, nil
}

// remove profile folder for FingerPrint item
func (m *FPManager) RemoveProfile(fp *FingerPrint) (bool, error) {
	return utils.DirRemove(fp.DirProfile)
}

// load FingerPrint items from local file
func (m *FPManager) LoadFPs(f string) ([]FingerPrint, error) {
	var res []FingerPrint

	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)

	err = decoder.Decode(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// save FingerPrint items to local file
func (m *FPManager) SaveFPs(f string, indent bool) (bool, error) {
	var b []byte
	var err error

	if indent {
		b, err = json.MarshalIndent(m.FingerPrints, "", "\t")
	} else {
		b, err = json.Marshal(m.FingerPrints)
	}

	if err != nil {
		return false, err
	}

	err = os.WriteFile(f, b, os.ModePerm)
	if err != nil {
		return false, err
	}
	return true, nil

}

// run shell command for FingerPrint
func (m *FPManager) StartProfile(fp *FingerPrint) (bool, error) {
	// exe := fmt.Sprintf(`"%s"`, m.Config.ExePath)
	// dir := fmt.Sprintf(`"%s"`, fp.DirProfile)
	args := []string{
		m.Config.ExePath,
		"-profile",
		fp.DirProfile,
	}
	return utils.RunCommand(args...)
}

// run shell command for Profile by id
func (m *FPManager) StartProfileByID(id string) (bool, error) {
	fp, err := m.GetFP(id)
	// fmt.Printf("get fp: %#v\n", fp)
	if err != nil {
		return false, err
	}
	return m.StartProfile(&fp)
}

// ---

type Chameleon struct {
	Settings chameleon.Settings
}

// init Chameleon
func (m *FPManager) InitChameleon() (Chameleon, error) {
	var res Chameleon

	if len(m.Assets) == 0 {
		return res, fmt.Errorf("no Assets available")
	}

	dir_profile := "" 
	for _, asset := range m.Assets {
		if asset.FileName == "chameleon_settings.json" {
			dir_profile = asset.DirParent
			break
		}
	} 
	if dir_profile == "" {
		return res, fmt.Errorf("asset.dir_profile error")
	}
	ch, err := chameleon.Load(dir_profile)
	if err != nil {
		return res, err
	}
	return Chameleon{Settings: ch}, nil

}


func (cs *Chameleon) UpdateByFP(fp *FingerPrint) bool {

	cs.Settings.SetEnable(true)
	cs.Settings.SetLanguage(true, fp.Device.Language)
	// dark theme for eyes
	cs.Settings.SetTheme("dark")
	// TODO: check..
	cs.Settings.SetProfile(
		fp.Device.OS.Code,
		fp.Device.Browser.Code,
	)
	cs.Settings.SetTimezone(fp.Location.GeoAddr.Timezone)
	cs.Settings.SetScreenSize(fp.Device.ScreenSize)

	return true

}