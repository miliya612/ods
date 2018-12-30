package dsif

// SSet represents the set filled with unordered and non-duplicated elements
type SSet interface {
	// Size returns the length of elements in the set
	Size() int
	// Add adds the given element to the set if it has not been included in it
	// if that is added successfully, return true, otherwise return false
	Add(interface{}) bool
	// Remove removes the given element from the set
	// if that is removed successfully, return that, otherwise return nil
	Remove(interface{}) interface{}
	// Find seeks and returns the given element in the set if it exists.
	// If not, find an element y in the set such that y equals x. Return y, or nil if no such element exists.
	Find(interface{}) *interface{}
}
