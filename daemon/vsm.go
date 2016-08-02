// Copyright 2016 CloudByte, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package daemon

import (
	"github.com/openebs/openebs/types"
)

// Vsms returns the list of VSMs to show given the user's filtering.
func (daemon *Daemon) Vsms(config *types.VSMListOptions) ([]*types.Vsm, error) {
	vsms := []*types.Vsm{}
	//TODO Fill with some data

	return vsms, nil
}

// This returns the newly created VSM.
func (daemon *Daemon) VsmCreate(opts *types.VSMListOptions) (*types.Vsm, error) {
	vsm := *types.Vsm{}
	//TODO Fill with some data

	return vsm, nil
}
