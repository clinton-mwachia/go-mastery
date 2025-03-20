## Hash Tables

A **hash table** is a data structure that maps keys to values, enabling efficient data retrieval. In Go, the built-in `map` type provides hash table functionality, but implementing a custom hash table can deepen your understanding of its mechanics.

**1. Define the Hash Table Structure**

We'll start by defining the structure of our hash table, which includes an array of buckets. Each bucket will hold key-value pairs to handle collisions using chaining.

```go
package main

import (
	"fmt"
)

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
```

**Explanation:**

- **KeyValuePair Struct:** Holds a key and its associated value.

- **HashTable Struct:** Contains a slice of buckets, where each bucket is a slice of `KeyValuePair`. The `size` field indicates the number of buckets.

**2. Initialize the Hash Table**

Next, we'll create a function to initialize the hash table with a specified number of buckets.

```go
// NewHashTable creates a new hash table with the given number of buckets
func NewHashTable(size int) *HashTable {
	buckets := make([][]KeyValuePair, size)
	return &HashTable{
		buckets: buckets,
		size:    size,
	}
}
```

**Explanation:**

- **NewHashTable Function:** Allocates memory for the specified number of buckets and returns a pointer to the initialized `HashTable`.

**3. Implement the Hash Function**

A hash function converts a key into an index corresponding to a bucket.

```go
// hashFunction computes an index for a given key
func (h *HashTable) hashFunction(key string) int {
	hash := 0
	for _, char := range key {
		hash += int(char)
	}
	return hash % h.size
}
```

**Explanation:**

- **hashFunction Method:** Calculates a simple hash by summing the ASCII values of the characters in the key and taking the modulus with the number of buckets to ensure the index is within bounds.

**4. Implement the Insert Method**

The `Insert` method adds a key-value pair to the hash table.

```go
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
```

**Explanation:**

- **Insert Method:** Computes the bucket index using the hash function, checks if the key exists to update its value, and appends a new key-value pair if the key is not found.

**5. Implement the Get Method**

The `Get` method retrieves the value associated with a given key.

```go
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
```

**Explanation:**

- **Get Method:** Calculates the bucket index, searches for the key within the bucket, and returns the associated value if found.

**6. Implement the Delete Method**

The `Delete` method removes a key-value pair from the hash table.

```go
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
```

**Explanation:**

- **Delete Method:** Finds the key in the appropriate bucket and removes it by slicing out the key-value pair.

**7. Example Usage**

Here's how you can use the hash table:

```go
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
```

**Explanation:**

- **main Function:** Demonstrates inserting key-value pairs, retrieving a value, and deleting a key from the hash table.

**Considerations:**

- **Hash Function:** The simple hash function used here may lead to collisions. In practice, more sophisticated hash functions are recommended.

- **Collision Handling:** This implementation uses chaining to handle collisions. Other methods include open addressing and linear probing.

- **Performance:** The efficiency of the hash table depends on factors like the quality of the hash function and the load factor.
