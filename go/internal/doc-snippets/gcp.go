// Copyright 2024 Google LLC
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

package snippets

import (
	"context"
	"log/slog"

	"github.com/firebase/genkit/go/plugins/googlecloud"
)

func gcpEx(ctx context.Context) error {
	//!+init
	if err := googlecloud.Init(
		ctx,
		googlecloud.Config{ProjectID: "your-google-cloud-project"},
	); err != nil {
		return err
	}
	//!-init

	_ = googlecloud.Config{
		ProjectID:      "your-google-cloud-project",
		ForceExport:    true,
		MetricInterval: 45e9,
		LogLevel:       slog.LevelDebug,
	}

	return nil
}