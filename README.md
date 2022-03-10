# Data Structures and Algorithms


## High level Checklist

- Recursion  
- Arrays  
- Strings  
- Matrices  
- Sparse Matrix  
- Polynomial Representation  
- Linked list  
- Stack  
- Queues  
- Trees  
- Binary Tree
- Binary Search Trees  
- AVL Trees  
- Search Trees  
- Heap  
- Hashing  
- Graphs  


- Searching and Sorting  
- Greedy algorithms  
- Dynamic programming  
- Pattern searching  
- String algorithms  
- Backtracking  
- Divide and Conquer  
- Geometric algorithms  
- Mathematical algorithms  
- Mathematical algorithms  
- Bit algorithms  
- Grph algorithms  
- Randomized algorithms  



## Types of data structures

data structure  =>  organization of data  in computer memory
- helps to minimize the storage space
- improve the ease of operating upon those data or easy retrival, processing etc


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

- Binary Tree 
- Binary Search Tree 
- Binary Heap 
- Hashing 

- Trees
- Tables
- Containers  


> Static data structures
---

In Static data structure the size of the structure is fixed.
The content of the data structure can be modified but without changing the memory space allocated to it.

---

- array

> Dynamic data structures
---

an organization or collection of data in memory that has the flexibility to grow or shrink in size,
enabling a programmer to control exactly how much memory is utilized

---
- dictionaries
- sequences
- treesets


> Structural Design Pattern  

Structures formed using classes and objects  



