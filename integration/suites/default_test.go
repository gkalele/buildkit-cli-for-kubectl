// Copyright (C) 2020 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0
package suites

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/vmware-tanzu/buildkit-cli-for-kubectl/integration/common"
)

type DefaultSuite struct{ common.BaseSuite }

func (s *DefaultSuite) SetupTest() {
	// For the "default" scenario, we rely on the initial test case to establish the builder with defaults
}

func TestDefaultSuite(t *testing.T) {
	common.Skipper(t)
	//t.Parallel() // TODO - tests fail if run in parallel, may be actual race bug
	suite.Run(t, &DefaultSuite{
		BaseSuite: common.BaseSuite{
			Name: "buildkit", // TODO pull this from the actual default name
		},
	})
}
