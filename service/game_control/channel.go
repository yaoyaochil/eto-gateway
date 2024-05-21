package game_control

import (
	"bufio"
	"fmt"
	"gateway/source/template"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	templateUtil "text/template"

	"github.com/gorilla/websocket"
)

type ChannelService struct{}

// CreateConfig 创建配置文件 siroco%d.cfg
func (cs *ChannelService) CreateConfig(config template.ChannelConfig, dir string) (err error) {
	filename := filepath.Join(dir, fmt.Sprintf("siroco%d.cfg", config.ChannelNo))
	tmpl, err := templateUtil.New("config").Parse(template.ConfigTemplate)
	if err != nil {
		return fmt.Errorf("error creating template: %v", err)
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	err = tmpl.Execute(writer, config)
	if err != nil {
		return fmt.Errorf("error executing template: %v", err)
	}
	writer.Flush()

	return nil
}

// GetPagedConfigs 获取分页配置 prefix: siroco pageSize: 10 page: 1
func (cs *ChannelService) GetPagedConfigs(dir string, prefix string, pageSize int, page int) ([]template.ChannelConfig, int64, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, 0, fmt.Errorf("error reading directory: %v", err)
	}

	var configs []template.ChannelConfig
	var total int
	for _, file := range files {
		if strings.HasPrefix(file.Name(), prefix) && strings.HasSuffix(file.Name(), ".cfg") {
			total++
			if total > (page-1)*pageSize && total <= page*pageSize {
				config, err := cs.ParseConfigFile(filepath.Join(dir, file.Name()))
				if err != nil {
					return nil, 0, err
				}
				configs = append(configs, config)
			}
		}
	}
	return configs, int64(total), nil
}

