#include <stdio.h>

int main(int argc, char const *argv[])
{
    int  a[5] = {1,2,3,4,5};
    
    int *p = &a[0];

    for (int i = 0; i < 5; i++)
    {
        
        printf("%d ",p[i]);
    }
    

    return 0;
}
