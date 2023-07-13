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

package kubernetes

import (
	"github.com/dapr/kit/logger"
	"k8s.io/client-go/kubernetes"
	ctlruntime "sigs.k8s.io/controller-runtime"

	appsv1alpha1 "github.com/apecloud/kubeblocks/apis/apps/v1alpha1"
	clientsetscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

// GetClientSet returns a kubernetes clientset.
func GetClientSet(logger logger.Logger) (*kubernetes.Clientset, error) {
	restConfig, err := ctlruntime.GetConfig()
	if err != nil {
		logger.Errorf("kubeconfig not found: %v", err)
	}
	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}

// GetClientForKubeBlocks returns a kubernetes restclient for KubeBlocks types.
func GetRESTClient(logger logger.Logger) (*rest.RESTClient, error) {
	restConfig, err := ctlruntime.GetConfig()
	if err != nil {
		logger.Errorf("kubeconfig not found: %v", err)
	}
	appsv1alpha1.AddToScheme(clientsetscheme.Scheme)
	restConfig.GroupVersion = &appsv1alpha1.GroupVersion
	restConfig.APIPath = "/apis"
	restConfig.NegotiatedSerializer = clientsetscheme.Codecs.WithoutConversion()
	client, err := rest.RESTClientFor(restConfig)
	if err != nil {
		return nil, err
	}

	return client, nil
}
