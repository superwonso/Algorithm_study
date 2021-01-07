#include <stdio.h>
#include <math.h>
int main(void)
{
double a,b;
long n;
scanf("%f %f %d",&a,&b,&n);
double c=a/b;
double d=c*pow(10,n);
double e=floor(d);
long f=(long)e;
printf("%d",f%(10));
return 0;
}