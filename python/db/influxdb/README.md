# InfluxDB 示例项目

这个项目演示了如何使用 InfluxDB 进行时序数据的存储和查询。

## 一、InfluxDB 的核心优势

### 1. **专为时间序列优化**
- 数据模型天然支持 **时间戳 + 指标 + 标签（Tags） + 字段（Fields）**，非常适合存储传感器数据、系统指标、日志事件等。
- 高效压缩算法（如 Gorilla 压缩），**存储成本低**，可节省 80%+ 空间。
- 自动按时间分区（Shard），支持快速按时间范围查询。

### 2. **高性能写入与查询**
- 支持 **每秒百万级数据点写入**（在合理硬件下）。
- 查询引擎针对时间窗口、降采样、聚合等操作高度优化。
- InfluxDB 2.x 引入 **Flux 语言**，支持复杂流式计算（如窗口函数、连接、跨表计算）。

### 3. **内置完整可观测性栈（TICK → InfluxDB 2.x）**
- **不再依赖外部组件**：InfluxDB 2.x 集成了：
  - 数据采集（Telegraf）
  - 存储（InfluxDB）
  - 可视化（Web UI）
  - 告警（Alerting）
  - 任务调度（Tasks）
- 一站式解决监控与分析需求。

### 4. **强大的数据生命周期管理**
- **Retention Policy（1.x） / Bucket Retention（2.x）**：自动过期删除旧数据。
- **Downsampling（降采样）**：通过任务将高精度数据聚合为低精度长期存储（如 1 秒 → 1 小时平均值）。

### 5. **灵活的权限与安全控制**
- 基于 **Token 的细粒度权限**（InfluxDB 2.x）：
  - 可限制 Token 对特定 Bucket 的读/写权限。
  - 支持组织（Organization）和用户隔离。
- 支持 TLS、认证、审计日志（企业版）。

### 6. **云原生与多部署模式**
- 支持：
  - **本地部署**（OSS 开源版）
  - **Docker / Kubernetes**
  - **InfluxDB Cloud**（托管服务，全球多区域）
- 提供 **OSS（免费）** 和 **Enterprise / Cloud（付费）** 版本。

---

## 二、InfluxDB 的核心功能

### 1. **时间序列数据模型**
- **Measurement**：类似“表”，如 `cpu_usage`
- **Tags**：索引的键值对（用于快速过滤），如 `host=server01, region=us-west`
- **Fields**：实际数值，如 `usage_idle=95.2`
- **Timestamp**：纳秒级时间戳（可自定义）

> 示例数据行（Line Protocol）：
> ```text
> cpu,host=server01,region=us-west usage_idle=95.2,usage_user=3.1 1700000000000000000
> ```

### 2. **查询语言**
- **InfluxQL**（1.x）：类 SQL 语法，简单易用。
  ```sql
  SELECT mean("usage_idle") FROM "cpu" WHERE time > now() - 1h GROUP BY time(10m)
  ```
- **Flux**（2.x 默认）：函数式、管道式语言，功能强大。
  ```js
  from(bucket: "metrics")
    |> range(start: -1h)
    |> filter(fn: (r) => r._measurement == "cpu" and r._field == "usage_idle")
    |> aggregateWindow(every: 10m, fn: mean)
  ```

### 3. **可视化与仪表盘**
- 内置 Web UI 支持：
  - 拖拽式创建图表（折线图、柱状图、热力图等）
  - 自定义变量（如 `$host` 下拉筛选）
  - 仪表盘共享与导出

### 4. **告警系统**
- 定义 **Checks**（检查规则）：
  - 阈值告警（如 CPU > 90%）
  - 死机检测（no data for 5m）
- 配置 **Notification Endpoints**：
  - Email、Slack、PagerDuty、Webhook 等
- 告警状态自动跟踪（OK → Critical → OK）

### 5. **自动化任务（Tasks）**
- 定时执行 Flux 脚本，用于：
  - 数据降采样
  - 数据清洗
  - 跨源聚合
  - 自动备份（导出到 CSV/其他系统）

### 6. **数据集成能力**
- **Telegraf**：官方代理，支持 200+ 输入插件（系统指标、数据库、MQTT、Kafka 等）
- **API 支持**：
  - HTTP 写入（Line Protocol / JSON）
  - Flux 查询 API
  - Client SDK（Python、Go、Java、JavaScript 等）

### 7. **高可用与扩展（企业版）**
- 集群部署（Enterprise 版）
- 多副本、故障转移
- 水平扩展（分片）

---

## 三、典型应用场景

| 场景                | 说明 |
|---------------------|------|
| **基础设施监控**     | 采集服务器 CPU、内存、磁盘、网络指标，实时告警 |
| **IoT 设备监控**     | 百万级传感器数据写入，按设备 ID（Tag）快速查询 |
| **应用性能监控（APM）** | 记录请求延迟、错误率、吞吐量 |
| **金融行情分析**     | 高频交易数据存储与实时计算 |
| **日志事件分析**     | 结构化日志的时间序列化处理（需配合 Telegraf） |

## 项目结构

```
influxdb_example/
├── requirements.txt    # Python 依赖
├── config.py          # InfluxDB 配置
├── influxdb_demo.py   # 主要示例代码
└── README.md          # 项目说明
```

## 安装依赖

```bash
conda create -n dev python=3.12
conda activate dev
pip install -r requirements.txt
```

## 运行示例

```bash
python influxdb_demo.py
```

## 功能特性

### 1. 数据写入
- **传感器数据**: 温度、湿度监控
- **系统监控**: CPU、内存、磁盘使用率
- **标签支持**: 传感器ID、位置、主机名等

### 2. 数据查询
- **原始数据查询**: 获取最新的时序数据
- **聚合查询**: 计算平均值、最大值等统计信息
- **时间范围过滤**: 查询特定时间段的数据

### 3. 示例场景
- IoT 设备监控
- 系统性能监控
- 业务指标分析

## InfluxDB 配置

项目使用以下 InfluxDB 实例:
- **URL**: http://10.0.1.195:8086
- **组织**: myorg
- **存储桶**: mybucket
- **数据保留**: 1周

## 核心概念

### Point (数据点)
```python
Point("measurement_name")
    .tag("tag_key", "tag_value")      # 标签 (索引字段)
    .field("field_key", field_value)  # 字段 (实际数据)
    .time(timestamp)                  # 时间戳
```

### Flux 查询语言
```flux
from(bucket: "mybucket")
|> range(start: -1h)                 # 时间范围
|> filter(fn: (r) => r._measurement == "sensor_data")  # 过滤
|> group(columns: ["location"])      # 分组
|> mean()                           # 聚合
```

## 最佳实践

1. **合理设计标签**: 标签用于索引，不要包含高基数数据
2. **批量写入**: 使用批量写入提高性能
3. **数据保留策略**: 根据需求设置合适的数据保留时间
4. **查询优化**: 使用时间范围过滤减少查询数据量
