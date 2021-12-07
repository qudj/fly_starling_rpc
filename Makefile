regenerate:
	cd models/proto; protoc --gogofaster_out=plugins=grpc:. fly_starling.proto; mkdir ../fly_starling_serv; mv fly_starling.pb.go ../fly_starling_serv;
	cp models/proto/fly_starling.proto ../fly_lib/proto/