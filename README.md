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


> Recursive Algorithms

	   	    1-f(4)
       	  2-f(3)	        7-f(3)
      3-f(3)    6-f(1)     8-f(1)    9-f(0)
 4-f(1)  5-f(0)

Flow of execution
1: f(4)
2: f(3)
3: f(2)
4: f(1)
5: f(0)
6: f(1)
7: f(2)
8: f(1)
9: f(0)

- Space complexity is not constant in recursion programs, because the function calls are stored in stack
- at any particular point of time no two function call at the same level of recursion will be in the stack at the same time
- only calls that are interlinked with each other will be in the stack at the same time
- space complexity is equal to the height of the tree
- the longest chain starting from root till leaf is the space complexity

There are 2 Types of Recursion
- Linear recursion
 F(n) = F(n-1) + F(n-2)
- Divide and conquer
 F(n) = F(n/2) + O(1)


Linear Recurences 

Form: 
F(x) = a'1.F(x-1) + a'2.F(x-2) + a'3.f(x-3) + ... + a'n.f(x-n)

Sma(i=1,k)[F(n)]     => sum (sigma) of all i till k  

F(x) = Sma(i=1,n)[ a'i.F(x-i) ] = 1  
		     => n : is the order of recurence

example : fibonacci number

F(n) = F(n-1) + F(n-2)

C => some constant

F(n) =  C^n 
     =  C^(n-1) + C^(n-2) 

C^n - C^(n-1) - C^(n-2) = 0
C^2 - C^1 - 1 = 0
 roots of this eqn
 C = [-b +- (b^2-4ac)^(1/2)] / 2a
 C = (1 +- 5^(1/2)) / 2

 C1 = (1 + 5^(1/2)) / 2
 C2 = (1 - 5^(1/2)) / 2

F(n) = C'1.(C1)^n + C'2.(C2)^n
F(n) = C'1.((1 + 5^(1/2)) / 2)^n + C'2.((1 - 5^(1/2)) / 2)^n

Based on the statement:
no.of roots = no.of answers already present
C1 and C2 are the 2 roots 
so we can substitute that in 


F(0) = 0
F(1) = 1

F(0) = C'1.((1 + 5^(1/2)) / 2)^0 + C'2.((1 - 5^(1/2)) / 2)^0
     = C'1 + C'2
C'1  = -C'2


F(1) = 1
F(1) = C'1.((1 + 5^(1/2)) / 2)^1 + C'2.((1 - 5^(1/2)) / 2)^1
1 = C'1.((1 + 5^(1/2)) / 2)^1 + C'2.((1 - 5^(1/2)) / 2)^1

since  C'1  = -C'2
1 = C'1.((1 + 5^(1/2)) / 2)^1 - C'1.((1 - 5^(1/2)) / 2)^1
C'1 = 1/5^(1/2)
C'1 = -1/5^(1/2)

F(n) = 1/5^(1/2).((1 + 5^(1/2)) / 2)^n + -1/5^(1/2).((1 - 5^(1/2)) / 2)^n
This is the formula for n'th fibonacci number

To get the time complexity from the above formula
F(n) = 1/5^(1/2).((1 + 5^(1/2)) / 2)^n - 1/5^(1/2).((1 - 5^(1/2)) / 2)^n

(1 - 5^(1/2)) / 2  can be ignored as n -> infinity
and constants can be ignored

Time complexity= O((1 + 5^(1/2)) / 2)^n  for nth fibonacci number
Time complexity= O(1.618)^n  for nth fibonacci number
1 + 5^(1/2)   => is also known as golden ratio
this results in the program not giving an output for even smaller numbers
and was in hanging state, its because the exponential time complexity is very bad



Solving to find complexity in recurrence relation of divide and conqure :

F(N) = F(N-1) + F(N-2)



Divide and conquer Recurences

Form: 
T(x) = a'1.T(b'1.x+F'1(x)) + a'2.T(b'2.x+F'2(x)) +.... + a'k(b'k.x+F'k(x)) + g(x)
		for x >= x'0        where x'0 is some constant

T(n) = 3T(N/2) + (1/2)T(N/6) + (4N^3-1)

a'1 = 3
b'1 = N/2
a'2 = 1/2
b'2 = N/6
g(n) = 	4N^3-1

T(n) = 2T(N/2) + (N-1)

a'1 = 2
b'1 = N/2
g(n) = N-1   [ time taken to operate on answer provided by the recursion call 2T(N/2) ]

[             ] array of lenth N
[ n/2 ] [ n/2 ] array is divided into 2 subarrays and when they are sorted by recursion
		then these arrays are merged in time N-1 comparisons
T(n) = T(N/2) + T(N/2) + (N-1)
     = 2T(N/2) + (N-1)           recurence relation of merge sort


Solving to find complexity in recurrence relation of divide and conqre :
- plug and check
- Masters theorem
- Akra Bazzi formula

Itg(1,x)[F(n)] 	=>  Integral of function F(n) where the lower limit is 1 and upper limit is x

Itg[u^p]du = u^(p+1) / (p+1)

Tta  => theta

T(x) =  Tta ( x^p + (x^p).Itg(1,x)[ G(u)/u^(p+1) ]du )

Sma(i=1,k)[F(n)]     => sum (sigma) of all i till k  

a'1.(b'1)^p + a'2.(b'2)^p + ..... = 1

