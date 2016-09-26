// Copyright 2016 AlexStocks(https://github.com/AlexStocks).
// All rights reserved.  Use of this source code is
// governed by a BSD-style license.

// packaeg gxos encapsulates os related functions.
package gxos

import (
	"os"
)

func CreateDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0777)
		if err != nil {
			return err
		}
	}

	return nil
}