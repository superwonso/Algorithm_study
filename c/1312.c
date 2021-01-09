#include <stdio.h>
#include <math.h>
int main(void)
{
double a,b;
long n;
scanf("%f %f %d",&a,&b,&n);
double c=a/b;
double d=c*pow(10,n); // c*(10^n), to extract digit number
double e=floor(d); // decimal abandon 
long f=(long)e; // convert to variable "e" integer
printf("%d",f%(10)); // to extract digit number
return 0;
}