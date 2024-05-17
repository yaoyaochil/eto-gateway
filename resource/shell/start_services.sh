#!/bin/bash

start_service() {
    cd $1
    chmod 755 *
    rm -f $1/pid/*.pid
    rm -rf $1/log/*.*
    LD_PRELOAD=/lib/libhook.so ./$2 start &
}

start_service "/home/neople/stun" "df_stun_r"
start_service "/home/neople/monitor" "df_monitor_r mnt_siroco"
start_service "/home/neople/manager" "df_manager_r manager"
start_service "/home/neople/relay" "df_relay_r relay_200"
start_service "/home/neople/bridge" "df_bridge_r bridge"
start_service "/home/neople/channel" "df_channel_r channel"
start_service "/home/neople/dbmw_guild" "df_dbmw_r dbmw_gld_siroco"
start_service "/home/neople/dbmw_mnt" "df_dbmw_r dbmw_mnt_siroco"
start_service "/home/neople/dbmw_stat" "df_dbmw_r dbmw_stat_siroco"
start_service "/home/neople/auction" "df_auction_r ./cfg/auction_siroco.cfg"
start_service "/home/neople/point" "df_point_r ./cfg/point_siroco.cfg"
start_service "/home/neople/guild" "df_guild_r gld_siroco"
start_service "/home/neople/statics" "df_statics_r stat_siroco"
start_service "/home/neople/coserver" "df_coserver_r coserver"
start_service "/home/neople/community" "df_community_r community"
start_service "/home/neople/secsvr/gunnersvr" "gunnersvr -t30 -i1"
start_service "/home/neople/secsvr/zergsvr" "zergsvr -t30 -i1"
start_service "/home/neople/secsvr/zergsvr" "secagent"
