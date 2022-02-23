#include <iostream>
#include <cstdio>
#include <algorithm>
#include <vector>
#include <malloc.h>

int main() {
    int N,res;
    scanf("%d", &N);
    while(N>0) {
        res+=(N%10);
		N/=10;        
    }
    printf("%d",res);
}