// Minimum Cost Spanning Tree / Prim's Algorithm
#include <stdio.h>
#include <limits.h>
int main() {
    int cost[10][10] = {0, };
    int visited[10] = {0, };
    int n, no_e = 1;
    int min, a ,b ,min_cost = 0;
    printf("Enter number of nodes : ");
    scanf("%d",&n);
    printf("Enter cost in form of adjacency matrix \n");
    for(int i=1; i<=n; i++)
    {
        for(int j=1; j<=n; j++)
        {
            scanf("%d",&cost[i][j]);
            // If cost is 0, init it to INT_MAX
            if(cost[i][j] == 0)
              cost[i][j] = INT_MAX;
        }
    }
    //Find MST to use Prim's Algorithm
    visited[1] = 1; // First Node is always visited
    printf("The cource of MST is : 1");
    while(no_e < n)
    {
        min = INT_MAX;
        // Find minimum cost 
        for(int i = 1; i <= n; i++)
        {
            for(int j = 1; j <= n; j++)
            {
                if(cost[i][j] < min)
                {
                    if(visited[i] != 0)
                    {
                        min = cost[i][j];
                        a = i;
                        b = j;
                    }
                }
            }
        }
        if(visited[b] == 0)
        {
            printf(" -> %d",b);
            min_cost = min_cost + min;
            no_e++;
        }
        visited[b] = 1;
        // Init to INT_MAX
        cost[a][b] = cost[b][a] = INT_MAX; 
    }
    printf("\n Minimum Cost = %d",min_cost);
    return 0;
}
