package config

type SiteInfo struct {
	CreateAt string `yaml:"create_at" json:"create_at"`
	Title    string `yaml:"title" json:"title"`
	Email    string `yaml:"email" json:"email"`
	Name     string `yaml:"name" json:"name"`
}
