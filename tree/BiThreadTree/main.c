#include <stdio.h>
#include <stdlib.h>


typedef char ElemType;

typedef enum {Link, Thread} PointerTag;

typedef struct BiThrNode 
{
    ElemType data;
    struct  BiThrNode *lchild, *rchild;
    PointerTag ltag;
    PointerTag rtag;
} BiThrNode, *BiThrTree;

// 指向刚刚访问过的结点 全局变量
BiThrTree pre;

// 提醒用户遵照前序遍历的方式输入数据创建二叉树
void CreateBiThrTree(BiThrTree *T) 
{
    char c;
    scanf("%c", &c);

    if ( ' ' == c) 
    {
        *T = NULL;
    }
    else 
    {
        *T = (BiThrNode *)malloc(sizeof(BiThrNode));
        (*T)->data = c ;
        (*T)->ltag = Link;
        (*T)->rtag = Link;
        
        CreateBiThrTree(&(*T)->lchild);
        CreateBiThrTree(&(*T)->rchild);
    }
}
// 中序线索化
void InThreading(BiThrTree T) 
{
    if ( T ) {
        InThreading(T->lchild); // 递归左孩子线索化

        // 节点处理
        if ( !T->lchild )
        {
            T->ltag = Thread;
            T->lchild = pre;
        }
        if ( !pre->rchild ) // 如果pre只有左孩子，那就让它左右都指向左孩子，不同是的ltag是Link，rtag是Thread
        {
            pre->rtag = Thread;
            pre->rchild = T;
        }
        pre = T;
        InThreading(T->rchild); // 递归右孩子线索化


    }
}

void InOrderThreading(BiThrTree *PRE, BiThrTree T)
{
    *PRE = (BiThrTree)malloc(sizeof(BiThrNode));
    (*PRE)->ltag = Link;
    (*PRE)->rtag = Thread;
    (*PRE)->rchild= *PRE;
    if ( !T )
    {
        (*PRE)->lchild = *PRE;
    }
    else 
    {
        (*PRE)->lchild = T;
        pre = *PRE;
        InThreading(T);
        pre->rchild = *PRE;
        pre->rtag = Thread;
        (*PRE)->rchild = pre;

        // (*PRE)->rchild = pre;
        // pre->rtag = Thread;
        // pre->rchild = *PRE;
        // pre = *PRE;

    }
}

void visit(char c)
{
    printf("%c", c);
}

void InOrderTraversal(BiThrTree T)
{
    BiThrTree tmp;
    tmp = T->lchild;

    while ( tmp != T ) {
        while ( tmp->ltag == Link )
        {
            tmp = tmp->lchild;  
        }
        visit(tmp->data);
        while ( tmp->rtag == Thread && tmp->rchild != T )
        {
            tmp = tmp->rchild;
            visit(tmp->data);
        }
       tmp = tmp->rchild; 

    }

}

int main() 
{
    BiThrTree T,PRE  = NULL;
    CreateBiThrTree(&T);
    InOrderThreading(&PRE,T);
    InOrderTraversal(PRE);
    return 0;
}