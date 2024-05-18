package template

var ConfigTemplate = `
gc_no = {{.GCNo}}
server_type = {{.ServerType}}
channel_no = {{.ChannelNo}}
channel_name = {{.ChannelName}}
ip = {{.IP}}
tcp_port = {{.TCPPort}}
udp_port = {{.UDPPort}}
udp_ip_of_monitor = {{.UDPIPOfMonitor}}
udp_port_of_monitor = {{.UDPPortOfMonitor}}
tcp_port_of_monitor = {{.TCPPortOfMonitor}}
udp_ip_of_doublecheck = {{.UDPIPOfDoublecheck}}
udp_port_of_doublecheck = {{.UDPPortOfDoublecheck}}
udp_ip_of_statistic = {{.UDPIPOfStatistic}}
udp_port_of_statistic = {{.UDPPortOfStatistic}}
udp_ip_of_guild = {{.UDPIPOfGuild}}
udp_port_of_guild = {{.UDPPortOfGuild}}
tcp_port_of_guild = {{.TCPPortOfGuild}}
udp_ip_of_channel = {{.UDPIPOfChannel}}
udp_port_of_channel = {{.UDPPortOfChannel}}
tcp_port_of_channel = {{.TCPPortOfChannel}}
exchange_server_ip = {{.ExchangeServerIP}}
exchange_server_port = {{.ExchangeServerPort}}
udp_ip_of_hades = {{.UDPIPOfHades}}
udp_port_of_hades = {{.UDPPortOfHades}}
tcp_port_of_hades = {{.TCPPortOfHades}}
auction_server_ip = {{.AuctionServerIP}}
auction_server_port = {{.AuctionServerPort}}
cera_auction_server_ip = {{.CeraAuctionServerIP}}
cera_auction_server_port = {{.CeraAuctionServerPort}}
ipg_ip = {{.IPGIP}}
nxj_ipg_ip = {{.NXJIPGIP}}
nxj_ipg_port = {{.NXJIPGPort}}
relay_ip = {{.RelayIP}}
relay_tcp_port = {{.RelayTCPPort}}
relay_udp_port = {{.RelayUDPPort}}
stun_ip = {{.StunIP}}
stun_port = {{.StunPort1}}
stun_port = {{.StunPort2}}
stun_port = {{.StunPort3}}
lls_keys = {{.LLSKeys}}
//chatting_server_ip = {{.ChattingServerIP}}
//chatting_server_port = {{.ChattingServerPort}}
//social_event_server_ip = {{.SocialEventServerIP}}
//social_event_server_tcp_port = {{.SocialEventServerTCPPort}}
//mobile_server_ip = {{.MobileServerIP}}
//mobile_server_port = {{.MobileServerPort}}
community_server_ip = {{.CommunityServerIP}}
community_server_port = {{.CommunityServerPort}}
//nexon_billing_ip = {{.NexonBillingIP}}
//nexon_billing_port = {{.NexonBillingPort}}
db_thread_count = {{.DBThreadCount}}
master_db_ip = {{.MasterDBIP}}
master_db_port = {{.MasterDBPort}}
master_db_acc = {{.MasterDBAcc}}
master_db_pwd = {{.MasterDBPwd}}
master_db_name = {{.MasterDBName}}
server_group = {{.ServerGroup}}
fatigue_time = {{.FatigueTime}}
avatar_time = {{.AvatarTime}}
max_user_num = {{.MaxUserNum}}
db_tbl_file = {{.DBTblFile}}
header_classification = {{.HeaderClassification}}
header_msg_no = {{.HeaderMsgNo}}
header_sLength = {{.HeaderSLength}}
//header_garvage = {{.HeaderGarvage}}
header_checksum = {{.HeaderChecksum}}
header_sequence = {{.HeaderSequence}}
header_load_interval = {{.HeaderLoadInterval}}
header_mangled_length = {{.HeaderMangledLength}}
libserverlib_config_name = {{.LibserverlibConfigName}}
script_dir = {{.ScriptDir}}
//stand_alone = {{.StandAlone}}
scriptpacks = {{.ScriptPacks}}
process_sequence = {{.ProcessSequence}}
`

type ChannelConfig struct {
	GCNo                     int
	ServerType               int
	ChannelNo                int
	ChannelName              string
	IP                       string
	TCPPort                  int
	UDPPort                  int
	UDPIPOfMonitor           string
	UDPPortOfMonitor         int
	TCPPortOfMonitor         int
	UDPIPOfDoublecheck       string
	UDPPortOfDoublecheck     int
	UDPIPOfStatistic         string
	UDPPortOfStatistic       int
	UDPIPOfGuild             string
	UDPPortOfGuild           int
	TCPPortOfGuild           int
	UDPIPOfChannel           string
	UDPPortOfChannel         int
	TCPPortOfChannel         int
	ExchangeServerIP         string
	ExchangeServerPort       int
	UDPIPOfHades             string
	UDPPortOfHades           int
	TCPPortOfHades           int
	AuctionServerIP          string
	AuctionServerPort        int
	CeraAuctionServerIP      string
	CeraAuctionServerPort    int
	IPGIP                    string
	NXJIPGIP                 string
	NXJIPGPort               int
	RelayIP                  string
	RelayTCPPort             int
	RelayUDPPort             int
	StunIP                   string
	StunPort1                int
	StunPort2                int
	StunPort3                int
	LLSKeys                  string
	ChattingServerIP         string
	ChattingServerPort       int
	SocialEventServerIP      string
	SocialEventServerTCPPort int
	MobileServerIP           string
	MobileServerPort         int
	CommunityServerIP        string
	CommunityServerPort      int
	NexonBillingIP           string
	NexonBillingPort         int
	DBThreadCount            int
	MasterDBIP               string
	MasterDBPort             int
	MasterDBAcc              string
	MasterDBPwd              string
	MasterDBName             string
	ServerGroup              int
	FatigueTime              int
	AvatarTime               int
	MaxUserNum               int
	DBTblFile                string
	HeaderClassification     string
	HeaderMsgNo              string
	HeaderSLength            string
	HeaderGarvage            string
	HeaderChecksum           string
	HeaderSequence           string
	HeaderLoadInterval       int
	HeaderMangledLength      int
	LibserverlibConfigName   string
	ScriptDir                string
	StandAlone               string
	ScriptPacks              string
	ProcessSequence          int
}
