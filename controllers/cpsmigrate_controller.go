/*
Copyright 2021 boxcube.

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
	sync "github.com/boxcube/zebra/controllers/sync_resource"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	zebrav1alpha1 "github.com/boxcube/zebra/api/v1alpha1"
)

// CpsMigrateReconciler reconciles a CpsMigrate object
type CpsMigrateReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=zebra.boxcube.io,resources=cpsmigrates,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=zebra.boxcube.io,resources=cpsmigrates/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=zebra.boxcube.io,resources=cpsmigrates/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the CpsMigrate object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *CpsMigrateReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)
	// set up signals so we handle the first shutdown signal gracefully
	//stopCh := uitls.SetupSignalHandler()

	//sync deployment
	sync.SyncDeploy(r)

	//sync statefulset

	//sync service

	//sync configmap

	//sync PersistentVolumeClaim

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CpsMigrateReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&zebrav1alpha1.CpsMigrate{}).
		Complete(r)
}
