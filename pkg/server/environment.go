// Copyright 2021 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package server

import (
	"github.com/gorilla/mux"
	"github.com/moov-io/base/config"
	"github.com/moov-io/base/log"
	"github.com/moov-io/base/stime"
)

// Environment - Contains everything thats been instantiated for this service.
type Environment struct {
	Logger       log.Logger
	Config       *Config
	TimeService  *stime.TimeService
	PublicRouter *mux.Router
	Shutdown     func()
}

// NewEnvironment - Generates a new default environment. Overrides can be specified via configs.
func NewEnvironment(env *Environment) (*Environment, error) {
	if env.Logger == nil {
		env.Logger = log.NewDefaultLogger()
	}

	if env.Config == nil {
		ConfigService := config.NewService(env.Logger)

		global := &GlobalConfig{}
		if err := ConfigService.Load(global); err != nil {
			return nil, err
		}

		env.Config = &global.ISO20022
	}

	if env.TimeService == nil {
		t := stime.NewSystemTimeService()
		env.TimeService = &t
	}

	// router
	if env.PublicRouter == nil {
		env.PublicRouter = mux.NewRouter()
	}

	// configure custom handlers
	ConfigureHandlers(env.PublicRouter)

	env.Shutdown = func() {}

	return env, nil
}
