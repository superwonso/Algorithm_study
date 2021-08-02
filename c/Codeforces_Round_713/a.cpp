#include <cstdio>
int main()
{
    int t;
    int an[100][100];
    scanf("%d",&t);
    for(int temp_col=0; temp_col<t; temp_col++)
    {
        for(int temp_row=0; temp_row<an[temp_col][temp_row]; temp_row++)
        {
            scanf("%d",an[temp_col][temp_row]);
        }
    }
    int flag=-1;
    for(int temp_col=0; temp_col<t; temp_col++)
    {
        for(int temp_row=0; temp_row<an[temp_col][temp_row]; temp_row++)
        {
            for(int check=1; check<an[temp_col][temp_row]-1; check++)
            {
                if(an[temp_col][temp_row]!=an[temp_col][temp_row-1]&&an[temp_col][temp_row]!=an[temp_col][temp_row+1]) flag=check; break;
                if(flag==-1)
                {
                    if(an[temp_col][0]!=an[temp_col][1]&&an[temp_col][1]==an[temp_col][2]) flag=0;
                }
                if(flag==-1)
                {
                    if(an[temp_col][temp_row-1]!=an[temp_col][temp_row-2]&&an[temp_col][temp_row-2]==an[temp_col][temp_row-3]) flag=temp_row-1;
                }
            }
            printf("%d",flag);
        }
    }

}