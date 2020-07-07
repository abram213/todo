package app

import (
	"github.com/AlecAivazis/survey/v2"
	"gopkg.in/yaml.v2"
	"os"
)

type sConfig struct {
	Port        int    `yaml:"Port"`
	DataMode    string `yaml:"DataMode"`
	DatabaseURI string `yaml:"DatabaseURI"`
}

func GenerateConfigFile() (string, error) {
	var qs = []*survey.Question{
		{
			Name: "Port",
			Prompt: &survey.Input{
				Message: "Enter web application port:",
				Default: "3030",
			},
			Validate: survey.Required,
		},
		{
			Name: "DataMode",
			Prompt: &survey.Select{
				Message: "Choose data mode:",
				Options: []string{"sql", "mongo", "file"},
				Default: "sql",
			},
			Validate: survey.Required,
		},
		{
			Name: "DatabaseURI",
			Prompt: &survey.Input{
				Message: "Enter URI to database:",
				Default: ":memory:",
				Help:    "SQLite example: 'your_filename'.db\n Mongo example: mongodb+srv://{your_username}:{your_password}@{cluster_name}.0fggh.mongodb.net/{table_name}?retryWrites=true&w=majority\n File example: 'path_to_file'",
			},
		},
	}

	var answers sConfig

	err := survey.Ask(qs, &answers)
	if err != nil {
		return "", err
	}
	fileName, err := createYamlConfig(answers)
	if err != nil {
		return "", err
	}
	return fileName, nil
}

func createYamlConfig(config sConfig) (string, error) {
	mConfig, err := yaml.Marshal(&config)
	if err != nil {
		return "", err
	}

	file, err := os.Create("config.yaml")
	if err != nil {
		return "", err
	}
	defer file.Close()
	if _, err := file.Write(mConfig); err != nil {
		return "", err
	}
	return "config.yaml", nil
}