we have to find p such that the below equation is satisfied

Sma(i=1,k)[ a'i.(b'i)^p ] = 1  


example: 

Binary search:
T(N) = T(N/2) + C	; C -> constant

Merge sort
T(N) = 2T(N/2) + (N-1)

a'1 = 2
b'1 = 1/2
g(n) = N-1


Sma(i=1,k)[ a'i.(b'i)^p ] = 1  
a'1.(b'1)^p = 1
2 . (1/2)^p = 1
p = 1
substituting p in Akra Bazzi Formula

T(x) =  Tta ( x^p + (x^p).Itg(1,x)[ G(u)/u^(p+1) ]du )
     =  Tta ( x^1 + (x^1).Itg(1,x)[ (u-1)/u^(1+1) ]du )
     =  Tta ( x + (x).Itg(1,x)[ (u-1)/u^2 ]du )
     =  Tta ( x + (x).Itg(1,x)[ (1/u - 1/u^2) ]du )
     =  Tta ( x + (x).Itg(1,x)[ (1/u - 1/u^2) ]du )
     =  Tta ( x + (x).(1,x)[ log(u) + 1/u ] )
     =  Tta ( x + (x).[ log(x) + 1/x - ( log(1) + 1/1 )]
     =  Tta ( x + x.[ log(x) + 1/x -  0 - 1 )]
     =  Tta ( x + x.log(x) + 1 - x )
     =  Tta ( x + x.log(x) + 1 - x )
     =  Tta ( x.log(x) + 1 )
     =  Tta ( x.log(x) )	# constants are ignored
     =  Tta ( x.log(x) )	# this is the time complexity

for an array of size N
merge sort complexity = O(N.log(N))


example:
T(N) = 4T(N/3) + (3/7).T(1N/4) + N^3

a'1 = 4
b'1 = 1/3
a'1 = 3/7
b'1 = 1/4
g(n)= N^3

Sma(i=1,k)[ a'i.(b'i)^p ] = 1  
4.(1/3)^p + (3/7).(1/4)^p = 1
when p = 0 ,  4 + 3/7 > 1
	      4.4 > 1
when p = 1 ,  4/3 + 3/28 > 1
	      1.3 + .1 > 1
	      1.4 > 1
when p = 2 ,  4/9 + 3/112
	      .44 + .02 
	      .42 < 1

Hence  1 < p < 2

when p < power of g(n), 
then answer is g(n)

here   g(n) = x^2
 p < 2   # that is the power of g(n)
hence, answer is  O(g(n))
		= O(x^2)

Proof:

substituting p in Akra Bazzi Formula

T(x) =  Tta ( x^p + (x^p).Itg(1,x)[ G(u)/u^(p+1) ]du )
     =  Tta ( x^p + (x^p).Itg(1,x)[ u^2/u^(p+1) ]du )
     =  Tta ( x^p + (x^p).Itg(1,x)[ u^(1-p) ]du )
     =  Tta ( x^p + (x^p).(1,x)[ u^(2-p)/(2-p) ]du )
     =  Tta ( x^p + (x^p).[ x^(2-p)/(2-p) - 1^(2-p)/(2-p) ] )
     =  Tta ( x^p + [ x^(2)/(2-p) - x^p ] )  
     =  Tta ( x^(2)/(2-p) )  	  # note: here since 1<p<2, x^p is a less dominating term copared to x^2 and can be ignored
     =  Tta ( x^(2) )     
 





