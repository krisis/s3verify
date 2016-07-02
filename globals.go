/*
 * Minio S3Verify Library for Amazon S3 Compatible Cloud Storage (C) 2016 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"github.com/minio/cli"
	"github.com/minio/mc/pkg/console"
)

var (
	globalDebug = false
)

// Separate out context.
func setGlobals(debug bool) {
	globalDebug = debug

	if globalDebug {
		// Allow printing of traces.
		console.DebugPrint = true
	}
}

// Set any global flags here.
func setGlobalsFromContext(ctx *cli.Context) error {
	debug := ctx.Bool("debug") || ctx.GlobalBool("debug")
	setGlobals(debug)
	return nil
}