package consts

type Category int32

const (
	TimeFormat = "2006-01-02 15:04:05"

	VeloGwAddPath                             = "/gateway/gatewayProvision"
	VeloCreateGwPoolPath                      = "/network/insertNetworkGatewayPool"
	VeloDeleteGwPoolPath                      = "/network/deleteNetworkGatewayPool"
	VeloInsertEnterprisePath                  = "/enterprise/insertEnterprise"
	VeloEnterpriseProxyGetPath                = "/network/getNetworkEnterpriseProxies"
	VeloEnterpriseProxyGetPoolPath            = "/enterpriseProxy/getEnterpriseProxyGatewayPools"
	VeloAssignEnterpriseToEnterpriseProxy     = "/network/assignEnterpriseToEnterpriseProxy"
	VeloUpdateEnterpriseProxyPath             = "/enterpriseProxy/updateEnterpriseProxy"
	VeloEnterpriseOfPartnerGetPath            = "/enterpriseProxy/getEnterpriseProxyEnterprises"
	VeloGwPoolGetPath                         = "/network/getNetworkGatewayPools"
	VeloGetEnterprisePoolPath                 = "/enterpriseProxy/getEnterpriseProxyGatewayPools"
	VeloAddGwToPoolPath                       = "/gatewayPool/addGatewayPoolGateway"
	VeloDelGwFromPoolPath                     = "/gatewayPool/removeGatewayPoolGateway"
	VeloEnterpriseMaximumSegmentGetPath       = "/enterprise/getEnterpriseMaximumSegments"
	VeloEnterpriseNetworkSegmentGetPath       = "/enterprise/getEnterpriseNetworkSegments"
	VeloInsertEnterpriseNetworkSegmentPath    = "/enterprise/insertEnterpriseNetworkSegment"
	VeloGwGetPath                             = "/network/getNetworkGateways"
	VeloUpdateGwAttrPath                      = "/gateway/updateGatewayAttributes"
	VeloInsertOrUpdateEnterpriseGwHandoffPath = "/enterprise/insertOrUpdateEnterpriseGatewayHandoff"
	VeloGetEnterpriseGwHandoffPath            = "/enterprise/getEnterpriseGatewayHandoff"
)

// 客户盒子相关操作path
const (
	VeloEdgeProvision               = "/edge/edgeProvision"
	VeloUpdateEdgeAdminPassword     = "/edge/updateEdgeAdminPassword"
	VeloGetEdgeConfigurationStack   = "/edge/getEdgeConfigurationStack"
	VeloSetEdgeHandOffGateways      = "/edge/setEdgeHandOffGateways"
	VeloDeleteEdge                  = "/edge/deleteEdge"
	VeloGetEdgeConfigurationModules = "/edge/getEdgeConfigurationModules"

	VeloGetNetworkEnterprises = "/network/getNetworkEnterprises"

	VeloUpdateConfigurationModule = "/configuration/updateConfigurationModule"
	VeloGetConfigurationModules   = "/configuration/getConfigurationModules"
	VeloInsertConfigurationModule = "/configuration/insertConfigurationModule"
	VeloCloneConfiguration        = "/configuration/cloneConfiguration"

	VeloGetEnterpriseEdges          = "/enterprise/getEnterpriseEdges"
	VeloGetEnterpriseConfigurations = "/enterprise/getEnterpriseConfigurations"

	VeloGetEnterpriseProxyEdgeInventory = "/enterpriseProxy/getEnterpriseProxyEdgeInventory"

	VeloGetEnterpriseEdgeLicenses = "/license/getEnterpriseEdgeLicenses"
)

const (
	VeloVlanID = 1
	// VeloGlobalSegmentID GlobalSegement id
	VeloGlobalSegmentID = 0

	VeloVlanLeaseTime15MinInSeconds = 900
	VeloVlanLeaseTime1HrInSeconds   = 3600
	VeloVlanLeaseTime4HrsInSeconds  = 14400
	VeloVlanLeaseTime12HrsInSeconds = 43200
	VeloVlanLeaseTime1DayInSeconds  = 86400
	VeloVlanLeaseTime1WeekInSeconds = 604800

	VeloEdgeStateConnected      = "CONNECTED"
	VeloEdgeStateDegraded       = "DEGRADED"
	VeloEdgeStateOffline        = "OFFLINE"
	VeloEdgeStateNeverActivated = "NEVER_ACTIVATED"

	VeloDefaultEnterpriseConfigurationName = "Quick Start Profile"

	VeloBoxVPNStraightThroughConfigHubSpk   = "HubSpk"
	VeloBoxVPNStraightThroughConfigFullMesh = "FullMesh"
)

