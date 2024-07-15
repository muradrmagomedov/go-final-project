.PHONY:

run:
	go run ./cmd/.
test1:
	go test -run ^TestApp$ ./tests
test2:
	go test -run ^TestDB$ ./tests
test3:
	go test -run ^TestNextDate$ ./tests
test4:
	go test -run ^TestAddTask$ ./tests
test5:
	go test -run ^TestTasks$ ./tests
test6:
	go test -run ^TestEditTask$ ./tests