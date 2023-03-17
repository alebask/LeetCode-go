package medium

import "strings"

//lint:ignore U1000 Ignore unused function temporarily for debugging
func restoreIpAddresses(s string) []string {
   
    res := make([]string, 0)

	var recursive func(s string, ip []string)
	recursive = func(s string, ip []string) {

		n := len(s)
        r := 4-len(ip)

		if r == 0 {
			if n == 0 {
				res = append(res, strings.Join(ip, "."))
			}
			return
		}
        
        //too many or too less digits left for proper IP
		if n < r || n > r*3{
			return
		}
        
        //1 digit
        if n > 0 {
            recursive(s[1:], append(ip, s[0:1]))
        }
            
        //2 digit
        if n > 1 && s[0] != '0' {
		    recursive(s[2:], append(ip, s[0:2]))
        }
            
        //3 digit
        if n > 2 && s[0] != '0' && s[0:3] < "256" {
		    recursive(s[3:], append(ip, s[0:3]))
        }
	}

    recursive(s, []string{})
	return res
}
