// Copyright 2012, 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package cleaner

import (
	"fmt"

	"launchpad.net/juju-core/log"
	"launchpad.net/juju-core/state"
	"launchpad.net/juju-core/state/api"
	"launchpad.net/juju-core/worker"
)

// Cleaner is responsible for cleaning up the state.
type Cleaner struct {
	st *state.State
}

// NewCleaner returns a worker.NotifyWorker that runs state.Cleanup()
// if the CleanupWatcher signals documents marked for deletion.
func NewCleaner(st *state.State) worker.NotifyWorker {
	return worker.NewNotifyWorker(&Cleaner{st: st})
}

func (c *Cleaner) String() string {
	return fmt.Sprintf("cleaner")
}

func (c *Cleaner) SetUp() (api.NotifyWatcher, error) {
	return c.st.WatchCleanups(), nil
}

func (c *Cleaner) Handle() error {
	if err := c.st.Cleanup(); err != nil {
		log.Errorf("worker/cleaner: cannot cleanup state: %v", err)
	}
	// We do not return the err from Cleanup, because we don't want to stop
	// the loop as a failure
	return nil
}

func (c *Cleaner) TearDown() error {
	// Nothing to cleanup, only state is the watcher
	return nil
}
