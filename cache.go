package main

// DO NOT CHANGE THIS CACHE SIZE VALUE
 const CACHE_SIZE int = 3
 var cache []map[int]int =  make([]map[int]int, CACHE_SIZE)
 
  // cache  = make(map[int]int)
func Set(key int, value int) {
	// TODO: add your code here!
    
   current := 0
	entry := map[int]int{key: value}
	for i := 0; i < CACHE_SIZE; i++ {
		page := cache[i]
		if page[key] != 0 { 
			current = 1
		}
	}

	if current == 0 {
		for i := 1; i<CACHE_SIZE; i++ {
			cache[i-1] = cache[i]
		}
		cache[CACHE_SIZE-1] = entry
	}
    
      
}

func Get(key int) int {
	// TODO: add your code here!
    for i := 0; i < CACHE_SIZE; i++ {
    entry := cache[i]
     if entry[key] != 0 {
     return entry[key]
    }
  }
  return -1
  
  
   }
  
