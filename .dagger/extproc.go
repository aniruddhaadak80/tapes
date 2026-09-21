package main

import "context"

// extprocPackages is the ext_proc adapter's own test surface: the service
// package, its header parser, and the replay server that exercises captured
// bundles through it.
var extprocPackages = []string{
	"./extproc/...",
	"./cli/extproc-replay-server/...",
}

// CheckExtproc runs the ext_proc adapter's test suite.
//
// It is separate from Test, and bound to no database, because none of these
// tests need one — extproc talks to ingest over HTTP and is tested against
// fixtures. That is what makes it affordable on every PR, which matters: Test
// binds Postgres and for that reason does not run on the PR path, so folding
// extproc's tests into it would have quietly retired the gate that used to run
// on every extproc change. This keeps that gate, including the wire-capture
// fidelity tests, which read the committed recordings and therefore always run
// rather than skipping.
//
// +check
func (t *Tapes) CheckExtproc(ctx context.Context) (string, error) {
	args := append([]string{"go", "test", "-count=1"}, extprocPackages...)

	return t.goContainer().
		WithExec(args).
		Stdout(ctx)
}
