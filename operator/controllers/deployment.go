package controllers

import (
	"context"
	"fmt"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	schedulev1 "replica-operator/api/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

func (r *PodSchedulerReconciler) Deployment(ctx context.Context, req ctrl.Request, scheduler *schedulev1.PodScheduler) (*v1.Deployment, error) {
	log := log.FromContext(ctx)
	replicas, err := r.GetExpectedReplica(ctx, req, scheduler)
	if err != nil {
		log.Error(err, "failed to get expected replica")

		return nil, err
	}

	labels := map[string]string{
		"app.kubernetes.io/name":       "podScheduler",
		"app.kubernetes.io/instance":   scheduler.Name,
		"app.kubernetes.io/version":    "v1",
		"app.kubernetes.io/part-of":    "podScheduler-operator",
		"app.kubernetes.io/created-by": "controller-manager",
	}

	dep := &v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      scheduler.Name,
			Namespace: scheduler.Namespace,
		},
		Spec: v1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image:           scheduler.Spec.Container.Image,
						Name:            scheduler.Name,
						ImagePullPolicy: corev1.PullIfNotPresent,
						Ports: []corev1.ContainerPort{{
							ContainerPort: int32(scheduler.Spec.Container.Port),
							Name:          "podScheduler",
						}},
					}},
				},
			},
		},
	}
	// Set the ownerRef for the Deployment
	if err := ctrl.SetControllerReference(scheduler, dep, r.Scheme); err != nil {
		log.Error(err, "failed to set controller owner reference")
		return nil, err
	}

	return dep, nil
}

func (r *PodSchedulerReconciler) DeploymentIfNotExist(ctx context.Context, req ctrl.Request, scheduler *schedulev1.PodScheduler) (bool, error) {
	log := log.FromContext(ctx)

	dep := &v1.Deployment{}

	err := r.Get(ctx, types.NamespacedName{Name: scheduler.Name, Namespace: scheduler.Namespace}, dep)
	if err != nil && apierrors.IsNotFound(err) {
		dep, err := r.Deployment(ctx, req, scheduler)
		if err != nil {
			log.Error(err, "Failed to create new Deployment for podScheduler")

			err = r.SetCondition(ctx, req, scheduler, TypeAvailable, fmt.Sprintf("Failed to create Deployment for podScheduler (%s): (%s)", scheduler.Name, err))
			if err != nil {
				return false, err
			}
		}

		log.Info("Creating a new Deployment", "Deployment.Namespace: ", dep.Namespace, "Deployment.Name: ", dep.Name)

		if err = r.Create(ctx, dep); err != nil {
			log.Error(err, "Failed to create new Deployment", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
			return false, err
		}

		if err = r.GetPodScheduler(ctx, req, scheduler); err != nil {
			log.Error(err, "Failed to re-fetch podScheduler")
			return false, err
		}

		if err = r.SetCondition(ctx, req, scheduler, TypeProgressing, fmt.Sprintf("Created Deployment for the TDSet: (%s)", scheduler.Name)); err != nil {
			return false, err
		}
		return true, nil
	}

	if err != nil {
		log.Error(err, "Failed to get Deployment")

		return false, err
	}

	return false, nil
}

func (r *PodSchedulerReconciler) UpdateDeploymentReplica(ctx context.Context, req ctrl.Request, scheduler *schedulev1.PodScheduler) error {
	log := log.FromContext(ctx)

	dep := &v1.Deployment{}

	if err := r.Get(ctx, types.NamespacedName{Name: scheduler.Name, Namespace: scheduler.Namespace}, dep); err != nil {
		log.Error(err, "Failed to get Deployment")
		return err
	}

	replicas, err := r.GetExpectedReplica(ctx, req, scheduler)
	if err != nil {
		log.Error(err, "failed to get expected replica")

		return err
	}

	if replicas == *dep.Spec.Replicas {
		return nil
	}

	log.Info("Updating a Deployment replica", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
	dep.Spec.Replicas = &replicas

	if err = r.Update(ctx, dep); err != nil {
		log.Error(err, "Failed to update Deployment", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)

		if err = r.GetPodScheduler(ctx, req, scheduler); err != nil {
			log.Error(err, "Failed to re-fetch podScheduler")
			return err
		}

		if err = r.SetCondition(ctx, req, scheduler, TypeProgressing, fmt.Sprintf("Failed to update replica for the TDSet (%s): (%s)", scheduler.Name, err)); err != nil {
			return err
		}

		return nil
	}

	if err = r.GetPodScheduler(ctx, req, scheduler); err != nil {
		log.Error(err, "Failed to re-fetch podScheduler")
		return err
	}

	if err = r.SetCondition(ctx, req, scheduler, TypeProgressing, fmt.Sprintf("Updated replica for the podScheduler (%s)", scheduler.Name)); err != nil {
		return err
	}

	return nil
}
