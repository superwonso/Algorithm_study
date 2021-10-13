#include <stdio.h>
int main()
{
    int n=0;
    scanf("%d",&n);
    int arr[n][n]; 
    int sum_row[n],sum_col[n],sum_diag_1=0,sum_diag_2=0;
    for(int y=0;y<n;y++) for(int z=0;z<n;z++) arr[y][z]=0;
    for(int l=0; l<n; l++)
    {
        for(int k=0; k<n; k++) scanf("%d",&arr[l][k]);
    }
    
    for(int aa=0;aa<n;aa++) {sum_row[aa]=0; sum_col[aa]=0;} 
    for(int o=0;o<n;o++) for(int m=0;m<n;m++) sum_row[o]+=arr[o][m]; 
    for(int p=0;p<n;p++) for(int q=0;q<n;q++) sum_col[p]+=arr[q][p];
    for(int s=0;s<n;s++) sum_diag_1+=arr[s][s];
    for(int t=n-1;t>=0;t--) sum_diag_2+=arr[t][n-t-1];
    
    int max1=0;
    int max2=0;
    int result=0;
    for(int u=0;u<n;u++) 
    {
        if(sum_row[u]>max1) max1=sum_row[u];
        if(sum_col[u]>max2) max2=sum_col[u];
    }
    if(max2>max1) result=max2; else result=max1;
    if(result>sum_diag_1) result=result; else result=sum_diag_1;
    if(result>sum_diag_2) result=result; else result=sum_diag_2;
    printf("%d\n",result); 
return 0;
} // 2n+2번의 계산