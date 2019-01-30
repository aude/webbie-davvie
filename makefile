.PHONY: all
all: webbie-davvie

.PHONY: clean
clean:
	rm -f webbie-davvie

webbie-davvie: main.go
	go build -o webbie-davvie main.go

.PHONY: run
run: webbie-davvie
	./webbie-davvie
