Get-ChildItem -Directory -Path .\proto | ForEach-Object {
  $name = $_.Name
  protoc -I .\proto\$name .\proto\$name\*.proto --go_out=.\gen\go\$name --go_opt=paths=source_relative --go-grpc_out=.\gen\go\$name --go-grpc_opt=paths=source_relative
}
