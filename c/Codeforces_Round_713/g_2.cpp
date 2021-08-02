#include <cstdio>
int main()
{
    const int max=1e7+10;
    long long int d[max];
    long long int ans[max];

    for(int i=1; i<=10000000;i++)
    {
        for(int j=i;j<=1000000;j+=i) d[j]+=i; 
    }
    for(int i=1; i<=1000000; i++)
    {
        if(d[i]<=1000000&&!ans[d[i]]) ans[d[i]]=i;
    }
    int t;
    scanf("%d",&t);
    while (t--)
    {
        int c;
        scanf("%d",&c);
        if(!ans[c]) printf("-1\n");
        else printf("%lld\n",ans[c]);
    }
}
