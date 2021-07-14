package podhealthcheck

import (
	"context"
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/labels"

	"github.com/riete/go-tools/notify"
	monitorv1 "github.com/riete/kubernetes-pod-health-check/pkg/apis/monitor/v1"
	"github.com/riete/requests"
	corev1 "k8s.io/api/core/v1"
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

var log = logf.Log.WithName("controller_podhealthcheck")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new PodHealthCheck Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcilePodHealthCheck{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("podhealthcheck-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource PodHealthCheck
	err = c.Watch(&source.Kind{Type: &monitorv1.PodHealthCheck{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner PodHealthCheck
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &monitorv1.PodHealthCheck{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcilePodHealthCheck implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcilePodHealthCheck{}

// ReconcilePodHealthCheck reconciles a PodHealthCheck object
type ReconcilePodHealthCheck struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a PodHealthCheck object and makes changes based on the state read
// and what is in the PodHealthCheck.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcilePodHealthCheck) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling PodHealthCheck")

	// Fetch the PodHealthCheck instance
	instance := &monitorv1.PodHealthCheck{}
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

	podList := &corev1.PodList{}
	labelSelector := labels.SelectorFromSet(instance.Spec.LabelSelector)
	listOps := &client.ListOptions{Namespace: instance.Spec.Namespace, LabelSelector: labelSelector}
	err = r.client.List(context.TODO(), listOps, podList)
	if err != nil {
		reqLogger.Error(err, fmt.Sprintf("[%s] Failed to list pods in namespace [%s] with label selector [%v]", instance.Name, instance.Spec.Namespace, instance.Spec.LabelSelector))
		return reconcile.Result{Requeue: true, RequeueAfter: instance.Spec.HealthCheck.Interval * time.Second}, nil
	}
	for _, pod := range podList.Items {
		go doHealthCheck(
			pod.Status.PodIP,
			instance.Spec.HealthCheck.Port,
			instance.Spec.HealthCheck.Path,
			pod.Namespace,
			pod.Name,
			instance.Spec.HealthCheck.Timeout,
			instance.Spec.DingTalk.Webhook,
			instance.Spec.DingTalk.Secret,
		)
	}
	reqLogger.Info(fmt.Sprintf("[%s] recheck after %d seconds", instance.Name, instance.Spec.HealthCheck.Interval))
	return reconcile.Result{Requeue: true, RequeueAfter: instance.Spec.HealthCheck.Interval * time.Second}, nil
}

func doHealthCheck(ip, port, path string, namespace, podName string, timeout time.Duration, webhook, secret string) {
	s := requests.NewSession()
	s.SetTimeout(timeout * time.Second)
	address := fmt.Sprintf("http://%s:%s%s", ip, port, path)
	err := s.Get(address)
	if err != nil {
		log.Info(fmt.Sprintf("GET %s with %s", address, err.Error()))
		message := fmt.Sprintf("GET %s Failed with %s, %s/%s", address, err.Error(), namespace, podName)
		sendAlert(message, webhook, secret)
	} else if s.Request.StatusCode < 200 || s.Request.StatusCode > 300 {
		log.Info(fmt.Sprintf("GET %s return %s", address, s.Request.Status))
		message := fmt.Sprintf("GET %s return [%s], %s/%s", address, s.Request.Status, namespace, podName)
		sendAlert(message, webhook, secret)
	} else {
		log.Info(fmt.Sprintf("GET %s return %s", address, s.Request.Status))
	}
}

func sendAlert(message, webhook, secret string) {
	log.Info(notify.SendDingTalkText("KubernetesPodHealthCheckFailed", message, webhook, secret, false))
}