// endpointPkiMode 字段, 默认CERTIFICATE_OPTIONAL
const (
	CertificateOptional = "CERTIFICATE_OPTIONAL"
	CertificateRequired = "CERTIFICATE_REQUIRED"
	CertificateDisabled = "CERTIFICATE_DISABLED"

	// EdgeSpecificProfile edge基于配置文件生成的configuration的名字
	EdgeSpecificProfile = "Edge Specific Profile"

	AccessRegion0 = "0"
	AccessRegion1 = "1"
	AccessRegion2 = "2"

	ModulesDeviceSetting = "deviceSettings"
	ModulesQOS           = "QOS"
	ModulesFirewall      = "firewall"

	WANTypeSTATIC = "STATIC"
	WANTypeDHCP   = "DHCP"
	WANTypePPPOE  = "PPPOE"

	DevCfgStatusNotConfigured = "10"
	DevCfgStatusConfigured    = "20"
	DevCfgStatusOnline        = "30"

	DevRunStatusActive   = "active"
	DevRunStatusInactive = "inactive"

	// DefaultBoxIP 设备默认的ip是 192.168.1.1， 如果不启用覆盖，这个字段 management 并不会被返回
	DefaultBoxIP = "192.168.1.1"

	BoxStatusInactive = "inactive"
	BoxStatusActive   = "active"

	BusiCodeBoxAdd    = "IO_AddBOXRes"
	BusiCodeBoxDelete = "IO_DelGWRes"
	BusiCodeBoxQuery  = "IO_QueryBOXRes"
	LANDhcpOn         = 1

	BoxAddFlag    = "add"
	BoxUpdateFlag = "update"

	RuleActionWhite = "white"
	RuleActionBlack = "black"

	RuleDirectionIn    = "in"
	RuleDirectionOut   = "out"
	RuleDirectionInout = "inout"

	ACLOutboundNegOne     = -1
	ACLOutboundDefaultMAC = "any"
	ACLOutboundPrefix     = "prefix"

	ACLAny         = "any"
	ACLDefaultMask = "255.255.255.255"

	StaticRouteValNegOne      = -1
	StaticRouteValZero        = 0
	StaticRouteValEmptyStr    = ""
	StaticRoutePreferredTrue  = true
	StaticRoutePreferredFalse = false
)

// 盒子更新subAction字段
const (
	BoxWANAdd         = "addWAN"
	BoxWANClose       = "closeWAN"
	BoxWANFullUpdate  = "updateWAN"
	BoxVPNSpeedUpdate = "updateSpeed"
	BoxVPNCfUpdate    = "updateVPNCf"
	BoxVPNACLOpen     = "openACL"
	BoxVPNACLUpdate   = "updateACL"
	BoxVPNACLClose    = "closeACL"
	BoxSRouteAdd      = "addSRoute"
	BoxSRouteRemove   = "rmSRoute"
	BoxNATAdd         = "addNAT"
	BoxNATClose       = "closeNAT"
)

const (
	// 接口返回结果
	IntfStatusFail    = 1
	IntfStatusSuccess = 0

	// 失败分类码
	ResCodeFormatErr      = "100"
	ResCodeBoxNotExist    = "101"
	ResCodeClientNotExist = "102"
	ResCodeRouteNotExist  = "103"
	ResCodeRepeatAdd      = "104"
)

// 告警级别
const (
	// CRITICAL 严重告警
	CRITICAL = "CRITICAL"
	// MAJOR 重大告警
	MAJOR = "MAJOR"
	// MINOR 次要告警
	MINOR = "MINOR"
	// WARNING 警告告警
	WARNING = "WARNING"
	// CLEARED 恢复告警
	CLEARED = "CLEARED"
)

// 联通的告警类别
const (
	// DevStatus 设备状态告警
	DevStatus = "devStatus"
	// IntfStatus 端口状态告警
	IntfStatus = "intfStatus"
	// TunnStatus 隧道状态告警
	TunnStatus = "tunnStatus"
	// BgpStatus BGP状态告警
	BgpStatus = "bgpStatus"
	// TestStatus 测试状态告警
	TestStatus = "testStatus"
)

