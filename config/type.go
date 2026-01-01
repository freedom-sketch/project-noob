package config

type SubscriptionMetaData struct {
	ProfileUpdateInterval int    `json:"profile-update-interval"`
	ProfileTitle          string `json:"profile-title"`
	SupportUrl            string `json:"support-url"`
	ProfileWebPageUrl     string `json:"profile-web-page-url"`
	Announce              string `json:"announce"`
}

type API struct {
	Host    string `json:"host"`
	WebPath string `json:"web_path"`
}

type Logging struct {
	FileName string `json:"file_name"`
	Level    string `json:"level"`
}

type Config struct {
	API          API                  `json:"api"`
	Subscription SubscriptionMetaData `json:"subscription"`
	Logging      Logging              `json:"logging"`
}
