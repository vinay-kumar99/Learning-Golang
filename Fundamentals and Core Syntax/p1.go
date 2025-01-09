/*
https://leetcode.com/problems/counting-words-with-a-given-prefix/

*/

package main
import (
	"fmt"
	"bufio"
	"strings"
	"os"
)
func prefixCount(words []string, pref string) int {
    
    var i int

    res := 0
    n := len(words)

    for i=0;i<n;i++{

        p:=0
        q:=0


        for p<len(words[i]) && q<len(pref){
            if words[i][p] != pref[q] {
                break;
            }
            p++
            q++
        }

        if q == len(pref){
            res++
        }
    }

    return res;
}

func main(){
	var n int;

	fmt.Scan(&n)

	words := make([]string, n)

	reader := bufio.NewReader(os.Stdin)

	for i:=0;i<n;i++{
		input, _ := reader.ReadString('\n')
		words[i] = strings.TrimSpace(input)
	}
	
	pref, _:= reader.ReadString('\n')

	fmt.Println(prefixCount(words, pref))
}