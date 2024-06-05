#include <stdio.h>

int main() {
int nums[5] = {2,4,7,8,9};
int i, total = 0, *n_pt;
n_pt = nums;
for (i=0; i<=4; ++i){
total += *(n_pt)++;
}
printf("%d\n", total);
}
