// Copyright (c) 2021-2024 Tenebris Technologies Inc.
// Use of this source code is governed by the MIT license.
// Please see the LICENSE file for details.

package easyconfig

import (
	"strconv"
)

// SetStr sets a string value
func (c *EasyConfig) SetStr(key string, value string) error {
	return c.set(key, value)
}

// SetInt sets an integer value
func (c *EasyConfig) SetInt(key string, value int) error {
	// Convert value to string
	v := strconv.Itoa(value)
	return c.set(key, v)
}

// SetInt64 sets an int64 value
func (c *EasyConfig) SetInt64(key string, value int64) error {
	// Convert value to string
	v := strconv.FormatInt(value, 10)
	return c.set(key, v)
}

// SetBool sets a boolean value
func (c *EasyConfig) SetBool(key string, value bool) error {
	b := "false"
	if value {
		b = "true"
	}
	return c.set(key, b)
}
