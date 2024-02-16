/*
 * Given a string like "aaaabbcd" compress it to "a4b2c1d1".
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char* Compress(char input[]) {
    int len = strlen(input);
    // Allocate potentially more space to handle multi-digit counts
    char* output = malloc(4 * len + 1); 
    if (!output) return NULL; // Check malloc success

    int i, j = 0;
    for (i = 0; i < len; i++) {
        int count = 1; // Start counting from 1 for the current character
        while (i + 1 < len && input[i] == input[i + 1]) {
            count++;
            i++;
        }
        // Write character and its count to output
        j += sprintf(output + j, "%c%d", input[i], count);
    }

    return output;
}
