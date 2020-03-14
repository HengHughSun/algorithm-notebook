#include <stdio.h>
#include <stdlib.h>

typedef  char ElemType;

typedef struct  BiTNode
{
    /* data */
    ElemType data;
    struct BiTNode *lchild, *rchild;
} BiTNode, *BiTree;

// 创建一棵二叉树
void CreateBiTree(BiTree *T)
{
    char c;
    scanf("%c", &c);

    if (' ' == c)
    {
        *T = NULL;
    }
    else 
    {
        *T = (BiTNode *)malloc(sizeof(BiTNode));
        (*T)->data = c;
        CreateBiTree(&(*T)->lchild);
        CreateBiTree(&(*T)->rchild);
    }
}

void visit(ElemType c, int level)
{
    printf("%c in %d level \n", c, level);
}

// 遍历二叉树
void PreOrderTraversal(BiTree T, int level)
{
    if (T) 
    {
        visit((*T).data, level);
        PreOrderTraversal(T->lchild, level+1);
        PreOrderTraversal(T->rchild, level+1);
    }
}

int main()
{
    int level = 1;
    BiTree T = NULL;
    CreateBiTree(&T);
    PreOrderTraversal(T, level);
    return 0;
}

// InOrderThreading 

