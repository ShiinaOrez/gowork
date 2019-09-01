// hello.c

#include <stdio.h>

void SayHello(const char* s) {
    int i;
    for(i=0; i<10; i++)
        printf("%s\n", s);
}
