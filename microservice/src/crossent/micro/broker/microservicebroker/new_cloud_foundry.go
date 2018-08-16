package microservicebroker

import (
	"os"
	"github.com/concourse/cf-resource/out"
	"os/exec"
	"bytes"
)

type ReCloudFoundry struct{
	out.CloudFoundry
	instanceID string
	env map[string]string
}

func NewCloudFoundry(instanceID string, env map[string]string) *ReCloudFoundry {
	return &ReCloudFoundry{
		instanceID: instanceID,
		env:env,
	}
}

func (cf *ReCloudFoundry) PushApp(manifest string, fpath string, currentAppName string) error {
	args := []string{}

	var buffer bytes.Buffer
	buffer.WriteString(currentAppName)
	buffer.WriteString(cf.instanceID)
	appName := buffer.String()

	args = append(args, "push", appName, "-f", manifest, "--hostname", appName, "--no-start")

	if fpath != "" {
		stat, err := os.Stat(fpath)
		if err != nil {
			return err
		}
		if stat.IsDir() {
			return chdir(fpath, cf.cf(args...).Run)
		}

		// path is a zip file, add it to the args
		args = append(args, "-p", fpath)
	}

	err := cf.cf(args...).Run()

	if err != nil {
		return err
	}

	for k, v := range cf.env {
		err := cf.cf("set-env", appName, k, v).Run()
		if err != nil {
			return err
		}
	}

	return nil
}

func chdir(path string, f func() error) error {
	oldpath, err := os.Getwd()
	if err != nil {
		return err
	}
	err = os.Chdir(path)
	if err != nil {
		return err
	}
	defer os.Chdir(oldpath)

	return f()
}

func (cf *ReCloudFoundry) cf(args ...string) *exec.Cmd {
	cmd := exec.Command("cf", args...)
	cmd.Stdout = os.Stderr
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(), "CF_COLOR=true", "CF_DIAL_TIMEOUT=30")
	return cmd
}
