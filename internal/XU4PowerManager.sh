#!/bin/bash

function readSensors(){

	governor_cup_curr=$(cat /sys/devices/system/cpu/cpu0/cpufreq/scaling_governor)
	governor_gup_curr=$(cat /sys/devices/platform/11800000.mali/devfreq/11800000.mali/governor)


	FAN=`cat $FAN_SPEED_FILE`
	FAN_MODE=$(cat $FAN_MODE_FILE)
	FAN=$((FAN * 100 / 255))

	t=`cat /sys/devices/virtual/thermal/thermal_zone0/temp`
	t0=$(( $t/1000))
	 
	t=`cat /sys/devices/virtual/thermal/thermal_zone1/temp`
	t1=$(( $t/1000))
	 
	t=`cat /sys/devices/virtual/thermal/thermal_zone2/temp`
	t2=$(( $t/1000))
	 
	t=`cat /sys/devices/virtual/thermal/thermal_zone3/temp`
	t3=$(( $t/1000))
	 
	t=`cat /sys/devices/virtual/thermal/thermal_zone4/temp`
	tg=$(( $t/1000))
	
 	#small cores
	f=`cat /sys/devices/system/cpu/cpu0/cpufreq/scaling_cur_freq`
	#f0=$(( $f/1000))	#usable
	f0=$(echo 'scale=2;' "$f/1000000" | bc -l )
	
	f=`cat /sys/devices/system/cpu/cpu1/cpufreq/scaling_cur_freq`
	f1=$(echo 'scale=2;' "$f/1000000" | bc -l )
	
	f=`cat /sys/devices/system/cpu/cpu2/cpufreq/scaling_cur_freq`
	f2=$(echo 'scale=2;' "$f/1000000" | bc -l )
	
	f=`cat /sys/devices/system/cpu/cpu3/cpufreq/scaling_cur_freq`
	f3=$(echo 'scale=2;' "$f/1000000" | bc -l )

 	#big cores
	f=`cat /sys/devices/system/cpu/cpu4/cpufreq/scaling_cur_freq`
	f4=$(echo 'scale=2;' "$f/1000000" | bc -l )
 
	f=`cat /sys/devices/system/cpu/cpu5/cpufreq/scaling_cur_freq`
	f5=$(echo 'scale=2;' "$f/1000000" | bc -l )
 
	f=`cat /sys/devices/system/cpu/cpu6/cpufreq/scaling_cur_freq`
	f6=$(echo 'scale=2;' "$f/1000000" | bc -l )
 
	f=`cat /sys/devices/system/cpu/cpu7/cpufreq/scaling_cur_freq`
	f7=$(echo 'scale=2;' "$f/1000000" | bc -l )
	
	
	f=`cat /sys/devices/platform/11800000.mali/devfreq/11800000.mali/cur_freq`
	fg=$(( $f/1000000))
	
	
	ar=($t0 $t1 $t2 $t3 $tg)
	IFS=$'\n'
	
	t_max=`echo "${ar[*]}" | sort -nr | head -n1`
}

function readSensorsLight(){

	# governor_cup_curr=$(cat /sys/devices/system/cpu/cpu0/cpufreq/scaling_governor)
	# governor_gup_curr=$(cat /sys/devices/platform/11800000.mali/devfreq/11800000.mali/governor)


	FAN=`cat $FAN_SPEED_FILE`
	# FAN_MODE=$(cat $FAN_MODE_FILE)
	# FAN=$((FAN * 100 / 255))

	t=`cat /sys/devices/virtual/thermal/thermal_zone0/temp`
	t0=$(( $t/1000))
	 
	t=`cat /sys/devices/virtual/thermal/thermal_zone1/temp`
	t1=$(( $t/1000))
	 
	t=`cat /sys/devices/virtual/thermal/thermal_zone2/temp`
	t2=$(( $t/1000))
	 
	t=`cat /sys/devices/virtual/thermal/thermal_zone3/temp`
	t3=$(( $t/1000))
	 
	t=`cat /sys/devices/virtual/thermal/thermal_zone4/temp`
	tg=$(( $t/1000))
	
 	# #small cores
	# f=`cat /sys/devices/system/cpu/cpu0/cpufreq/scaling_cur_freq`
	# #f0=$(( $f/1000))	#usable
	# f0=$(echo 'scale=2;' "$f/1000000" | bc -l )
	
	# f=`cat /sys/devices/system/cpu/cpu1/cpufreq/scaling_cur_freq`
	# f1=$(echo 'scale=2;' "$f/1000000" | bc -l )
	
	# f=`cat /sys/devices/system/cpu/cpu2/cpufreq/scaling_cur_freq`
	# f2=$(echo 'scale=2;' "$f/1000000" | bc -l )
	
	# f=`cat /sys/devices/system/cpu/cpu3/cpufreq/scaling_cur_freq`
	# f3=$(echo 'scale=2;' "$f/1000000" | bc -l )

 	# #big cores
	# f=`cat /sys/devices/system/cpu/cpu4/cpufreq/scaling_cur_freq`
	# f4=$(echo 'scale=2;' "$f/1000000" | bc -l )
 
	# f=`cat /sys/devices/system/cpu/cpu5/cpufreq/scaling_cur_freq`
	# f5=$(echo 'scale=2;' "$f/1000000" | bc -l )
 
	# f=`cat /sys/devices/system/cpu/cpu6/cpufreq/scaling_cur_freq`
	# f6=$(echo 'scale=2;' "$f/1000000" | bc -l )
 
	# f=`cat /sys/devices/system/cpu/cpu7/cpufreq/scaling_cur_freq`
	# f7=$(echo 'scale=2;' "$f/1000000" | bc -l )
	
	
	# f=`cat /sys/devices/platform/11800000.mali/devfreq/11800000.mali/cur_freq`
	# fg=$(( $f/1000000))
	
	
	ar=($t0 $t1 $t2 $t3 $tg)
	IFS=$'\n'
	
	t_max=`echo "${ar[*]}" | sort -nr | head -n1`
}



