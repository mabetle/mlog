
all: drop-table run-demo view-table-vertical

run-demo:
	go run demo/mlog_demo/mlog_demo_main.go

view-table:
	echo "select * from common_wlog;" | mysql demo -uroot -p

view-table-vertical:
	echo "select * from common_wlog \G " | mysql demo -uroot -p

drop-table:
	echo "drop table common_wlog;" | mysql demo -uroot -p


