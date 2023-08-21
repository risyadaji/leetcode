# [1768. Merge Strings Alternately](https://leetcode.com/problems/merge-strings-alternately/description/?envType=study-plan-v2&envId=leetcode-75)

### Description

You are given two strings `word1` and `word2`. Merge the strings by adding letters in alternating order, starting with `word1`. If a string is longer than the other, append the additional letters onto the end of the merged string.

Return the merged string.

### Explanation

in Go, Slice has both **length** and **capacity**.

>  length -> number of available elements in the slice.
> 
> capacity -> number of element in the backing array.

i write a
```
ans := make([]byte, 0, m+n)
```
to prepare a slice with total capacity from sum of two words. To prevent allocate new memory when one of two words larger than other. And in the loop parts

```
for i := 0; i < m || i < n; i++ {
    if i < m {
        ans = append(ans, word1[i])
    }
    if i < n {
        ans = append(ans, word2[i])
    }
}
```
contains multiple condition to make it loop as long the character still exists. But, it should be validated every loop when we want to add from `word1` or `word2` not exceed each of word's length to prevent panic.
