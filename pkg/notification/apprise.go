package notification

import (
	"fmt"
	"os/exec"
)

type ApprisePayload struct {
	AppriseUrl string
}

type Apprise struct {
	Payload ApprisePayload
}

func (a *Apprise) Send(title string, message string) error {
	cmd := exec.Command("apprise", "-vv", "-b", message, a.Payload.AppriseUrl)
	if title != "" {
		cmd.Args = append(cmd.Args, "-t", title)
	}
	output, err := cmd.Output()
	if err != nil {
		return err
	}
	if string(output) == "ERROR: maybe apprise not found" {
		return fmt.Errorf("ERROR: maybe apprise not found")
	}
	return nil
}
