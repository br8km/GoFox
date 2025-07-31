package firefox

var (
	// profile folder: `{prefix}{seperator}{name}`
	// profile folder: `{group}{seperator}{id}`
	PROFILE_SEPERATOR string = "--"

	// firefox extensions file names
	EXTENSIONS = map[string]string{
		"Chameleon":         "{3579f63b-d8ee-424f-bbb6-6d0ce3285e6a}.xpi",
		"Dark Reader":       "addon@darkreader.org.xpi",
		"Simple Translate":  "simple-translate@sienori.xpi",
		"HypeStat Analyzer": "{6d930f30-aa68-421f-83b8-71322461bdaa}.xpi",
		"Vimium":            "vimium-c@gdh1995.cn.xpi",
	}
)
