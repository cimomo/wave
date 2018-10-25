/*
Copyright 2018 Pusher Ltd.

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

package deployment

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pusher/wave/test/utils"
	appsv1 "k8s.io/api/apps/v1"
)

var _ = Describe("Wave required annotation Suite", func() {
	var deployment *appsv1.Deployment

	BeforeEach(func() {
		deployment = utils.ExampleDeployment.DeepCopy()
	})

	Context("hasRequiredAnnotation", func() {
		It("returns true when the annotation has value true", func() {
			annotations := deployment.GetAnnotations()
			if annotations == nil {
				annotations = make(map[string]string)
			}
			annotations[requiredAnnotation] = "true"
			deployment.SetAnnotations(annotations)

			Expect(hasRequiredAnnotation(deployment)).To(BeTrue())
		})

		It("returns false when the annotation has value other than true", func() {
			annotations := deployment.GetAnnotations()
			if annotations == nil {
				annotations = make(map[string]string)
			}
			annotations[requiredAnnotation] = "false"
			deployment.SetAnnotations(annotations)

			Expect(hasRequiredAnnotation(deployment)).To(BeFalse())
		})

		It("returns false when the annotation is not set", func() {
			Expect(hasRequiredAnnotation(deployment)).To(BeFalse())
		})

	})
})