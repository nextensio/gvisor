// automatically generated by stateify.

// +build arm64
// +build arm64
// +build arm64

package arch

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *context64) StateTypeName() string {
	return "pkg/sentry/arch.context64"
}

func (x *context64) StateFields() []string {
	return []string{
		"State",
		"sigFPState",
	}
}

func (x *context64) beforeSave() {}

func (x *context64) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.State)
	m.Save(1, &x.sigFPState)
}

func (x *context64) afterLoad() {}

func (x *context64) StateLoad(m state.Source) {
	m.Load(0, &x.State)
	m.Load(1, &x.sigFPState)
}

func init() {
	state.Register((*context64)(nil))
}