#include <cstdio>
#include <malloc.h>
#include <algorithm>

using namespace std;

int main() {
int T;
scanf("%d",&T);
int inp[11];
int *output = (int*)malloc((sizeof(int) * T)+1);
for (int i=1;i<=T;i++) {
for (int j=1; j<=10; j++) {
    scanf("%d",&inp[j]);
    output[i]=max(output[i],inp[j]);
}
}
for(int i=1; i<=T; i++) {
    printf("#%d %d\n",i,output[i]);
}
return 0;
}