### Data Structures in Depth  
```
####################    Array

Array is a collection of items stored at contiguous memory locations.
Store multiple items of the same type together.
Calculate the position of each element by simply adding an offset to a base value, the memory location of the first element of the array 


if size of array = n.

Accessing Time: O(1) [This is possible because elements
                      are stored at contiguous locations]   
Search Time:   O(n) for Sequential Search: 
               O(log n) for Binary Search [If Array is sorted]
Insertion Time: O(n) [The worst case occurs when insertion 
                     happens at the Beginning of an array and 
                     requires shifting all of the elements]
Deletion Time: O(n) [The worst case occurs when deletion 
                     happens at the Beginning of an array and 
                     requires shifting all of the elements]


####################    Linked List

Linear data structure, in which the elements are not stored at contiguous memory locations. The elements in a linked list are linked using pointers.


Types of Linked List : 

======= Singly Linked List: 
every node stores address or reference of the next node in the list and the last node has the next address or reference as NULL

======= Doubly Linked List: 
two references are associated with each node, One of the reference points to the next node and one to the previous node. 
we can traverse in both directions and for deletion, we don’t need to have explicit access to the previous node.

======= Circular Linked List: 
all nodes are connected to form a circle.
there is no NULL at the end. 
it can be a singly circular linked list or a doubly circular linked list. 
any node can be made as starting node. 



Accessing time of an element : O(n)
Search time of an element : O(n)
Insertion of an Element : O(1) [If we are at the position 
                                where we have to insert 
                                an element] 
Deletion of an Element : O(1) [If we know address of node
                               previous the node to be 
                               deleted] 


####################    Stack

Linear data structure which follows a particular order in which the operations are performed. 
The order may be LIFO(Last In First Out) or FILO(First In Last Out)


serves as a collection of elements two principal operations: 
push, which adds an element to the collection
pop, which removes the last element that was added.

In stack both the operations of push and pop take place at the same end that is top of the stack. 
It can be implemented by using both array and linked list.  

Insertion : O(1)
Deletion :  O(1)
Access Time : O(n) [Worst Case]
Insertion and Deletion are allowed on one end. 



####################    Queue

Linear structure which follows a particular order in which the operations are performed. The order is First In First Out (FIFO). 

serves as a collection of elements, with two principal operations: 
enqueue, the process of adding an element to the collection. 
(The element is added from the rear side) 
dequeue the process of removing the first element that was added. 
(The element is removed from the front side). 
It can be implemented by using both array and linked list. 

Insertion : O(1)
Deletion  : O(1)
Access Time : O(n) [Worst Case]

======== Circular Queue 
reduces wastage of space in case of array implementation, as the insertion of the (n+1),’th element is done at the 0’th index if it is empty. 


####################    Binary Tree

Tree whose elements have at most 2 children is called a binary tree. Since each element in a binary tree can have only 2 children, we typically name them the left and right child

A Binary Tree node contains 
 Data
 Pointer to left child
 Pointer to right child


A Binary Tree can be traversed in two ways: 

Depth First Traversal: 
Inorder (Left-Root-Right), Preorder (Root-Left-Right), and Postorder (Left-Right-Root) 

Breadth-First Traversal: 
Level Order Traversal 




The maximum number of nodes at level 'l' = 2^l

Maximum number of nodes = 2^(h + 1) - 1
Here h is height of a tree. Height is considered 
as the maximum number of edges on a path from root to leaf.

Minimum possible height =  ceil(Log2(n+1)) - 1  

In Binary tree, number of leaf nodes is always one 
more than nodes with two children.

Time Complexity of Tree Traversal: O(n)


####################    Binary Search Tree

Node-based binary tree data structure which has the following properties:

    left subtree of a node contains only nodes with keys lesser than the node’s key.
    right subtree of a node contains only nodes with keys greater than the node’s key.
    left and right subtree each must also be a binary search tree.



Search :  O(h)
Insertion : O(h)
Deletion : O(h)
Extra Space : O(n) for pointers

h: Height of BST
n: Number of nodes in BST

If Binary Search Tree is Height Balanced, 
then h = O(Log n) 

Self-Balancing BSTs such as AVL Tree, Red-Black
Tree and Splay Tree make sure that height of BST 
remains O(Log n)


BST provides moderate access/search (quicker than Linked List and slower than arrays). 
BST provides moderate insertion/deletion (quicker than Arrays and slower than Linked Lists). 


####################    Heap

Special Tree-based data structure in which the tree is a complete binary tree. 
That means all levels are completely filled except possibly the last level and the last level has all keys as left as possible
This property of makes them suitable to be stored in an array. 

Generally, Heaps can be of two types:

    Max-Heap: In a Max-Heap the key present at the root node must be greatest among the keys present at all of it’s children. The same property must be recursively true for all sub-trees in that Binary Tree.
    Min-Heap: In a Min-Heap the key present at the root node must be minimum among the keys present at all of it’s children. The same property must be recursively true for all sub-trees in that Binary Tree.



Get Minimum in Min Heap: O(1) [Or Get Max in Max Heap]
Extract Minimum Min Heap: O(Log n) [Or Extract Max in Max Heap]
Decrease Key in Min Heap: O(Log n)  [Or Decrease Key in Max Heap]
Insert: O(Log n) 
Delete: O(Log n)


####################    Hashing

Process of mapping keys, values into the hash table by using a hash function. 
Provides faster access to elements. 
Efficiency of mapping depends on the efficiency of the hash function used.


Hash Function: 
A function that converts a given big input key to a small practical integer value.
The mapped integer value is used as an index in the hash table. 
A good hash function should have the following properties 

- Efficiently computable. 
- Should uniformly distribute the keys (Each table position equally likely for each key) 


Hash Table: 
An array that stores pointers to records corresponding to a given phone number. 
An entry in a hash table is NIL if no existing phone number has a hash function value equal to the index for the entry. 

Collision Handling: 
Since a hash function gets us a small number for a key which is a big integer or string, 
there is the possibility that two keys result in the same value. 
The situation where a newly inserted key maps to an already occupied slot in the hash table is called collision 
and must be handled using some collision handling technique. 
Following are the ways to handle collisions: 

- Chaining: 
The idea is to make each cell of the hash table point to a linked list of records that have the same hash function value. 
Chaining is simple but requires additional memory outside the table. 
- Open Addressing: 
In open addressing, all elements are stored in the hash table itself. 
Each table entry contains either a record or NIL. 
When searching for an element, we one by one examine table slots until the desired element is found 
or it is clear that the element is not in the table. 
 

Space : O(n)
Search    : O(1) [Average]    O(n) [Worst case]
Insertion : O(1) [Average]    O(n) [Worst Case]
Deletion  : O(1) [Average]    O(n) [Worst Case]

Hashing seems better than BST for all the operations. 
But in hashing, elements are unordered and in BST elements are stored in an ordered manner. 
Also, BST is easy to implement but hash functions can sometimes be very complex to generate. 
In BST, we can also efficiently find floor and ceil of values. 



####################    Graph

Non-linear data structure consisting of nodes and edges. 
Nodes are also referred to as vertices 
Edges are lines or arcs that connect any two nodes in the graph, a finite set of ordered pair of the form (u, v)
The pair is ordered because (u, v) is not same as (v, u) in case of directed graph(di-graph)
The pair of form (u, v) indicates that there is an edge from vertex u to vertex v
The edges may contain weight/value/cost


V -> Number of Vertices.
E -> Number of Edges.

Graph can be classified on the basis of many things, below are the two most common classifications :

Direction :
    Undirected Graph : The graph in which all the edges are bidirectional.
    Directed Graph   : The graph in which all the edges are unidirectional.
Weight :
    Weighted Graph   : The Graph in which weight is associated with the edges.
    Unweighted Graph : The Graph in which there is no weight associated to the edges.

Graph can be represented in many ways, below are the two most common representations :

Adjacency Matrix Representation
Adjacency List Representation

V -> Number of Vertices.
E -> Number of Edges.


Time Complexities in case of Adjacency Matrix :
Traversal :(By BFS or DFS) O(V^2)
Space : O(V^2)

Time Complexities in case of Adjacency List :
Traversal :(By BFS or DFS) O(V + E)
Space : O(V+E)



####################    Trie:

a type of k-ary search tree, a tree data structure used for locating specific keys from within a set.
these keys are most often strings, with links between nodes defined not by the entire key, but by individual characters.

efficient for searching words in dictionaries, 
search complexity with Trie is linear in terms of word (or key) length to be searched. 
if we store keys in binary search tree, 
a well balanced BST will need time proportional to M * log N, 
where M is maximum string length and N is number of keys in tree. 
Using trie, we can search the key in O(M) time. So it is much faster than BST.

Hashing also provides word search in O(n) time on average. 
But the advantages of Trie are there are no collisions (like hashing) so worst case time complexity is O(n). 
With Trie, we can find all words beginning with a prefix (This is not possible with Hashing). 
The only problem with Tries is they require a lot of extra space. 




Insert time : O(M) where M is the length of the string.
Search time : O(M) where M is the length of the string.
Space : O(ALPHABET_SIZE * M * N) where N is number of 
        keys in trie, ALPHABET_SIZE is 26 if we are 
        only considering upper case Latin characters.
Deletion time: O(M)


####################    Segmnet tree:

a tree data structure used for storing information about intervals, or segments


used when there are a lot of queries on a set of values. 
The queries involve minimum, maximum, sum, update etc on a input range of given set. 
These are implemented using array.



Construction of segment tree : O(N)
Query : O(log N)
Update : O(log N)
Space : O(N) [Exact space = 2 * N-1]


####################    Suffix tree:


a suffix tree is a compressed trie containing all the suffixes of the given text as their keys and positions in the text as their values. Suffix trees allow particularly fast implementations of many important string operations.


its mainly used to search a pattern in a text. 
we preprocess the text so that search operation can be done in time linear in terms of pattern length. 
Other pattern searching algorithms like KMP, Z, etc take time proportional to text length. 
This is really a great improvement because length of pattern is generally much smaller than text.

Imagine we have stored complete work of William Shakespeare and preprocessed it. 
You can search any string in the complete work in time just proportional to length of the pattern. 
But using Suffix Tree may not be a good idea when text changes frequently like text editor, etc.

Suffix Tree is compressed trie of all suffixes, so following are very abstract steps to build a suffix tree from given text.

- Generate all suffixes of given text.
- Consider all suffixes as individual words and build a compressed trie.



####################    Advanced Data Structure


========  Fenwick tree:
A Fenwick tree or binary indexed tree is a data structure that can efficiently update elements and calculate prefix sums in a table of numbers.

========  Suffix array:
a suffix array is a sorted array of all suffixes of a string. It is a data structure used in, among others, full text indices, data compression algorithms, and the field of bibliometrics.

========  AVL tree:
an AVL tree is a self-balancing binary search tree. It was the first such data structure to be invented. In an AVL tree, the heights of the two child subtrees of any node differ by at most one; if at any time they differ by more than one, rebalancing is done to restore this property.

========  Splay tree:
A splay tree is a binary search tree with the additional property that recently accessed elements are quick to access again. Like self-balancing binary search trees, a splay tree performs basic operations such as insertion, look-up and removal in O(log n) amortized time.

========  B-tree:
a B-tree is a self-balancing tree data structure that maintains sorted data and allows searches, sequential access, insertions, and deletions in logarithmic time. The B-tree generalizes the binary search tree, allowing for nodes with more than two children.

========  Red-blcak tree:
a red–black tree is a kind of self-balancing binary search tree. Each node stores an extra bit representing "color", used to ensure that the tree remains balanced during insertions and deletions

========  K Dimensional Tree:
a k-d tree is a space-partitioning data structure for organizing points in a k-dimensional space. k-d trees are a useful data structure for several applications, such as searches involving a multidimensional search key and creating point clouds. 


========  Treap (A Randomized Binary Search Tree)
========  Ternary Search Tree
========  Interval Tree
========  Implement LRU Cache
========  Sort numbers stored on different machines
========  Find the k most frequent words from a file
========  Given a sequence of words, print all anagrams together
========  Tournament Tree (Winner Tree) and Binary Heap
========  Decision Trees – Fake (Counterfeit) Coin Puzzle (12 Coin Puzzle)
========  Spaghetti Stack
========  Data Structure for Dictionary and Spell Checker?
========  Cartesian Tree
========  Cartesian Tree Sorting
========  Sparse Set
========  Centroid Decomposition of Tree
========  Gomory-Hu Treek



####################    Abstract Data type (ADT)


its a type (or class) for objects whose behaviour is defined by a set of value and a set of operations.
its called abstract because it gives an implementation-independent view

The definition of ADT only mentions what operations are to be performed but not how these operations will be implemented. 
It does not specify how data will be organized in memory and what algorithms will be used for implementing the operations. 
The process of providing only the essentials and hiding the details is known as abstraction.

- List ADT
- Stack ADT
- Queue ADT

```



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
characteristic equation of recurrence  
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



