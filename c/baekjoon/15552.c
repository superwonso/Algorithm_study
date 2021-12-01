#include <stdio.h>
int main()
{
    int t;
    scanf("%d",&t);
    int a[t],b[t];
    
    for(int temp=0; temp<t; temp++)
    scanf("%d %d",&a[temp],&b[temp]);
    
    for(int temp0=0; temp0<t; temp0++)
    printf("%d\n",a[temp0]+b[temp0]);

}