 #!/bin/sh
    SERVICE=$2
    CMD="./$3 -conf $4"
 
    start(){
        echo "starting $SERVICE..."

        num=`ps -ef | grep $SERVICE | grep -v grep | wc -l`
        if [ $num -eq 0 ]
        then
            nohup $CMD > /dev/null 2>&1 &  
            # nohup $CMD >output.log 2>&1 &
            if [ $? -ne 0 ]
            then
                echo "start failed, please check the log!"
                exit $?
            else
                echo $! > $SERVICE.pid 
                echo "start success"
            fi
        else
            echo "$SERVICE is already running"
        fi


    }
    stop(){
        echo "stopping $SERVICE..."
        kill -9 `cat $SERVICE.pid`
        if [ $? -ne 0 ]
        then
            echo "stop failed, may be $SERVICE isn't running"
            exit $?
        else
            rm -rf $SERVICE.pid 
            echo "stop success"
        fi
    }
    restart(){
        stop&&start
    }
    status(){
        num=`ps -ef | grep $SERVICE | grep -v grep | wc -l`
        if [ $num -eq 0 ]
        then
            echo "$SERVICE isn't running"
        else
            echo "$SERVICE is running"
        fi
    }
    case $1 in    
        start)      start ;;  
        stop)      stop ;;  
        restart)  restart ;;
        status)  status ;; 
        *)          echo "Usage: $0 {start|stop|restart|status}" ;;     
    esac  
 
    exit 0
