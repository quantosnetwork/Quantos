package set

import (
	"fmt"
	"sync"
)

// Set objects are collections of strings. A value in the Set may only occur once,
// it is unique in the Set's collection.
type Set struct {
	set   map[string]struct{}
	mutex *sync.Mutex
}

// Add appends a new element with the given value to the Set object.
// It returns an error if the value already in set.
func (s *Set) Add(v string) error {
	// Check for mutex
	s.init()

	// Lock this function
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Init the set map if set is empty.
	if s.set == nil {
		s.set = make(map[string]struct{})
	}

	// Check for value exist.
	if _, ok := s.set[v]; ok {
		return fmt.Errorf("Value already in set.")
	}

	s.set[v] = struct{}{}
	return nil
}

// Clear removes all elements from the Set object.
func (s *Set) Clear() {
	// Check for mutex
	s.init()

	// Lock this function
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.set = nil
}

// Values returns a new list object that contains the values for each element
// in the Set object.
func (s *Set) Values() (keys []string) {
	// Check for mutex
	s.init()

	// Lock this function
	s.mutex.Lock()
	defer s.mutex.Unlock()

	keys = make([]string, 0, len(s.set))
	for k := range s.set {
		keys = append(keys, k)
	}

	return
}

// Has returns a boolean asserting whether an element is present with the
// given value in the Set object or not.
func (s *Set) Has(v string) (ok bool) {
	// Check for mutex
	s.init()

	// Lock this function
	s.mutex.Lock()
	defer s.mutex.Unlock()

	_, ok = s.set[v]

	return
}

// init the set mutex
func (s *Set) init() {
	// Check for mutex
	if s.mutex == nil {
		s.mutex = &sync.Mutex{}
	}
}
