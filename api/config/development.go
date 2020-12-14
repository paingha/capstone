// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"fmt"

	"bitbucket.com/irb/api/plugins"
	"bitbucket.com/irb/api/utils"
)

//BuildDevDBConfig - Builds DB Config for dev environment
func BuildDevDBConfig() *DBConfig {
	var cfg SystemConfig
	err := InitConfig(&cfg)
	if err != nil {
		plugins.LogFatal("API", "Wrong Dev System config", err)
	}
	port, _ := utils.ConvertStringToInt(cfg.DevDBPort)
	dbConfig := DBConfig{
		Host:     cfg.DevDBHost,
		Port:     port,
		User:     cfg.DevDBUser,
		DBName:   cfg.DevDBDatabase,
		Password: cfg.DevDBPass,
		SSL:      cfg.DevDBSSL,
	}
	return &dbConfig
}

//DevDbURL - returns connection string for dev database
func DevDbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
		dbConfig.SSL,
	)
}
