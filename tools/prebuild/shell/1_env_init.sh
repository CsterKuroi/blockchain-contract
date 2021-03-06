#!/bin/bash

# --------------------------------------------------
#current_path=$(cd `dirname $0`; pwd)
current_path=`pwd`
#默认的要写入环境变量的文件
env_file=$HOME/.bashrc
#默认的go安装路径
go_root="/usr/local"
#默认的proto安装路径
proto_root="/opt/protoc"
#默认的第三方安装包的路径
dep_path=$current_path/tools/prebuild/dependent
#默认的go安装包名字
go_name="go1.8.linux-amd64.tar.gz"
#默认的proto安装包的名字
proto_name="protoc-3.2.0-linux-x86_64.zip"
#默认的云端（七牛云）的go安装包位置
download_go="http://ooekkyv04.bkt.clouddn.com/go1.8.linux-amd64.tar.gz"
#默认的云端（七牛云）的proto安装包位置
download_proto="http://ooekkyv04.bkt.clouddn.com/protoc-3.2.0-linux-x86_64.zip"
download_success=false
#默认的程序日志写入路径
log_path=$current_path/log
#默认的程序配置文件寻找路径
config_path=$current_path/conf
#没有安装过go时，默认的第一个GOPATH的路径
first_gopath=$HOME/go
# --------------------------------------------------

# --------------------------------------------------
# init go
echo "init go"
version=`go version 2>/dev/null`
if [ -z "$version" ]
then
    if [ ! -e $dep_path/$go_name ]
    then
        download_success=false 
        sudo apt-get install -y axel 1>/dev/null 2>&1
        axel $download_go -o $dep_path 1>/dev/null 2>&1 && download_success=true
    fi

    if [ ! $download_success ]
    then
        echo "download go failed"
        exit 1
    fi

    sudo tar -C $go_root -xzf $dep_path/$go_name
    echo export GOROOT=$go_root/go >> $env_file
    export GOROOT=$go_root/go
    echo export PATH=$PATH:$GOROOT/bin >> $env_file
    export PATH=$PATH:$GOROOT/bin
    if [ -z $GOPATH ]
    then
        echo export GOPATH=$first_gopath:$current_path >> $env_file
        export GOPATH=$first_gopath:$current_path
        echo export PATH=$PATH:$first_gopath/bin >> $env_file
        export PATH=$PATH:$first_gopath/bin
    else
        echo export GOPATH=$GOPATH:$current_path >> $env_file
        export GOPATH=$GOPATH:$current_path
    fi
    #. $env_file
else
    result=`echo $GOPATH | grep $current_path`
    if [ -z "$result" ]
    then
        tmp_GOPATH=$GOPATH
        sed -i -e '/export GOPATH=/d' $env_file
        echo export GOPATH=$tmp_GOPATH:$current_path >> $env_file
        export GOPATH=$tmp_GOPATH:$current_path
    fi
fi
# --------------------------------------------------

# --------------------------------------------------
#init proto
echo "init proto"
version=`protoc --version 2>/dev/null`
if [ -z "$version" ]
then
    if [ ! -d $proto_root ]
    then
        sudo mkdir $proto_root
    fi

    if [ ! -e $dep_path/$proto_name ]
    then
        download_success=false 
        sudo apt-get install -y axel 1>/dev/null 2>&1
        axel $download_proto -o $dep_path 1>/dev/null 2>&1 && download_success=true
    fi

    if [ ! $download_success ]
    then
        echo "download proto failed"
        exit 1
    fi

    sudo apt-get -y install zip 1>/dev/null 2>&1
    sudo unzip $dep_path/$proto_name -d $proto_root
    sudo chmod 755 $proto_root -R
    echo export PATH=$PATH:$proto_root/bin >> $env_file
    export PATH=$PATH:$proto_root/bin
    #. $env_file
fi
# --------------------------------------------------

# --------------------------------------------------
#set env var
echo "set env var"
if [ ! -d $log_path ]
then
    mkdir $log_path
fi

if [ -z "$LOGPATH" ]
then
    echo export LOGPATH=$log_path >> $env_file
    export LOGPATH=$log_path
fi

if [ -z "$CONFIGPATH" ]
then
    echo export CONFIGPATH=$config_path >> $env_file
    export CONFIGPATH=$config_path
fi
#. $env_file
# --------------------------------------------------

# --------------------------------------------------
echo "ok"
# --------------------------------------------------
