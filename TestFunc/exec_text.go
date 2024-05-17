package main

import (
	"errors"
	"fmt"
	"gateway/global"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/songzhibin97/gkit/cache/local_cache"
)

type ProcessInfo struct {
	Name      string
	Command   []string
	IsRunning bool
}

func main() {
	global.GatewayCache = local_cache.NewCache()

	// 检查主服务是否已经启动
	// service_status, found := global.GatewayCache.Get("main_service")
	// if found {
	// 	if service_status == "running" {
	// 		fmt.Println("主服务已全部运行中")
	// 		return
	// 	}
	// }

	// runMainService()
	// if len(checkProcess()) >= 15 {
	// 	fmt.Println("主服务运行正常")
	// 	// 把主服务的状态写入缓存
	// 	global.GatewayCache.Set("main_service", "running", 0)
	// } else {
	// 	fmt.Println("主服务未全部运行")
	// }

	stopService()
}

func checkProcess() []string {
	cmd := `ps aux | grep -E "df_stun_r|df_monitor_r|df_manager_r|relay_200|df_bridge_r|df_channel_r|df_dbmw_r|df_auction_r|df_point_r|df_guild_r|df_statics_r|df_coserver_r|df_community_r|gunnersvr|secagent|zergsvr" | grep -v grep`
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return nil
	}
	var processNames []string
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) < 11 {
			fmt.Println("Unexpected line format:", line)
			continue
		}

		command := fields[10]
		if len(fields) > 11 {
			command = strings.Join(fields[10:], " ")
		}

		baseCommand := filepath.Base(command)
		processNames = append(processNames, baseCommand)
	}
	return processNames
}

func runMainService() (err error) {
	// 检查主服务是否已经启动
	if len(checkProcess()) >= 15 {
		fmt.Println("主服务已全部运行中")
		return nil
	}

	// 清理日志文件
	firstCommands := [][]string{
		{"find", "/home/neople/", "-name", "*.log", "-type", "f", "-print", "-exec", "rm", "-f", "{}", "+"},
		{"find", "/home/neople/", "-name", "*.pid", "-type", "f", "-print", "-exec", "rm", "-f", "{}", "+"},
		{"find", "/home/neople/", "-name", "core.*", "-type", "f", "-print", "-exec", "rm", "-f", "{}", "+"},
	}

	for _, command := range firstCommands {
		cmd := exec.Command(command[0], command[1:]...)
		err := cmd.Run()
		if err != nil {
			continue
		}
	}

	serviceList := []ProcessInfo{
		{
			Name:      "df_stun_r",
			Command:   []string{"/bin/bash", "-c", "pushd /home/neople/stun && chmod 755 * && ./df_stun_r start && popd"},
			IsRunning: false,
		},
		{
			Name:      "df_monitor_r",
			Command:   []string{"/bin/bash", "-c", "pushd /home/neople/monitor && chmod 755 * && LD_PRELOAD=/lib/libhook.so ./df_monitor_r mnt_siroco start && popd"},
			IsRunning: false,
		},
		{
			Name:      "df_manager_r",
			Command:   []string{"/bin/bash", "-c", "pushd /home/neople/manager && chmod 755 * && LD_PRELOAD=/lib/libhook.so ./df_manager_r manager start && popd"},
			IsRunning: false,
		},
		{
			Name:      "df_relay_r",
			Command:   []string{"/bin/bash", "-c", "pushd /home/neople/relay && chmod 755 * && ./df_relay_r relay_200 start && popd"},
			IsRunning: false,
		},
		{
			Name:      "df_bridge_r",
			Command:   []string{"/bin/bash", "-c", "pushd /home/neople/bridge && chmod 755 * && LD_PRELOAD=/lib/libhook.so ./df_bridge_r bridge start && popd"},
			IsRunning: false,
		},
		{
			Name:      "df_channel_r",
			Command:   []string{"/bin/bash", "-c", "pushd /home/neople/channel && chmod 755 * && LD_PRELOAD=/lib/libhook.so ./df_channel_r channel start && popd"},
			IsRunning: false,
		},
		{
			Name:      "df_dbmw_r",
			Command:   []string{"/bin/bash", "-c", "pushd /home/neople/dbmw_guild && chmod 755 * && LD_PRELOAD=/lib/libhook.so ./df_dbmw_r dbmw_gld_siroco start && popd"},
			IsRunning: false,
		},
		{
			Name:      "df_dbmw_r",
			Command:   []string{"/bin/bash", "-c", "pushd /home/neople/dbmw_mnt && chmod 755 * && LD_PRELOAD=/lib/libhook.so ./df_dbmw_r dbmw_mnt_siroco start && popd"},
			IsRunning: false,
		},
		{
			Name:      "df_dbmw_r",
			Command:   []string{"/bin/bash", "-c", "pushd /home/neople/dbmw_stat && chmod 755 * && LD_PRELOAD=/lib/libhook.so ./df_dbmw_r dbmw_stat_siroco start && popd"},
			IsRunning: false,
		},
		{
			Name:      "df_auction_r",
			Command:   []string{"/bin/bash", "-c", "pushd /home/neople/auction && chmod 755 * && ./df_auction_r ./cfg/auction_siroco.cfg start ./df_auction_r && popd"},
			IsRunning: false,
		},
		{
			Name:      "df_point_r",
			Command:   []string{"/bin/bash", "-c", "pushd /home/neople/point && chmod 755 * && ./df_point_r ./cfg/point_siroco.cfg start ./df_point_r && popd"},
			IsRunning: false,
		},
		{
			Name:      "df_guild_r",
			Command:   []string{"/bin/bash", "-c", "pushd /home/neople/guild && chmod 755 * && LD_PRELOAD=/lib/libhook.so ./df_guild_r gld_siroco start && popd"},
			IsRunning: false,
		},
		{
			Name:      "df_statics_r",
			Command:   []string{"/bin/bash", "-c", "pushd /home/neople/statics && chmod 755 * && LD_PRELOAD=/lib/libhook.so ./df_statics_r stat_siroco start && popd"},
			IsRunning: false,
		},
		{
			Name:      "df_coserver_r",
			Command:   []string{"/bin/bash", "-c", "pushd /home/neople/coserver && chmod 755 * && LD_PRELOAD=/lib/libhook.so ./df_coserver_r coserver start && popd"},
			IsRunning: false,
		},
		{
			Name:      "df_community_r",
			Command:   []string{"/bin/bash", "-c", "pushd /home/neople/community && chmod 755 * && ./df_community_r community start && popd"},
			IsRunning: false,
		},
		{
			Name:      "gunnersvr",
			Command:   []string{"/bin/bash", "-c", "pushd /home/neople/secsvr/gunnersvr && chmod 755 * && exec ./gunnersvr -t30 -i1 && popd"},
			IsRunning: false,
		},
		{
			Name:      "zergsvr",
			Command:   []string{"/bin/bash", "-c", "pushd /home/neople/secsvr/zergsvr && chmod 755 * && exec ./secagent && exec ./zergsvr -t30 -i1 && popd"},
			IsRunning: false,
		},
	}

	for _, service := range serviceList {
		if service.IsRunning {
			continue
		}

		cmd := exec.Command(service.Command[0], service.Command[1:]...)
		// 运行命令 并防止阻塞
		err := cmd.Start()
		if err != nil {
			fmt.Println("Error executing command1:", err)
			return err
		}
	}

	return nil
}

