package service

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"skynx.io/s-lib/pkg/errors"
	"skynx.io/s-lib/pkg/k8s/config"
)

func (a *API) Create(s *corev1.Service) (*corev1.Service, error) {
	if err := a.Delete(s.ObjectMeta.Namespace, s.ObjectMeta.Name); err != nil {
		return nil, errors.Wrapf(err, "[%v] function a.Get()", errors.Trace())
	}

	if a.clientset == nil {
		clientset, err := config.NewClient(a.KubeConfig)
		if err != nil {
			return nil, errors.Wrapf(err, "[%v] function config.NewClient()", errors.Trace())
		}
		a.clientset = clientset
	}

	ctx := context.TODO()

	createOpts := metav1.CreateOptions{}

	s, err := a.clientset.CoreV1().Services(s.ObjectMeta.Namespace).Create(ctx, s, createOpts)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function a.clientset.CoreV1().Services().Create()", errors.Trace())
	}

	return s, nil
}
