#include <stdio.h>
#include <stdlib.h>

int main() {
int t;
scanf("%d",&t);
for (int i=1; i<=t; i++) {
int n;
scanf("%d", &n);
int *seq = malloc(n*sizeof(int));
for (int j = 0; j < n; j++) {
scanf("%d", &seq[j]);
}
int ans=0;
for (int j=0; j<n; j++) {
    for(int k=j; k<n; k++) {
        ans += seq[j]%seq[k];
    }
}
for (int j=0; j<n; j++) {
    for(int k=j; k<n; k++) {
        ans += seq[k]%seq[j];
    }
}
printf("#%d %d\n",i,ans);
free(seq);
}
return 0;
}