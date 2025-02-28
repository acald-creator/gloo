// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// MultiPoolConfigApplyConfiguration represents a declarative configuration of the MultiPoolConfig type for use
// with apply.
type MultiPoolConfigApplyConfiguration struct {
	Priorities []PriorityApplyConfiguration `json:"priorities,omitempty"`
}

// MultiPoolConfigApplyConfiguration constructs a declarative configuration of the MultiPoolConfig type for use with
// apply.
func MultiPoolConfig() *MultiPoolConfigApplyConfiguration {
	return &MultiPoolConfigApplyConfiguration{}
}

// WithPriorities adds the given value to the Priorities field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Priorities field.
func (b *MultiPoolConfigApplyConfiguration) WithPriorities(values ...*PriorityApplyConfiguration) *MultiPoolConfigApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPriorities")
		}
		b.Priorities = append(b.Priorities, *values[i])
	}
	return b
}
