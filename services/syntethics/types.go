package syntehics

//SyntethicMonitors describes monitors in new relic
type SyntethicMonitors struct {
	Monitors []struct {
		ID           string   `json:"id"`
		Name         string   `json:"name"`
		Type         string   `json:"type"`
		Frequency    int      `json:"frequency"`
		URI          string   `json:"uri"`
		Locations    []string `json:"locations"`
		Status       string   `json:"status"`
		SLAThreshold float32  `json:"slaThreshold"`
		Options      struct {
		} `json:"options"`
		ModifiedAt string `json:"modifiedAt"`
		CreatedAt  string `json:"createdAt"`
		UserID     int    `json:"userId"`
		APIVersion string `json:"apiVersion"`
	} `json:"monitors"`
	Count int `json:"count"`
}
