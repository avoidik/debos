package main

type AptAction struct {
	*BaseAction
	Recommends bool
	Packages   []string
}

func (apt *AptAction) Run(context *YaibContext) {
	aptOptions := []string{"-y"}

	if !apt.Recommends {
		aptOptions = append(aptOptions, "--no-install-recommends")
	}

	aptOptions = append(aptOptions, "install")
	aptOptions = append(aptOptions, apt.Packages...)

	options := []string{"-q", "-EDEBIAN_FRONTEND=noninteractive",
		"-D", context.rootdir, "apt-get"}
	options = append(options, aptOptions...)

	q := NewQemuHelper(*context)
	q.Setup()
	defer q.Cleanup()

	RunCommand("apt", "systemd-nspawn", options...)
}