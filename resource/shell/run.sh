#!/bin/bash
#"************************************"
#"*           _ __   _______         *"
#"*          | '_ \ |__   __|        *"
#"*          | |_) |   | |           *"
#"*          |_,__/    |_|           *"
#"*                                  *"
#"************************************"
#"*   By:大唐一键端 | 大唐免费登陆器   *"
#"*          Q 群:289254980          *"
#"************************************"
cd /home/neople/stun
chmod 755 *
rm -f  /home/neople/stun/pid/*.pid
rm -rf /home/neople/stun/log/*.*
./df_stun_r start &

cd /home/neople/monitor
chmod 755 *
rm -f  /home/neople/monitor/pid/*.pid
rm -rf  /home/neople/monitor/log/*.*
LD_PRELOAD=/lib/libhook.so ./df_monitor_r mnt_siroco start &

cd /home/neople/manager
chmod 755 *
rm -f  /home/neople/manager/pid/*.pid
rm -rf  /home/neople/manager/log/*.*
LD_PRELOAD=/lib/libhook.so ./df_manager_r manager start &

cd /home/neople/relay
chmod 755 *
rm -f  /home/neople/relay/pid/*.pid
rm -rf  /home/neople/relay/log/*.*
./df_relay_r relay_200 start &

cd /home/neople/bridge
chmod 755 *
rm -f  /home/neople/bridge/pid/*.pid
rm -rf  /home/neople/bridge/log/*.*
LD_PRELOAD=/lib/libhook.so ./df_bridge_r bridge start &

cd /home/neople/channel
chmod 755 *
rm -f  /home/neople/channel/pid/*.pid
rm -rf  /home/neople/channel/log/*.*
LD_PRELOAD=/lib/libhook.so ./df_channel_r channel start &

cd /home/neople/dbmw_guild
chmod 755 *
rm -f  /home/neople/dbmw_guild/pid/*.pid
rm -rf  /home/neople/dbmw_guild/log/*.*
LD_PRELOAD=/lib/libhook.so ./df_dbmw_r dbmw_gld_siroco start &

cd /home/neople/dbmw_mnt
chmod 755 *
rm -f  /home/neople/dbmw_mnt/pid/*.pid
rm -rf  /home/neople/dbmw_mnt/log/*.*
LD_PRELOAD=/lib/libhook.so ./df_dbmw_r dbmw_mnt_siroco start &

cd /home/neople/dbmw_stat
chmod 755 *
rm -f  /home/neople/dbmw_stat/pid/*.pid
rm -rf  /home/neople/dbmw_stat/log/*.*
LD_PRELOAD=/lib/libhook.so ./df_dbmw_r dbmw_stat_siroco start &

cd /home/neople/auction
chmod 755 *
rm -f  /home/neople/auction/pid/*.pid
rm -rf  /home/neople/auction/log/*.*
./df_auction_r ./cfg/auction_siroco.cfg start ./df_auction_r &

cd /home/neople/point
chmod 755 *
rm -f  /home/neople/point/pid/*.pid
rm -rf  /home/neople/point/log/*.*
./df_point_r ./cfg/point_siroco.cfg start df_point_r &

cd /home/neople/guild
chmod 755 *
rm -f  /home/neople/guild/pid/*.pid
rm -rf  /home/neople/guild/log/*.*
LD_PRELOAD=/lib/libhook.so ./df_guild_r gld_siroco start &

cd /home/neople/statics
chmod 755 *
rm -f  /home/neople/statics/pid/*.pid
rm -rf  /home/neople/statics/log/*.*
LD_PRELOAD=/lib/libhook.so ./df_statics_r stat_siroco start &

cd /home/neople/coserver
chmod 755 *
rm -f  /home/neople/coserver/pid/*.pid
rm -rf  /home/neople/coserver/log/*.*
LD_PRELOAD=/lib/libhook.so ./df_coserver_r coserver start &

cd /home/neople/community
chmod 755 *
rm -f /home/neople/community/pid/*.pid
rm -rf /home/neople/community/log/*.*
./df_community_r community start &

cd /home/neople/secsvr/gunnersvr
chmod 755 *
rm -f /home/neople/secsvr/gunnersvr/*.pid
./gunnersvr -t30 -i1  &

cd /home/neople/secsvr/zergsvr
chmod 755 *
rm -f /home/neople/secsvr/zergsvr/*.pid

./secagent  &
./zergsvr -t30 -i1  &

cd /home/neople/game
chmod 755 *
rm -rf /home/neople/game/log/*
sleep 1
LD_PRELOAD=./frida.so ./df_game_r siroco11 start &
sleep 3
LD_PRELOAD=./frida.so ./df_game_r siroco56 start &