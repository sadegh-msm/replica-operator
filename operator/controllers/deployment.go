package controllers

import (
	"context"
	v1 "k8s.io/api/apps/v1"
	schedulev1 "replica-operator/api/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

func (r *PodSchedulerReconciler) Deployment(ctx context.Context, req ctrl.Request, scheduler *schedulev1.PodScheduler) (*v1.Deployment, error) {
	log := log.FromContext(ctx)
	replicas, err := r.GetExpectedReplica(ctx, req, tdSet)
	if err != nil {
		log.Error(err, "failed to get expected replica")

		return nil, err
	}
}
