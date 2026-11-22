package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/estratocloud/manifesting/internal"
	"github.com/estratocloud/manifesting/internal/deprecations"
)

type Args struct {
	configPath         internal.PathInterface
	generatedDirectory internal.PathInterface
	workingDirectory   internal.WorkingDirectoryInterface
	deprecations       *deprecations.Checker
	variables          Variables
}

type inputArgs struct {
	configPath         string
	generatedDirectory string
	workingDirectory   string
	failOnDeprecations bool
	variables          Variables
}

func GetArgs(argv []string) (*Args, error) {
	input := defineArgs(argv)

	args, err := validateArgs(input)
	if err != nil {
		return nil, err
	}

	return args, nil
}

func defineArgs(argv []string) *inputArgs {

	flags := flag.NewFlagSet("Manifesting", flag.ExitOnError)

	configPath := flags.String("config", "", "The location of the manifesting config file")
	generatedDirectory := flags.String("generated-dir", ".generated", "The location to write generated manifests to")
	workingDirectory := flags.String("working-dir", "", "Run as if manifesting was started in this path")
	failOnDeprecations := flags.Bool("fail-on-deprecations", false, "Throw an error if any deprecated functionality is used")

	var variables Variables
	flags.Var(&variables, "var", "Set a default project variable (eg --var=Replicas=1)")

	_ = flags.Parse(argv)

	return &inputArgs{
		configPath:         *configPath,
		generatedDirectory: *generatedDirectory,
		workingDirectory:   *workingDirectory,
		failOnDeprecations: *failOnDeprecations,
		variables:          variables,
	}
}

func validateArgs(input *inputArgs) (*Args, error) {

	args := &Args{}

	var err error

	workingDirectory := input.workingDirectory
	if workingDirectory == "" {
		workingDirectory, err = os.Getwd()
		if err != nil {
			return nil, err
		}
	}
	args.workingDirectory, err = internal.NewWorkingDirectory(workingDirectory)
	if err != nil {
		return nil, fmt.Errorf("unable to use the --working-dir '%s': %w", workingDirectory, err)
	}

	if input.configPath == "" {
		input.configPath = "manifesting.yaml"
	}
	args.configPath = args.workingDirectory.NewPath(input.configPath)
	err = args.configPath.ExistsOrError("unable to find the manifesting --config file at '%s'")
	if err != nil {
		return nil, err
	}

	args.generatedDirectory = args.workingDirectory.NewPath(input.generatedDirectory)

	args.deprecations = deprecations.New(input.failOnDeprecations)

	args.variables = input.variables

	return args, nil
}
