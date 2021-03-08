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
	Screenshots
)

const (
	unknownDownloadTypeStr   = "unknown-download-type"
	imageStr                 = "image"
	boxArtStr                = "box-art"
	backgroundImageStr       = "background-image"
	galaxyBackgroundImageStr = "galaxy-background-image"
	logoStr                  = "logo"
	iconStr                  = "icon"
	screenshotsStr           = "screenshots"
)

func (dt DownloadType) String() string {
	switch dt {
	case Image:
		return imageStr
	case BoxArt:
		return boxArtStr
	case BackgroundImage:
		return backgroundImageStr
	case GalaxyBackgroundImage:
		return galaxyBackgroundImageStr
	case Logo:
		return logoStr
	case Icon:
		return iconStr
	case Screenshots:
		return screenshotsStr
	default:
		return unknownDownloadTypeStr
	}
}

func ParseDownloadType(downloadType string) DownloadType {
	switch downloadType {
	case imageStr:
		return Image
	case boxArtStr:
		return BoxArt
	case backgroundImageStr:
		return BackgroundImage
	case galaxyBackgroundImageStr:
		return GalaxyBackgroundImage
	case logoStr:
		return Logo
	case iconStr:
		return Icon
	case screenshotsStr:
		return Screenshots
	default:
		return UnknownDownloadType
	}
}

func ValidDownloadType(dt DownloadType) bool {
	switch dt {
	case Image:
		fallthrough
	case BoxArt:
		fallthrough
	case BackgroundImage:
		fallthrough
	case GalaxyBackgroundImage:
		fallthrough
	case Logo:
		fallthrough
	case Icon:
		fallthrough
	case Screenshots:
		return true
	default:
		return false
	}
}
