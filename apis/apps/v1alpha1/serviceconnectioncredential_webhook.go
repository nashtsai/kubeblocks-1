/*
Copyright (C) 2022-2023 ApeCloud Co., Ltd

This file is part of KubeBlocks project

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var serviceconnectioncredentiallog = logf.Log.WithName("serviceconnectioncredential-resource")

func (r *ServiceConnectionCredential) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-apps-kubeblocks-io-v1alpha1-serviceconnectioncredential,mutating=true,failurePolicy=fail,sideEffects=None,groups=apps.kubeblocks.io,resources=serviceconnectioncredentials,verbs=create;update,versions=v1alpha1,name=mserviceconnectioncredential.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &ServiceConnectionCredential{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *ServiceConnectionCredential) Default() {
	serviceconnectioncredentiallog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-apps-kubeblocks-io-v1alpha1-serviceconnectioncredential,mutating=false,failurePolicy=fail,sideEffects=None,groups=apps.kubeblocks.io,resources=serviceconnectioncredentials,verbs=create;update,versions=v1alpha1,name=vserviceconnectioncredential.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &ServiceConnectionCredential{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *ServiceConnectionCredential) ValidateCreate() error {
	serviceconnectioncredentiallog.Info("validate create", "name", r.Name)

	return r.validate()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *ServiceConnectionCredential) ValidateUpdate(old runtime.Object) error {
	serviceconnectioncredentiallog.Info("validate update", "name", r.Name)

	return r.validate()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *ServiceConnectionCredential) ValidateDelete() error {
	serviceconnectioncredentiallog.Info("validate delete", "name", r.Name)

	return r.validate()
}

func (r *ServiceConnectionCredential) validate() error {
	return nil
}
