package fake

type AppInstall struct {
	Data    AppInstallMetaData `json:"data"`
	Message string             `json:"message"`
	Status  int                `json:"status"`
}
type AppInstallMetaData struct {
	IsUpdated bool `json:"is_updated"`
}
