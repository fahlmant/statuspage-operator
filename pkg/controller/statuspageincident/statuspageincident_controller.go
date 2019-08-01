package statuspageincident

import (
	"context"

	hackdayv1alpha1 "github.com/fahlmant/statuspage-operator/pkg/apis/hackday/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_statuspageincident")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new StatusPageIncident Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileStatusPageIncident{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("statuspageincident-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource StatusPageIncident
	err = c.Watch(&source.Kind{Type: &hackdayv1alpha1.StatusPageIncident{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileStatusPageIncident implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileStatusPageIncident{}

// ReconcileStatusPageIncident reconciles a StatusPageIncident object
type ReconcileStatusPageIncident struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a StatusPageIncident object and makes changes based on the state read
// and what is in the StatusPageIncident.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileStatusPageIncident) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling StatusPageIncident")

	// Fetch the StatusPageIncident instance
	instance := &hackdayv1alpha1.StatusPageIncident{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// If the type is scheduled create a scheduled incident aka Maintainance
	if instance.Spec.Type == hackdayv1alpha1.Scheduled {
		// Check data

		// Call api and create scheduled incident

		// Update the status of ourbject
		if true {
			reqLogger.Info("Scheduled Incident Created ", "Incident Name", instance.Spec.Name)
			instance.Status.Scheduled == true
			r.Client.Status().Update(context.TODO(), instance)
			if err != nil {
				return reconcile.Result{}, err
			}
		}

	}

	return reconcile.Result{}, nil
}
