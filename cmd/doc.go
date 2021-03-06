/*
Copyright 2015 All rights reserved.
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

package cmd

import "github.com/cloudwatt/vault-sync/pkg/api"

var (
	// Author is the author of the program
	Author = "Félix Cantournet"
	// Email is the email of the author
	Email = "felix.cantournet@gmail.com"
)

type resources struct {
	// a collection of auths
	auths []*api.Auth
	// a collection of users
	users []*api.User
	// a collection of backends
	backends []*api.Backend
	// a collection of secrets
	secrets []*api.Secret
}
