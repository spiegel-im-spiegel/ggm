package config

import (
	"io"
	"io/ioutil"

	"github.com/BurntSushi/toml"
	"github.com/spiegel-im-spiegel/ggm/errs"
)

//Config is configuration class
type Config struct {
	Node map[string]interface{} `toml:"node"`
	Edge map[string]interface{} `toml:"edge"`
}

//Decode returns Config instance from stream
func Decode(r io.Reader) (*Config, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, errs.Wrap(err, "error in config.Decode() function")
	}
	c := &Config{}
	if err := toml.Unmarshal(data, c); err != nil {
		return nil, errs.Wrap(err, "error in config.Decode() function")
	}
	return c, nil
}

/* Copyright 2019 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
