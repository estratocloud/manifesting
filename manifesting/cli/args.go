package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/estratocloud/manifesting/internal"
)

type Args struct {
	configPath       internal.PathInterface
	workingDirectory internal.WorkingDirectoryInterface
}

type inputArgs struct {
	configPath       string
	workingDirectory string
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
	workingDirectory := flags.String("working-dir", "", "Run as if manifesting was started in this path")

	_ = flags.Parse(argv)

	return &inputArgs{
		configPath:       *configPath,
		workingDirectory: *workingDirectory,
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

	return args, nil
}
