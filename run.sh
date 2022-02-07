export PORT="2021"
export MYSQL_CONN_STRING="nevad_user:nevad_password@tcp(127.0.0.1:3307)/nevad_db?charset=utf8mb4&parseTime=True&loc=Local"
export SYSTEM_SECRET="NutxNARUywmdjHc9FdGQrITDG2bua/WY"

go run ./server.go