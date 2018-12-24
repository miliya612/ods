package dsif

// USet represents the set filled with unordered and non-duplicated elements
// it can be utilized to implement Dictionary and Map interface
type USet interface {
	// Size returns the length of elements in the set
	Size() int
	// Add adds the given element to the set if it has not been included in it
	// if that is added successfully, return true, otherwise return false
	Add(interface{}) bool
	// Remove removes the given element from the set
	// if that is removed successfully, return that, otherwise return nil
	Remove(interface{}) interface{}
	// Find seeks and returns the given element in the set
	// If that is found, return that, otherwise return nil
	Find(interface{}) interface{}
}
