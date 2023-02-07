/*
Copyright ApeCloud, Inc.

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

package cluster

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	dynamicfakeclient "k8s.io/client-go/dynamic/fake"
	"k8s.io/client-go/kubernetes/scheme"
	cmdtesting "k8s.io/kubectl/pkg/cmd/testing"

	dbaasv1alpha1 "github.com/apecloud/kubeblocks/apis/dbaas/v1alpha1"
	"github.com/apecloud/kubeblocks/internal/cli/create"
	"github.com/apecloud/kubeblocks/internal/cli/types"
	testutil "github.com/apecloud/kubeblocks/internal/testutil/k8s"
	"github.com/apecloud/kubeblocks/test/testdata"
)

func NewFakeOperationsOptions(ns, cName string, opsType dbaasv1alpha1.OpsType, objs ...runtime.Object) (*cmdtesting.TestFactory, *OperationsOptions) {
	streams, _, _, _ := genericclioptions.NewTestIOStreams()
	tf := cmdtesting.NewTestFactory().WithNamespace(ns)
	o := &OperationsOptions{
		BaseOptions: create.BaseOptions{
			IOStreams: streams,
			Name:      cName,
			Namespace: ns,
		},
		TTLSecondsAfterSucceed: 30,
		OpsType:                opsType,
	}

	err := dbaasv1alpha1.AddToScheme(scheme.Scheme)
	if err != nil {
		panic(err)
	}

	// TODO using GroupVersionResource of FakeKubeObjectHelper
	listMapping := map[schema.GroupVersionResource]string{
		types.ClusterDefGVR():       types.KindClusterDef + "List",
		types.ClusterVersionGVR():   types.KindClusterVersion + "List",
		types.ClusterGVR():          types.KindCluster + "List",
		types.ConfigConstraintGVR(): types.KindConfigConstraint + "List",
		types.BackupGVR():           types.KindBackup + "List",
		types.RestoreJobGVR():       types.KindRestoreJob + "List",
		types.OpsGVR():              types.KindOps + "List",
	}
	o.Client = dynamicfakeclient.NewSimpleDynamicClientWithCustomListKinds(scheme.Scheme, listMapping, objs...)
	return tf, o
}

func NewFakeClusterResource(namer testutil.ResourceNamer, componentName, componentType string, options ...testdata.ResourceOptions) testutil.CreateResourceObject {
	return func() runtime.Object {
		cluster := &dbaasv1alpha1.Cluster{
			TypeMeta: metav1.TypeMeta{
				APIVersion: dbaasv1alpha1.APIVersion,
				Kind:       dbaasv1alpha1.ClusterKind,
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      namer.ClusterName,
				Namespace: namer.NS,
			},
			Spec: dbaasv1alpha1.ClusterSpec{
				ClusterDefRef:     namer.CDName,
				ClusterVersionRef: namer.CVName,
				Components: []dbaasv1alpha1.ClusterComponent{{
					Name: componentName,
					Type: componentType,
				}},
			},
		}
		for _, option := range options {
			option(cluster)
		}
		return cluster
	}
}
