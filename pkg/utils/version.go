package utils

var Version string //will be set by -ldflags
/*go build|run|install [...] -ldflags "-X github.com/cspengl/conmake/pkg/utils.Version=<version>" [...]*/