// 告警关键值，DOWN或者UP
const (
	DOWN        = "DOWN"
	UP          = "UP"
	DevActive   = "active"
	DevInactive = "inactive"
)

// 设备状态
const (
	CLOSED  = "CLOSED"
	PENDING = "PENDING"
	ACTIVE  = "ACTIVE"
)

// 网关告警事件类型
const (
	GatewayBfdNeighborDown   = "GATEWAY_BFD_NEIGHBOR_DOWN"
	GatewayBgpNeighborDown   = "GATEWAY_BGP_NEIGHBOR_DOWN"
	GatewayNatServiceFailed  = "GATEWAY_NAT_SERVICE_FAILED"
	GatewayDown              = "GATEWAY_DOWN"
	GatewayServiceFailed     = "GATEWAY_SERVICE_FAILED"
	GatewayActivationFailure = "GATEWAY_ACTIVATION_FAILURE"
)

const (
	// VeloEnterpriseProxyEnterprisesURI 获取enterpriseProxy下所有的edge设备以及客户enterpriseID信息
	VeloEnterpriseProxyEnterprisesURI = "/enterpriseProxy/getEnterpriseProxyEnterprises"
	// VeloEdgeConfigurationModulesURI 获取edge设备对应的managementIP信息接口
	VeloEdgeConfigurationModulesURI = "/edge/getEdgeConfigurationModules"
)

const (
	// VeloEnterpriseAlertsURI 获取客户下触发的所有告警事件
	VeloEnterpriseAlertsURI = "/enterprise/getEnterpriseAlerts"
	// VeloGatewayAlertEventsURI 获取网关下一段时间的告警事件
	VeloGatewayAlertEventsURI = "/event/getOperatorEvents"
)

const (
	// VeloEdgeStatusMetricsURI 设备CPU，Mem,flow设备性能接口
	VeloEdgeStatusMetricsURI = "/metrics/getEdgeStatusMetrics"
	// VeloGatewayStatusMetricsURI 网关CPU，Mem,flow网关性能接口
	VeloGatewayStatusMetricsURI = "/metrics/getGatewayStatusMetrics"
	// VeloEdgeLinkMetricsURI 设备端口性能接口
	VeloEdgeLinkMetricsURI = "/metrics/getEdgeLinkMetrics"
	// VeloEdgeAppMetricsURI 设备应用性能接口
	VeloEdgeAppMetricsURI = "/metrics/getEdgeAppMetrics"
)

// 联通告警数据上报接口
const (
	CuccNmsAlertURI = "/sdwan/nms/alarm"
)

// 联通性能上报接口
const (
	// CuccDeviceStatusPerfURI 联通设备性能上报接口
	CuccDeviceStatusPerfURI = "/sdwan/nms/perf/devPerf"
	// CuccLinkStatusPerfURI 联通端口性能上报接口
	CuccLinkStatusPerfURI = "/sdwan/nms/perf/intfPerf"
	// CuccTunlStatusPerfURI 联通隧道性能上报接口
	CuccTunlStatusPerfURI = "/sdwan/nms/perf/tunnelPerf"
	// CuccDPIStatusPerfURI 联通应用识别上报接口
	CuccDPIStatusPerfURI = "/sdwan/nms/perf/dpiPerf"
)

// 联通接口BusiCode
const (
	CuccNmsDevPerfBusiCode     = "IO_nmsDevPerf"
	CuccNmsIntfPerfBusiCode    = "IO_nmsIntfPerf"
	CuccNmsTunlPerfBusiCode    = "IO_nmsTunlPerf"
	CuccNmsDPIPerfBusiCode     = "IO_nmsDPIPerf"
	CuccNmsDevStatPerfBusiCode = "IO_nmsDevStat"
	CuccNmsAlertBusiCode       = "IO_nmsAlert"
)

// 设备型号
const (
	Edge510 = "edge510"
	Edge520 = "edge520"
	Edge840 = "edge840"
	Gateway = "gateway"
)

// 应用类型category
const (
	Unclassified Category = iota
	AnonymizersProxies
	RealTimeAudioOrVideo
	Authentication
	BusinessApplication
	BusinessCollaboration
	Email
	FileSharing
	Gaming
	Infrastructure
	InternetInstantMessaging
	Management
	Media
	NetworkService
	Peer2Peer
	RemoteDesktop
	SocialNetworking
	StorageBackup
	TunnelingVPN
	Web
	SDWAN
	OtherTCPUDP
)
