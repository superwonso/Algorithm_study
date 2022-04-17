#include <stdio.h>

int f[40];
int fibo(int n);

int main() {
int n;
scanf("%d",&n);
printf("%d %d",fibo(n),n-2);
}

int fibo(int n) {
    if(n==1)
        return 1;
    else if(n==2)
        return 1;
    else
        return fibo(n-1)+fibo(n-2);
}
