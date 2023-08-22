// reverseWords: https://leetcode.com/problems/reverse-words-in-a-string
// #medium
func reverseWords(s string) string {
	func reverseWords(s string) string {
		i, n, res := 0, len(s), ""
	
		for i < n {
			for i < n && s[i] == ' ' {
				i++
			}
			if i == n {
				break
			}
	
			j := i
			for j < n && s[j] != ' ' {
				j++
			}
			if len(res) == 0 {
				res = s[i:j]
			} else {
				res = s[i:j]+ " " + res 
			}
			i = j+1
		}
		return res
	}
}


func reverseWordsSolution1(s string) string {
	var (
        w []string
        n = len(s)
        p1, p2 = n-1, n-1
    )

    if n == 1 && s[0] != ' ' {
        return string(s[0])
    }

    for p2 > 0 && p1 > 0 {
        for p1 > 0 && s[p1] == ' ' {
            p1--
        }
        if p1 == 0 && s[p1] == ' ' {
            break
        }

        p2 = p1
        for p2 > 0 && s[p2] != ' ' {
            p2--
        }
        
        if p2 == 0 && s[p2] != ' ' {
            w = append(w, s[p2:p1+1])
        } else if p1 == p2 && s[p1] != ' ' {
            w = append(w, s[p1:])
        } else {
            w = append(w, s[p2+1:p1+1])
        }
        p1 = p2
    }

    return strings.Join(w, " ")
}
