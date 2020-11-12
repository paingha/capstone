// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"bitbucket.com/irb/api/config"
)

func TestSendMail(t *testing.T) {
	err := config.InitConfig(&cfg)
	if err != nil {
		t.Error(err)
	}

}
