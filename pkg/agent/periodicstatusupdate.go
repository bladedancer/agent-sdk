package agent

import (
	"time"

	"github.com/Axway/agent-sdk/pkg/jobs"
	"github.com/Axway/agent-sdk/pkg/util/errors"
	hc "github.com/Axway/agent-sdk/pkg/util/healthcheck"
	"github.com/Axway/agent-sdk/pkg/util/log"
)

const defaultCheckInterval time.Duration = 30000000000 // 30s

type periodicStatusUpdate struct {
	jobs.Job
	previousActivityTime time.Time
	currentActivityTime  time.Time
}

var statusUpdate *periodicStatusUpdate

func (su *periodicStatusUpdate) Ready() bool {
	if runStatusUpdateCheck() != nil {
		return false
	}
	log.Debug("Periodic status update is ready")
	su.currentActivityTime = time.Now()
	su.previousActivityTime = su.currentActivityTime
	return true
}

func (su *periodicStatusUpdate) Status() error {
	// error out if the agent name does not exist
	err := runStatusUpdateCheck()
	if err != nil {
		return err
	}
	return nil
}

func (su *periodicStatusUpdate) Execute() error {
	// error out if the agent name does not exist
	err := runStatusUpdateCheck()
	if err != nil {
		log.Error(errors.ErrStatusUpdate)
		log.Debug("Status update failed")
		return err
	} else {
		// if the last timestamp for an event has changed, update the status
		if time.Time(su.currentActivityTime).After(time.Time(su.previousActivityTime)) {
			log.Debug("Activity change detected, updating status")
			UpdateStatus(AgentRunning, "")
			su.previousActivityTime = su.currentActivityTime
		}
	}
	return nil
}

//StartPeriodicStatusUpdate - starts a job that runs the periodic status updates
func StartPeriodicStatusUpdate() {
	interval := defaultCheckInterval

	if hc.GetStatusConfig() != nil {
		interval = hc.GetStatusConfig().GetHealthCheckInterval()
	}
	statusUpdate = &periodicStatusUpdate{}
	_, err := jobs.RegisterIntervalJob(statusUpdate, interval)

	if err != nil {
		log.Error(errors.Wrap(ErrStartingPeriodicStatusUpdate, err.Error()))
	}
}

//runStatusUpdateCheck - returns an error if agent name is blank
func runStatusUpdateCheck() error {
	if agent.cfg.GetAgentName() == "" {
		return ErrStartingPeriodicStatusUpdate
	}
	return nil
}

//UpdateLocalActivityTime - updates the local activity timestamp for the event to compare against
func UpdateLocalActivityTime() {
	statusUpdate.currentActivityTime = time.Now()
}
