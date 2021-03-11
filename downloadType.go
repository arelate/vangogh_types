package vangogh_types

type DownloadType int

const (
	UnknownDownloadType DownloadType = iota
	Image
	BoxArt
	Logo
	Icon
	BackgroundImage
	GalaxyBackgroundImage
	Screenshots
)

var downloadTypeStrings = map[DownloadType]string{
	UnknownDownloadType:   "unknown-download-type",
	Image:                 "image",
	BoxArt:                "box-art",
	Logo:                  "logo",
	Icon:                  "icon",
	BackgroundImage:       "background-image",
	GalaxyBackgroundImage: "galaxy-background-image",
	Screenshots:           "screenshots",
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
	return ok && dt != UnknownDownloadType
}
