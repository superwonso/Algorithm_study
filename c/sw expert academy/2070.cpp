#include <cstdio>

int main() {
int T;
scanf("%d",&T);
int o1,o2;
int res[3];
char result[3] = {'<', '>', '='};
for (int i=1; i<=T; i++) {
scanf("%d %d",&o1,&o2);
if (o1 > o2) {
res[i]=1;
} else if (o1 < o2) {
res[i]=0;
} else {
    res[i]=2;
    }
}

for (int i=1; i<=T; i++) {
printf("#%d %c\n",i,result[res[i]]);
}
return 0;
}