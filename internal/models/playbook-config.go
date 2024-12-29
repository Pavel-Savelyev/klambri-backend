package models

type PlaybookConfig struct {
	PlaybookName string   `json:"playbook_name"`
	Hosts        []string `json:"hosts"`
}
