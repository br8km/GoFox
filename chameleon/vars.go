package chameleon

/*

# TODO
TimeZone_Codes = [...]string{} // -> add popular ones

*/

var (
	// export from chameleon extension options page

	OS_Families     = [...]string{"win", "lin", "mac", "ios", "and"}
	BR_Families     = [...]string{"Chrome", "Edge", "Firefox", "Safari"}
	Device_Types = [...]string{"desktop", "mobile", "tablet", "iPhone", "iPad"}

	Lang_Codes     = [...]string{"en-US", "en-GB", "en-AU", "zh-CN"}
	Theme_Codes    = [...]string{"dark", "light"}
	TimeZone_Codes = [...]string{"US/Central"} // TODO
)

// https://sereneblue.github.io/chameleon/wiki/developer-guide/

var CHAMELEON_OS = map[string]string{
	"win1": "Windows 7",
	"win2": "Windows 8",
	"win3": "Windows 8.1",
	"win4": "Windows 10",
	"mac1": "macOS 10.13 (or 2 behind latest)",
	"mac2": "macOS 10.14 (or 1 behind latest)",
	"mac3": "macOS 10.15 (or latest)",
	"lin1": "Linux",
	"lin2": "Ubuntu Linux",
	"lin3": "Fedora Linux",
	"ios1": "iOS 11 (or 2 behind latest)",
	"ios2": "iOS 12 (or 1 behind latest)",
	"ios3": "iOS 13 (or latest)",
	"and1": "Android 6 (or 4 behind latest)",
	"and2": "Android 7 (or 3 behind latest)",
	"and3": "Android 8 (or 2 behind latest)",
	"and4": "Android 9 (or 1 behind latest)",
}

var CHAMELEON_BROWSER = map[string]string{
	"gcr":  "Chrome-desktop",
	"gcrm": "Chrome-mobile",
	"gcrt": "Chrome-tablet",
	"edg":  "Edge-desktop",
	"edgm": "Edge-mobile",
	"ff":   "Firefox-desktop",
	"ffm":  "Firefox-mobile",
	"fft":  "Firefox-tablet",
	"sf":   "Safari-desktop",
	"sfm":  "Safari-iPhone",
	"sft":  "Safari-iPad",
}

// screen size ratio
// https://gs.statcounter.com/screen-resolution-stats/desktop/worldwide

var DESKTOP_RESOLUTIONS = []string{
	"1366x768",
	"1440x900",
	"1600x900",
	"1920x1080",
	"1920x1200",
	"2560x1440",
	"2560x1600",
	"3840x2160",
}

var MAC_RESOLUTIONS = []string{
	"1920x1080",
	"2560x1600",
	"4096x2304",
	"5120x288",
}

var RATIO_SCREEN_DESKTOP = map[string]float64{
	"1920x1080": 22.27,
	"1366x768":  15.84,
	"1536x864":  10.43,
	"1280x720":  6.47,
	"1440x900":  6.02,
	"1600x900":  3.09,
}

var RATIO_SCREEN_MOBILE = map[string]float64{
	"360x800": 11.17,
	"390x844": 6.84,
	"414x896": 5.87,
	"412x915": 5.56,
	"393x873": 4.87,
	"360x780": 4.1,
}

var RATIO_SCREEN_TABLET = map[string]float64{
	"768x1024": 28.17,
	"810x1080": 8.55,
	"1280x800": 7.16,
	"800x1280": 6.48,
	"601x962":  4.74,
	"820x1180": 3.65,
}

var TABLET_RESOLUTIONS = []string{
	"768x1024",
	"810x1080",
	"1280x800",
	"800x1280",
	"601x962",
	"820x1180",
}

var MOBILE_RESOLUTIONS = []string{
	"360x800",
	"390x844",
	"414x896",
	"412x915",
	"393x873",
	"360x780",
}
