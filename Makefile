all:
	go build -o aromahousekpg .

dev:
	CGO_ENABLED=0 go build -a -installsuffix cgo -o aromahousekpg .
	- docker image rm localhost:32000/homestay/aromahousekpg
	docker build -t localhost:32000/homestay/aromahousekpg .
	docker push localhost:32000/homestay/aromahousekpg

prod:
	CGO_ENABLED=0 go build -a -installsuffix cgo -o aromahousekpg .
	- docker image rm reg.urantiatech.com/homestay/aromahousekpg
	docker build -t reg.urantiatech.com/homestay/aromahousekpg .
	docker push reg.urantiatech.com/homestay/aromahousekpg
