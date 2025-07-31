package gofox

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"strings"

	"github.com/br8km/gofox/chameleon"
	"github.com/br8km/gofox/utils"
)

type Browser struct {
	Family string // `json:"Family"`
	Code   string // `json:"Code"`
}

// this code are corrsponding to ChameleonSettings -> Profile.Selected, from key of CHAMELEON_BROWSER.
// type ChameleonProfile string

type OperationSystem struct {
	Family string // `json:"Family"`
	Code   string // `json:"Code"`
}

type GeoAddr struct {
	IPAddr   string // `json:"IPAddr"`
	Country  string // `json:"Country"`
	State    string // `json:"State"`
	City     string // `json:"City"`
	Timezone string // `json:"Timezone"`
}


type Location struct {
	GeoAddr GeoAddr     // `json:"GeoAddr"`
	Proxy   utils.Proxy // `json:"Proxy"`
}

type Device struct {
	Browser    Browser         // `json:"Browser"`
	OS         OperationSystem // `json:"OS"`
	Type       string          // `json:"Type"`
	Language   string          // `json:"Language"`
	ScreenSize string          // `json:"ScreenSize"`
}

type FingerPrint struct {
	Group      string // `json:"Group"`
	Id         string // `json:"Id"`

	DirProfile string // `json:"DirProfile"`

	Device   Device   // `json:"Device"`
	Location Location // `json:"Location"`
	// Person   Person   // `json:"Person"`
}

// get random Browser by family, eg: `Firefox`, `Edge`, etc.
// see package chameleon for details
func GetRandomBrowser(
	family string,
	device_type string,
) (Browser, error) {
	var b Browser
	if !slices.Contains(chameleon.BR_Families[:], family) {
		return b, fmt.Errorf("browser family error @ %s", family)
	}

	if !slices.Contains(chameleon.Device_Types[:], device_type) {
		return b, fmt.Errorf("device error @ %s", device_type)
	}

	browser := new(Browser)
	for key, value := range chameleon.CHAMELEON_BROWSER {
		word := fmt.Sprintf("%s-%s", family, device_type)
		if word == value {
		// if key == family && value == device_type {
			browser.Family = family
			browser.Code = key
			break
		}
	}
	if browser.Code == "" || browser.Family == "" {
		return b, fmt.Errorf("get Browser `%s`-`%s` failed", family, device_type)
	}

	return *browser, nil
}

// get random OS by family abbr, eg: `win`, `lin`, etc.
// see package chameleon for details
func GetRandomOS(family string) (OperationSystem, error) {
	var o OperationSystem

	if !slices.Contains(chameleon.OS_Families[:], family) {
		return o, fmt.Errorf("os family error @ %s", family)
	}

	var codes []string
	for code := range chameleon.CHAMELEON_OS {
		if strings.HasPrefix(code, family) {
			codes = append(codes, code)
		}
	}
	o = OperationSystem{
		Family: family,
		Code:   codes[rand.IntN(len(codes)-1)],
	}
	return o, nil
}

func GetRandomScreenSize(device_type string) (string, error) {
	var s string
	if !slices.Contains(chameleon.Device_Types[:], device_type) {
		return s, fmt.Errorf("device error @ %s", device_type)
	}

	var codes []string
	switch device_type {
		case "desktop": {
			codes = chameleon.DESKTOP_RESOLUTIONS
		}
		case "tablet": {
			codes = chameleon.TABLET_RESOLUTIONS
		}
		default: {
			codes = chameleon.MOBILE_RESOLUTIONS
		}
	}
	return codes[rand.IntN(len(codes)-1)], nil
}

// TODO
// func (fp *FingerPrint) ToYaml() string {
// 	return ""
// }

