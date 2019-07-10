/*
 * Copyright (c) 2019.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"github.com/Flaque/filet"
	_ "github.com/Flaque/filet"
	"os"
	"testing"
)

func TestParseCommandLine(t *testing.T) {
	defer filet.CleanUp(t)

	// Creates a temporary file with string "some content"
	file := filet.TmpFile(t, "", "some content")

	os.Args = []string{"command", "--file=" + file.Name(), "--version=1.2.3"}

	m := &Config{}
	m.ParseCommandLine()

	if m.file != file.Name() {
		t.Errorf("The file is not correct specified %s.", file.Name())
	}
	if m.version != "1.2.3" {
		t.Errorf("The content of the file %s is not correct specified (1.2.3).", file.Name())
	}
}

func TestFileExist(t *testing.T) {
	defer filet.CleanUp(t)

	// Creates a temporary file with string "some content"
	file := filet.TmpFile(t, "", "some content")

	md := &Config{}
	md.file = file.Name()

	if !md.FileExists() {
		t.Errorf("File is there but not correct identified. File is %s", file.Name())
	}
}
