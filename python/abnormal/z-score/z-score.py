import pandas as pd
import numpy as np
from sklearn.preprocessing import StandardScaler
import matplotlib.pyplot as plt
from matplotlib import font_manager
import os

def setup_font():
    potential_fonts = [
        '/usr/share/fonts/truetype/wqy/wqy-microhei.ttc',
        '/usr/share/fonts/wqy-microhei/wqy-microhei.ttc',
    ]
    
    for font_path in potential_fonts:
        if os.path.exists(font_path):
            font_manager.fontManager.addfont(font_path)
            plt.rcParams['font.family'] = font_manager.FontProperties(fname=font_path).get_name()
            break
    plt.rcParams['axes.unicode_minus'] = False

def load_data(file_path):
    data = pd.read_csv(file_path, header=None)
    return data.iloc[0].values

def detect_anomalies_zscore(data, threshold=3):
    scaler = StandardScaler()
    z_scores = scaler.fit_transform(data.reshape(-1, 1))
    return (abs(z_scores) > threshold).ravel()

def plot_results(data, anomalies, output_path='zscore_anomalies.png'):
    plt.figure(figsize=(15, 5))
    plt.plot(range(len(data)), data, label='原始数据')
    anomaly_indices = np.where(anomalies)[0]
    anomaly_values = data[anomaly_indices]
    plt.scatter(anomaly_indices, anomaly_values, color='red', label='异常值')
    plt.title('使用Z-score方法检测的异常值')
    plt.legend()
    plt.tight_layout()
    plt.savefig(output_path)
    plt.close()

def main():
    setup_font()
    data = load_data('data.csv')
    
    anomalies = detect_anomalies_zscore(data, threshold=3)
    print(f'Z-score方法检测到 {sum(anomalies)} 个异常值')
    
    plot_results(data, anomalies)
    
    print("\n详细的异常值信息：")
    anomaly_indices = np.where(anomalies)[0]
    anomaly_values = data[anomaly_indices]
    for idx, value in zip(anomaly_indices, anomaly_values):
        print(f"索引: {idx}, 值: {value:.2f}")

if __name__ == "__main__":
    main() 
