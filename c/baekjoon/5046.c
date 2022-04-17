#include <stdio.h>

int main() {
int N,B,H,W;
scanf("%d %d %d %d",&N,&B,&H,&W);
int min, p, sum;
min = B+1;
for(int i=0; i<H; i++){
    scanf("%d",&p);
    sum = 0;
    for(int j=0,tmp; j<W; j++){
        scanf("%d",&tmp);
        if(tmp>=N){
            sum = p*N;
        }
    }
    if(min>sum&&sum!=0){
        min = sum;
    }
}
if(min>B){
    printf("stay home");
}else{
    printf("%d",min);
}
}