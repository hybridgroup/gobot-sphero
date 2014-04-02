package gobotSphero

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SpheroDriver", func() {
	var (
		driver *SpheroDriver
	)

	BeforeEach(func() {
		driver = NewSphero(new(SpheroAdaptor))
	})

	PIt("Must be able to Start", func() {
		Expect(driver.Start()).To(Equal(true))
	})
	PIt("Must be able to Init", func() {
		Expect(driver.Init()).To(Equal(true))
	})
	PIt("Must be able to Halt", func() {
		Expect(driver.Halt()).To(Equal(true))
	})
})
