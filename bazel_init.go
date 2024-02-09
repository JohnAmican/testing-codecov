package main

import (
	"strings"
	"text/template"
)

type Language string

type httpArchive struct {
	Name   string
	Sha256 string
	Urls   []string
}

func (a *httpArchive) String() (string, error) {
	tmpl, err := template.New("http_archive").Parse(`http_archive(
		Name = "{{.Name}}",
		Sha256 = "{{.Sha256}}",
		Urls = [
			{{range .Urls}}{{.}},
			{{end}}
		],
	)`)
	if err != nil {
		return "", err
	}
	b := new(strings.Builder)
	if err = tmpl.Execute(b, a); err != nil {
		return "", err
	}
	return b.String(), nil
}

//goland:noinspection GoSnakeCaseUsage
var (
	rules_go_0_45_1 = &httpArchive{
		Name:   "io_bazel_rules_go",
		Sha256: "6734a719993b1ba4ebe9806e853864395a8d3968ad27f9dd759c196b3eb3abe8",
		Urls: []string{
			"https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.45.1/rules_go-v0.45.1.zip",
			"https://github.com/bazelbuild/rules_go/releases/download/v0.45.1/rules_go-v0.45.1.zip",
		},
	}
)

func main() {
	_, _ = template.New("WORKSPACE").Parse(`load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    Name = "io_bazel_rules_go",
    Sha256 = "6734a719993b1ba4ebe9806e853864395a8d3968ad27f9dd759c196b3eb3abe8",
    Urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.45.1/rules_go-v0.45.1.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.45.1/rules_go-v0.45.1.zip",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(version = "1.21.6")
`)
}
