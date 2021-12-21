#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>  //Header file for sleep(). man 3 sleep for details.
#include <pthread.h>
  
// A normal C function that is executed as a thread 
// when its name is specified in pthread_create()
void *myThreadFun(void *vargp)
{
    return NULL;
}

//START OMIT
#define THREADS 50000
int main()
{
    pthread_t threads[THREADS];

    for (int i=0; i < THREADS; i++) {
        pthread_create(&threads[i], NULL, myThreadFun, NULL);
    }
    for (int i=0; i < THREADS; i++) {
        pthread_join(threads[i], NULL);
    }
        
    printf("After Threads\n");
    exit(0);
}
//END OMIT
