#include <stdio.h>
#include "awesome.h"

int main() {
    printf("Using awesome lib from C:\n");

    
  GoInt a = 12;
    GoInt b = 99;
    printf("awesome.Add(12,99) = %d\n", (int)Add(a, b)); 
    printf("awesome.Cosine(1) = %f\n", (float)(Cosine(1.0)));


    GoString msg = {"Hello from C!", 13};
    Log(msg);
}

