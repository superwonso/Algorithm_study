#include <stdio.h>
#include <stdlib.h>

int cases[1000001] = {0, };
int results[1000001] = {0, };

int main() {
    int n;
    int i, j;

    // 에라토스테네스의 체로 소수 판별
    for (i = 2; i * i < 1000001; i++) {
        if (!cases[i]) {
            for (j = i * i; j < 1000001; j += i)
                cases[j] = 1;
        }
    }

    // 소수 배열 생성
    int primes[1000001];
    int prime_count = 0;
    for (i = 2; i < 1000001; i++) {
        if (!cases[i]) {
            primes[prime_count++] = i;
        }
    }

    while (scanf("%d", &n) && n) {
        int found = 0;
        for (i = 0; i < prime_count && primes[i] <= n / 2; i++) {
            int complement = n - primes[i];
            if (!cases[complement]) {
                printf("%d = %d + %d\n", n, primes[i], complement);
                found = 1;
                break;
            }
        }
        if (!found) {
            printf("Goldbach's conjecture is wrong.\n");
        }
    }

    return 0;
}