/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	v1 "k8s.io/api/apps/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	schedulev1 "replica-operator/api/v1"
)

const (
	DefaultReconciliationIntervalInMinute = 5
)

// PodSchedulerReconciler reconciles a PodScheduler object
type PodSchedulerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=schedule.my.domain,resources=podschedulers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=schedule.my.domain,resources=podschedulers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=schedule.my.domain,resources=podschedulers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the PodScheduler object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *PodSchedulerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	log.Info("starting reconciliation")

	scheduler := &schedulev1.PodScheduler{}

	err := r.GetPodScheduler(ctx, req, scheduler)
	if err != nil {
		if apierrors.IsNotFound(err) {
			log.Info("podScheduler resource not found. Ignoring since object must be deleted")

			return ctrl.Result{}, nil
		}
		log.Error(err, "Failed to get podScheduler")

		return ctrl.Result{}, err
	}

	err = r.SetInitialCondition(ctx, req, scheduler)
	if err != nil {
		log.Error(err, "failed to set initial condition")

		return ctrl.Result{}, err
	}

	ok, err := r.DeploymentIfNotExist(ctx, req, scheduler)
	if err != nil {
		log.Error(err, "failed to deploy deployment for podScheduler")

		return ctrl.Result{}, err
	}
	if ok {
		return ctrl.Result{RequeueAfter: time.Minute}, nil
	}

	err = r.UpdateDeploymentReplica(ctx, req, scheduler)
	if err != nil {
		log.Error(err, "failed to update deployment for podScheduler")

		return ctrl.Result{}, err
	}

	interval := DefaultReconciliationIntervalInMinute
	if scheduler.Spec.IntervalMint != 0 {
		interval = int(scheduler.Spec.IntervalMint)
	}

	log.Info("ending reconciliation")

	return ctrl.Result{RequeueAfter: time.Minute * time.Duration(interval)}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *PodSchedulerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&schedulev1.PodScheduler{}).
		Owns(&v1.Deployment{}).
		Complete(r)
}
