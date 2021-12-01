#include <cstdio>
const int max=(123456*2)+1;
int arr[max];

int main()
{
    arr[1]=arr[0]=1;
    for(int i=2; i<=max; i++)
    {
        if(!arr[i])
        {
            for (int j=i+i; j<max; j+=i) arr[j]=1;
        }
    }
//Sieve of Eratosthenes
    while(1)
    {
        int n,result=0;
        scanf("%d",&n);
        if(n==0) break;
        for(int i=n+1; i<=(2*n); i++) if(!arr[i]) result++;
        printf("%d\n",result);
    }

}