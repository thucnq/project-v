package asynq

type Config struct {
	Redis       Redis
	Concurrency int `json:"concurrency" mapstructure:"concurrency"`
}

type Redis struct {
	Address   string `json:"address" mapstructure:"address"`
	Namespace string `json:"namespace" mapstructure:"namespace"`
	Database  int    `json:"database" mapstructure:"database"`
	Password  string `json:"password" mapstructure:"password"`
}
