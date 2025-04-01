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

// qemu-system-x86_64 --enable-kvm -hda /home/mathias/vm/winxp.img -m 6144 \
//     -net user -cdrom /home/mathias/vm/winxp.iso -boot d \
//     -rtc base=localtime,clock=host -smp cores=4,threads=4 \
//     -usb -device usb-tablet \
//     -net user,smb=$HOME \
//     -vga qxl -device virtio-serial-pci -spice port=5930,disable-ticketing=on -device virtserialport,chardev=spicechannel0,name=com.redhat.spice.0 -chardev spicevmc,id=spicechannel0,name=vdagent

func (m *Quatrevm) Run(ctx context.Context, directoryArg *dagger.Directory) *dagger.Service {
	base, err := directoryArg.Name(ctx)
	if err != nil {
		panic(err)
	}
	disk := base + "/winxp.img"
	// cdrom := base + "/winxp.iso"
	return dag.Container().
		From("ubuntu:16.04").
		WithMountedDirectory("/mnt", directoryArg).
		WithWorkdir("/mnt").
		WithExec(
			[]string{"apt-get", "update", "-y"},
		).WithExec(
		[]string{"apt-get", "install", "qemu", "-y"}, // We hope it gets cached
	).WithExposedPort(5930).
		AsService(dagger.ContainerAsServiceOpts{Args: []string{
			"qemu-system-x86_64",
			"--enable-kvm",
			"-hda", disk,
			"-m", "6144",
			"-net", "user",
			// "-cdrom", cdrom,
			"-boot", "d",
			"-rtc", "base=localtime,clock=host",
			"-smp", "cores=1,threads=1",
			"-usb",
			"-device", "usb-tablet",
			"-vga", "qxl",
			"-device", "virtio-serial-pci",
			"-spice", "port=5930,disable-ticketing=on",
			"-device", "virtserialport,chardev=spicechannel0,name=com.redhat.spice.0",
			"-chardev", "spicevmc,id=spicechannel0,name=vdagent",
		}, InsecureRootCapabilities: true})

}
