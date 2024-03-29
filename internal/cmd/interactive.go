// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: Copyright 2023 Chainguard Inc

package cmd

import (
	"fmt"

	"github.com/chainguard-dev/bomshell/pkg/shell"
	"github.com/chainguard-dev/bomshell/pkg/ui"
	"github.com/spf13/cobra"
	"sigs.k8s.io/release-utils/version"
)

func interactiveCommand() *cobra.Command {
	type interactiveOpts = struct {
		commandLineOptions
		sboms []string
	}
	opts := &interactiveOpts{
		sboms: []string{},
	}
	execCmd := &cobra.Command{
		PersistentPreRunE: initLogging,
		Short:             "Launch bomshell interactive workbench (experimental)",
		Long: `bomshell interactive sbom.spdx.json → Launch the bomshell interactive workbench

The interactive subcommand launches the bomshell interactive workbench
`,
		Use:           "interactive [sbom.spdx.json...] ",
		SilenceUsage:  true,
		SilenceErrors: true,
		Version:       version.GetVersionInfo().GitVersion,
		RunE: func(cmd *cobra.Command, args []string) error {
			return launchInteractive(commandLineOpts)
		},
	}

	commandLineOpts.AddFlags(execCmd)
	opts.commandLineOptions = *commandLineOpts

	return execCmd
}

func launchInteractive(_ *commandLineOptions) error {
	i, err := ui.NewInteractive(
		shell.Options{
			SBOMs:  commandLineOpts.sboms,
			Format: shell.DefaultFormat,
		},
	)
	if err != nil {
		return fmt.Errorf("creating interactive env: %w", err)
	}

	// Start the interactive shell
	if err := i.Start(); err != nil {
		return fmt.Errorf("executing interactive mode: %w", err)
	}
	return nil
}
