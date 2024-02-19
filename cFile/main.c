#include <stdio.h>
#include <time.h>
#include <math.h>

#define NUM_LINES 90000

int main() {
    clock_t start, end;
    double cpu_time_used;

    start = clock();

    FILE *f = fopen("cFile.txt", "w");
    if (f == NULL) {
        printf("Error opening file!\n");
        return -1;
    }

    int buffer_size = snprintf(NULL, 0, "This is line %d\n", NUM_LINES) + 1;
    char buffer[buffer_size];
    for(int i = 0; i < NUM_LINES; i++) {
        sprintf(buffer, "This is line %d\n", i);
        fputs(buffer, f);
    }

    fclose(f);

    end = clock();
    cpu_time_used = ((double) (end - start)) / CLOCKS_PER_SEC;

    printf("Time taken: %f milliseconds\n", cpu_time_used * 1000);

    return 0;
}