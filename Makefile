
all: drop-table run-demo view-table-vertical

run-demo:
	go run demo/mlog_demo/mlog_demo_main.go

view-table:
	echo "select * from common_wlog;" | mysql demo -u root -p -h db.demo.com

view-table-vertical:
	echo "select * from common_wlog \G " | mysql demo -u root -p -h db.demo.com

drop-table:
	echo "drop table common_wlog;" |\
		mysql demo -u root -p -h db.demo.com

