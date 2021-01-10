all: 	install run

install:
	npm init;
	npm install express --save;
	sudo npm install pm2 -g;

run:
	pm2 start ejer6.js -i 4;

stop:
	pm2 stop ejer6;
