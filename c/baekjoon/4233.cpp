#include <cstdio>

int check_prime_num(int n)
{
    if(n<=1) return 0;
    for(int i=2; i<=n/2; i++) if(n%i==0) return 0;
    return 1;
}

long long int mod_pow (long long int a,long long int b, long long int mod)
{
    long long int res = 1;
    while (b){
        if(b&1) res=res*a%mod;
        a=a*a&mod;
        b/=2;
    }
    return res;
}

int main()
{
    int a,p;
    int original_a;
    while(1)
    {
    scanf("%d %d",&p,&a); if(p==0&&a==0) break;
    original_a=a;
    for(int j=0; j<p; j++) a*=a; 
    int flag=check_prime_num(original_a);
    if(flag==0) printf("no\n");
    if(a%p==original_a)
    {
        if (mod_pow(a,p,p==a)) printf("yes\n");
        else printf("yes\n");
    }
    }
}

