#include <stdio.h>

int main()
{
    int n=0;
    scanf("%d",&n);
        for(int y=1;y<=n;y++)
        {
            for(int z=1;z<=n-y;z++)
            printf(" ");
            for(int x=1;x<=y;x++)
            printf("*");
        printf("\n");
        }
}