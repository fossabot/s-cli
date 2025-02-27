package serviceaccount

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"skynx.io/s-lib/pkg/errors"
	"skynx.io/s-lib/pkg/k8s/config"
)

func (a *API) Create(ns, name string) (*corev1.ServiceAccount, error) {
	sa, err := a.Get(ns, name)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function a.Get()", errors.Trace())
	}
	if sa != nil { // already found
		return sa, nil
	}

	if a.clientset == nil {
		clientset, err := config.NewClient(a.KubeConfig)
		if err != nil {
			return nil, errors.Wrapf(err, "[%v] function config.NewClient()", errors.Trace())
		}
		a.clientset = clientset
	}

	ctx := context.TODO()

	sa = a.New(ns, name)

	createOpts := metav1.CreateOptions{}

	sa, err = a.clientset.CoreV1().ServiceAccounts(ns).Create(ctx, sa, createOpts)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function a.clientset.CoreV1().ServiceAccounts().Create()", errors.Trace())
	}

	return sa, nil
}
