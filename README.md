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


> Memory Allocation in Golang
- Go prefers to allocate memory in stack
- Heap contains values referenced outside of a function (go structs, static constants at start of a fuction)
- Go has a stack per go routine
- non-generational concurrent, tri-color mark and sweep garbage collector


## Time and Space Complexity

- Time complexity is calculated irrespective of the machine the coder runs on and the time taken
- The relation or graph between size of input data and the time taken respectively
- Slope of the graph varies with machine but the relationship remains same
- Time complexity => Function of relationship between size of input data and time taken
- Relation of how time grows with respect to growth in input data
- O(1) < O(log(n)) < O(n)

## Math

- Computer Science includes topics like Operating Systems, Databases, Networking, Artificial intelligence, Embedded systems, Data analytics etc.

Boolean Algebra => Logic Gates  
Set Theory  
Combinatorics  
Calculus  
Linear Algebra  
Graph Theory  
Probability  => Networking and communication 
Number Theory => Cryptography  
Relational Algebra => Databases  
Algorithm  

> Logics
- BigInteger
- Bit Manipulation
- Modulo Arithmetic, Modulo exponentiation and Modulo inverse 
- Sieve of Eratosthenes and Segmented Sieve
- LCM,GCD (Euclid Algorithm) and Extended Euclid Algorithm
- Catalan Numbers
- Lucas Theorem
- Chinese Remainder Theorem
- Series and Sequences 
- Catalan Numbers 
- Inclusion Exclusion Principle 
- Advanced counting techniques – Polya counting, burnsides lemma
- Stirling, eulerian, harmonic, Bernoulli, Fibonacci numbers
- Basic probability and Conditional probability
- Random variables, probability generating functions
- Bernoulli, Binomial, Poisson, normal distribution
- Counting
- Basic principles – Pigeon hole principle, addition, multiplication rules
- Inclusion-exclusion

> Calculus  

Applications in Physics and Engineering

- perfect shapes like square, circle, cubes can be solved using geometry and algebra concepts. But in nature, everything is not in perfect shape and size. Suppose you need to find the area of an irregular shape, with the help of calculus concepts easily you can find it.
- calculating weight of structures
- astronomy
- patient diagnosis
- determining behavioural patterns
- determining the amount of the necessary materials to construct curved shape constructions (e.g dome over a sports arena) and also to measure the weight of that structure. 
- to determine the exact length of power cable needed to connect two substations, which are miles away from each other
- to launch an exploratory probe, they must consider the different orbiting velocities of the Earth and the planet the probe is targeted for, as well as other gravitational influences like the sun and the moon
- to determine the exact rate of growth in a bacterial culture when different variables such as temperature and food source are changed
- to calculate the Centre of Mass, Centre of Gravity and Mass Moment of Inertia of a sports utility vehicle
- to calculate the velocity and trajectory of an object, predict the position of planets, and understand electromagnetism
- to evaluate survey data to help develop business plans for different companies
- to determine the rate of a chemical reaction and to determine some necessary information of Radioactive decay reaction



## Links

- [https://pkg.go.dev/std](https://pkg.go.dev/std)  Golang Standard library
- [https://pkg.go.dev/reflect](https://pkg.go.dev/reflect)
- [https://devdocs.io/go/](https://devdocs.io/go/)
- [https://go.dev/doc/](https://go.dev/doc/)
- [https://medium.com/safetycultureengineering/an-overview-of-memory-management-in-go-9a72ec7c76a8](https://medium.com/safetycultureengineering/an-overview-of-memory-management-in-go-9a72ec7c76a8)
- [https://en.wikipedia.org/wiki/Memory_management](https://en.wikipedia.org/wiki/Memory_management)
- [https://www.memorymanagement.org/glossary/](https://www.memorymanagement.org/glossary/)
- [https://www.educative.io/blog/become-golang-developer](https://www.educative.io/blog/become-golang-developer)
- [https://www.educative.io/blog/the-7-most-important-software-design-patterns](https://www.educative.io/blog/the-7-most-important-software-design-patterns)
- [https://www.educative.io/blog/google-coding-interview](https://www.educative.io/blog/google-coding-interview)
- [https://www.educative.io/blog/google-coding-interview-questions](https://www.educative.io/blog/google-coding-interview-questions)
- [https://www.educative.io/blog/behavioral-interviews-how-to-prepare-and-ace-interview-questions](https://www.educative.io/blog/behavioral-interviews-how-to-prepare-and-ace-interview-questions)
- [https://www.geeksforgeeks.org/math-in-competitive-programming/](https://www.geeksforgeeks.org/math-in-competitive-programming/)
- [https://www.toppr.com/bytes/calculus-in-everyday-life/](https://www.toppr.com/bytes/calculus-in-everyday-life/)
- [https://math.mit.edu/~djk/calculus_beginners/chapter00/section02.html](https://math.mit.edu/~djk/calculus_beginners/chapter00/section02.html)
- [https://www.researchgate.net/publication/259669422_The_Roles_of_Mathematics_in_Computer_Science](https://www.researchgate.net/publication/259669422_The_Roles_of_Mathematics_in_Computer_Science)
- [https://math.stackexchange.com/questions/1488418/what-is-the-use-of-calculus](https://math.stackexchange.com/questions/1488418/what-is-the-use-of-calculus)
- [https://allusesof.com/math/51-amazing-uses-of-calculus-in-real-life/](https://allusesof.com/math/51-amazing-uses-of-calculus-in-real-life/)
- [https://medium.com/@john_marsh7/uses-of-calculus-in-real-life-1bee5ea9f32d](https://medium.com/@john_marsh7/uses-of-calculus-in-real-life-1bee5ea9f32d)
- [https://discuss.codechef.com/t/programming-contest-detailed-syllabus-along-with-example-problems/17791](https://discuss.codechef.com/t/programming-contest-detailed-syllabus-along-with-example-problems/17791)
- [https://www.codingninjas.com/blog/2020/07/27/must-do-math-for-competitive-programming/](https://www.codingninjas.com/blog/2020/07/27/must-do-math-for-competitive-programming/)