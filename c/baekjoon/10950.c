#include <stdio.h>
int main(void)
{
    int t=0;
    scanf("%d",&t);
    int a[t],b[t];
    for(int temp1=0; temp1<t; temp1++)
    {
      scanf("%d %d",&a[temp1],&b[temp1]);  
    }
    
    for(int temp2=0; temp2<t; temp2++)
    {
      printf("%d\n",a[temp2]+b[temp2]);  
    }

}