function formatTemps(){
#clear;
formattedTemps=$(
	
	
	echo "FAN1: $FAN%	 auto:$FAN_MODE" 
	#small cores
	echo "CPU 0		" $f0" GHz"
	echo "CPU 1		" $f1" GHz"
	echo "CPU 2		" $f2" GHz"
	echo "CPU 3		" $f3" GHz"
 
	#big cores
	echo "CPU 4	" $t0"C,	" $f4" GHz"
	echo "CPU 5	" $t1"C,	" $f5" GHz"
	echo "CPU 6	" $t2"C,	" $f6" GHz"
	echo "CPU 7	" $t3"C,	" $f7" GHz	"$governor_cup_curr
	#gpu
	echo "GPU 0	" $tg"C,	" $fg" MHz	"$governor_gup_curr
	
	echo "MAXtemp	" $t_max"C	 trip points:	$trip_point_tr0 $trip_point_tr1 $trip_point_tr2"	
)
}

function finish {
	if [ $DO_FAN -gt 0 ];then
	echo "set fan to auto"
		echo 255 | sudo tee $FAN_SPEED_FILE
		echo 1 | sudo tee $FAN_MODE_FILE
	fi
	
	if [ $DO_SERVICE -gt 0 ];then
		sudo rm -f $mypidfile
	fi
}

function detectSudo(){
	hasSudo=1
	if [ "$EUID" -ne 0 ];then
		hasSudo=0
		echo "Please run as root"
		exit
	fi
}

function setGovernor(){
	detectSudo
	echo $1 | sudo tee /sys/devices/system/cpu/cpu{0,4}/cpufreq/scaling_governor
	echo $2 | sudo tee cat /sys/devices/platform/11800000.mali/devfreq/11800000.mali/governor
}

function printHelp(){
cat << EOF

Usage: XU4PowerManager.sh [OPTION] 

blank = -l -pa

-l	run as loop
-pa	print for human: temperature, fan speed, freq ...
-ph print updating pa
-pm	print max temp
-pf	print fan speed
-ps	print current settings

-sf	set fan [-s] (s=speed. Leave empty to set automatic)
-sg	set governor [cpu_gov] [gpu_gov]

-service run as service, get max temp and adjust fan.

settings in /usr/local/etc/XU4PowerManager.cfg
EOF


}


last_log_time=0
log_time_delay=60
logStats(){
	log_time_curr=`date +%s`
	if [ "$log_time_curr" -gt "$last_log_time" ]; then
		echo $t_max, $log_time_curr | tee -a "/var/log/XU4Stats/temp.log" #>> "/media/XU4_SSD/server/www/ondsi.net/public_html/log/temp.log"
		echo $FAN, $log_time_curr  | tee -a "/var/log/XU4Stats/fan.log" #>> "/media/XU4_SSD/server/www/ondsi.net/public_html/log/fan.log"
		#echo $t_max, $log_time_curr >> /var/log/XU4Stats/temp.log
		last_log_time=$(($log_time_curr+$log_time_delay))
	fi
}


#########################MAIN#######################################
trap finish EXIT


FAN_MODE_FILE="/sys/devices/platform/pwm-fan/hwmon/hwmon0/automatic"
FAN_SPEED_FILE="/sys/devices/platform/pwm-fan/hwmon/hwmon0/pwm1"

governor_cpu_def=powersave #conservative userspace powersave ondemand performance schedutil
governor_gpu_def=powersave #userspace powersave performance simple_ondemand

trip_point_0=45
trip_point_1=55
trip_point_2=65

trip_point_tr0=$((trip_point_0 * 85 / 100))
trip_point_tr1=$((trip_point_1 * 85 / 100))
trip_point_tr2=$((trip_point_2 * 85 / 100))

fan_speed_min=50
fan_speed_0=255
fan_speed_1=255
fan_speed_2=255


DO_LOG=0
source /usr/local/etc/XU4PowerManager.cfg



#		-pa -l

#-l		run as loop
#-pa	print for human: temperature, fan speed, freq
#-pm	print max temp
#-pf	print fan speed
#-ps	print current settings

