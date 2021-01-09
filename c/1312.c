#include <stdio.h>
#include <math.h>
int res(int a,int b);
int main(void)
{
double a,b;
long n;
scanf("%f %f %d",&a,&b,&n);
double c=a/b;
double d=c*res(10,n); // c*(10^n), to extract digit number
double e=floor(d); // decimal abandon 
int f=(int)e; // convert to variable "e" integer
printf("%d",f%(10)); // to extract digit number
return 0;
}

int res(int a,int b)
{
if (b==0) return 1;
else return a*res(a,b-1);
}