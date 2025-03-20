package main

import "fmt"

// KeyValuePair represents a key-value pair stored in the hash table
type KeyValuePair struct {
	Key   string
	Value interface{}
}

// HashTable represents the hash table with an array of buckets
type HashTable struct {
	buckets [][]KeyValuePair
	size    int
}

// NewHashTable creates a new hash table with the given number of buckets
func NewHashTable(size int) *HashTable {
	buckets := make([][]KeyValuePair, size)
	return &HashTable{
		buckets: buckets,
		size:    size,
	}
}

// hashFunction computes an index for a given key
func (h *HashTable) hashFunction(key string) int {
	hash := 0
	for _, char := range key {
		hash += int(char)
	}
	return hash % h.size
}

// Insert adds a key-value pair to the hash table
func (h *HashTable) Insert(key string, value interface{}) {
	index := h.hashFunction(key)
	bucket := h.buckets[index]

	// Check if the key already exists and update its value
	for i, kv := range bucket {
		if kv.Key == key {
			h.buckets[index][i].Value = value
			return
		}
	}

	// If the key doesn't exist, append the new key-value pair
	h.buckets[index] = append(bucket, KeyValuePair{Key: key, Value: value})
}

// Get retrieves the value associated with the given key
func (h *HashTable) Get(key string) (interface{}, bool) {
	index := h.hashFunction(key)
	bucket := h.buckets[index]

	for _, kv := range bucket {
		if kv.Key == key {
			return kv.Value, true
		}
	}
	return nil, false
}

// Delete removes the key-value pair associated with the given key
func (h *HashTable) Delete(key string) bool {
	index := h.hashFunction(key)
	bucket := h.buckets[index]

	for i, kv := range bucket {
		if kv.Key == key {
			h.buckets[index] = append(bucket[:i], bucket[i+1:]...)
			return true
		}
	}
	return false
}

func main() {
	ht := NewHashTable(10)

	ht.Insert("apple", "fruit")
	ht.Insert("carrot", "vegetable")
	ht.Insert("pear", "fruit")

	value, found := ht.Get("carrot")
	if found {
		fmt.Println("Found:", value)
	} else {
		fmt.Println("Not found")
	}

	ht.Delete("carrot")
	value, found = ht.Get("carrot")
	if found {
		fmt.Println("Found:", value)
	} else {
		fmt.Println("Not found")
	}
}
