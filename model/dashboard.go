package model

type Dashboard struct {
	CustomerID            int                   `json:"customerId"`
	UserID                int                   `json:"userId"`
	Domain                string                `json:"domain"`
	Widgets               []Widget              `json:"widgets"`
	Version               int                   `json:"version"`
	Shared                bool                  `json:"shared"`
	Schema                int                   `json:"schema"`
	Name                  string                `json:"name"`
	ID                    int                   `json:"id"`
	Public                bool                  `json:"public"`
	CreatedOn             int64                 `json:"createdOn"`
	Username              string                `json:"username"`
	ModifiedOn            int64                 `json:"modifiedOn"`
	AutoCreated           bool                  `json:"autoCreated"`
	TeamID                int                   `json:"teamId"`
	PublicToken           string                `json:"publicToken"`
	Favorite              bool                  `json:"favorite"`
	EventsOverlaySettings EventsOverlaySettings `json:"eventOverlaySettings"`
	ScopeExpressionList   []ScopeExpression     `json:"scopeExpressionList"`
}

type Metric struct {
	ID               string `json:"id"`
	PropertyName     string `json:"propertyName"`
	TimeAggregation  string `json:"timeAggregation"`
	GroupAggregation string `json:"groupAggregation"`
}

type Axis struct {
	From int `json:"from"`
	To   int `json:"to"`
}

type ValueLimit struct {
	Count     int    `json:"count"`
	Direction string `json:"direction"`
}

type Histogram struct {
	NumberOfBuckets int `json:"numberOfBuckets"`
}

type CustomDisplayOptions struct {
	ValueLimit       ValueLimit `json:"valueLimit"`
	Histogram        `json:"histogram"`
	YAxisScale       string `json:"yAxisScale"`
	YAxisLeftDomain  Axis   `json:"yAxisLeftDomain"`
	YAxisRightDomain Axis   `json:"yAxisRightDomain"`
	XAxis            Axis   `json:"xAxis"`
}

type GridConfiguration struct {
	Col   int `json:"col"`
	Row   int `json:"row"`
	SizeX int `json:"size_x"`
	SizeY int `json:"size_y"`
}

type Widget struct {
	ShowAs               string               `json:"showAs"`
	Name                 string               `json:"name"`
	GridConfigurations   []GridConfiguration  `json:"gridConfigurations"`
	CustomDisplayOptions CustomDisplayOptions `json:"customDisplayOptions"`
	Scope                string               `json:"scope"`
	OverrideScope        bool                 `json:"overrideScope"`
	Metrics              []Metric             `json:"metrics"`
	CompareToConfig      interface{}          `json:"compareToConfig"`
	ColorCoding          ColorCoding          `json:"colorCoding,omitempty"`
}

type ColorCoding struct {
	Active     bool       `json:"active"`
	Thresholds Thresholds `json:"thresholds"`
}

type Thresholds struct {
	Worst int `json:"worst"`
	Best  int `json:"best"`
}

type ScopeExpression struct {
	Operand     string   `json:"operand"`
	Operator    string   `json:"operator"`
	DisplayName string   `json:"displayName"`
	Value       []string `json:"value"`
	IsVariable  bool     `json:"isVariable"`
}

type EventsOverlaySettings struct {
	ShowNotificationsEnabled           bool   `json:"showNotificationsEnabled"`
	FilterNotificationsScopeFilter     bool   `json:"filterNotificationsScopeFilter"`
	FilterNotificationsUserInputFilter string `json:"filterNotificationsUserInputFilter"`
	EventOverlayLimit                  int    `json:"eventOverlayLimit"`
	FilterNotificationsTypeFilter      string `json:"filterNotificationsTypeFilter"`
	FilterNotificationsStateFilter     string `json:"filterNotificationsStateFilter"`
	FilterNotificationsSeverityFilter  string `json:"filterNotificationsSeverityFilter"`
	FilterNotificationsResolvedFilter  string `json:"filterNotificationsResolvedFilter"`
}
