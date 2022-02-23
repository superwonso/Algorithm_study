#include <cstdio>
#include <malloc.h>

int main() {
int T;
scanf("%d",&T);
int input[11];
int* output = (int*)malloc((sizeof(int) * T)+1);
int i,j,k,l,result=0;
for (j=1; j<=T; j++) {
for (i=1;i<=10;i++) {
    scanf("%d ",&input[i]); 
    }
for (k=1;k<=10;k++) {
    if(input[k]%2==1){
        result+=input[k];
        }
    }
    output[j]=result;
    result=0;
}
for(l=1; l<=T; l++){
    printf("#%d %d\n",l,output[l]);
}
free(output);
return 0;
}
