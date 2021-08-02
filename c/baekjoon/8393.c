#include <stdio.h>
int sigma(int n);

int main(void)
{
int n=0;
scanf("%d",&n);
printf("%d",sigma(n));
return 0;
}

int sigma(int n)
{
int temp,value=0;
for(temp=1; temp<=n; temp++)
{
value=temp+value;
}
return value;
}