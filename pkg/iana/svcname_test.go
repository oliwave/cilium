// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

//go:build !privileged_tests
// +build !privileged_tests

package iana

import (
	"testing"

	. "gopkg.in/check.v1"
)

type IANATestSuite struct{}

var _ = Suite(&IANATestSuite{})

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

func (s *IANATestSuite) TestIsSvcName(c *C) {
	c.Assert(IsSvcName(""), Equals, false)                 // Too short
	c.Assert(IsSvcName("1234567890abcdef"), Equals, false) // Too long
	c.Assert(IsSvcName("1"), Equals, false)                // Missing letter
	c.Assert(IsSvcName("1a"), Equals, true)
	c.Assert(IsSvcName("Z"), Equals, true)
	c.Assert(IsSvcName("a9"), Equals, true)
	c.Assert(IsSvcName("a-9"), Equals, true)
	c.Assert(IsSvcName("a--9"), Equals, false) // Two consecutive dashes
	c.Assert(IsSvcName("-a9"), Equals, false)  // Begins with a dash
	c.Assert(IsSvcName("a9-"), Equals, false)  // Ends with a dash
	c.Assert(IsSvcName("a-b9-1"), Equals, true)
	c.Assert(IsSvcName("1-a-9"), Equals, true)
	c.Assert(IsSvcName("a-b-c-d-e-f"), Equals, true)
	c.Assert(IsSvcName("1-2-3-4"), Equals, false) // No letter(s)
}
