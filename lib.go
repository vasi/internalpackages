package internalpackages

import "github.com/vasi/internalpackages/internal"

// A user of this library can't get PrivateData,
// and can't import internal.
type MyFile interface {
	GetPublicInfoAboutFile() string
}

func ParseFile(s string) MyFile {
	// Assume we're doing some calculations to figure out what PrivateData and
	// PublicData should be
	return &internal.MyFileImpl{
		PrivateData: "private",
		PublicData:  "public",
	}
}
