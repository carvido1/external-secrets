/*
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

package v1beta1

import (
	esmeta "github.com/external-secrets/external-secrets/apis/meta/v1"
)

// Configures a store to push secrets to Github Actions.
type GithubProvider struct {
	// URL configures the Github instance URL. Defaults to https://github.com/.
	URL string `json:"url,omitempty"`

	// auth configures how secret-manager authenticates with a Github instance.
	Auth GithubAppAuth `json:"auth"`

	// appID specifies the Github APP that will be used to authenticate the client
	AppID int64 `json:"appID"`

	// installationID specifies the Github APP installation that will be used to authenticate the client
	InstallationID int64 `json:"installationID"`

	// organisationName will be used to fetch secrets from the Github organization
	OrganisationName string `json:"organisationName"`
}

type GithubAppAuth struct {
	PrivateKey esmeta.SecretKeySelector `json:"privateKey"`
}