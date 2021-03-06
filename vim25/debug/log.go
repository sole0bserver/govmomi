/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

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

package debug

import (
	"io"
	"log"
)

type LogWriterCloser struct {
}

func NewLogWriterCloser() *LogWriterCloser {
	return &LogWriterCloser{}
}

func (lwc *LogWriterCloser) Write(p []byte) (n int, err error) {
	log.Print(string(Scrub(p)))
	return len(p), nil
}

func (lwc *LogWriterCloser) Close() error {
	return nil
}

type LogProvider struct {
}

func (s *LogProvider) NewFile(p string) io.WriteCloser {
	log.Print(p)
	return NewLogWriterCloser()
}

func (s *LogProvider) Flush() {
}
