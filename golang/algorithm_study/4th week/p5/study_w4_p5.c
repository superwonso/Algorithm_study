#include <stdio.h>
const int MAX=100000;
int top=-1;
int stack[MAX];

void push(int data){
    if(top>=MAX-1) return;
    stack[++top]=data;
}

void pop(){
    if(top<0) return;
    stack[top--]=0;
}

int main(){
    int N,temp,sum=0;    
    scanf("%d",&N);

    while(N--){
        scanf("%d",&temp);
        if (temp==0) pop();
        else push(temp);
    }
    for(int i=0; i<=top; i++) sum+=stack[i];
    printf("%d\n",sum);
}
