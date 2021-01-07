#include <stdio.h>
int main(void)
{
    int a;
    scanf("%d",&a);
    for(int temp=a; temp>0; temp--)
    {
    printf("%d\n",temp); 
    }
    return 0;
}