package vangogh_types

type DownloadType int

const (
	UnknownDownloadType DownloadType = iota
	ProductImage
	BoxArt
	Background
	Logo
	Screenshots
)

const (
	unknownDownloadTypeStr = "unknown-download-type"
	productImageStr        = "product-image"
	boxArtStr              = "box-art"
	backgroundStr          = "background"
	logoStr                = "logo"
	screenshotsStr         = "screenshots"
)

func (dt DownloadType) String() string {
	switch dt {
	case ProductImage:
		return productImageStr
	case BoxArt:
		return boxArtStr
	case Background:
		return backgroundStr
	case Logo:
		return logoStr
	case Screenshots:
		return screenshotsStr
	default:
		return unknownDownloadTypeStr
	}
}

func ParseDownloadType(downloadType string) DownloadType {
	switch downloadType {
	case productImageStr:
		return ProductImage
	case boxArtStr:
		return BoxArt
	case backgroundStr:
		return Background
	case logoStr:
		return Logo
	case screenshotsStr:
		return Screenshots
	default:
		return UnknownDownloadType
	}
}

func ValidDownloadType(dt DownloadType) bool {
	switch dt {
	case ProductImage:
		fallthrough
	case BoxArt:
		fallthrough
	case Background:
		fallthrough
	case Logo:
		fallthrough
	case Screenshots:
		return true
	default:
		return false
	}
}
