#include <stdio.h>
int main()
{
    int a=0;
    int b,c;
    int next;
    int count=0;
    scanf("%d",&a);
    while(next==a)
{
    if(a<10) 
    {
    next=11*a;
    count=count+1;
    }
    else
    {
        b=a%10;
        c=(a-b)/10;
        next=b+c;
        count=count+1;
    }  
printf("%d",count);
}
}