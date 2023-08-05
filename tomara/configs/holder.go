package configs

import _ "embed"

//go:embed mysql.yaml
var MySqlConfigYamlString []byte
