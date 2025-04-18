/*
Copyright [2018] [jc3wish]

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

package mongo

import inputDriver "github.com/brokercap/Bifrost/input/driver"

const (
	VERSION         string = "v2.4.0"
	BIFROST_VERSION string = "v2.4.0"
)

const OutPutName = "mongo"

func init() {
	inputDriver.Register("mongo", NewInputPlugin, VERSION, BIFROST_VERSION)
}
