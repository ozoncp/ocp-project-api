package ocp_repo_api_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOcpRepoApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "OcpRepoApi Suite")
}
