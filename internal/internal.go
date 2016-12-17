package internal

type MyFileImpl struct {
	PrivateData string
	PublicData  string
}

func (f *MyFileImpl) GetPublicInfoAboutFile() string {
	return f.PublicData + " " + string(len(f.PrivateData))
}
