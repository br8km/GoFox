package chameleon

// Chameleon Settings
type Settings struct {
	Config    Config        `json:"config"`
	Excluded  []interface{} `json:"excluded"`
	Headers   Headers       `json:"headers"`
	IPRules   []interface{} `json:"ipRules"`
	Profile   Profile       `json:"profile"`
	Options   Options       `json:"options"`
	Whitelist Whitelist     `json:"whitelist"`
	Version   string        `json:"version"`
}

type Config struct {
	Enabled              bool   `json:"enabled"`
	NotificationsEnabled bool   `json:"notificationsEnabled"`
	Theme                string `json:"theme"`
	ReloadIPStartupDelay int64  `json:"reloadIPStartupDelay"`
	HasPrivacyPermission bool   `json:"hasPrivacyPermission"`
}

type Headers struct {
	BlockEtag       bool            `json:"blockEtag"`
	EnableDNT       bool            `json:"enableDNT"`
	Referer         Referer         `json:"referer"`
	SpoofAcceptLang SpoofAcceptLang `json:"spoofAcceptLang"`
	SpoofIP         SpoofIP         `json:"spoofIP"`
}

type Referer struct {
	Disabled bool  `json:"disabled"`
	Xorigin  int64 `json:"xorigin"`
	Trimming int64 `json:"trimming"`
}

type SpoofAcceptLang struct {
	Enabled bool   `json:"enabled"`
	Value   string `json:"value"`
}

type SpoofIP struct {
	Enabled   bool   `json:"enabled"`
	Option    int64  `json:"option"`
	RangeFrom string `json:"rangeFrom"`
	RangeTo   string `json:"rangeTo"`
}

type Options struct {
	CookieNotPersistent    bool                 `json:"cookieNotPersistent"`
	CookiePolicy           string               `json:"cookiePolicy"`
	BlockMediaDevices      bool                 `json:"blockMediaDevices"`
	BlockCSSExfil          bool                 `json:"blockCSSExfil"`
	DisableWebRTC          bool                 `json:"disableWebRTC"`
	FirstPartyIsolate      bool                 `json:"firstPartyIsolate"`
	LimitHistory           bool                 `json:"limitHistory"`
	ProtectKBFingerprint   ProtectKBFingerprint `json:"protectKBFingerprint"`
	ProtectWinName         bool                 `json:"protectWinName"`
	ResistFingerprinting   bool                 `json:"resistFingerprinting"`
	ScreenSize             string               `json:"screenSize"`
	SpoofAudioContext      bool                 `json:"spoofAudioContext"`
	SpoofClientRects       bool                 `json:"spoofClientRects"`
	SpoofFontFingerprint   bool                 `json:"spoofFontFingerprint"`
	SpoofMediaDevices      bool                 `json:"spoofMediaDevices"`
	TimeZone               string               `json:"timeZone"`
	TrackingProtectionMode string               `json:"trackingProtectionMode"`
	WebRTCPolicy           string               `json:"webRTCPolicy"`
	WebSockets             string               `json:"webSockets"`
}

type ProtectKBFingerprint struct {
	Enabled bool  `json:"enabled"`
	Delay   int64 `json:"delay"`
}

type Profile struct {
	Selected          string   `json:"selected"`
	Interval          Interval `json:"interval"`
	ShowProfileOnIcon bool     `json:"showProfileOnIcon"`
}

type Interval struct {
	Option int64 `json:"option"`
	Min    int64 `json:"min"`
	Max    int64 `json:"max"`
}

type Whitelist struct {
	EnabledContextMenu bool          `json:"enabledContextMenu"`
	DefaultProfile     string        `json:"defaultProfile"`
	Rules              []interface{} `json:"rules"`
}
