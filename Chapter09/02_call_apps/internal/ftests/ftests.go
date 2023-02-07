package ftests

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

var errNoChanges = errors.New("no changes found")

var excludedDirs = regexp.MustCompile(`((^|/)vendor(/|$))|((^|/)testdata(/|$))`)

type Coordinator struct{}

func (c *Coordinator) Run(ctx context.Context, dir string) error {
	changeList, err := c.getChanges(ctx, dir)
	if err != nil {
		return err
	}

	pkgs, err := c.buildListOfChangedPackages(changeList)
	if err != nil {
		return err
	}

	pkgs, err = c.filterUnwantedPackages(pkgs)
	if err != nil {
		return err
	}

	baseDir, err := c.getGitBaseDir(ctx, dir)
	if err != nil {
		return err
	}

	c.runTests(ctx, baseDir, pkgs)

	return nil
}

func (c *Coordinator) getChanges(ctx context.Context, dir string) ([]byte, error) {
	// Ask Git for a list of all of the files we have changed.
	cmd := exec.CommandContext(ctx, "git", "diff", "--name-only", "-M100%", "master")

	// run the command in the supplied directory
	cmd.Dir = dir

	// use the current Environment
	cmd.Env = os.Environ()

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve changes from Git with err: %w", err)
	}

	return output, nil
}

func (c *Coordinator) buildListOfChangedPackages(changeList []byte) ([]string, error) {
	input := strings.TrimSpace(string(changeList))
	if input == "" {
		// short cut when there are no changes
		return nil, errNoChanges
	}

	// convert Git's output to a series of lines
	lines := strings.Split(input, "\n")

	var output []string
	dedupe := map[string]struct{}{}

	// convert the lines to directories and deduplicate
	for _, line := range lines {
		pkg := filepath.Dir(line)

		_, found := dedupe[pkg]
		if found {
			continue
		}

		dedupe[pkg] = struct{}{}
		output = append(output, pkg)
	}

	return output, nil
}

func (c *Coordinator) filterUnwantedPackages(pkgs []string) ([]string, error) {
	var output []string

	for _, pkg := range pkgs {
		// Remove directories like vendor/ and testdata/
		if excludedDirs.MatchString(pkg) {
			continue
		}

		output = append(output, pkg)
	}

	// short cut when we filtered out all of the changes
	if len(output) == 0 {
		return nil, errNoChanges
	}

	return output, nil
}

func (c *Coordinator) getGitBaseDir(ctx context.Context, dir string) (string, error) {
	// Ask Git for its base directory
	cmd := exec.CommandContext(ctx, "git", "rev-parse", "--show-toplevel")

	// run the command in the supplied directory
	cmd.Dir = dir

	// use the current Environment
	cmd.Env = os.Environ()

	// ignore errors so that test errors do not break this code
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}
func (c *Coordinator) runTests(ctx context.Context, baseDir string, pkgs []string) {
	for _, pkgDir := range pkgs {
		// Ask Git for a list of all of the files we have changed.
		cmd := exec.CommandContext(ctx, "go", "test", "-v", ".")

		// run the command in the supplied directory
		cmd.Dir = path.Join(baseDir, pkgDir)

		// use the current Environment
		cmd.Env = os.Environ()

		// ignore errors so that test errors do not break this code
		output, _ := cmd.CombinedOutput()

		_, _ = fmt.Fprint(os.Stdout, string(output))
	}
}