func stopService() (err error) {
	// 检查主服务是否已经启动
	if len(checkProcess()) == 0 {
		fmt.Println("主服务已全部停止")
		return nil
	}

	// 要执行的killall命令
	killallCommands := []string{
		"df_stun_r", "df_monitor_r", "df_manager_r", "df_relay_r", "df_bridge_r",
		"df_channel_r", "df_dbmw_r", "df_auction_r", "df_point_r", "df_guild_r",
		"df_statics_r", "df_coserver_r", "df_community_r", "gunnersvr", "zergsvr",
		"df_game_r", "secagent", "df_monitor_P_r", "df_guild_P_r", "df_statics_P_r",
	}

	for _, process := range killallCommands {
		cmd := exec.Command("killall", "-9", process)
		cmd.Run()
	}

	// 清理文件的find命令
	dnfDir := "/path/to/dnf_dir" // 替换为实际的DNF_DIR路径
	findCommands := [][]string{
		{"find", dnfDir, "-name", "*.log", "-type", "f", "-print", "-exec", "rm", "-rf", "{}", ";"},
		{"find", dnfDir, "-name", "*.pid", "-type", "f", "-print", "-exec", "rm", "-rf", "{}", ";"},
		{"find", dnfDir, "-name", "core.*", "-type", "f", "-print", "-exec", "rm", "-rf", "{}", ";"},
		{"find", dnfDir, "-name", "*.cri", "-type", "f", "-print", "-exec", "rm", "-rf", "{}", ";"},
		{"find", dnfDir, "-name", "*.debug", "-type", "f", "-print", "-exec", "rm", "-rf", "{}", ";"},
		{"find", dnfDir, "-name", "*.error", "-type", "f", "-print", "-exec", "rm", "-rf", "{}", ";"},
		{"find", dnfDir, "-name", "*.init", "-type", "f", "-print", "-exec", "rm", "-rf", "{}", ";"},
		{"find", dnfDir, "-name", "*.snap", "-type", "f", "-print", "-exec", "rm", "-rf", "{}", ";"},
	}

	for _, command := range findCommands {
		cmd := exec.Command(command[0], command[1:]...)
		cmd.Run()
	}

	// 检查主服务是否已经启动
	if len(checkProcess()) == 0 {
		fmt.Println("主服务已全部停止")
		return nil
	}
	return errors.New("主服务未全部停止")
}
