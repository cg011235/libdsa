#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <unistd.h>
#include <arpa/inet.h>

char** deserialize(unsigned char* buffer, size_t bufferSize, int n) {
    char** strings = (char**)malloc(n * sizeof(char*));
    if (!strings) {
        fprintf(stderr, "Memory allocation failed for string array\n");
        return NULL;
    }

    size_t offset = 0;
    int i;
    for (i = 0; i < n && offset < bufferSize; i++) {
        if (offset + sizeof(uint16_t) > bufferSize) {
            fprintf(stderr, "Buffer overflow detected during deserialization\n");
            goto cleanup;
        }

        uint16_t length;
        memcpy(&length, buffer + offset, sizeof(length));
        offset += sizeof(length);

        if (offset + length > bufferSize) {
            fprintf(stderr, "Buffer overflow detected during deserialization\n");
            goto cleanup;
        }

        strings[i] = (char*)malloc(length + 1); // +1 for null terminator
        if (!strings[i]) {
            fprintf(stderr, "Memory allocation failed for string\n");
            goto cleanup;
        }

        memcpy(strings[i], buffer + offset, length);
        strings[i][length] = '\0'; // Null-terminate the string
        offset += length;
    }

    return strings;

cleanup:
    // Free allocated memory before returning NULL in case of error
    for (int j = 0; j < i; j++) {
        free(strings[j]);
    }
    free(strings);
    return NULL;
}

unsigned char* receiveData(int serverSocket, size_t* dataSize) {
    // Buffer to store the serialized data
    unsigned char* buffer = malloc(65536); // Adjust size as needed
    if (!buffer) {
        perror("Failed to allocate buffer");
        return NULL;
    }

    ssize_t bytesReceived = recv(serverSocket, buffer, 65536, 0);
    if (bytesReceived < 0) {
        perror("Receive failed");
        free(buffer);
        return NULL;
    }

    *dataSize = bytesReceived;
    return buffer;
}

int main() {
    int clientSocket;
    struct sockaddr_in serverAddr;
    socklen_t addrSize;

    // Create the socket
    clientSocket = socket(PF_INET, SOCK_STREAM, 0);

    // Configure settings of the server address struct
    serverAddr.sin_family = AF_INET;
    serverAddr.sin_port = htons(7799);
    serverAddr.sin_addr.s_addr = inet_addr("127.0.0.1");
    memset(serverAddr.sin_zero, '\0', sizeof serverAddr.sin_zero);

    // Connect the socket to the server using the address struct
    addrSize = sizeof serverAddr;
    connect(clientSocket, (struct sockaddr *) &serverAddr, addrSize);

    // Receive data from the server
    size_t dataSize;
    unsigned char* data = receiveData(clientSocket, &dataSize);

    if (data) {
        // Assuming the number of strings 'n' is known or sent by the server
        int n = 2; // Example, replace with actual value or logic to determine 'n'

        // Deserialize the data
        char** strings = deserialize(data, dataSize, n);

        if (strings) {
            // Process each string
            for (int i = 0; i < n; ++i) {
                printf("Received string %d: %s\n", i, strings[i]);
                free(strings[i]);  // Free each individual string
            }
            free(strings);  // Free the array of string pointers
        }
        free(data);
    }

    // Close the socket
    close(clientSocket);

    return 0;
}
