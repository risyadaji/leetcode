package main

// ref: https://leetcode.com/problems/ransom-note/description/

func canConstruct(ransomNote string, magazine string) bool {
	magzChar := make(map[rune]int)
	for _, char := range magazine {
		magzChar[char]++
	}

	for _, char := range ransomNote {
		if magzChar[char] == 0 {
			return false
		}
		magzChar[char]--
	}

	return true
}

// nice solution
// just init 26 arr (because of total alphabet)
// and then count total letter in magazines, and substract in every letter in ransomNote
// if less then 0 (means not enought letter) then return false.

// func canConstruct(ransomNote string, magazine string) bool {
//     magazin := make([]int, 26)
//     for _, v := range magazine {
//         magazin[v - 'a']++
//     }

//     for _, v := range ransomNote {
//         magazin[v - 'a']--
//         if magazin[v - 'a' ] < 0 {
//             return false
//         }
//     }

//     return true
// }
