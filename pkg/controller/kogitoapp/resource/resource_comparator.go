// Copyright 2019 Red Hat, Inc. and/or its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package resource

import (
	"github.com/RHsyseng/operator-utils/pkg/resource/compare"
	"github.com/kiegroup/kogito-cloud-operator/pkg/framework"
	appsv1 "github.com/openshift/api/apps/v1"
	buildv1 "github.com/openshift/api/build/v1"
	routev1 "github.com/openshift/api/route/v1"
	v1 "k8s.io/api/core/v1"
	"reflect"
)

// GetComparator returns the comparator for the kubernetes resources of the KogitoApp
func GetComparator() compare.MapComparator {
	resourceComparator := compare.DefaultComparator()

	resourceComparator.SetComparator(
		framework.NewComparatorBuilder().
			WithType(reflect.TypeOf(appsv1.DeploymentConfig{})).
			WithCustomComparator(framework.CreateDeploymentConfigComparator()).
			UseDefaultComparator().
			Build())

	resourceComparator.SetComparator(
		framework.NewComparatorBuilder().
			WithType(reflect.TypeOf(buildv1.BuildConfig{})).
			WithCustomComparator(framework.CreateBuildConfigComparator()).
			UseDefaultComparator().
			Build())

	resourceComparator.SetComparator(
		framework.NewComparatorBuilder().
			WithType(reflect.TypeOf(v1.Service{})).
			WithCustomComparator(framework.CreateServiceComparator()).
			UseDefaultComparator().
			Build())

	resourceComparator.SetComparator(
		framework.NewComparatorBuilder().
			WithType(reflect.TypeOf(routev1.Route{})).
			WithCustomComparator(framework.CreateRouteComparator()).
			UseDefaultComparator().
			Build())

	return compare.MapComparator{Comparator: resourceComparator}
}
