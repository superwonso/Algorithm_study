#include <stdio.h>

int main(){
    int n=0;
    int i,j;
    int cases[10001]={0, };
    
    for (i=2; i<10001/i; i++) {
      if(!cases[i]){
          for(j=i*i;j<10001; j+=i)
          if (j%i==0) cases[j]=1;
      }
    }
    
    while(1){
        scanf("%d",&n);
        if (n==0) return 0;
        for(i=n/2; i>1; i--) {
            if(cases[i]==1) continue;
            for(j=n/2; j<=n; j++){
                if(cases[j]==1) continue;
                if(i+j==n) goto OUT;
            }
        }
        OUT:printf("%d = %d + %d\n",i+j, i,j);
    }
    return 0;
}
