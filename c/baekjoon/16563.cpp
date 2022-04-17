#include <cstdio>

const int max=5000000;
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
    int n=0;
    scanf("%d",&n);
    for (int i=0; i<=
}
