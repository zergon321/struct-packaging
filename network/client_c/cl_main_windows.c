#include <winsock2.h>
#include <stdlib.h>
#include <stdio.h>
#include "data.h"

#pragma comment (lib, "Ws2_32.lib")

typedef struct addrinfo addrinfo;

int __cdecl main() {
    WSADATA wsaData;
    int iResult;
    SOCKET connectSocket = INVALID_SOCKET;
    addrinfo *result = NULL, *ptr = NULL, hints;

    iResult = WSAStartup(MAKEWORD(2,2), &wsaData);

    if (iResult != 0) {
        printf("WSAStartup failed with error: %d\n", iResult);
        return 1;
    }

    ZeroMemory(&hints, sizeof(hints));
    hints.ai_family = AF_INET;
    hints.ai_socktype = SOCK_STREAM;
    hints.ai_protocol = IPPROTO_TCP;

    iResult = getaddrinfo("v1215758.hosted-by-vdsina.ru", "9828", &hints, &result);

    if (iResult != 0) {
        printf("getaddrinfo failed with error: %d\n", iResult);
        WSACleanup();
        return 1;
    }

    for(ptr=result; ptr != NULL ;ptr=ptr->ai_next) {
        connectSocket = socket(ptr->ai_family, ptr->ai_socktype, 
            ptr->ai_protocol);

        if (connectSocket == INVALID_SOCKET) {
            printf("socket failed with error: %ld\n", WSAGetLastError());
            WSACleanup();
            return 1;
        }

        iResult = connect(connectSocket, ptr->ai_addr, (int)ptr->ai_addrlen);

        if (iResult == SOCKET_ERROR) {
            closesocket(connectSocket);
            connectSocket = INVALID_SOCKET;
        }
    }

    freeaddrinfo(result);

    if (connectSocket == INVALID_SOCKET) {
        printf("Unable to connect to server!\n");
        WSACleanup();
        return 1;
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
        int total = 0;

        while (total < sizeof(Packet)) {
            int written = send(connectSocket, &arr[total],
                sizeof(Packet) - total, 0);
            
            if (written == SOCKET_ERROR) {
                printf("send failed with error: %d\n", WSAGetLastError());
                closesocket(connectSocket);
                WSACleanup();
                return 1;
            }
            total += written;
        }
    }

    return 0;
}