# Data Structures and Algorithms


data structure  =>  organization of data  in computer memory
- helps to minimize the storage space
- improve the ease of operating upon those data or easy retrival, processing etc

> Types of data structures

- linear   ----[memory structure is linear]
- nonlinear   ----[memory structure is non linear]
- homogeneous  ----[all data elements are of same data type]
- hetrogeneous  ----[data elements can have different data type]
- dynamic   ----[size of sturcture modifiable during operation]
- static   ----[size of structure fixed, content can be modified]

Solving a problem involves
- chosing best suited data structure and its design pattern
- chosing best suited algorithm
- doing performance analysis in terms of space time complexity


> Hertogeneous data structures  
---
---
- Linked lists
- Ordered lists
- unordered lists
> Homogeneous data structures
---
---
- 2D arrays
- multi arrays  
> Linear data structures
---

computer memory arranged linearly  
traverse all elements in single run  
elements arranged sequentially  

---
- Array
- Lists  [sequence of elements connected to another via links]
- Sets
- Tuples  [immutable sequential collection]
- Queues
- Stacks
- Heaps  [heap property: value stored on each node >= to its children(max heap)]
> Non-Linear data structures
---

not arranged sequentially but in a hirarchical manner  
cannot traversal all elements in a single run  
memory utilization efficiency higher than linear type  

---
- Trees
- Tables
- Containers  
> Dynamic data structures
---
---
- dictionaries
- sequences
- treesets


> Structural Design Pattern  

Structures formed using classes and objects  

> Complexity and Performance analysis

- CPU time
- memory
- disk
- network  
<pre>
processing time  
resource cosumption  
complexity: how the algorithms scale when no.of input parameters increase  

Big O notation  
T(n)=O(n)             #linear time complexity  
T(n)=O(log n)         #logartihmic time complexity  
T(n)=O(n^2)           #quadratic time complexity  
T(n)=O(n^3)           #cubic time complexity  

</pre>

- Brute force algorithms
- Divide and conquer algorithms
- Backtracking algorithms



> Sorting Algorithms
- output is monotonic (each element < or > than previous element)
- output is a permutation (reorder but retain all elements)
- input data should be stored in a data structure which allows random access rather than one that allows only sequential access

Classification parameters  
- computational complexity
- memory
- recursion
- stability (stable sorting algos maintain relative order of records with equal keys)

Comparison  of algos  

n => no of records to be sorted  
comparison sorts cant perform > O(nlog(n))

---
comparison sorts
<pre>
Name        best    average     worst   memory  stable  method          description
quicksort   nlog(n) nlog(n)     n^2     log(n)  no      partitioning    
megesort    nlog(n) nlog(n)     nlog(n) n       yes     merging    
heapsort    nlog(n) nlog(n)     nlog(n) 1       no      selection    
bubblesort  n       n^2         n^2     1       yes     exchanging
cyclesort   n^2     n^2         n^2     1       no      selection
</pre>
---
## Links

- [https://pkg.go.dev/std](https://pkg.go.dev/std)  Golang Standard library
- [https://pkg.go.dev/reflect](https://pkg.go.dev/reflect)
- [https://devdocs.io/go/](https://devdocs.io/go/)
- [https://go.dev/doc/](https://go.dev/doc/)
- [https://medium.com/safetycultureengineering/an-overview-of-memory-management-in-go-9a72ec7c76a8](https://medium.com/safetycultureengineering/an-overview-of-memory-management-in-go-9a72ec7c76a8)
- [https://en.wikipedia.org/wiki/Memory_management](https://en.wikipedia.org/wiki/Memory_management)