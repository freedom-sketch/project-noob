package config

type PanelCredentials struct {
	URL      string `json:"url"`
	Username int    `json:"username"`
	Password string `json:"password"`
}

type SubscriptionData struct {
	ProfileUpdateInterval int    `json:"profile-update-interval"`
	ProfileTitle          string `json:"profile-title"`
	SubscriptionUserinfo  string `json:"subscription-userinfo"`
	SupportUrl            string `json:"support-url"`
	ProfileWebPageUrl     string `json:"profile-web-page-url"`
	Announce              string `json:"announce"`
}
