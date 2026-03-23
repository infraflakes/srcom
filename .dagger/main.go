package main

import (
	"context"
	"srcom/internal/dagger"
)

type Srcom struct{}

func (m *Srcom) Build(ctx context.Context, source *dagger.Directory) *dagger.File {
	return dag.Container().
		From("ubuntu:22.04").
		WithEnvVariable("DEBIAN_FRONTEND", "noninteractive").
		WithExec([]string{"apt-get", "update"}).
		WithExec([]string{"apt-get", "install", "-y",
			"meson", "ninja-build", "gcc", "pkg-config",
			"libx11-dev", "libxcb1-dev", "libxcb-composite0-dev", "libxcb-damage0-dev",
			"libxcb-dpms0-dev", "libxcb-glx0-dev", "libxcb-image0-dev", "libxcb-present-dev",
			"libxcb-randr0-dev", "libxcb-render0-dev", "libxcb-render-util0-dev",
			"libxcb-shape0-dev", "libxcb-xfixes0-dev", "libxcb-xinerama0-dev",
			"libxcb-ewmh-dev", "libxcb-icccm4-dev",
			"libgl-dev", "libegl-dev", "libepoxy-dev",
			"libev-dev", "libpcre2-dev", "libpixman-1-dev",
			"libconfig-dev", "libdbus-1-dev", "uthash-dev", "libx11-xcb-dev",
			"libxcb-util-dev", "cmake", "git",
		}).
		WithDirectory("/src", source.WithoutDirectory("build")).
		WithWorkdir("/src").
		WithExec([]string{"meson", "setup", "build", "--buildtype=release", "-Ddefault_library=static", "-Ddbus=false"}).
		WithExec([]string{"ninja", "-C", "build"}).
		File("build/src/srcom")
}
