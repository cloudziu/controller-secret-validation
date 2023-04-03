package main

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"

	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

// podValidator validates Pods
type secretValidator struct{}

// validate admits a pod if a specific annotation exists.
func (v *secretValidator) validate(ctx context.Context, obj runtime.Object) error {
	log := logf.FromContext(ctx)
	
	secret, ok := obj.(*corev1.Secret)
	if !ok {
		return fmt.Errorf("expected a Secret but got a %T", obj)
	}

	log.Info("validating secret: ", "name: ", secret.Name, "namespace: ", secret.Namespace)
	if err := validate(secret.Data); err != nil {
		log.Info(err.Error())
		return err
	}

	return nil
}

func (v *secretValidator) ValidateCreate(ctx context.Context, obj runtime.Object) error {
	return v.validate(ctx, obj)
}

func (v *secretValidator) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) error {
	return v.validate(ctx, newObj)
}

func (v *secretValidator) ValidateDelete(ctx context.Context, obj runtime.Object) error {
	return v.validate(ctx, obj)
}

func validate(data map[string][]byte) error {
	if string(data["manifest"]) != "valid" {
		return fmt.Errorf("manifest field is not valid")
	}
	return nil
}