Equal roots  

F(n) = 2F(n-1) + F(n-2)

F(n) = C^n = 2C^(n-1) + C^(n-2)
     
   C^2 - 2.C^n + 1 = 0	double root C=1

if C is repeated r times then 
C^n, n.C^n, ....  n^(r-1).C^n     are all solutions 

hence we can take 2 roots as solutions  
1, n.C^n  =  1, n

F(n) = C'1.C^n + C'2.n.C^n  
     = C'1 + C'2.n 

F(0) = 0 = C'1
F(1) = 1 = C'1 + C'2
F(n) = n    => Time complexity  O(n)



Non homogenious linear recurrences

F(n) = C'1.F(n-1) + C'2.F(n-2) + C'3.F(n-3) +..... +  C'd.F(n-d) + g(n)

the extra function g(n) makes this recurrence non homogenious

To solve: 
replace g(n) by 0 and solve usually   

F(n) = 4F(n-1) + 3^n

F(1) = 1

3^n = 0

C^n = 4C^(n-1) 
C^n - 4C^(n-1) = 0
C - 4 = 0
C = 4 

Homogeneous solution  
F(n) = C'1.(C^n)
F(n) = C'1.(4^n)


put g(n) on one side and find particular solutions

F(n) - 4F(n-1) = 3^n

