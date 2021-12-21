.PHONY: mthread
mthread:
	gcc -pthread multithread.c -o mthread

.PHONY: mproc
mproc:
	gcc multiprocess.c -o mproc

