package vangogh_types

func AllImageDownloadTypes() []DownloadType {
	downloadTypes := make([]DownloadType, 0, len(downloadTypeStrings))
	for dt, _ := range downloadTypeStrings {
		if dt == UnknownDownloadType {
			continue
		}
		downloadTypes = append(downloadTypes, dt)
	}
	return downloadTypes
}
