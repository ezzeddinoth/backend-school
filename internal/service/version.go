package service

import (
	"school/internal/build"
	"fmt"
)

func GetVersion() string {
	v := build.Get()
	return fmt.Sprintf("%s_%s_%s", v.Version, v.Commit, v.BuildDate)
}
