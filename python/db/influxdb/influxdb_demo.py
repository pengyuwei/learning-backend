#!/usr/bin/env python3
"""
InfluxDB 示例代码
演示如何连接、写入和查询时序数据
"""

import time
import random
from datetime import datetime, timezone
from influxdb_client import InfluxDBClient, Point, WritePrecision
from influxdb_client.client.write_api import SYNCHRONOUS
from config import INFLUXDB_URL, INFLUXDB_TOKEN, INFLUXDB_ORG, INFLUXDB_BUCKET


class InfluxDBDemo:
    def __init__(self):
        """初始化 InfluxDB 客户端"""
        self.client = InfluxDBClient(
            url=INFLUXDB_URL,
            token=INFLUXDB_TOKEN,
            org=INFLUXDB_ORG
        )
        self.write_api = self.client.write_api(write_options=SYNCHRONOUS)
        self.query_api = self.client.query_api()
        
    def write_sensor_data(self, sensor_id: str, temperature: float, humidity: float, location: str):
        """写入传感器数据"""
        point = Point("sensor_data") \
            .tag("sensor_id", sensor_id) \
            .tag("location", location) \
            .field("temperature", temperature) \
            .field("humidity", humidity) \
            .time(datetime.now(timezone.utc), WritePrecision.NS)
        
        self.write_api.write(bucket=INFLUXDB_BUCKET, org=INFLUXDB_ORG, record=point)
        print(f"✅ 写入数据: {sensor_id} - 温度: {temperature}°C, 湿度: {humidity}%")
    
    def write_system_metrics(self, host: str, cpu_usage: float, memory_usage: float, disk_usage: float):
        """写入系统监控数据"""
        point = Point("system_metrics") \
            .tag("host", host) \
            .field("cpu_usage", cpu_usage) \
            .field("memory_usage", memory_usage) \
            .field("disk_usage", disk_usage) \
            .time(datetime.now(timezone.utc), WritePrecision.NS)
        
        self.write_api.write(bucket=INFLUXDB_BUCKET, org=INFLUXDB_ORG, record=point)
        print(f"✅ 写入系统指标: {host} - CPU: {cpu_usage}%, 内存: {memory_usage}%, 磁盘: {disk_usage}%")
    
    def generate_sample_data(self, count: int = 10):
        """生成示例数据"""
        print(f"\n📊 生成 {count} 条示例数据...")
        
        sensors = ["sensor_001", "sensor_002", "sensor_003"]
        locations = ["办公室", "机房", "仓库"]
        hosts = ["web-server-01", "db-server-01", "cache-server-01"]
        
        for i in range(count):
            # 生成传感器数据
            sensor_id = random.choice(sensors)
            location = random.choice(locations)
            temperature = round(random.uniform(18.0, 35.0), 1)
            humidity = round(random.uniform(30.0, 80.0), 1)
            self.write_sensor_data(sensor_id, temperature, humidity, location)
            
            # 生成系统监控数据
            host = random.choice(hosts)
            cpu_usage = round(random.uniform(10.0, 90.0), 1)
            memory_usage = round(random.uniform(20.0, 85.0), 1)
            disk_usage = round(random.uniform(15.0, 75.0), 1)
            self.write_system_metrics(host, cpu_usage, memory_usage, disk_usage)
            
            time.sleep(0.1)  # 避免时间戳冲突
    
    def query_sensor_data(self, limit: int = 10):
        """查询传感器数据"""
        query = f'''
        from(bucket: "{INFLUXDB_BUCKET}")
        |> range(start: -1h)
        |> filter(fn: (r) => r["_measurement"] == "sensor_data")
        |> limit(n: {limit})
        '''
        
        print(f"\n🔍 查询最近 {limit} 条传感器数据:")
        tables = self.query_api.query(query, org=INFLUXDB_ORG)
        
        for table in tables:
            for record in table.records:
                print(f"  📍 {record.get_time()} | {record.get_field()}: {record.get_value()} | "
                      f"传感器: {record.values.get('sensor_id')} | 位置: {record.values.get('location')}")
    
    def query_system_metrics(self, limit: int = 10):
        """查询系统监控数据"""
        query = f'''
        from(bucket: "{INFLUXDB_BUCKET}")
        |> range(start: -1h)
        |> filter(fn: (r) => r["_measurement"] == "system_metrics")
        |> limit(n: {limit})
        '''
        
        print(f"\n🖥️ 查询最近 {limit} 条系统监控数据:")
        tables = self.query_api.query(query, org=INFLUXDB_ORG)
        
        for table in tables:
            for record in table.records:
                print(f"  📊 {record.get_time()} | {record.get_field()}: {record.get_value()}% | "
                      f"主机: {record.values.get('host')}")
    
    def query_aggregated_data(self):
        """查询聚合数据 - 计算平均值"""
        query = f'''
        from(bucket: "{INFLUXDB_BUCKET}")
        |> range(start: -1h)
        |> filter(fn: (r) => r["_measurement"] == "sensor_data")
        |> filter(fn: (r) => r["_field"] == "temperature")
        |> group(columns: ["location"])
        |> mean()
        '''
        
        print(f"\n📈 各位置平均温度:")
        tables = self.query_api.query(query, org=INFLUXDB_ORG)
        
        for table in tables:
            for record in table.records:
                location = record.values.get('location')
                avg_temp = round(record.get_value(), 2)
                print(f"  🌡️ {location}: {avg_temp}°C")
    
    def close(self):
        """关闭连接"""
        self.client.close()
        print("\n✅ InfluxDB 连接已关闭")


def main():
    """主函数"""
    print("🚀 InfluxDB 示例程序启动")
    print(f"📡 连接到: {INFLUXDB_URL}")
    print(f"🏢 组织: {INFLUXDB_ORG}")
    print(f"🪣 存储桶: {INFLUXDB_BUCKET}")
    
    demo = InfluxDBDemo()
    
    try:
        # 1. 生成示例数据
        demo.generate_sample_data(5)
        
        # 2. 查询原始数据
        demo.query_sensor_data(10)
        demo.query_system_metrics(10)
        
        # 3. 查询聚合数据
        demo.query_aggregated_data()
        
    except Exception as e:
        print(f"❌ 错误: {e}")
    finally:
        demo.close()


if __name__ == "__main__":
    main()
