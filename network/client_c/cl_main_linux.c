#include <arpa/inet.h>
#include <netinet/in.h>
#include <string.h>
#include <netdb.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include "data.h"

typedef struct sockaddr_in sockaddr_in;
typedef struct sockaddr sockaddr;

int main() {
    int sock = socket(AF_INET, SOCK_STREAM, 0);

    if (sock < 0) {
        puts("couldn;t create a socket\n");
        return -1;
    }

    sockaddr_in serverAddress;
    serverAddress.sin_family = AF_INET;
    serverAddress.sin_port = htons(9828);

    if (inet_pton(AF_INET, "176.124.202.165", &serverAddress.sin_addr)<=0) {
        puts("inet_pton error occured\n");
        return -1;
    }

    if (connect(sock, (sockaddr*)&serverAddress, sizeof(serverAddress)) < 0) {
        puts("connect error occured\n");
        return -1;
    }

    char characterID[16];
    readRandomData(characterID, 16);
    Packet packet;
    packet.Opcode = 32;
    memcpy(packet.CharacterID, characterID, 16);
    packet.X = 13.34;
    packet.Y = 20.36;
    packet.Z = 45.13;
    byte* arr = (byte*)&packet;

    while (1) {
        ssize_t total = 0;

        while (total < sizeof(Packet)) {
            ssize_t written = send(sock, &arr[total],
                sizeof(Packet) - total, 0);
            total += written;
        }
    }
}