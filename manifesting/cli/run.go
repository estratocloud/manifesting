package cli

import (
	"fmt"
	"os"

	"github.com/estratocloud/manifesting/manifesting"
	"github.com/estratocloud/manifesting/manifesting/config"
	"sigs.k8s.io/yaml"
)

func Run() error {
	printProgramIntro()

	args, err := GetArgs(os.Args[1:])
	if err != nil {
		return err
	}

	data, err := args.configPath.ReadFile()
	if err != nil {
		return err
	}

	var conf config.Config
	if err := yaml.Unmarshal(data, &conf); err != nil {
		return err
	}
	printConfig(args)

	for _, environment := range conf.Environments {
		printEnvironmentDetails(args, environment)
		err := manifesting.GenerateManifest(environment, &conf, args.workingDirectory)
		if err != nil {
			return err
		}
		fmt.Println()
	}

	return nil
}

func printProgramIntro() {
	fmt.Println("🧘️ Manifesting", manifesting.Version, "🪄")
	fmt.Println()

	fmt.Println("©️  2026 Estrato Cloud Ltd")
	fmt.Println()
}

func printConfig(args *Args) {
	fmt.Println("Loading config from", args.configPath.GetFullyQualifiedPath())
	fmt.Println()
}

func printEnvironmentDetails(args *Args, environment *config.Environment) {
	fmt.Println("Generating manifest for...")
	fmt.Println("\tEnvironment:", environment.Name)
	fmt.Println("\tFile Path:  ", environment.GetOutputPath(args.workingDirectory).GetPath())
	if environment.EnvFrom != "" {
		fmt.Println("\tEnv Vars:   ", environment.EnvFrom)
	}
}
