// Copyright (c) 2019 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package script

import (
	"fmt"
	"strings"
)

// CopyCommand represents a COPY directive which may have
// one of the following two forms:
//
//     COPY path0 path1 ... pathN
//     COPY paths:"path0 path1 ... pathN"
//
// The former uses no named parameters while the latter uses
// named parameter (i.e. paths)
type CopyCommand struct {
	cmd
}

// NewCopyCommand returns *CopyCommand
func NewCopyCommand(index int, rawArgs string) (*CopyCommand, error) {
	if err := validateRawArgs(CmdCopy, rawArgs); err != nil {
		return nil, err
	}

	// determine shape of directive
	var argMap map[string]string
	if strings.Contains(rawArgs, "paths:") {
		args, err := mapArgs(rawArgs)
		if err != nil {
			return nil, fmt.Errorf("COPY: %v", err)
		}
		argMap = args
	} else {
		argMap = map[string]string{"paths": rawArgs}
	}
	cmd := &CopyCommand{cmd: cmd{index: index, name: CmdCopy, args: argMap}}
	if err := validateCmdArgs(CmdCopy, argMap); err != nil {
		return nil, err
	}
	return cmd, nil
}

// Index is the position of the command in the script
func (c *CopyCommand) Index() int {
	return c.cmd.index
}

// Name represents the name of the command
func (c *CopyCommand) Name() string {
	return c.cmd.name
}

// Paths returned a []string of paths to copy
func (c *CopyCommand) Paths() []string {
	return spaceSep.Split(c.cmd.args["paths"], -1)
}

// Args returns a slice of raw command arguments
func (c *CopyCommand) Args() map[string]string {
	return c.cmd.args
}
