execute: libhello
	 go run main.go

a.out: libhello
	gcc -o main main.c -L. -lhello

libhello:
	gcc -c hello.c
	ar rusv libhello.a hello.o

clean:
	rm -f libhello.a hello.o main
