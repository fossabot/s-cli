package secret

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"skynx.io/s-lib/pkg/errors"
	"skynx.io/s-lib/pkg/k8s/config"
	"skynx.io/s-lib/pkg/xlog"
)

func (a *API) Get(ns, name string) (*corev1.Secret, error) {
	if a.clientset == nil {
		clientset, err := config.NewClient(a.KubeConfig)
		if err != nil {
			return nil, errors.Wrapf(err, "[%v] function config.NewClient()", errors.Trace())
		}
		a.clientset = clientset
	}

	ctx := context.TODO()

	getOpts := metav1.GetOptions{}

	s, err := a.clientset.CoreV1().Secrets(ns).Get(ctx, name, getOpts)
	if k8sErrors.IsNotFound(err) {
		xlog.Debugf("Secret %s/%s not found", ns, name)
		return nil, nil
	} else if statusError, isStatus := err.(*k8sErrors.StatusError); isStatus {
		xlog.Errorf("Unable to get secret %v", statusError.ErrStatus.Message)
		return nil, errors.Wrapf(err, "[%v] function a.clientset.CoreV1().Secrets().Get()", errors.Trace())
	} else if err != nil {
		return nil, errors.Wrapf(err, "[%v] function a.clientset.CoreV1().Secrets().Get()", errors.Trace())
	}

	return s, nil
}
