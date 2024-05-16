package iamclient

// repositoryImpl ...
type repositoryImpl struct {
	cfg           Config
	mapServiceKey map[string]bool
	listServices  []*ServiceInfo
	req           *httprequest.HttpClient
}

// Config ...
type Config struct {
	BaseUrl   string `json:"base_url" mapstructure:"base_url"`
	SecretKey string `json:"api_key" mapstructure:"api_key"`
	Skip      bool   `json:"skip" mapstructure:"skip"`
	DebugKey  string `json:"debug_key" mapstructure:"debug_key"`
}

// ServiceInfo ...
type ServiceInfo struct {
	ID     string   `json:"id"`
	APIKey string   `json:"apiKey"`
	Scopes []string `json:"scopes"`
}

// LoadServices ...
func (o *repositoryImpl) LoadServices() error {
	return o.loadServices()
}

// IsSkip ...
func (o *repositoryImpl) IsSkip() bool {
	return o.cfg.Skip
}

// NewRepositoryImpl creates a new instance of repositoryImpl, contains whole common functions
// for a service
func NewRepositoryImpl(config Config) gapoiam.IRepository {
	httpClient := httprequest.NewClient()
	return &repositoryImpl{
		cfg: config,
		req: httpClient,
	}
}
