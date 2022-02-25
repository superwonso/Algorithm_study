#include <stdio.h>
int nCk(int N, int K, int M){
    int A = 1, B = 1;
    for (int i = N; i > N-K; i--)
        A = A*i %M;
	
    for (int i = K; i >= 2; i--)
        B = B*i %M;

    N = 1, K = M-2;
    while (K){
        if (K & 1) N = N*B %M;
        K >>= 1;
        B = B*B %M;
    }
    return A*N %M;
}
int main(){
    long long N, K;
    int M;
    scanf("%lld %lld %d", &N, &K, &M);

    int result = 1;
    while (N || K){
        result = result*nCk(N%M, K%M, M) %M;
        N /= M, K /= M;
    }
    printf("%d", result);
}