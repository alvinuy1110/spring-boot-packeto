/*
 * Copyright 2018-2020 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package boot_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/buildpacks/libcnb"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"

	"github.com/paketo-buildpacks/spring-boot/v5/boot"
)

func testDetect(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		ctx    libcnb.DetectContext
		detect boot.Detect
	)

	it("always passes for standard build", func() {
		Expect(os.RemoveAll(filepath.Join(ctx.Application.Path, "META-INF", "native-image"))).To(Succeed())
		Expect(detect.Detect(ctx)).To(Equal(libcnb.DetectResult{
			Pass: true,
			Plans: []libcnb.BuildPlan{
				{
					Provides: []libcnb.BuildPlanProvide{
						{Name: "spring-boot"},
					},
					Requires: []libcnb.BuildPlanRequire{
						{Name: "jvm-application"},
						{Name: "spring-boot"},
					},
				},
			},
		}))
	})

	it("always passes for native build", func() {
		Expect(os.MkdirAll(filepath.Join(ctx.Application.Path, "META-INF", "native-image"), 0755)).To(Succeed())
		Expect(os.WriteFile(filepath.Join(ctx.Application.Path, "META-INF", "native-image", "argfile"), []byte("file-data"), 0644)).To(Succeed())
		Expect(detect.Detect(ctx)).To(Equal(libcnb.DetectResult{
			Pass: true,
			Plans: []libcnb.BuildPlan{
				{
					Provides: []libcnb.BuildPlanProvide{
						{Name: "spring-boot"},
						{Name: "native-image-argfile"},
					},
					Requires: []libcnb.BuildPlanRequire{
						{Name: "spring-boot"},
						{Name: "jvm-application"},
					},
				},
			},
		}))
	})

}
