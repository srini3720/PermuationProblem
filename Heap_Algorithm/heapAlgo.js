function printArr(a,n)
{
    console.log("Result ","a", a, "n", n);
     
}
 
// Generating permutation using Heap Algorithm
function heapPermutation(a,size,n)
{
        // if size becomes 1 then prints the obtained
        // permutation
        if (size == 1)
            printArr(a, n);
  
        for (let i = 0; i < size; i++) {
            console.log("array",a,"i", i, "size", size,"n", n);
            heapPermutation(a, size - 1, n);
  
            // if size is odd, swap 0th i.e (first) and
            // (size-1)th i.e (last) element
            if (size % 2 == 1) {
                let temp = a[0];
                a[0] = a[size - 1];
                a[size - 1] = temp;
            }
  
            // If size is even, swap ith
            // and (size-1)th i.e last element
            else {
                let temp = a[i];
                a[i] = a[size - 1];
                a[size - 1] = temp;
            }
        }
}
 
// Driver code
let a=[1, 2, 3];
heapPermutation(a, a.length,3);
 
 