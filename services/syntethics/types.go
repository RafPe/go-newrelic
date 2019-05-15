package syntehics

//SyntethicMonitors describes monitors in new relic
type SyntethicMonitors struct {
	Monitors []SyntethicMonitor `json:"monitors"`
	Count    int                `json:"count"`
}

// SyntethicMonitor represents single monitor from New Relic
type SyntethicMonitor struct {
	ID           string   `json:"id,omitempty"`
	Name         string   `json:"name,omitempty"`
	Type         string   `json:"type,omitempty"`
	Frequency    int      `json:"frequency,omitempty"`
	URI          string   `json:"uri,omitempty"`
	Locations    []string `json:"locations,omitempty"`
	Status       string   `json:"status,omitempty"`
	SLAThreshold float32  `json:"slaThreshold,omitempty"`
	Options      struct {
	} `json:"options,omitempty"`
	ModifiedAt string `json:"modifiedAt,omitempty"`
	CreatedAt  string `json:"createdAt,omitempty"`
	UserID     int    `json:"userId,omitempty"`
	APIVersion string `json:"apiVersion,omitempty"`
}
