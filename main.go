// A generated module for Quatrevm functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/quatrevm/internal/dagger"
)

type Quatrevm struct{}


// Returns a container that echoes whatever string argument is provided
func (m *Quatrevm) ContainerEcho(stringArg string) *dagger.Container {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", stringArg})
}

// Returns lines that match a pattern in the files of the provided Directory
func (m *Quatrevm) GrepDir(ctx context.Context, directoryArg *dagger.Directory) (string, error) {
	return dag.Container().
		From("ubuntu:16.04").
		WithMountedDirectory("/mnt", directoryArg).
		WithWorkdir("/mnt").
		WithExec([]string{"ls", "-l", "."}).
		Stdout(ctx)
}

func (m *Quatrevm) Run(ctx context.Context, directoryArg *dagger.Directory) *dagger.Service {
	return dag.Container().
		From("ubuntu:16.04").
		WithMountedDirectory("/mnt", directoryArg).
		WithWorkdir("/mnt").
		WithExec(
		[]string{"apt-get", "update", "-y"},
	).WithExec(
		[]string{"apt-get", "install", "qemu", "-y"}, // We hope it gets cached
	).WithExec(
		[]string{"qemu-system-i386", "-hda", "winxp.img", "-cdrom", "STARCRAFT.ISO", "-cpu", "pentium3", "-m", "512", "-vga", "cirrus", "-net", "nic,model=pcnet", "-net", "user"}, dagger.ContainerWithExecOpts{
			InsecureRootCapabilities: true,
		},
	).AsService(dagger.ContainerAsServiceOpts{Args: []string{"python", "-m", "http.server", "8080"}})


}

