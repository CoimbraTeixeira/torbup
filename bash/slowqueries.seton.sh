#!/bin/bash
LOG=/tmp/logfile.log
#alterar destinatarios
USERS=none@email.com
USERS_ERRO=none@email.com
 
mv $LOG $LOG.old  2> /dev/null
 
rm -rf /tmp/logfile.log.old
 
cp /tmp/logfile.log /tmp/logfile.log.old
 
echo "" >  /tmp/logfile.log
 
echo "Begin $0 -`date`" > $LOG
 
echo "" >> $LOG
echo "Conected to MySQL IP" >> $LOG
echo "Enable slow queries" >> $LOG
 
/usr/local/mysql/bin/mysql -h <IP> -u root -p<PASSWD> -t -v -v << !  >> $LOG 2>&1
 
select @@hostname,now();
set global slow_query_log_file='/tmp/logfile.log';
set global log_queries_not_using_indexes=1;
set global slow_query_log = 1;
set global long_query_time=2;
set global min_examined_row_limit = 5;
set global log = 0;
set global general_log=0;
set global log_slow_queries = 1;
set global log_warnings = 2;
show global variables;
 
!
 
echo "" >> $LOG
echo "Logs" >> $LOG
echo "Finish $0 -`date`" >> $LOG
 
grep -qi "ERROR 1" $LOG
if test $? -eq 0
then
   mail -s "Script has erros $0" $USUARIOS_ERRO < $LOG
else
   mail -s "Script has no errors $0" $USUARIOS < $LOG
fi