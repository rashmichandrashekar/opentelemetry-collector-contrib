// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/filter"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
)

// AttributeClass specifies the a value class attribute.
type AttributeClass int

const (
	_ AttributeClass = iota
	AttributeClassDatapath
	AttributeClassServices
)

// String returns the string representation of the AttributeClass.
func (av AttributeClass) String() string {
	switch av {
	case AttributeClassDatapath:
		return "datapath"
	case AttributeClassServices:
		return "services"
	}
	return ""
}

// MapAttributeClass is a helper map of string to AttributeClass attribute value.
var MapAttributeClass = map[string]AttributeClass{
	"datapath": AttributeClassDatapath,
	"services": AttributeClassServices,
}

// AttributeDirection specifies the a value direction attribute.
type AttributeDirection int

const (
	_ AttributeDirection = iota
	AttributeDirectionReceived
	AttributeDirectionTransmitted
)

// String returns the string representation of the AttributeDirection.
func (av AttributeDirection) String() string {
	switch av {
	case AttributeDirectionReceived:
		return "received"
	case AttributeDirectionTransmitted:
		return "transmitted"
	}
	return ""
}

// MapAttributeDirection is a helper map of string to AttributeDirection attribute value.
var MapAttributeDirection = map[string]AttributeDirection{
	"received":    AttributeDirectionReceived,
	"transmitted": AttributeDirectionTransmitted,
}

// AttributeDiskState specifies the a value disk_state attribute.
type AttributeDiskState int

const (
	_ AttributeDiskState = iota
	AttributeDiskStateUsed
	AttributeDiskStateAvailable
)

// String returns the string representation of the AttributeDiskState.
func (av AttributeDiskState) String() string {
	switch av {
	case AttributeDiskStateUsed:
		return "used"
	case AttributeDiskStateAvailable:
		return "available"
	}
	return ""
}

// MapAttributeDiskState is a helper map of string to AttributeDiskState attribute value.
var MapAttributeDiskState = map[string]AttributeDiskState{
	"used":      AttributeDiskStateUsed,
	"available": AttributeDiskStateAvailable,
}

// AttributePacketType specifies the a value packet.type attribute.
type AttributePacketType int

const (
	_ AttributePacketType = iota
	AttributePacketTypeDropped
	AttributePacketTypeErrored
	AttributePacketTypeSuccess
)

// String returns the string representation of the AttributePacketType.
func (av AttributePacketType) String() string {
	switch av {
	case AttributePacketTypeDropped:
		return "dropped"
	case AttributePacketTypeErrored:
		return "errored"
	case AttributePacketTypeSuccess:
		return "success"
	}
	return ""
}

// MapAttributePacketType is a helper map of string to AttributePacketType attribute value.
var MapAttributePacketType = map[string]AttributePacketType{
	"dropped": AttributePacketTypeDropped,
	"errored": AttributePacketTypeErrored,
	"success": AttributePacketTypeSuccess,
}

