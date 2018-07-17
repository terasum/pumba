package kubernetes

import (
	"context"

	"github.com/alexei-led/pumba/pkg/chaos"
	log "github.com/sirupsen/logrus"
)

// PatchCommand `kubernetes patch` command
type PatchCommand struct {
	dryRun bool
}

// NewPatchCommand create new Patch command instance
func NewPatchCommand(dryRun bool) (chaos.Command, error) {
	patch := &PatchCommand{dryRun}
	return patch, nil
}

// Run patch command
func (p *PatchCommand) Run(ctx context.Context, random bool) error {
	log.Debug("removing all matching containers")
	log.WithFields(log.Fields{
		"names": p,
	}).Debug("listing matching kube resources")
	// containers, err := container.ListNContainers(ctx, r.client, r.names, r.pattern, r.limit)
	// if err != nil {
	// 	log.WithError(err).Error("failed to list containers")
	// 	return err
	// }
	// if len(containers) == 0 {
	// 	log.Warning("no containers to remove")
	// 	return nil
	// }

	// select single random container from matching container and replace list with selected item
	if random {
		log.Debug("selecting single random whatever")
		// if c := container.RandomContainer(containers); c != nil {
		// 	containers = []container.Container{*c}
		// }
	}

	// for _, container := range containers {
	// 	log.WithFields(log.Fields{
	// 		"container": container,
	// 		"force":     r.force,
	// 		"links":     r.links,
	// 		"volumes":   r.volumes,
	// 	}).Debug("removing container")
	// 	err := r.client.RemoveContainer(ctx, container, r.force, r.links, r.volumes, r.dryRun)
	// 	if err != nil {
	// 		log.WithError(err).Error("failed to remove container")
	// 		return err
	// 	}
	// }
	return nil
}