// 解析配置文件
func (cs *ChannelService) ParseConfigFile(filename string) (template.ChannelConfig, error) {
	file, err := os.Open(filename)
	if err != nil {
		return template.ChannelConfig{}, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	config := template.ChannelConfig{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//") || line == "" {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		switch key {
		case "gc_no":
			fmt.Sscanf(value, "%d", &config.GCNo)
		case "server_type":
			fmt.Sscanf(value, "%d", &config.ServerType)
		case "channel_no":
			fmt.Sscanf(value, "%d", &config.ChannelNo)
		case "channel_name":
			config.ChannelName = value
		case "ip":
			config.IP = value
		case "tcp_port":
			fmt.Sscanf(value, "%d", &config.TCPPort)
		case "udp_port":
			fmt.Sscanf(value, "%d", &config.UDPPort)
		case "udp_ip_of_monitor":
			config.UDPIPOfMonitor = value
		case "udp_port_of_monitor":
			fmt.Sscanf(value, "%d", &config.UDPPortOfMonitor)
		case "tcp_port_of_monitor":
			fmt.Sscanf(value, "%d", &config.TCPPortOfMonitor)
		case "udp_ip_of_doublecheck":
			config.UDPIPOfDoublecheck = value
		case "udp_port_of_doublecheck":
			fmt.Sscanf(value, "%d", &config.UDPPortOfDoublecheck)
		case "udp_ip_of_statistic":
			config.UDPIPOfStatistic = value
		case "udp_port_of_statistic":
			fmt.Sscanf(value, "%d", &config.UDPPortOfStatistic)
		case "udp_ip_of_guild":
			config.UDPIPOfGuild = value
		case "udp_port_of_guild":
			fmt.Sscanf(value, "%d", &config.UDPPortOfGuild)
		case "tcp_port_of_guild":
			fmt.Sscanf(value, "%d", &config.TCPPortOfGuild)
		case "udp_ip_of_channel":
			config.UDPIPOfChannel = value
		case "udp_port_of_channel":
			fmt.Sscanf(value, "%d", &config.UDPPortOfChannel)
		case "tcp_port_of_channel":
			fmt.Sscanf(value, "%d", &config.TCPPortOfChannel)
		case "exchange_server_ip":
			config.ExchangeServerIP = value
		case "exchange_server_port":
			fmt.Sscanf(value, "%d", &config.ExchangeServerPort)
		case "udp_ip_of_hades":
			config.UDPIPOfHades = value
		case "udp_port_of_hades":
			fmt.Sscanf(value, "%d", &config.UDPPortOfHades)
		case "tcp_port_of_hades":
			fmt.Sscanf(value, "%d", &config.TCPPortOfHades)
		case "auction_server_ip":
			config.AuctionServerIP = value
		case "auction_server_port":
			fmt.Sscanf(value, "%d", &config.AuctionServerPort)
		case "cera_auction_server_ip":
			config.CeraAuctionServerIP = value
		case "cera_auction_server_port":
			fmt.Sscanf(value, "%d", &config.CeraAuctionServerPort)
		case "ipg_ip":
			config.IPGIP = value
		case "nxj_ipg_ip":
			config.NXJIPGIP = value
		case "nxj_ipg_port":
			fmt.Sscanf(value, "%d", &config.NXJIPGPort)
		case "relay_ip":
			config.RelayIP = value
		case "relay_tcp_port":
			fmt.Sscanf(value, "%d", &config.RelayTCPPort)
		case "relay_udp_port":
			fmt.Sscanf(value, "%d", &config.RelayUDPPort)
		case "stun_ip":
			config.StunIP = value
		case "stun_port":
			// 解析多个 STUN 端口
			ports := strings.Fields(value)
			if len(ports) > 0 {
				fmt.Sscanf(ports[0], "%d", &config.StunPort1)
			}
			if len(ports) > 1 {
				fmt.Sscanf(ports[1], "%d", &config.StunPort2)
			}
			if len(ports) > 2 {
				fmt.Sscanf(ports[2], "%d", &config.StunPort3)
			}
		case "lls_keys":
			config.LLSKeys = value
		case "community_server_ip":
			config.CommunityServerIP = value
		case "community_server_port":
			fmt.Sscanf(value, "%d", &config.CommunityServerPort)
		case "db_thread_count":
			fmt.Sscanf(value, "%d", &config.DBThreadCount)
		case "master_db_ip":
			config.MasterDBIP = value
		case "master_db_port":
			fmt.Sscanf(value, "%d", &config.MasterDBPort)
		case "master_db_acc":
			config.MasterDBAcc = value
		case "master_db_pwd":
			config.MasterDBPwd = value
		case "master_db_name":
			config.MasterDBName = value
		case "server_group":
			fmt.Sscanf(value, "%d", &config.ServerGroup)
		case "fatigue_time":
			fmt.Sscanf(value, "%d", &config.FatigueTime)
		case "avatar_time":
			fmt.Sscanf(value, "%d", &config.AvatarTime)
		case "max_user_num":
			fmt.Sscanf(value, "%d", &config.MaxUserNum)
		case "db_tbl_file":
			config.DBTblFile = value
		case "header_classification":
			config.HeaderClassification = value
		case "header_msg_no":
			config.HeaderMsgNo = value
		case "header_sLength":
			config.HeaderSLength = value
		case "header_checksum":
			config.HeaderChecksum = value
		case "header_sequence":
			config.HeaderSequence = value
		case "header_load_interval":
			fmt.Sscanf(value, "%d", &config.HeaderLoadInterval)
		case "header_mangled_length":
			fmt.Sscanf(value, "%d", &config.HeaderMangledLength)
		case "libserverlib_config_name":
			config.LibserverlibConfigName = value
		case "script_dir":
			config.ScriptDir = value
		case "stand_alone":
			config.StandAlone = value
		case "scriptpacks":
			config.ScriptPacks = value
		case "process_sequence":
			fmt.Sscanf(value, "%d", &config.ProcessSequence)
		}
	}

	if err := scanner.Err(); err != nil {
		return template.ChannelConfig{}, fmt.Errorf("error reading file: %v", err)
	}

	return config, nil
}

// DeleteConfig 删除配置文件
func (cs *ChannelService) DeleteConfig(dir string, gcNo int) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("error reading directory: %v", err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".cfg") {
			config, err := cs.ParseConfigFile(filepath.Join(dir, file.Name()))
			if err != nil {
				return err
			}
			if config.GCNo == gcNo {
				return os.Remove(filepath.Join(dir, file.Name()))
			}
		}
	}

	return fmt.Errorf("config with gc_no %d not found", gcNo)
}

// UpdateConfig 更新配置文件
func (cs *ChannelService) UpdateConfig(dir string, gcNo int, newConfig template.ChannelConfig) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("error reading directory: %v", err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".cfg") {
			config, err := cs.ParseConfigFile(filepath.Join(dir, file.Name()))
			if err != nil {
				return err
			}
			if config.GCNo == gcNo {
				newConfig.ChannelNo = config.ChannelNo // 保留原来的 ChannelNo
				return cs.CreateConfig(newConfig, dir)
			}
		}
	}

	return fmt.Errorf("config with gc_no %d not found", gcNo)
}

