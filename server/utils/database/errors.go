package database

import "errors"

// EOD (End Of Data) is the error added to channels after Find operations and all documents have been read.
var EOD error = errors.New("EOD")
