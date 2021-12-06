// Copyright 2016-2021, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"github.com/stretchr/testify/assert"
)

func TestGetPlugin(t *testing.T) {
	cases := []struct {
		Name      string
		Mod       *modInfo
		Expected  *pulumirpc.PluginDependency
		ShouldErr bool
		json      *plugin.PulumiPluginJSON
	}{
		{
			Name: "valid-pulumi-mod",
			Mod: &modInfo{
				Path:    "github.com/pulumi/pulumi-aws/sdk",
				Version: "v1.29.0",
			},
			Expected: &pulumirpc.PluginDependency{
				Name:    "aws",
				Version: "v1.29.0",
			},
		},
		{
			Name: "pulumi-pseduo-version-plugin",
			Mod: &modInfo{
				Path:    "github.com/pulumi/pulumi-aws/sdk",
				Version: "v1.29.1-0.20200403140640-efb5e2a48a86",
			},
			Expected: &pulumirpc.PluginDependency{
				Name:    "aws",
				Version: "v1.29.0",
			},
		},
		{
			Name: "non-pulumi-mod",
			Mod: &modInfo{
				Path:    "github.com/moolumi/pulumi-aws/sdk",
				Version: "v1.29.0",
			},
			ShouldErr: true,
		},
		{
			Name: "invalid-version-module",
			Mod: &modInfo{
				Path:    "github.com/pulumi/pulumi-aws/sdk",
				Version: "42-42-42",
			},
			ShouldErr: true,
		},
		{
			Name: "pulumi-pulumi-mod",
			Mod: &modInfo{
				Path:    "github.com/pulumi/pulumi/sdk",
				Version: "v1.14.0",
			},
			ShouldErr: true,
		},
		{
			Name: "beta-pulumi-module",
			Mod: &modInfo{
				Path:    "github.com/pulumi/pulumi-aws/sdk",
				Version: "v2.0.0-beta.1",
			},
			Expected: &pulumirpc.PluginDependency{
				Name:    "aws",
				Version: "v2.0.0-beta.1",
			},
		},
		{
			Name: "non-zero-patch-module", Mod: &modInfo{
				Path:    "github.com/pulumi/pulumi-kubernetes/sdk",
				Version: "v1.5.8",
			},
			Expected: &pulumirpc.PluginDependency{
				Name:    "kubernetes",
				Version: "v1.5.8",
			},
		},
		{
			Name: "pulumiplugin",
			Mod: &modInfo{
				Path:    "github.com/me/myself/i",
				Version: "invalid-Version",
			},
			Expected: &pulumirpc.PluginDependency{
				Name:    "thing1",
				Version: "v1.2.3",
				Server:  "myserver.com",
			},
			json: &plugin.PulumiPluginJSON{
				Resource: true,
				Name:     "thing1",
				Version:  "v1.2.3",
				Server:   "myserver.com",
			},
		},
		{
			Name:      "non-resource",
			Mod:       &modInfo{},
			ShouldErr: true,
			json: &plugin.PulumiPluginJSON{
				Resource: false,
			},
		},
		{
			Name: "missing-pulumiplugin",
			Mod: &modInfo{
				Dir: "/not/real",
			},
			ShouldErr: true,
			json: &plugin.PulumiPluginJSON{
				Name:    "thing2",
				Version: "v1.2.3",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			cwd := t.TempDir()
			if c.json != nil {
				if c.Mod.Dir == "" {
					c.Mod.Dir = cwd
				}
				bytes, err := c.json.JSON()
				assert.NoError(t, err, "Failed to setup test pulumiplugin.json")
				err = os.WriteFile(filepath.Join(cwd, "pulumiplugin.json"), bytes, 0600)
				assert.NoError(t, err, "Failed to write pulumiplugin.json")
			}
			actual, err := c.Mod.getPlugin()
			if c.ShouldErr {
				assert.Error(t, err)
			} else {
				// Kind must be resource. We can thus exclude it from the test.
				if c.Expected.Kind == "" {
					c.Expected.Kind = "resource"
				}
				assert.NoError(t, err)
				assert.Equal(t, c.Expected, actual)
			}
		})
	}
}
