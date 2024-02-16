#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <unistd.h>
#include <stdint.h>

unsigned char* serialize(char* strings[], int n, size_t* bufferSize) {
    size_t size = 0;
    for (int i = 0; i < n; i++) {
        size += 2 + strlen(strings[i]); // 2 bytes for length + string length
    }

    unsigned char* buffer = (unsigned char*)malloc(size);
    if (!buffer) {
        fprintf(stderr, "Memory allocation failed for serialization buffer\n");
        return NULL;
    }

    size_t offset = 0;
    for (int i = 0; i < n; i++) {
        uint16_t length = (uint16_t)strlen(strings[i]);
        memcpy(buffer + offset, &length, sizeof(length));
        offset += sizeof(length);

        memcpy(buffer + offset, strings[i], length);
        offset += length;
    }

    *bufferSize = size;
    return buffer;
}

void sendData(int clientSocket, unsigned char* data, size_t dataSize) {
    send(clientSocket, data, dataSize, 0);
}

int main() {
    int serverSocket, clientSocket;
    struct sockaddr_in serverAddr, clientAddr;
    socklen_t addrSize;

    // Create the socket
    serverSocket = socket(PF_INET, SOCK_STREAM, 0);

    // Configure settings of the server address struct
    serverAddr.sin_family = AF_INET;
    serverAddr.sin_port = htons(7799);
    serverAddr.sin_addr.s_addr = INADDR_ANY;
    memset(serverAddr.sin_zero, '\0', sizeof serverAddr.sin_zero);

    // Bind the address struct to the socket
    bind(serverSocket, (struct sockaddr *) &serverAddr, sizeof(serverAddr));

    // Listen on the socket
    if(listen(serverSocket, 5) == 0)
        printf("Listening\n");
    else
        printf("Error\n");

    // Accept call creates a new socket for the incoming connection
    addrSize = sizeof clientAddr;
    clientSocket = accept(serverSocket, (struct sockaddr *) &clientAddr, &addrSize);

    // Example data to send
    char* strings[] = {"Hello", "World"};
    size_t bufferSize;
    unsigned char* buffer = serialize(strings, 2, &bufferSize);

    // Send serialized data
    if (buffer) {
        sendData(clientSocket, buffer, bufferSize);
        free(buffer);
    }

    // Close the socket
    close(clientSocket);
    close(serverSocket);

    return 0;
}