#-sf	set fan [-s] (s=speed. Leave empty to set automatic)
#-sg	set governor [cpu_gov] [gpu_gov]

#-service		run as service, get max temp and adjust fan (every 5 sec)


{
DO_SERVICE=0
DO_PRINTALL=0
DO_PRINTMAX=0
DO_PRINTFAN=0
DO_PRINTSETT=0
DO_LOOP=0
DO_FAN=0
DO_GOVERNOR=0

hasSudo=0

mypidfile=/var/run/XU4PowerManager.pid

trip_point_curr=0


if [ $# -eq 0 ]; then 
	#echo "asd "[ "$EUID" -ne 0 ] " " [ -f $mypidfile ]
	#if [ "$EUID" -ne 0 ] || [ -f $mypidfile ];then 
		DO_PRINTALL=1;DO_LOOP=1
	#else
	#	DO_SERVICE=1
	#fi
fi


while test $# -gt 0
do
	case "$1" in
		-l) DO_LOOP=1
			;;
		-pa) DO_PRINTALL=1; 
			;;
		-pm) DO_PRINTMAX=1
			;;
		-pf) DO_PRINTFAN=1
			;;
		-ps) DO_PRINTSETT=1
			;;
		-ph) DO_PRINTALL=1;DO_LOOP=1
			;;
		-sf) DO_FAN=1
			;;
		-sg) DO_GOVERNOR=1
			;;
		-service) DO_SERVICE=1
			;;
		-h) printHelp
			;;
		-v)	echo "XU4PowerManager Version 1.2 (April 15th, 2020)"; exit 0
			;;
		-*) echo "bad option $1"; exit
			;;
		*) break;
			;;
	esac
	shift
done
}


if [ $DO_SERVICE -gt 0 ];then
	detectSudo
	echo "service"
	if [ -f $mypidfile ]; then
		echo "process already running"
		DO_SERVICE=0
		exit
	fi
	echo $$ > "$mypidfile"

	DO_FAN=1 
	setGovernor $governor_cpu_def $governor_gpu_def
	echo 0 | sudo tee $FAN_MODE_FILE
	echo 255 | sudo tee $FAN_SPEED_FILE
	while true :
	do
		#clear
	
		readSensorsLight
		
		if [ $t_max -lt $trip_point_tr0 ]; then trip_point_curr=0; fi
		if [ $t_max -gt $trip_point_0 ]; then trip_point_curr=1; fi
		if [ $t_max -gt $trip_point_1 ]; then trip_point_curr=2; fi
		if [ $t_max -gt $trip_point_2 ]; then trip_point_curr=3; fi
		

		if [ $trip_point_curr -eq 0 ]; 
		then 
		echo $fan_speed_min | sudo tee $FAN_SPEED_FILE 
		#elif 
		else
		echo $fan_speed_1 | sudo tee $FAN_SPEED_FILE
		fi
		
		# echo "service. trip_point=$trip_point_curr max=$t_max "
		# echo "$trip_point_0 $trip_point_tr0"
		# echo "$trip_point_1 $trip_point_tr1"
		# echo "$trip_point_2 $trip_point_tr2"
		
		
		logStats
		
		sleep 4
	done 
	exit
fi

 
#if not running as service:

newline=$(clear;);

formattedTemps=""
while [ $DO_LOOP -gt 0 ]
do
	readSensors
	if [ $DO_PRINTALL -eq 1 ]; then formatTemps; fi
	if [ $DO_PRINTMAX -eq 1 ]; then formattedTemps=$(echo "MAXtemp	" $t_max"C"); fi
	if [ $DO_PRINTFAN -eq 1 ]; then formattedTemps=$(clear; echo "FAN1: $FAN %" ); fi
	if [ $DO_PRINTSETT -eq 1 ]; then formattedTemps=$(clear; echo "MAXtemp	" $t_max"C"); fi

	echo $newline"$formattedTemps"
	sleep 0.2
done

#run once
readSensors
if [ $DO_PRINTALL -eq 1 ]; then formatTemps; fi
if [ $DO_PRINTMAX -eq 1 ]; then formattedTemps=$(echo "MAXtemp	" $t_max"C"); fi
if [ $DO_PRINTFAN -eq 1 ]; then formattedTemps=$(echo "FAN1: $FAN %" ); fi
if [ $DO_PRINTSETT -eq 1 ]; then formattedTemps=$(echo "MAXtemp	" $t_max"C"); fi

if [ ! -z "$formattedTemps" ]; then echo "$formattedTemps"; fi


if [ $DO_GOVERNOR -eq 1 ]; then detectSudo; setGovernor $1 $2; fi
if [ $DO_FAN -eq 1 ]; then 
	detectSudo; 
	if [ -z $1 ]; then exit; fi
	DO_FAN=0
	echo 0 | sudo tee $FAN_MODE_FILE;
	echo $1 | sudo tee $FAN_SPEED_FILE; 
fi

exit

#cpufreq-info -o
