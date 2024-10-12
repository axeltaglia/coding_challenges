/*
Problem:
721. Accounts Merge (Medium)

Given a list of accounts where each account contains a name and a list of email addresses, merge the accounts. Two accounts belong to the same person if there is some common email in both accounts. Return a list of merged accounts where each account is in the format [name, email1, email2, ...]. The accounts themselves can be returned in any order.

Example 1:
Input:

perl
Copy code
accounts = [
    ["John", "johnsmith@mail.com", "john00@mail.com"],
    ["John", "johnnybravo@mail.com"],
    ["John", "johnsmith@mail.com", "john_newyork@mail.com"],
    ["Mary", "mary@mail.com"]
]
Output:

perl
Copy code
[
    ["John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"],
    ["John", "johnnybravo@mail.com"],
    ["Mary", "mary@mail.com"]
]
Explanation:

The first and third John's accounts have johnsmith@mail.com in common, so they are merged into one account. The second John account has johnnybravo@mail.com, which is unique, so it remains a separate account. The Mary account remains unchanged.

Constraints:
1 <= accounts.length <= 1000
2 <= accounts[i].length <= 10
1 <= accounts[i][j].length <= 30
accounts[i][0] consists of English letters.
accounts[i][1:] are emails of the account.
All email addresses are guaranteed to be unique within a particular account.

*/

package main

import "fmt"

func main() {
	input := [][]string{
		{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"John", "johnnybravo@mail.com"},
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"Mary", "mary@mail.com"},
	}
	fmt.Println(accountsMerge(input))
}

func accountsMerge(accounts [][]string) [][]string {
	output := make([][]string, 0)
	for i := 0; i < len(accounts); i++ {
		account := accounts[i]
		base := make([]string, 0)
		name := account[0]
		emails := account[1:]
		dic := make(map[string]bool)

		for _, email := range emails {
			ok, _ := dic[email]
			if !ok {
				dic[email] = true
				base = append(base, email)
			}

		}

		noCommons := make([]string, 0)
		for j := i + 1; j < len(accounts); j++ {
			thereAreCommons := false
			subNoCommons := make([]string, 0)
			for c := 1; c < len(accounts[j]); c++ {
				if dic[accounts[j][c]] {
					thereAreCommons = true
				} else {
					subNoCommons = append(noCommons, accounts[j][c])
				}
			}
			if thereAreCommons {
				noCommons = append(noCommons, subNoCommons...)
				accounts = append(accounts[:j], accounts[j+1:]...)
			}
		}

		resul := make([]string, 0)
		resul = append([]string{name}, base...)
		resul = append(resul, noCommons...)
		output = append(output, resul)
	}

	return output
}
