// Copyright 2021 Authors of KubeArmor
// SPDX-License-Identifier: Apache-2.0

package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	securityv1 "github.com/kubearmor/KubeArmor/pkg/KubeArmorAuditPolicy/api/v1"
)

// KubeArmorAuditPolicyReconciler reconciles a KubeArmorAuditPolicy object
type KubeArmorAuditPolicyReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=security.kubearmor.com,resources=kubearmorauditpolicies,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=security.kubearmor.com,resources=kubearmorauditpolicies/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=security.kubearmor.com,resources=kubearmorauditpolicies/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the KubeArmorAuditPolicy object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *KubeArmorAuditPolicyReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *KubeArmorAuditPolicyReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&securityv1.KubeArmorAuditPolicy{}).
		Complete(r)
}
