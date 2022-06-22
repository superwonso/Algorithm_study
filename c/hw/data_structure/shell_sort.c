#include <stdio.h>
#include <string.h>
#include <stdlib.h>

void inc_insertion_sort(int list[], int from, int to, int gap);
void shell_sort(int list[], int n);
void string_shell_sort(char a[], int n);
void string_print_data(char a[], int n);

int main(void) {
  int i;
  int n;
  
  scanf("%d", &n);
  
  char *char_list = malloc(sizeof(char)*n);
  int *int_list = malloc(sizeof(int)*n); 
  
  printf("Enter the int elements. \n");
  for (int i=0; i<n; i++) {    
    scanf("%d", &int_list[i]);
  }

  printf("Enter the Alphabet elements.\n");
  for (int i=0; i<n; i++) {
    scanf("%c", &char_list[i]);
  }

  // 정수형 셸 정렬 수행
  shell_sort(int_list, n);

  // 정수형 정렬 결과 출력
  for(i=0; i<n; i++){
    printf("%d ", int_list[i]);
  }
  printf("\n");

  // 알파뱃 셸 정렬 수행
  string_shell_sort(char_list, strlen(char_list));
  // 정수형 셸 정렬 결과 출력
  string_print_data(char_list, strlen(char_list));
  
  return 0;
}

// 정렬의 범위는 from에서 to까지
void inc_insertion_sort(int list[], int from, int to, int gap){
  int i, j, key;
  for (i=from + gap; i <=to; i+=gap){
    key = list[i]; // 현재 삽입될 숫자인 i번째 정수를 key 변수로 복사

    // 현재 정렬된 배열은 i-gap까지이므로 i-gap번째부터 역순으로 조사한다.
    // j 값은 from 이상이어야 하고
    // key 값보다 정렬된 배열에 있는 값이 크면 j번째를 j+gap번째로 이동
    for(j=i-gap; j>=from && list[j]>key; j-=gap){
      list[j+gap] = list[j]; // 레코드를 gap만큼 오른쪽으로 이동
    }
    list[j+gap] = key;
  }
}

// 정수형 셸 정렬
// gap만큼 떨어진 요소들을 삽입 정렬
void shell_sort(int list[], int n){
  int i, gap;
  for(gap=n/2; gap>0; gap/=2){
    if((gap%2) == 0){
      gap++; // gap을 홀수로 만든다.
    }
    // 부분 리스트의 개수는 gap과 같다.
    for(i=0; i<gap; i++){
    // 부분 리스트에 대한 삽입 정렬 수행
    inc_insertion_sort(list, i, n-1, gap);
    }
  }
}

void string_shell_sort(char a[], int n) {
    int i, j, k, gap, v;
    for (gap = n/2; gap > 0; gap /=2) {
        for (i=0; i < gap; i++) {
            for (j = i+gap; j < n; j += gap) { /* Gap 에 있는 데이터 삽입 정렬 */
                v = a[j];
                k = j;
                while (k > gap-1 && a[k-gap] > v) {
                    a[k] = a[k-gap];
                    k -= gap;
                }
                a[k] = v;
            }
        }
    }
    
}

void string_print_data(char a[], int n) {
    int i;
    for (i = 0; i < n; i++) {
        printf ("%c ", a[i]);
    }
}