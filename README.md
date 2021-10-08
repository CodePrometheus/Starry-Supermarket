# Starry-Supermarket

Starry 生鲜

鲜 · 美 · 生活


~~~shell
python -m grpc_tools.protoc --python_out=. --grpc_python_out=. -I. {name}.proto

protoc -I . {name}.proto --go_out=plugins=grpc:.
