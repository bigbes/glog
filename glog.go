// Go support for leveled logs, analogous to https://code.google.com/p/google-glog/
//
// Copyright 2013 Google Inc. All Rights Reserved.
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

// Package bigbes/glog implements soveren/log wrapper for glog (to use in packages that use other logging libraries)

package glog

import (
	"sync"

	"github.com/insolar/vanilla/throw"
	"github.com/soverenio/log"
	"github.com/soverenio/log/global"
)

var logWrapperOnce sync.Once
var logWrapper log.Logger

func logWrapperInit() {
	logWrapperOnce.Do(func() {
		logWrapper = global.Logger().Copy().
			WithoutInheritedFields().
			WithField("loginstance", "glog").
			WithSkipFrameCount(1).
			MustBuild()
	})
}

func Flush() { global.Flush() }

func CopyStandardLogTo(name string) { panic(throw.NotImplemented()) }

type Level int32
type Verbose bool

func V(level Level) Verbose                  { return true }
func (v Verbose) Info(args ...interface{})   { logWrapperInit(); logWrapper.Info(args...) }
func (v Verbose) Infoln(args ...interface{}) { logWrapperInit(); logWrapper.Info(args...) }
func (v Verbose) Infof(format string, args ...interface{}) {
	logWrapperInit()
	logWrapper.Infof(format, args...)
}

func Info(args ...interface{})                 { logWrapperInit(); logWrapper.Info(args...) }
func Infoln(args ...interface{})               { logWrapperInit(); logWrapper.Info(args...) }
func Infof(format string, args ...interface{}) { logWrapperInit(); logWrapper.Infof(format, args...) }
func Warning(args ...interface{})              { logWrapperInit(); logWrapper.Warn(args...) }
func Warningln(args ...interface{})            { logWrapperInit(); logWrapper.Warn(args...) }
func Warningf(format string, args ...interface{}) {
	logWrapperInit()
	logWrapper.Warnf(format, args...)
}

func Error(args ...interface{})                 { logWrapperInit(); logWrapper.Error(args...) }
func Errorln(args ...interface{})               { logWrapperInit(); logWrapper.Error(args...) }
func Errorf(format string, args ...interface{}) { logWrapperInit(); logWrapper.Errorf(format, args...) }
func Fatal(args ...interface{})                 { logWrapperInit(); logWrapper.Fatal(args...) }
func Fatalln(args ...interface{})               { logWrapperInit(); logWrapper.Fatal(args...) }
func Fatalf(format string, args ...interface{}) { logWrapperInit(); logWrapper.Fatalf(format, args...) }
func Exit(args ...interface{})                  { logWrapperInit(); logWrapper.Fatal(args...) }
func Exitln(args ...interface{})                { logWrapperInit(); logWrapper.Fatal(args...) }
func Exitf(format string, args ...interface{})  { logWrapperInit(); logWrapper.Fatalf(format, args...) }

func InfoDepth(depth int, args ...interface{}) {
	logWrapperInit()
	localWrapper := logWrapper.Copy().
		WithSkipFrameCount(1 + depth).
		MustBuild()
	localWrapper.Info(args...)
}

func WarningDepth(depth int, args ...interface{}) {
	logWrapperInit()
	localWrapper := logWrapper.Copy().
		WithSkipFrameCount(1 + depth).
		MustBuild()
	localWrapper.Warn(args...)
}

func ErrorDepth(depth int, args ...interface{}) {
	logWrapperInit()
	localWrapper := logWrapper.Copy().
		WithSkipFrameCount(1 + depth).
		MustBuild()
	localWrapper.Error(args...)
}

func FatalDepth(depth int, args ...interface{}) {
	logWrapperInit()
	localWrapper := logWrapper.Copy().
		WithSkipFrameCount(1 + depth).
		MustBuild()
	localWrapper.Fatal(args...)
}

func ExitDepth(depth int, args ...interface{}) {
	logWrapperInit()
	localWrapper := logWrapper.Copy().
		WithSkipFrameCount(1 + depth).
		MustBuild()
	localWrapper.Fatal(args...)
}
