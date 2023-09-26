buffer: 
	go test -bench buffer_test.go

builder: 
	go test -bench builder_test.go

concat: 
	go test -bench concat_test.go
