package controllers

import (
	"context"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schedulev1 "replica-operator/api/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// Status defines podScheduler condition status.
type Status string

// Defines podScheduler condition status.
const (
	TypeAvailable   Status = "Available"
	TypeProgressing Status = "Progressing"
	TypeDegraded    Status = "Degraded"
)

// GetPodScheduler gets the podScheduler from api server.
func (r *PodSchedulerReconciler) GetPodScheduler(ctx context.Context, req ctrl.Request, scheduler *schedulev1.PodScheduler) error {
	err := r.Get(ctx, req.NamespacedName, scheduler)
	if err != nil {
		return err
	}

	return nil
}

// SetInitialCondition sets the status condition of the TDSet to available initially
// when no condition exists yet.
func (r *PodSchedulerReconciler) SetInitialCondition(ctx context.Context, req ctrl.Request, scheduler *schedulev1.PodScheduler) error {
	if scheduler.Status.Conditions != nil || len(scheduler.Status.Conditions) != 0 {
		return nil
	}

	err := r.SetCondition(ctx, req, scheduler, TypeAvailable, "Starting reconciliation")

	return err
}

// SetCondition sets the status condition of the TDSet.
func (r *PodSchedulerReconciler) SetCondition(ctx context.Context, req ctrl.Request, scheduler *schedulev1.PodScheduler, condition Status, message string) error {
	log := log.FromContext(ctx)

	meta.SetStatusCondition(
		&scheduler.Status.Conditions,
		metav1.Condition{
			Type:   string(condition),
			Status: metav1.ConditionUnknown, Reason: "Reconciling",
			Message: message,
		},
	)

	if err := r.Status().Update(ctx, scheduler); err != nil {
		log.Error(err, "Failed to update podScheduler status")
		return err
	}

	if err := r.Get(ctx, req.NamespacedName, scheduler); err != nil {
		log.Error(err, "Failed to re-fetch podScheduler")
		return err
	}

	return nil
}
