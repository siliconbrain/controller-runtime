/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package manager_test

import (
	"log"

	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/runtime/signals"
)

var mgr manager.Manager

// This example creates a new Manager that can be used with controller.New to create Controllers.
func ExampleNew() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	mgr, err := manager.New(cfg, manager.Options{})
	if err != nil {
		log.Fatal(err)
	}
	log.Print(mgr)
}

// This example adds a Runnable for the Manager to Start.
func ExampleManager_Add() {
	err := mgr.Add(manager.RunnableFunc(func(<-chan struct{}) error {
		// Do something
		return nil
	}))
	if err != nil {
		log.Fatal(err)
	}
}

// This example starts a Manager that has had Runnables added.
func ExampleManager_Start() {
	err := mgr.Start(signals.SetupSignalHandler())
	if err != nil {
		log.Fatal(err)
	}
}
