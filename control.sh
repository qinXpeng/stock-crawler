#!/bin/bash

cur_file_abs_path=$(readlink -f "${BASH_SOURCE[0]}")
project_dir_abs_path=$(dirname "${cur_file_abs_path}")
setting_file_abs_path="${project_dir_abs_path}/conf/setting.toml"

get_setting_http_listen_port(){
	local setting_file_abs_path=$1

	local http_listen_port=$(cat "${setting_file_abs_path}" | grep -P 'port=' | grep -Po '(?<=")\d*(?=")')
	if [ -z "${http_listen_port}" ]
	then
		echo ""
		return 0
	fi

	echo "${http_listen_port}"
	return 0
}

get_pid_by_listen_port(){
	local listen_port=$1
	if [ -z "${listen_port}" ]
	then
		echo ""
		return 1
	fi

	local process_id=$(lsof -t -i ":${listen_port}")
	if [ -z "${process_id}" ]
	then
		echo ""
		return 1
	fi

	echo "${process_id}"
	return 0
}

start(){
	local project_dir_abs_path=$1

	local process_name=$(basename "${project_dir_abs_path}")

	local run_dir_abs_path="${project_dir_abs_path}/output"

	cd "${run_dir_abs_path}"
	echo "${run_dir_abs_path}"
	nohup ./bin/"${process_name}" > nemo.log 2>&1 &
}

stop(){
	local pid=$1

	echo "kill pid: ${pid}"
	kill -9 "${pid}"
}

case $1 in

	start)
		start "${project_dir_abs_path}"

		listen_port=$(get_setting_http_listen_port "${setting_file_abs_path}")
		count=10
		while [ -z "${run_pid}" ]
		do
			run_pid=$(get_pid_by_listen_port "${listen_port}")

			count=$((count-1))
			if [ $count -le 0 ]
			then
				break
			fi

			sleep 1
		done
		echo "running pid: ${run_pid}"
		;;

	stop)
		listen_port=$(get_setting_http_listen_port "${setting_file_abs_path}")
		echo "listen port: ${listen_port}"
		if [ -z "${listen_port}" ]
		then
			echo "http listen port not specified"
			exit 1
		fi

		pid=$(get_pid_by_listen_port "${listen_port}")
		echo "running pid: ${pid}"
		if [ -z "${pid}" ]
		then
			echo "process of pid:${pid} not running"
			exit 0
		fi

		stop "${pid}"
		;;

	*)
		echo "usage:"
		echo "       sh control.sh /start/stop/"
		;;
esac
