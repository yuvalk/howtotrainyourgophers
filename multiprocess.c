#include <stdio.h>
#include <sys/types.h>
#include <unistd.h>

//START OMIT
int main() {
    int is_fork = 0;
    for (int i=0; i < 50000; i++) {
        if (fork() == 0) {
            is_fork = 1;
            break;
        };

    }

    return 0;
}
//END OMIT
