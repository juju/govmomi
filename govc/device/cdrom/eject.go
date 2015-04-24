/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cdrom

import (
	"flag"

	"github.com/juju/govmomi/govc/cli"
	"github.com/juju/govmomi/govc/flags"
	"golang.org/x/net/context"
)

type eject struct {
	*flags.VirtualMachineFlag

	device string
}

func init() {
	cli.Register("device.cdrom.eject", &eject{})
}

func (cmd *eject) Register(f *flag.FlagSet) {
	f.StringVar(&cmd.device, "device", "", "CD-ROM device name")
}

func (cmd *eject) Process() error { return nil }

func (cmd *eject) Description() string {
	return `Eject media from CD-ROM device.

If device is not specified, the first CD-ROM device is used.`
}

func (cmd *eject) Run(f *flag.FlagSet) error {
	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	if vm == nil {
		return flag.ErrHelp
	}

	devices, err := vm.Device(context.TODO())
	if err != nil {
		return err
	}

	c, err := devices.FindCdrom(cmd.device)
	if err != nil {
		return err
	}

	return vm.EditDevice(context.TODO(), devices.EjectIso(c))
}
