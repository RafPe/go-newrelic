package syntehics

//MonitorType defines allowed types of Syntethic monitors.
type MonitorType string

type MonitorStatus string

type MonitorFrequency int

const (
	SyntethicMonitorEnabled  MonitorStatus = "ENABLED"
	SyntethicMonitorMuted    MonitorStatus = "MUTED"
	SyntethicMonitorDisabled MonitorStatus = "DISABLED"

	SyntethicMonitorSimple        MonitorType = "SIMPLE"
	SyntethicMonitorBrowser       MonitorType = "BROWSER"
	SyntethicMonitorScriptAPI     MonitorType = "SCRIPT_API"
	SyntethicMonitorScriptBrowser MonitorType = "SCRIPT_BROWSER"

	SyntethicMonitorFreq1m    MonitorFrequency = 1
	SyntethicMonitorFreq5m    MonitorFrequency = 5
	SyntethicMonitorFreq10m   MonitorFrequency = 10
	SyntethicMonitorFreq15m   MonitorFrequency = 15
	SyntethicMonitorFreq30m   MonitorFrequency = 30
	SyntethicMonitorFreq60m   MonitorFrequency = 60
	SyntethicMonitorFreq360m  MonitorFrequency = 360
	SyntethicMonitorFreq720m  MonitorFrequency = 720
	SyntethicMonitorFreq1440m MonitorFrequency = 1440
)

//SyntethicMonitors describes monitors in new relic
type SyntethicMonitors struct {
	Monitors []SyntethicMonitor `json:"monitors"`
	Count    int                `json:"count"`
}

// SyntethicMonitor represents single monitor from New Relic
type SyntethicMonitor struct {
	ID           string           `json:"id,omitempty"`
	Name         string           `json:"name,omitempty"`
	Type         MonitorType      `json:"type,omitempty"`
	Frequency    MonitorFrequency `json:"frequency,omitempty"`
	URI          string           `json:"uri,omitempty"`
	Locations    []string         `json:"locations,omitempty"`
	Status       MonitorStatus    `json:"status,omitempty"`
	SLAThreshold float32          `json:"slaThreshold,omitempty"`
	Options      struct {
	} `json:"options,omitempty"`
	ModifiedAt string `json:"modifiedAt,omitempty"`
	CreatedAt  string `json:"createdAt,omitempty"`
	UserID     int    `json:"userId,omitempty"`
	APIVersion string `json:"apiVersion,omitempty"`
}
