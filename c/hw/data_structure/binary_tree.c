// binary tree
#include <stdio.h>
#include <stdlib.h>
#include <limits.h> // to define INT_MIN, INT_MAX

typedef struct node {
    int data;
    struct node *left;
    struct node *right;
} node;

node *root = NULL;

node *create_node(int data);
node *insert_node(node *root, int data);
int check_BST(node* root);
int check_BST_util(node* root, int min, int max);

int main() {
    node *example1=NULL; // non BST
    example1 = insert_node(example1, 3);
    example1 = insert_node(example1, 2);
    example1 = insert_node(example1, 5);
    example1 = insert_node(example1, 2);
    example1 = insert_node(example1, 2);
    
    node *example2=NULL; // BST
    example2 = insert_node(example2, 10);
    example2 = insert_node(example2, 5);
    example2 = insert_node(example2, 17);
    example2 = insert_node(example2, 1);
    example2 = insert_node(example2, 6);
    example2 = insert_node(example2, 14);
    example2 = insert_node(example2, 21);
    
    printf("%d\n",example1->data);
    printf("%d\n",example2->data);
    
    if (check_BST(example1)) {
        printf("example1 is BST\n");
    } else {
        printf("example1 is not BST\n");
    }
    
    if (check_BST(example2)) {
        printf("example2 is BST\n");
    } else {
        printf("example2 is not BST\n");
    }
    
    return 0;
}

node *create_node(int data) {
    node *new_node = (node *)malloc(sizeof(node));
    new_node->data = data;
    new_node->left = NULL;
    new_node->right = NULL;
    return new_node;
}

node *insert_node(node *root, int data) {
    if (root == NULL) {
        root = create_node(data);
    } else if (data < root->data) {
        root->left = insert_node(root->left, data);
    } else {
        root->right = insert_node(root->right, data);
    }
    return root;
}

int check_BST(node* root)
{
  return(check_BST_util(root, INT_MIN, INT_MAX));
}
 
int check_BST_util(node* root, int min, int max)
{
  if (root==NULL)
     return 1;
       
  if (root->data < min || root->data > max)
     return 0;
 
  return
    check_BST_util(root->left, min, root->data-1) && check_BST_util(root->right, root->data+1, max); 
}
