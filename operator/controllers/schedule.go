package controllers

import (
	"context"
	v1 "replica-operator/api/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"time"
)

func (r *PodSchedulerReconciler) GetExpectedReplica(ctx context.Context, req ctrl.Request, scheduler *v1.PodScheduler) (int32, error) {
	log := log.FromContext(ctx)

	if scheduler.Spec.SchedulingConfig != nil {
		if len(scheduler.Spec.SchedulingConfig) != 0 {
			now := time.Now()
			hour := now.Hour()

			log.Info("current server", "hour: ", hour, "time: ", now)
			for _, config := range scheduler.Spec.SchedulingConfig {
				if hour >= config.StartTime && hour < config.EndTime {
					return int32(config.Replica), nil
				}
			}
		}
	}

	return scheduler.Spec.DefaultReplica, nil
}
