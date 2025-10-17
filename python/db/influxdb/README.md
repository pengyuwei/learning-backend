# InfluxDB 示例项目

这个项目演示了如何使用 InfluxDB 进行时序数据的存储和查询。

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
