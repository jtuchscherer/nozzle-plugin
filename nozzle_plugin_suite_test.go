package main_test

import (
	"github.com/cloudfoundry/cli/testhelpers/plugin_builder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestNozzlePlugin(t *testing.T) {
	RegisterFailHandler(Fail)
	plugin_builder.BuildTestBinary(".", "main")
	RunSpecs(t, "NozzlePlugin Suite")
}
