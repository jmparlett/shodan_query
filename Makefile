SHODANDIR = shodan/*

shodan: main.go $(SHODANDIR)
	go build main.go 

install:
	make shodan 
	mv main /usr/bin/shodan
