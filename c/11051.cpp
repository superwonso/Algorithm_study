#include <cstdio>

const int mod = 10007;
int dp[1001][1001]={1, };

int pascal(int m,int k)
{
    for(int i=1; i<=m; ++i)
    {
        dp[i][0]=dp[i][i]=1;
        for(int j=1;j<i;++j) dp[i][j]=((dp[i-1][j]%mod)+(dp[i-1][j-1]%mod))%mod;
    }
    return dp[m][k];
}

int main()
{
    int m,k;
    scanf("%d %d",&m,&k);
    printf("%d\n",pascal(m,k));
}
