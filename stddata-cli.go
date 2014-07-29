package main

import "os"

func main() {
	// os.Exit will terminate the program at the place of call without running
	// any deferred cleanup statements. It might cause unintended effects. To
	// be safe, we wrap the program in run() and only os.Exit() outside the
	// wrapper. Be careful not to indirectly trigger os.Exit() in the program,
	// notably via log.Fatal() and on flag.Parse() where the default behavior
	// is ExitOnError.
	os.Exit(run())
}

// Run the program and return exit code.
func run() int {
	// TODO: this would be the right place to load a config file. The config
	// file could specify which stddata providers to load. So far, there's
	// not much else to configure.
	// TODO: what's with the flags variable?
	flags, err := config()
	if err != nil {
		errf("config error: %v\n", err)
		return 1
	}

	switch cmd := flags.Arg(0); cmd {
	case "launch":
		return cmdLaunch()
	case "stop":
		return cmdStop()
	case "reload":
		return cmdReload()
	case "help":
		flags.Usage()
		return 0
	case "":
		usageShort()
		return 0
	default:
		errf("Unknown command %q\n", cmd)
		usageShort()
		return 1
	}
}