## Time and Space Complexity - Design and Analysis of Algorithms

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


> Memory Allocation in Golang

- Go prefers to allocate memory in stack
- Heap contains values referenced outside of a function (go structs, static constants at start of a fuction)
- Go has a stack per go routine
- non-generational concurrent, tri-color mark and sweep garbage collector


> Calculations

- Time complexity is calculated irrespective of the machine the coder runs on and the time taken
- time complexity != total time taken
- The relation or graph between size of input data and the time taken respectively
- Slope of the graph varies with machine but the relationship remains same
- Time complexity => Function of relationship between size of input data and time taken
- Relation of how time grows with respect to growth in input data
- In time complexity we always care about worst case when your input data grow large in size what happens
- Always look at complexity for large and infinite data
	 consider O(n^5+log(n))
	 for n= 1000000= 1million
	Time taken
	 = 1000000^5 + log(1000000)
	 = 1000000^5 + 6    [ since 6 sec is too small compartively it can be ingored]
	~= 1000000^5 
- O(1) < O(log(n)) < O(n) < O(nlog(n)) <  O(2^n)
- cases when y=x, y=2x, y=3x, y=5x+2 etc.. here the values of actual time is different but all grows linearly
- here we do not care about the constants so we ignore all constants in time complexity
	O(5n^4 + 4n^3+ n^2 + 2) 
	~= O(n^4 + n^3 + n^2)
	~= O(n^4)
- Big O notation gives the uppers bound of the time complexity relationship
	f(n) = O(g(n))
	ifty ~ infinity
	lim n->ifty [ f(n) / g(n) ] < ifty
	O(5n^4 + 4n^3+ n^2 + 2)  =~ O(n^4)
	f(n) = 5n^4 + 4n^3+ n^2 + 2 
	g(n) = n^4

	 lim n->ifty [ 5n^4 + 4n^3+ n^2 + 2 / n^4 ] 
	= lim n->ifty [ 5 + 4/n + 1/n^2 + 2/n^4 ] 
	= [ 5 + 4/ifty + 1/ifty + 2/ifty ] 
	= [ 5 + 0 + 0 + 0 ] 
	= 5 < ifty

- Big Omega Om() notation is the opposite of O(n) , it gives the lower bound
	lim n->ifty [ f(n) / g(n) ] > 0
- Theta notation is the combination of both Big O and Big Omega
	0 < lim n->ifty [ f(n) / g(n) ] < ifty
- little O notation gives the upper bound which is not a strict upper bound but a loose upper bound
	Big O	 	f=O(g)  f <= g  [ f grows slower than g ]
	Little O 	f=O(g)  f <  g  [ f grows strictly slower than g ]

	lim n->ifty [ f(n) / g(n) ] = 0
- little Omega notation gives the lower bound which is not a strict upper bound but a loose upper bound
	Big Om	 	f=Om(g)  f >= g  [ f grows faster than g ]
	Little Om 	f=Om(g)  f >  g  [ f grows strictly faster than g ]

	lim n->ifty [ f(n) / g(n) ] = ifty


- Space complexity = Auxilary space (extra or temporary space used by an algorithm) + Input space
	space complexity of all sorting algorithms is O(n)
	in sorting algos no new arrays are getting created so its taking constant space complexity
	merge sort uses O(n) auxiliary space
	insertion sort and heap sort uses O(1) auxiliary space
- calculating O(n) for loops
for i:=1; i<=n {
	for j:=1; j<=k; j++ {
		// operation taking time t
	}
	i=i+k
}

O(kt * no of iterations in outer loop)
when i=1,  1+1k
when i=2,  1+2k
when i=3,  1+3k ...
when i=x,  1+xk 
1+xk <= n
xk <= n-1
x  =  (n-1)/k   [ no of iteration in outer loop ]

O(kt * (n-1)/k) 
= O(nt) 
= O(n)





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