// StartChannelServiceByGCNo 根据 GCNo 启动频道服务 pushd /home/neople/game && LD_PRELOAD=./frida.so ./df_game_r siroco11 start; popd
func (cs *ChannelService) StartChannelServiceByGCNo(gcNo int) error {
	cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("pushd /home/neople/game && LD_PRELOAD=./frida.so ./df_game_r siroco%d start; popd", gcNo))
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("error executing command: %v", err)
	}
	go func(cmd *exec.Cmd) {
		err := cmd.Wait()
		if err != nil {
			fmt.Println("Error waiting for command:", err)
		}
	}(cmd)
	return nil
}

// GetChannelStatusByGCNo 根据 GCNo 获取频道服务状态
func (cs *ChannelService) GetChannelStatusByGCNo(gcNo int) (bool, error) {
	cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("ps aux | grep df_game_r | grep siroco%d | grep -v grep", gcNo))
	output, err := cmd.Output()
	if err != nil {
		return false, nil
	}
	return len(output) > 0, nil
}

// StopChannelServiceByGCNo 根据 GCNo 停止频道服务
func (cs *ChannelService) StopChannelServiceByGCNo(gcNo int) error {
	cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("ps aux | grep df_game_r | grep siroco%d | grep -v grep | awk '{print $2}' | xargs kill -9", gcNo))
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("error executing command: %v", err)
	}
	go func(cmd *exec.Cmd) {
		err := cmd.Wait()
		if err != nil {
			fmt.Println("Error waiting for command:", err)
		}
	}(cmd)
	return nil
}

// RestartChannelServiceByGCNo 根据 GCNo 重启频道服务
func (cs *ChannelService) RestartChannelServiceByGCNo(gcNo int) error {
	err := cs.StopChannelServiceByGCNo(gcNo)
	if err != nil {
		return err
	}
	return cs.StartChannelServiceByGCNo(gcNo)
}

type ChannelStatus struct {
	GCNo      int
	IsRunning bool
}

// 批量查询频道状态
func (cs *ChannelService) GetChannelStatusByGCNos(gcNos []int) ([]ChannelStatus, error) {
	var statuses []ChannelStatus
	for _, gcNo := range gcNos {
		status, err := cs.GetChannelStatusByGCNo(gcNo)
		if err != nil {
			return []ChannelStatus{}, err
		}
		statuses = append(statuses, ChannelStatus{GCNo: gcNo, IsRunning: status})
	}
	return statuses, nil
}

// GetChannelRunLog 获取频道运行日志
func (cs *ChannelService) GetChannelRunLog(logFilePath string, conn *websocket.Conn) {
	// 使用tail -n +1 -f logFilePath 实时获取日志
	cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("tail -n +1 -f %s", logFilePath))

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Println("Error getting stdout pipe:", err)
		return
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Println("Error getting stderr pipe:", err)
		return
	}

	if err := cmd.Start(); err != nil {
		log.Println("Error starting command:", err)
		return
	}
	defer cmd.Process.Kill()

	stdoutReader := bufio.NewReader(stdout)
	stderrReader := bufio.NewReader(stderr)

	// 创建一个关闭信号通道
	done := make(chan struct{})

	// 启动一个 goroutine 监控连接状态
	go func() {
		for {
			_, _, err := conn.NextReader()
			if err != nil {
				log.Println("WebSocket connection closed")
				cmd.Process.Kill()
				return
			}
		}
	}()

	// 处理标准输出
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				line, err := stdoutReader.ReadString('\n')
				if err != nil {

					//

					return
				}
				if err := conn.WriteMessage(websocket.TextMessage, []byte(line)); err != nil {
					log.Println("Error writing message:", err)
					return
				}
			}
		}
	}()

	// 处理标准错误
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				line, err := stderrReader.ReadString('\n')
				if err != nil {
					// 错误就关闭连接
					conn.Close()
					return
				}
				if err := conn.WriteMessage(websocket.TextMessage, []byte(line)); err != nil {
					log.Println("Error writing message:", err)
					conn.Close()
					return
				}
			}
		}
	}()

	// 检查连接是否已关闭
	for {
		if _, _, err := conn.NextReader(); err != nil {
			log.Println("Connection closed by client")
			close(done) // 发送关闭信号
			cmd.Process.Kill()
			return
		}
	}
}
