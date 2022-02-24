#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int main() {
    int T;
    vector<int> res;
    cin>>T;
    for(int i=0; i<T; i++){
        vector<int> nums;
        long int N;
        cin>>N;
        for(int j=0; j<N; j++){
            scanf("%d",&nums[j]);
            }
        for(int j=nums.size()-1; j>0; j--){
            int MAX=nums[nums.size()-1];
            if(nums[j]>nums[j-1])
            {
                res[i]+=MAX-nums[j];
            } else {
                MAX=nums[j];
            }
        }
    printf("#%d %d",i+1,res[i]);
    }
    }