guessing something similar to g(n)
if g(n) = n^2, guess a polynomial of degree 2 

guess1 : F(n) = C.3^n

C.3^n - 4.C.3^(n-1) = 3^n

C = -3 

solution  F(n) = -3^(n+1)

Add both solutions together  

F(n) = C'1.4^(n) + -3^(n+1)

F(1) = 1

(C'1).4 - 3^2 = 1 

C'1 = 5/2


Solution:    F(n) = (5/2).4^n - 3^(n+1)


To guess a particular solution:

If g(n) is exponential, guess something similar or same
ex: g(n) = 2^n + 3^n

F(n) = a.2^n + b.3^n


If g(n) is polynonmial of some degree  
g(n) = n^2 - 1 

F(n) = a.n^2 + b.n + c  


If g(n) is a combination
g(n) = 2^n + n

F(n) = a.2^n + (bn + c) 

If in some case when you guess F(n) = a.2^n
and if it fails, then try  F(n) = (a.n + b).2^n
and if that fails increase the degree   F(n) = ((a^2).n + b.n + c).2^n


ex: 
F(n) = 2F(n-1) + 2^n
F(0) = 1 

put 2^n = 0 
F(n) = 2F(n-1) 
F(n) = C^n

C^n - 2.C^(n-1) = 0 
C = 2 

Guess particular solution  
g(n) = 2^n

guess: 
F(n) = a.2^n

a.2^n = 2.a.2^(n-1) + 2^n

2^(n-1) will get cancelled  

a = a + 1 , wrong since no answer coming so go for the next guess  

F(n) = (a.n + b).2^n

(a.n + b).2^n = 2(a.(n-1) + b).2^(n-1) + 2^n

a.n + b = a.n - a + b + 1

a = 1

F(n) = (1.n + b).2^n
     = (n).2^n		# here be is discarded

F(n) = (n).2^n		

Taking some of both the equations 

F(n) = (C'1).2^n + (n).2^n
F(0) = 1 
     = C'1 + 0 

C'1 = 1


F(n) = 2^n + (n).2^n

Complexity of this equation = O( n.2^n )







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
 




CAP Theorem (cap theorem)

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

ifty => infinity  

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





## Maths  

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


prime numbers  
square root  
Factors  
HCF/GCD  
LCM  
Permutaitons  
Combinations  
Harmonic series 
probability 
Fermics theorems  
Chinese reminder theorem 
Eulers product formula










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



> Bitwise Operators with number system  
 



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



## OOPS  

> notes about golang regarding OOPS  

Go has no classes, no objects, no exceptions, and no templates  
Go has no type hierarchy  
Go has garbage collection and built-in concurrency  


Structs  

type Book struct {  
  Name string  
  Good bool  
}  



## Links

- [https://pkg.go.dev/std](https://pkg.go.dev/std)  Golang Standard library
- [https://pkg.go.dev/reflect](https://pkg.go.dev/reflect)
- [https://devdocs.io/go/](https://devdocs.io/go/)
- [https://go.dev/doc/](https://go.dev/doc/)
- [https://code.tutsplus.com/tutorials/lets-go-object-oriented-programming-in-golang--cms-26540](https://code.tutsplus.com/tutorials/lets-go-object-oriented-programming-in-golang--cms-26540)
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
