#include <stdio.h>
int main()
{
    int n,x;
    scanf("%d %d",&n,&x);
    int a[10000];
    for(int temp=0; temp<n; temp++)
    scanf("%d",&a[temp]); // scanf can receive values by typing data horizontally
    for(int temp=0; temp<n; temp++)
    {
        if(x>a[temp])
        {
            printf("%d",a[temp]);
        }
    }
}