type metricNsxtNodeCPUUtilization struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills nsxt.node.cpu.utilization metric with initial data.
func (m *metricNsxtNodeCPUUtilization) init() {
	m.data.SetName("nsxt.node.cpu.utilization")
	m.data.SetDescription("The average amount of CPU being used by the node.")
	m.data.SetUnit("%")
	m.data.SetEmptyGauge()
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricNsxtNodeCPUUtilization) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64, classAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
	dp.Attributes().PutStr("class", classAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricNsxtNodeCPUUtilization) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricNsxtNodeCPUUtilization) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricNsxtNodeCPUUtilization(cfg MetricConfig) metricNsxtNodeCPUUtilization {
	m := metricNsxtNodeCPUUtilization{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricNsxtNodeFilesystemUsage struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills nsxt.node.filesystem.usage metric with initial data.
func (m *metricNsxtNodeFilesystemUsage) init() {
	m.data.SetName("nsxt.node.filesystem.usage")
	m.data.SetDescription("The amount of storage space used by the node.")
	m.data.SetUnit("By")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricNsxtNodeFilesystemUsage) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, diskStateAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("state", diskStateAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricNsxtNodeFilesystemUsage) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricNsxtNodeFilesystemUsage) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricNsxtNodeFilesystemUsage(cfg MetricConfig) metricNsxtNodeFilesystemUsage {
	m := metricNsxtNodeFilesystemUsage{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricNsxtNodeFilesystemUtilization struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills nsxt.node.filesystem.utilization metric with initial data.
func (m *metricNsxtNodeFilesystemUtilization) init() {
	m.data.SetName("nsxt.node.filesystem.utilization")
	m.data.SetDescription("The percentage of storage space utilized.")
	m.data.SetUnit("%")
	m.data.SetEmptyGauge()
}

func (m *metricNsxtNodeFilesystemUtilization) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricNsxtNodeFilesystemUtilization) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricNsxtNodeFilesystemUtilization) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricNsxtNodeFilesystemUtilization(cfg MetricConfig) metricNsxtNodeFilesystemUtilization {
	m := metricNsxtNodeFilesystemUtilization{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricNsxtNodeMemoryCacheUsage struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills nsxt.node.memory.cache.usage metric with initial data.
func (m *metricNsxtNodeMemoryCacheUsage) init() {
	m.data.SetName("nsxt.node.memory.cache.usage")
	m.data.SetDescription("The size of the node's memory cache.")
	m.data.SetUnit("KBy")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricNsxtNodeMemoryCacheUsage) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricNsxtNodeMemoryCacheUsage) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricNsxtNodeMemoryCacheUsage) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricNsxtNodeMemoryCacheUsage(cfg MetricConfig) metricNsxtNodeMemoryCacheUsage {
	m := metricNsxtNodeMemoryCacheUsage{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricNsxtNodeMemoryUsage struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills nsxt.node.memory.usage metric with initial data.
func (m *metricNsxtNodeMemoryUsage) init() {
	m.data.SetName("nsxt.node.memory.usage")
	m.data.SetDescription("The memory usage of the node.")
	m.data.SetUnit("KBy")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricNsxtNodeMemoryUsage) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricNsxtNodeMemoryUsage) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricNsxtNodeMemoryUsage) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricNsxtNodeMemoryUsage(cfg MetricConfig) metricNsxtNodeMemoryUsage {
	m := metricNsxtNodeMemoryUsage{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricNsxtNodeNetworkIo struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills nsxt.node.network.io metric with initial data.
func (m *metricNsxtNodeNetworkIo) init() {
	m.data.SetName("nsxt.node.network.io")
	m.data.SetDescription("The number of bytes which have flowed through the network interface.")
	m.data.SetUnit("By")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricNsxtNodeNetworkIo) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, directionAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("direction", directionAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricNsxtNodeNetworkIo) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricNsxtNodeNetworkIo) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricNsxtNodeNetworkIo(cfg MetricConfig) metricNsxtNodeNetworkIo {
	m := metricNsxtNodeNetworkIo{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricNsxtNodeNetworkPacketCount struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills nsxt.node.network.packet.count metric with initial data.
func (m *metricNsxtNodeNetworkPacketCount) init() {
	m.data.SetName("nsxt.node.network.packet.count")
	m.data.SetDescription("The number of packets which have flowed through the network interface on the node.")
	m.data.SetUnit("{packets}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricNsxtNodeNetworkPacketCount) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, directionAttributeValue string, packetTypeAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("direction", directionAttributeValue)
	dp.Attributes().PutStr("type", packetTypeAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricNsxtNodeNetworkPacketCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricNsxtNodeNetworkPacketCount) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricNsxtNodeNetworkPacketCount(cfg MetricConfig) metricNsxtNodeNetworkPacketCount {
	m := metricNsxtNodeNetworkPacketCount{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user config.
type MetricsBuilder struct {
	config                              MetricsBuilderConfig // config of the metrics builder.
	startTime                           pcommon.Timestamp    // start time that will be applied to all recorded data points.
	metricsCapacity                     int                  // maximum observed number of metrics per resource.
	metricsBuffer                       pmetric.Metrics      // accumulates metrics data before emitting.
	buildInfo                           component.BuildInfo  // contains version information.
	resourceAttributeIncludeFilter      map[string]filter.Filter
	resourceAttributeExcludeFilter      map[string]filter.Filter
	metricNsxtNodeCPUUtilization        metricNsxtNodeCPUUtilization
	metricNsxtNodeFilesystemUsage       metricNsxtNodeFilesystemUsage
	metricNsxtNodeFilesystemUtilization metricNsxtNodeFilesystemUtilization
	metricNsxtNodeMemoryCacheUsage      metricNsxtNodeMemoryCacheUsage
	metricNsxtNodeMemoryUsage           metricNsxtNodeMemoryUsage
	metricNsxtNodeNetworkIo             metricNsxtNodeNetworkIo
	metricNsxtNodeNetworkPacketCount    metricNsxtNodeNetworkPacketCount
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(mbc MetricsBuilderConfig, settings receiver.CreateSettings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		config:                              mbc,
		startTime:                           pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                       pmetric.NewMetrics(),
		buildInfo:                           settings.BuildInfo,
		metricNsxtNodeCPUUtilization:        newMetricNsxtNodeCPUUtilization(mbc.Metrics.NsxtNodeCPUUtilization),
		metricNsxtNodeFilesystemUsage:       newMetricNsxtNodeFilesystemUsage(mbc.Metrics.NsxtNodeFilesystemUsage),
		metricNsxtNodeFilesystemUtilization: newMetricNsxtNodeFilesystemUtilization(mbc.Metrics.NsxtNodeFilesystemUtilization),
		metricNsxtNodeMemoryCacheUsage:      newMetricNsxtNodeMemoryCacheUsage(mbc.Metrics.NsxtNodeMemoryCacheUsage),
		metricNsxtNodeMemoryUsage:           newMetricNsxtNodeMemoryUsage(mbc.Metrics.NsxtNodeMemoryUsage),
		metricNsxtNodeNetworkIo:             newMetricNsxtNodeNetworkIo(mbc.Metrics.NsxtNodeNetworkIo),
		metricNsxtNodeNetworkPacketCount:    newMetricNsxtNodeNetworkPacketCount(mbc.Metrics.NsxtNodeNetworkPacketCount),
		resourceAttributeIncludeFilter:      make(map[string]filter.Filter),
		resourceAttributeExcludeFilter:      make(map[string]filter.Filter),
	}
	if mbc.ResourceAttributes.DeviceID.MetricsInclude != nil {
		mb.resourceAttributeIncludeFilter["device.id"] = filter.CreateFilter(mbc.ResourceAttributes.DeviceID.MetricsInclude)
	}
	if mbc.ResourceAttributes.DeviceID.MetricsExclude != nil {
		mb.resourceAttributeExcludeFilter["device.id"] = filter.CreateFilter(mbc.ResourceAttributes.DeviceID.MetricsExclude)
	}
	if mbc.ResourceAttributes.NsxtNodeID.MetricsInclude != nil {
		mb.resourceAttributeIncludeFilter["nsxt.node.id"] = filter.CreateFilter(mbc.ResourceAttributes.NsxtNodeID.MetricsInclude)
	}
	if mbc.ResourceAttributes.NsxtNodeID.MetricsExclude != nil {
		mb.resourceAttributeExcludeFilter["nsxt.node.id"] = filter.CreateFilter(mbc.ResourceAttributes.NsxtNodeID.MetricsExclude)
	}
	if mbc.ResourceAttributes.NsxtNodeName.MetricsInclude != nil {
		mb.resourceAttributeIncludeFilter["nsxt.node.name"] = filter.CreateFilter(mbc.ResourceAttributes.NsxtNodeName.MetricsInclude)
	}
	if mbc.ResourceAttributes.NsxtNodeName.MetricsExclude != nil {
		mb.resourceAttributeExcludeFilter["nsxt.node.name"] = filter.CreateFilter(mbc.ResourceAttributes.NsxtNodeName.MetricsExclude)
	}
	if mbc.ResourceAttributes.NsxtNodeType.MetricsInclude != nil {
		mb.resourceAttributeIncludeFilter["nsxt.node.type"] = filter.CreateFilter(mbc.ResourceAttributes.NsxtNodeType.MetricsInclude)
	}
	if mbc.ResourceAttributes.NsxtNodeType.MetricsExclude != nil {
		mb.resourceAttributeExcludeFilter["nsxt.node.type"] = filter.CreateFilter(mbc.ResourceAttributes.NsxtNodeType.MetricsExclude)
	}

	for _, op := range options {
		op(mb)
	}
	return mb
}

// NewResourceBuilder returns a new resource builder that should be used to build a resource associated with for the emitted metrics.
func (mb *MetricsBuilder) NewResourceBuilder() *ResourceBuilder {
	return NewResourceBuilder(mb.config.ResourceAttributes)
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pmetric.ResourceMetrics) {
	if mb.metricsCapacity < rm.ScopeMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.ScopeMetrics().At(0).Metrics().Len()
	}
}

// ResourceMetricsOption applies changes to provided resource metrics.
type ResourceMetricsOption func(pmetric.ResourceMetrics)

// WithResource sets the provided resource on the emitted ResourceMetrics.
// It's recommended to use ResourceBuilder to create the resource.
func WithResource(res pcommon.Resource) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		res.CopyTo(rm.Resource())
	}
}

// WithStartTimeOverride overrides start time for all the resource metrics data points.
// This option should be only used if different start time has to be set on metrics coming from different resources.
func WithStartTimeOverride(start pcommon.Timestamp) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		var dps pmetric.NumberDataPointSlice
		metrics := rm.ScopeMetrics().At(0).Metrics()
		for i := 0; i < metrics.Len(); i++ {
			switch metrics.At(i).Type() {
			case pmetric.MetricTypeGauge:
				dps = metrics.At(i).Gauge().DataPoints()
			case pmetric.MetricTypeSum:
				dps = metrics.At(i).Sum().DataPoints()
			}
			for j := 0; j < dps.Len(); j++ {
				dps.At(j).SetStartTimestamp(start)
			}
		}
	}
}

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead.
// Resource attributes should be provided as ResourceMetricsOption arguments.
func (mb *MetricsBuilder) EmitForResource(rmo ...ResourceMetricsOption) {
	rm := pmetric.NewResourceMetrics()
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol/nsxtreceiver")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricNsxtNodeCPUUtilization.emit(ils.Metrics())
	mb.metricNsxtNodeFilesystemUsage.emit(ils.Metrics())
	mb.metricNsxtNodeFilesystemUtilization.emit(ils.Metrics())
	mb.metricNsxtNodeMemoryCacheUsage.emit(ils.Metrics())
	mb.metricNsxtNodeMemoryUsage.emit(ils.Metrics())
	mb.metricNsxtNodeNetworkIo.emit(ils.Metrics())
	mb.metricNsxtNodeNetworkPacketCount.emit(ils.Metrics())

	for _, op := range rmo {
		op(rm)
	}
	for attr, filter := range mb.resourceAttributeIncludeFilter {
		if val, ok := rm.Resource().Attributes().Get(attr); ok && !filter.Matches(val.AsString()) {
			return
		}
	}
	for attr, filter := range mb.resourceAttributeExcludeFilter {
		if val, ok := rm.Resource().Attributes().Get(attr); ok && filter.Matches(val.AsString()) {
			return
		}
	}

	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user config, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(rmo ...ResourceMetricsOption) pmetric.Metrics {
	mb.EmitForResource(rmo...)
	metrics := mb.metricsBuffer
	mb.metricsBuffer = pmetric.NewMetrics()
	return metrics
}

// RecordNsxtNodeCPUUtilizationDataPoint adds a data point to nsxt.node.cpu.utilization metric.
func (mb *MetricsBuilder) RecordNsxtNodeCPUUtilizationDataPoint(ts pcommon.Timestamp, val float64, classAttributeValue AttributeClass) {
	mb.metricNsxtNodeCPUUtilization.recordDataPoint(mb.startTime, ts, val, classAttributeValue.String())
}

// RecordNsxtNodeFilesystemUsageDataPoint adds a data point to nsxt.node.filesystem.usage metric.
func (mb *MetricsBuilder) RecordNsxtNodeFilesystemUsageDataPoint(ts pcommon.Timestamp, val int64, diskStateAttributeValue AttributeDiskState) {
	mb.metricNsxtNodeFilesystemUsage.recordDataPoint(mb.startTime, ts, val, diskStateAttributeValue.String())
}

// RecordNsxtNodeFilesystemUtilizationDataPoint adds a data point to nsxt.node.filesystem.utilization metric.
func (mb *MetricsBuilder) RecordNsxtNodeFilesystemUtilizationDataPoint(ts pcommon.Timestamp, val float64) {
	mb.metricNsxtNodeFilesystemUtilization.recordDataPoint(mb.startTime, ts, val)
}

// RecordNsxtNodeMemoryCacheUsageDataPoint adds a data point to nsxt.node.memory.cache.usage metric.
func (mb *MetricsBuilder) RecordNsxtNodeMemoryCacheUsageDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricNsxtNodeMemoryCacheUsage.recordDataPoint(mb.startTime, ts, val)
}

// RecordNsxtNodeMemoryUsageDataPoint adds a data point to nsxt.node.memory.usage metric.
func (mb *MetricsBuilder) RecordNsxtNodeMemoryUsageDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricNsxtNodeMemoryUsage.recordDataPoint(mb.startTime, ts, val)
}

// RecordNsxtNodeNetworkIoDataPoint adds a data point to nsxt.node.network.io metric.
func (mb *MetricsBuilder) RecordNsxtNodeNetworkIoDataPoint(ts pcommon.Timestamp, val int64, directionAttributeValue AttributeDirection) {
	mb.metricNsxtNodeNetworkIo.recordDataPoint(mb.startTime, ts, val, directionAttributeValue.String())
}

// RecordNsxtNodeNetworkPacketCountDataPoint adds a data point to nsxt.node.network.packet.count metric.
func (mb *MetricsBuilder) RecordNsxtNodeNetworkPacketCountDataPoint(ts pcommon.Timestamp, val int64, directionAttributeValue AttributeDirection, packetTypeAttributeValue AttributePacketType) {
	mb.metricNsxtNodeNetworkPacketCount.recordDataPoint(mb.startTime, ts, val, directionAttributeValue.String(), packetTypeAttributeValue.String())
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}
