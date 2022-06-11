#include <stdlib.h>
#include <time.h>
#include "data.h"

void readRandomData(byte* arr, size_t size) {
    srand(time(NULL));

    for (int i = 0; i < size; i++) {
        arr[i] = rand();
    }
}