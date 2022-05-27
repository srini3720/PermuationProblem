//Heap Algorithm 

Problem

Notes on Random Subsection Assignment

       1.Doing alternate assignment (1,2,1,2,1,2,1….) is less than optimal because it is not random.

       2.However, we do want some kind of block randomization because time of login it known to be an important “nuisance” factor characterizing students (such as logging in early vs. near deadline).

       3.Therefore we want to block for login time.  As a proxy for time, we will use login order. Of course this is a non-uniform measure of time, though it is monotonic.

We will sequentially assign students to subsection as they login. Subsection assignment will occur via the following:  Assume we have N subsections that we want to which we want to (block) randomly assign all students. Let us make the blocks of size N.  To begin with the first N students, first create (or access from a file) the possible N! permutations of assignment sequences, then randomly select one of the permutations, and then use that sequence for N students. Once N students are assigned, then again randomly pick (with replacement) a possible permutation, and assign the next N students. Repeat as needed.

For example, for
N=2, the possible permutations of sequences are [1,2] and [2,1].
N= 3 the possible permutations are [1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1]
These can be generated at https://numbergenerator.org/permutations-and-combinations
Also with code: https://www.geeksforgeeks.org/heaps-algorithm-for-generating-permutations/

Google sheet for runtime memory complexity
https://docs.google.com/spreadsheets/d/1aITT37ixEpQsGtspOClXVbBsvqLFwuMwVyv5AA3weyk/edit#gid=0
