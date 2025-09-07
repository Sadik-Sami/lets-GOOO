package main

import "fmt"

func main() {
	// ----------------------------
	// 1. Array → Slice
	// ----------------------------
	arr := [6]string{"This", "is", "a", "go", "interview", "question"}
	fmt.Println(arr) // full array
	// Slice s1 = arr[1:4] → references part of array (not copy)
	// Slice is a descriptor (struct) with 3 fields:
	//   1. Pointer → address of first element in underlying array
	//   2. Length  → number of elements visible in slice
	//   3. Capacity → how far slice can grow from starting index
	s1 := arr[1:4] // indexes 1,2,3 = ["is", "a", "go"]
	fmt.Println("s1 (slice from array):")
	fmt.Println("emp:", s1, "len:", len(s1), "cap:", cap(s1))
	// len(s1) = 3 → ["is", "a", "go"]	// cap(s1) = 5 → from arr[1] → arr[5] (5 elements possible)

	// ----------------------------
	// 2. Slice from Slice
	// ----------------------------
	s2 := s1[1:2] // from s1 = ["is", "a", "go"], take index 1 only → ["a"]
	fmt.Println("s2 (slice from slice):")
	fmt.Println("emp:", s2, "len:", len(s2), "cap:", cap(s2))
	// len(s2) = 1
	// cap(s2) = 4 (because capacity counts from arr[2] → arr[5])

	// ----------------------------
	// 3. Slice Literal
	// ----------------------------
	s := []int{1, 2, 5} // Slice literal (not array literal!)
	fmt.Println("s (slice literal):")
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))
	// Here Go creates a hidden array [1,2,5] and s points to it.

	// ----------------------------
	// 4. Slice using make(len)
	// ----------------------------
	s3 := make([]int, 3) // allocates array [0,0,0], len=3, cap=3
	fmt.Println("s3 (make with len):")
	fmt.Println("emp:", s3, "len:", len(s3), "cap:", cap(s3))
	s3[0] = 1
	s3[1] = 3
	s3[2] = 6
	fmt.Println("updated s3:", s3)

	// ----------------------------
	// 5. Slice using make(len, cap)
	// ----------------------------
	s4 := make([]int, 3, 5) // allocates array [0,0,0,?,?], len=3, cap=5
	s4[0] = 2
	s4[1] = 4
	s4[2] = 8
	fmt.Println("s4 (make with len & cap):")
	fmt.Println("emp:", s4, "len:", len(s4), "cap:", cap(s4))
	// s4 can grow up to 5 before new array allocation is needed.

	// ----------------------------
	// 6. Nil Slice (zero value of slice)
	// ----------------------------
	var s5 []int // nil slice → pointer=nil, len=0, cap=0
	fmt.Println("s5 (var, nil slice):")
	fmt.Println("emp:", s5, "len:", len(s5), "cap:", cap(s5))

	// Append creates underlying array on demand if nil
	fmt.Println("append to s5:")
	s5 = append(s5, 1, 2, 3)
	fmt.Println("emp:", s5, "len:", len(s5), "cap:", cap(s5))
}

/*
----------------------------
THEORY: SLICES
----------------------------
- Slice is not an array. It's a small struct with 3 fields:
    type slice struct {
        ptr *Element  // pointer to first element of underlying array
        len int       // length (visible part)
        cap int       // capacity (size from start to end of array)
    }

- Arrays are value types (copied when passed).
- Slices are reference-like: they point to an underlying array.

----------------------------
COMPILATION PHASE
----------------------------
- Compiler creates code segment entries:
    main = func() { ... }
- arr = [6]string{"This", "is", ...} → static array literal in data segment.
- Slice descriptors (s1, s2, etc.) are stack variables containing
  pointer/len/cap values.

----------------------------
EXECUTION PHASE
----------------------------
- Stack frame for main() is created.
- arr is allocated on stack (6 strings).
- s1 is created → pointer points to arr[1].
- s2 is created → pointer points to arr[2].
- s, s3, s4 → cause hidden arrays to be created (stack or heap, depends on escape).
- s5 starts as nil slice → (ptr=nil, len=0, cap=0).
- append(s5, 1) → Go allocates new array [1], updates slice header.

----------------------------
STACK vs HEAP
----------------------------
- arr: on stack (local).
- s1, s2: slice headers on stack, pointing to arr’s stack array.
- s: slice header on stack, hidden array may go heap if needed.
- s3, s4: slice headers on stack, hidden arrays may go heap if they escape.
- s5: initially nil, after append → new array allocated on heap.

----------------------------
KEY RULES TO REMEMBER
----------------------------
1. Arrays = fixed length, value type (copied).
2. Slices = flexible view into arrays (pointer+len+cap).
3. Slices share same underlying array → changing slice elements changes array.
4. Capacity determines how far slice can grow before reallocation.
5. `append()` may allocate new underlying array if cap exceeded.
6. Nil slice is valid (len=0, cap=0), appending works.

*/
