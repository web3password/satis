#!/usr/bin/env bash

ET_HOME=`pwd`

web3_satis_temp_conf_file=$ET_HOME/conf-template/satis.yaml
web3_satis_conf_file=$ET_HOME/conf/satis.yaml
web3_ares_conf_file=$ET_HOME/conf/ares.yaml
web3_storage_conf_file=$ET_HOME/conf/storage.yaml
web3_index_conf_file=$ET_HOME/conf/index.yaml
w3p_satis_service_file=$ET_HOME/scripts/w3p-satis.service
w3p_ares_service_file=$ET_HOME/scripts/w3p-ares.service
w3p_storage_service_file=$ET_HOME/scripts/w3p-storage.service
w3p_index_service_file=$ET_HOME/scripts/w3p-index.service
w3p_kvrocks_service_file = $ET_HOME/scripts/kvrocks.service

function display_menu() {
    echo "============================================================================="
    echo "web3password requires golang environment(above 1.20) and kvrocks(127.0.0.1:6666)."
    echo "web3password uses ports:8893 8894 8895 8896 8897,please make sure they aren't be used."
    echo "8893 is access port,others ports is used by internal."
    echo "============================================================================="
    echo "Please choose you server's config:"
    echo "1) install kvrocks"
    echo "2) enable local storage"
    echo "3) set workplace path"
    echo "4) register service to systemd"
    echo "5) start service"
#    echo "5) storage_service"
#    echo "6) msg"
    echo "q) quit"
}

func_local_config ()
{
    read -p "enter enable local storage (yes/no):" localStorage
}

func_workplace_config ()
{
    read -p "enter workplace path:" dictionary
}

func_register_service ()
{
    read -p "confirm register service to systemd (yes/no):" register
}
func_start_service ()
{
    read -p "start service (yes/no):" start
}

func_kvrocks_path ()
{
    read -p "input kvrocks path:" kvpath
}

func_kvrocks_db_path ()
{
    read -p "input kvrocks db path:" kvdbpath
}

while true; do
    display_menu
    read -s -n 1 key
    case $key in
      1)
        echo "setting kvrocks running path:"
        func_kvrocks_path
        echo "you choose:" $kvpath
        if [ -d $kvpath ]; then
          echo "db path is exist"
        else
          mkdir -p $kvpath
        fi
        cd $kvpath; wget "https://github.com/web3password/kvrocks/releases/download/v2.6.0/kvrocks.tar.gz"
        tar -xvf $kvpath/kvrocks.tar.gz
        echo "setting kvrocks db path:"
        func_kvrocks_db_path
        if [ -d $kvdbpath ]; then
          echo "db path is exist"
        else
          mkdir -p $kvdbpath
        fi
        sed -i '/^dir /!b;c\dir\ \'"${kvdbpath}"  $kvpath/kvrocks.conf
        c=" -c "
        bin="/bin/kvrocks"
        conf="/kvrocks.conf"
        startCommand=$kvpath$bin$c$kvpath$conf
        sed -i '/^ExecStart=/!b;c\ExecStart=\ \'"${startCommand}"  $w3p_kvrocks_service_file
        cp $w3p_kvrocks_service_file /etc/systemd/system/
        systemctl daemon-reload
        systemctl start kvrocks.service
        systemctl enable kvrocks.service
        echo "kvrocks running success. checking details by systemctl!"
      2)
        echo "choose local storage:"
        func_local_config
        echo "you choose:" $localStorage
        result=$([ "$localStorage" == yes ] && echo -n "local" || echo -n "audit")
        sed -i '/^running_mode:/!b;c\running_mode:\ \'$result''  $web3_satis_temp_conf_file
        ;;
      3)
        echo "setting workplace:"
        func_workplace_config
        echo "your workplace:" $dictionary
        if [ -d $dictionary ]; then
          echo "workplace is exist"
         else
          mkdir -p $dictionary
        fi
        rm -rf $ET_HOME/conf
        cp -r $ET_HOME/conf-template $ET_HOME/conf
        c=" --conf "
        bin="/bin/satis"
        yaml="/conf/satis.yaml"
        startCommand=$dictionary$bin$c$dictionary$yaml
#        echo $startCommand
        sed -i '/^ExecStart=/!b;c\ExecStart=\ \'"${startCommand}"  $w3p_satis_service_file
        bin="/bin/ares"
        yaml="/conf/ares.yaml"
        startCommand=$dictionary$bin$c$dictionary$yaml
#        echo $startCommand
        sed -i '/^ExecStart=/!b;c\ExecStart=\ \'"${startCommand}"  $w3p_ares_service_file
        bin="/bin/storage"
        yaml="/conf/storage.yaml"
        startCommand=$dictionary$bin$c$dictionary$yaml
#        echo $startCommand
        sed -i '/^ExecStart=/!b;c\ExecStart=\ \'"${startCommand}"  $w3p_storage_service_file
        bin="/bin/index"
        yaml="/conf/index.yaml"
        startCommand=$dictionary$bin$c$dictionary$yaml
#        echo $startCommand
        sed -i '/^ExecStart=/!b;c\ExecStart=\ \'"${startCommand}"  $w3p_index_service_file
#        set log dir
        log=$(echo -n "/logs/")
        logDir=$dictionary$log
        sed -i '/^log:/!b;c\log:\ \'"${logDir}"  $web3_satis_conf_file
        sed -i '/^log_dir:/!b;c\log_dir:\ \'"${logDir}"  $web3_ares_conf_file
        sed -i '/^log_dir:/!b;c\log_dir:\ \'"${logDir}"  $web3_storage_conf_file
        sed -i '/^  file_path:/!b;c\  file_path:\ \'"${logDir}"  $web3_index_conf_file
#        set storage dir path
        store=$(echo -n "/files/")
        storeDir=$dictionary$store
        sed -i '/^file_path:/!b;c\file_path:\ \'"${storeDir}"  $web3_storage_conf_file
        cp -r $ET_HOME/bin $ET_HOME/conf $dictionary
        ;;
      4)
        echo "register service to systemd:"
        func_register_service
        if [ $register  == "no" ]; then
          echo "register cancel"
        else
          echo "register begin"
          cp $w3p_satis_service_file /etc/systemd/system/
          cp $w3p_ares_service_file /etc/systemd/system/
          cp $w3p_storage_service_file /etc/systemd/system/
          cp $w3p_index_service_file /etc/systemd/system/
          systemctl daemon-reload
          echo "register done"
        fi
        ;;
      4)
        func_start_service
        if [ $start  == "no" ]; then
          echo "start cancel"
        else
          echo "begin to start service...."
          systemctl start w3p-satis.service
          systemctl start w3p-ares.service
          systemctl start w3p-storage.service
          systemctl start w3p-index.service
          echo "start service success! you can manage service by systemctl!"
        fi
        ;;
      5)
        log=$(echo -n "/logs/")
        echo $log
        ;;
      q|Q)
        echo "Exiting..."
        exit 0
        ;;
      *)
        echo "Invalid input. Please choose a valid option."
        ;;
    esac
#    echo "done"
done
