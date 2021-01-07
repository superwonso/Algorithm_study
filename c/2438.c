#include <stdio.h>
int main(void)
{
    int n=0;
    scanf("%d",&n);
        for(int y=1;y<=n;y++)
        {
            for(int x=1;x<=y;x++)
            {
            printf("*");
            }
        printf("\n");
        }
}