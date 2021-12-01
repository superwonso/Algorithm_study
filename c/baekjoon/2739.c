#include <stdio.h>
int main(void)
{
    int input=0;
    scanf("%d",&input);
    for(int x=1;x<=9;x++)
    printf("%d * %d = %d\n",input,x,x*input);
    return 0;
}