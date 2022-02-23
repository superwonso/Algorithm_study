#include <cstdio>
#include <malloc.h>

int main() {
    int T;
    float input[11];
    scanf("%d",&T);
    float* output = (float*)malloc((sizeof(float) * T)+1);
    int i,j,k;
    float temp;
    for(i=1; i<=T; i++) {
        for(j=1; j<=10; j++) {
            scanf("%f ",&input[j]);
        }
        for(k=1; k<=10; k++) {
            temp+=input[k];
        }
        output[i]=temp/10;
        temp=0;
    }
    int l;
    for (l=1; l<=T; l++) {
        printf("#%d %.f\n",l,output[l]);
    }
    return 0;
}