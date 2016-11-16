package config

type Config struct {
	// 应用Id
	appId string

	// 应用名称
	appName string

	// 应用Key，用于加密
	appKey string

	// ManageCenterAPI
	manageCenterAPI string

	// 合作商API
	partnerListAPI string

	// 服务器API
	serverListAPI string

	// 获取服务器组列表的API地址
	serverGroupListAPI string

	// 服务器组类型
	groupType string

	// ChatServerCenter监听的rpc地址
	chatServerCenterRpcAddress string

	// ChatServerCenter监听的web地址
	chatServerCenterWebAddress string

	// 获取玩家信息的API
	playerInfoAPI string

	// 消息最大长度
	maxMessageLength int

	// 最大的客户端数量
	maxClientCount int

	// 是否记录API日志
	ifRecordAPILog bool
}

func (config *Config) GetAppId() string {
	return config.appId
}

func (config *Config) GetAppName() string {
	return config.appName
}

func (config *Config) GetAppKey() string {
	return config.appKey
}

func (config *Config) GetManageCenterAPI() string {
	return config.manageCenterAPI
}

func (config *Config) GetPartnerListAPI() string {
	return config.partnerListAPI
}

func (config *Config) GetServerListAPI() string {
	return config.serverListAPI
}

func (config *Config) GetServerGroupListAPI() string {
	return config.serverGroupListAPI
}

func (config *Config) GetGroupType() string {
	return config.groupType
}

func (config *Config) GetChatServerCenterRpcAddress() string {
	return config.chatServerCenterRpcAddress
}

func (config *Config) GetChatServerCenterWebAddress() string {
	return config.chatServerCenterWebAddress
}

func (config *Config) GetPlayerInfoAPI() string {
	return config.playerInfoAPI
}

func (config *Config) GetMaxMessageLength() int {
	return config.maxMessageLength
}

func (config *Config) GetMaxClientCount() int {
	return config.maxClientCount
}

func (config *Config) GetIfRecordAPILog() bool {
	return config.ifRecordAPILog
}

func NewConfig(_appId, _appName, _appKey,
	_manageCenterAPI, _partnerListAPI, _serverListAPI, _serverGroupListAPI, _groupType,
	_chatServerCenterRpcAddress, _chatServerCenterWebAddress, _playerInfoAPI string,
	_maxMessageLength, _maxClientCount int, _ifRecordAPILog bool) *Config {
	return &Config{
		appId:                      _appId,
		appName:                    _appName,
		appKey:                     _appKey,
		manageCenterAPI:            _manageCenterAPI,
		partnerListAPI:             _partnerListAPI,
		serverListAPI:              _serverListAPI,
		serverGroupListAPI:         _serverGroupListAPI,
		groupType:                  _groupType,
		chatServerCenterRpcAddress: _chatServerCenterRpcAddress,
		chatServerCenterWebAddress: _chatServerCenterWebAddress,
		playerInfoAPI:              _playerInfoAPI,
		maxMessageLength:           _maxMessageLength,
		maxClientCount:             _maxClientCount,
		ifRecordAPILog:             _ifRecordAPILog,
	}
}
