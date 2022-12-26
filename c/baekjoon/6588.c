#include <stdio.h>

int cases[1000001] = {0, };

int main() {
    int n;
    int k = 0;
    int i,j;
    int smallest;

    for (i = 2; i*i<1000001; i++) {
      if (!cases[i])
        for (j = i*i; j<1000001; j+=i) 
            if (j%i == 0)
                cases[j]=1;
    }

    while (scanf("%d",&n) && n>5) {
        for (i = n/2; i>1; i--) {
            if(cases[i] == 1) continue;
            for(j = n/2; j<=n; j++) {
                if(cases[j] == 1) continue;
                if(i+j == n) smallest = i;
            }
        }
    printf("%d = %d + %d\n", n, smallest, n-smallest);
    }

    return 0;
}

