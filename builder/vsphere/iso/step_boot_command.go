package iso

import (
	"context"
	"fmt"
	"github.com/hashicorp/packer/builder/vsphere/driver"
	"github.com/hashicorp/packer/common/bootcommand"
	"github.com/hashicorp/packer/helper/multistep"
	"github.com/hashicorp/packer/packer"
	"github.com/hashicorp/packer/template/interpolate"
	"golang.org/x/mobile/event/key"
	"time"
)

type BootConfig struct {
	bootcommand.BootConfig `mapstructure:",squash"`
	HTTPIP                 string `mapstructure:"http_ip"`
}

type bootCommandTemplateData struct {
	HTTPIP   string
	HTTPPort int
	Name     string
}

func (c *BootConfig) Prepare(ctx *interpolate.Context) []error {
	var errs []error

	if c.BootWait == 0 {
		c.BootWait = 10 * time.Second
	}

	c.BootConfig.Prepare(ctx)

	return errs
}

type StepBootCommand struct {
	Config *BootConfig
	VMName string
	Ctx    interpolate.Context
}

func (s *StepBootCommand) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	debug := state.Get("debug").(bool)
	ui := state.Get("ui").(packer.Ui)
	vm := state.Get("vm").(*driver.VirtualMachine)

	if s.Config.BootCommand == nil {
		return multistep.ActionContinue
	}

	// Wait the for the vm to boot.
	if int64(s.Config.BootWait) > 0 {
		ui.Say(fmt.Sprintf("Waiting %s for boot...", s.Config.BootWait.String()))
		select {
		case <-time.After(s.Config.BootWait):
			break
		case <-ctx.Done():
			return multistep.ActionHalt
		}
	}

	var pauseFn multistep.DebugPauseFn
	if debug {
		pauseFn = state.Get("pauseFn").(multistep.DebugPauseFn)
	}

	port := state.Get("http_port").(int)
	if port > 0 {
		ip := state.Get("http_ip").(string)
		s.Ctx.Data = &bootCommandTemplateData{
			ip,
			port,
			s.VMName,
		}
		ui.Say(fmt.Sprintf("HTTP server is working at http://%v:%v/", ip, port))
	}

	sendCodes := func(code key.Code, down bool) error {
		var keyAlt, keyCtrl, keyShift bool

		switch code {
		case key.CodeLeftAlt:
			// <leftAltOn>
			keyAlt = down
		case key.CodeLeftControl:
			// <leftCtrlOn>
			keyCtrl = down
		default:
			keyShift = down
		}

		_, err := vm.TypeOnKeyboard(driver.KeyInput{
			Scancode: code,
			Ctrl:     keyCtrl,
			Alt:      keyAlt,
			Shift:    keyShift,
		})
		if err != nil {
			return fmt.Errorf("error typing a boot command: %v", err)
		}
		return nil
	}
	d := bootcommand.NewUSBDriver(sendCodes, s.Config.BootGroupInterval)

	ui.Say("Typing boot command...")
	flatBootCommand := s.Config.FlatBootCommand()
	command, err := interpolate.Render(flatBootCommand, &s.Ctx)
	if err != nil {
		err := fmt.Errorf("Error preparing boot command: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	seq, err := bootcommand.GenerateExpressionSequence(command)
	if err != nil {
		err := fmt.Errorf("Error generating boot command: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	if err := seq.Do(ctx, d); err != nil {
		err := fmt.Errorf("Error running boot command: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	if pauseFn != nil {
		pauseFn(multistep.DebugLocationAfterRun, fmt.Sprintf("boot_command: %s", command), state)
	}

	return multistep.ActionContinue
}

func (s *StepBootCommand) Cleanup(_ multistep.StateBag) {}
