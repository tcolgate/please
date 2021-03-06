package intellij

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"core"
)

func TestProjectLocation(t *testing.T) {
	assert.True(t, strings.HasSuffix(projectLocation(), "plz-out/intellij/.idea"))
}

func TestModuleFileLocation(t *testing.T) {
	target := &core.BuildTarget{
		Label: core.BuildLabel{
			PackageName: "some/package", Name: "target", Subrepo: "",
		},
	}

	assert.True(t, strings.HasSuffix(moduleFileLocation(target), "plz-out/intellij/some/package/some_package_target.iml"))
}
