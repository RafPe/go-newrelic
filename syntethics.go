package newrelic

// SyntethicService represents exposed services to manage Synthethics
type SyntethicService struct {
	client *Client
}

//SyntethicService returns all monitors in your account
type SyntethicMonitors struct {
	Monitors []struct {
		ID           string   `json:"id"`
		Name         string   `json:"name"`
		Type         string   `json:"type"`
		Frequency    int      `json:"frequency"`
		URI          string   `json:"uri"`
		Locations    []string `json:"locations"`
		Status       string   `json:"status"`
		SLAThreshold int      `json:"slaThreshold"`
		Options      struct {
		} `json:"options"`
		ModifiedAt string `json:"modifiedAt"`
		CreatedAt  string `json:"createdAt"`
		UserID     int    `json:"userId"`
		APIVersion string `json:"apiVersion"`
	} `json:"monitors"`
	Count int `json:"count"`
}

func getAllMonitors() {
	// curl -v \
	//  -H 'X-Api-Key:{Admin_User_Key}' https://synthetics.newrelic.com/synthetics/api/v3/monitors
}
