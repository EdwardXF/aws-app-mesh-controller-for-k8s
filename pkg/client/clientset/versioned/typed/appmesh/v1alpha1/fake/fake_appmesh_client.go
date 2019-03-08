// Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/aws/aws-app-mesh-controller-for-k8s/pkg/client/clientset/versioned/typed/appmesh/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeAppmeshV1alpha1 struct {
	*testing.Fake
}

func (c *FakeAppmeshV1alpha1) Meshes(namespace string) v1alpha1.MeshInterface {
	return &FakeMeshes{c, namespace}
}

func (c *FakeAppmeshV1alpha1) VirtualNodes(namespace string) v1alpha1.VirtualNodeInterface {
	return &FakeVirtualNodes{c, namespace}
}

func (c *FakeAppmeshV1alpha1) VirtualServices(namespace string) v1alpha1.VirtualServiceInterface {
	return &FakeVirtualServices{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeAppmeshV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
