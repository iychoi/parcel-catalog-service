/*
Copyright 2020 CyVerse
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

package service

import "fmt"

// Config holds the parameters list which can be configured
type Config struct {
	Address     string // service address
	Port        int    // service port
	RequestRoot string // service request root "/"
}

// MakeServiceAddress returns service address
func (service *Service) MakeServiceAddress() string {
	return MakeServiceAddress(service.config)
}

// MakeServiceAddress returns service address
func MakeServiceAddress(conf *Config) string {
	return fmt.Sprintf("%s:%d", conf.Address, conf.Port)
}
