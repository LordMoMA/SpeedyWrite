#include <stdio.h>
#include <pthread.h>

#define NUM_THREADS 1000

pthread_mutex_t lock;

void* writeToFile(void* threadid) {
    long tid;
    tid = (long)threadid;

    pthread_mutex_lock(&lock);

    FILE *f = fopen("cFile.txt", "a");
    if (f == NULL) {
        printf("Error opening file!\n");
        return NULL;
    }

    fprintf(f, "This is line %ld\n", tid);

    fclose(f);

    pthread_mutex_unlock(&lock);

    pthread_exit(NULL);
}

int main() {
    pthread_t threads[NUM_THREADS];
    int rc;
    long t;

    pthread_mutex_init(&lock, NULL);

    for(t = 0; t < NUM_THREADS; t++) {
        rc = pthread_create(&threads[t], NULL, writeToFile, (void *)t);
        if (rc) {
            printf("ERROR; return code from pthread_create() is %d\n", rc);
            return -1;
        }
    }

    for(t = 0; t < NUM_THREADS; t++) {
        pthread_join(threads[t], NULL);
    }

    pthread_mutex_destroy(&lock);

    return 0;
}