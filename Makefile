generate-grpc:
	protoc -I ../proto --go_out=plugins=grpc:../proto ../proto/*.proto
	python -m grpc_tools.protoc -I ../proto --python_out=../web-api --grpc_python_out=../web-api ../proto/*.proto
