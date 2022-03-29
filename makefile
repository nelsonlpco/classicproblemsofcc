install-githooks:
	chmod +x ./scripts/pre-commit; \
	cp ./scripts/pre-commit ./.git/hooks/pre-commit

fibonacci-b:
	go test -bench=. -run=^# -count 1 ./fib | tee ./graphics/out.dat; \
	awk '/Benchmark/{count ++; gsub(/Benchmark_/,""); printf("%d,%s,%s,%s\n",count,$$1,$$2,$$3)}' ./graphics/out.dat > ./graphics/pout.dat; \
	gnuplot -e "file_path='./graphics/pout.dat" -e "graphic_file_name='./graphics/operations_fibonacci.png'" -e "y_label='number of operations'" -e "y_range_min='000000000'" -e "y_range_max='100000000'" -e "column_1=1" -e "column_2=3" ./graphics/performance.gp; \
	gnuplot -e "file_path='./graphics/pout.dat" -e "graphic_file_name='./graphics/time_fibonacci.png'" -e "y_label='operations in nanoseconds'" -e "y_range_min='000'" -e "y_range_max='150'" -e "column_1=1" -e "column_2=4" ./graphics/performance.gp; \
	echo "'graphics/operations_fibonacci.png' and 'graphics/time_fibonacci.png' graphics were generated."

fibonacci-t:
	go test -cover ./fib

fibonacci: fibonacci-t fibonacci-b