/*
 * You have a string S and an array roll.
 * Each element in the array roll corresponds to an operation to be applied to the string S.
 * The operation, referred to as "rolling," involves modifying the characters of the string based on their ASCII values.
  When you "roll" a character, you increase its ASCII value by 1.  
For example, rolling a (ASCII 97) would result in b (ASCII 98).
If the character is z (ASCII 122), rolling it wraps around to a (ASCII 97), maintaining the cyclic nature of the alphabet.
For each roll[i], the operation is applied to the first roll[i] characters of the string S.
The task is to apply each roll operation defined in the roll array to the string S sequentially and output the final state of the string after all operations have been applied.
 */

#include <iostream>
#include <vector>
#include <string>

using namespace std;

string rollString(string s, vector<int>& roll) {
    int n = s.length();
    // Store the roll count for each char in given string
    // So we apply them in one go.
    vector<int> rollCount(n, 0);

    for (int i = 0; i < roll.size(); ++i) {
        rollCount[0] += 1; // Every roll includes the first character
        if (roll[i] < n) rollCount[roll[i]] -= 1;
    }

    // Compute the final roll count for each character
    for (int i = 1; i < n; ++i) {
        rollCount[i] += rollCount[i - 1];
    }

    // Apply the roll operations to the string
    for (int i = 0; i < n; ++i) {
        s[i] = 'a' + (s[i] - 'a' + rollCount[i]) % 26; // 'a' ensures we start from 'a', and % 26 handles wrap-around
    }

    return s;
}
