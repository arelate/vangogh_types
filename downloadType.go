package vangogh_types

type DownloadType int

const (
	UnknownDownloadType DownloadType = iota
	Image
	BoxArt
	BackgroundImage
	GalaxyBackgroundImage
	Logo
	Icon
)

var downloadTypeStrings = map[DownloadType]string{
	UnknownDownloadType:   "unknown-download-type",
	Image:                 "image",
	BoxArt:                "box-art",
	BackgroundImage:       "background-image",
	GalaxyBackgroundImage: "galaxy-background-image",
	Logo:                  "logo",
	Icon:                  "icon",
}

func (dt DownloadType) String() string {
	str, ok := downloadTypeStrings[dt]
	if ok {
		return str
	}

	return downloadTypeStrings[UnknownDownloadType]
}

func ParseDownloadType(downloadType string) DownloadType {
	for dt, str := range downloadTypeStrings {
		if str == downloadType {
			return dt
		}
	}
	return UnknownDownloadType
}

func ValidDownloadType(dt DownloadType) bool {
	_, ok := downloadTypeStrings[dt]
	return ok
}
