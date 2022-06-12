#include <stdlib.h>

typedef struct Packet {
    int Opcode;
    char CharacterID[16];
    double X;
    double Y;
    double Z;
} Packet;

typedef unsigned char byte;

void readRandomData(byte* arr, size_t size);