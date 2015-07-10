// +build !debug

package linode

import "net/url"

func debug(output string) {}

func record(action string, vals url.Values, resp []byte) {}
