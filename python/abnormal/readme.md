# 异常数据检测

## 统计学方法

- 基于分布的方法： 假设数据遵循某种已知的分布（例如正态分布），然后通过计算数据点与分布的偏差来判断是否异常。常用的方法包括：
    - Z-score/标准差方法： 计算数据点与均值的偏差，如果偏差超过某个阈值（例如 3 个标准差），则认为是异常值。
    - 箱线图方法 (IQR 方法)： 使用四分位数间距（IQR）来定义正常范围，超出范围的数据点被认为是异常值。
- 基于邻近度的方法： 考察数据点周围的邻居分布情况，如果一个数据点的邻居很少或者距离很远，则认为是异常值。
- 时间序列分析方法： 针对时间序列数据，分析数据的趋势、季节性、周期性等特征，检测偏离这些特征的异常点。常用的方法包括：
    - 移动平均/指数平滑： 比较实际值与预测值的偏差。
    - ARIMA 模型： 预测时间序列的未来值，并检测预测误差较大的点。

## 机器学习方法

- 监督学习： 需要标记好的数据集（包含正常和异常样本），训练一个分类器来区分正常和异常数据。常用的算法包括：
    - 支持向量机 (SVM)： 寻找一个最优的超平面来分隔正常和异常数据。
    - 决策树/随机森林： 构建决策树或随机森林来对数据进行分类。
- 非监督学习： 不需要标记数据，通过学习数据的内在结构来发现异常。常用的算法包括：
    - 聚类算法 (例如 K-means, DBSCAN)： 将数据分成不同的簇，远离簇中心的数据点被认为是异常值。
    - 自编码器 (Autoencoder)： 训练一个神经网络来重构输入数据，重构误差较大的数据点被认为是异常值。
    - One-Class SVM： 训练一个只包含正常数据的 SVM 模型，用于检测与正常数据分布不同的异常点。
    - 孤立森林 (Isolation Forest)： 通过随机切割数据来构建树，异常值更容易被孤立出来。
- 半监督学习： 使用部分标记的数据进行训练，通常使用正常数据进行训练，然后检测与正常数据分布不同的异常点。

## 深度学习方法

深度学习方法在处理复杂数据和高维数据方面表现出色。常用的方法包括：

- 自编码器 (Autoencoder)： 如前所述，深度学习中的自编码器可以学习数据的低维表示，并用于异常检测。
- 循环神经网络 (RNN) 和长短期记忆网络 (LSTM)： 适用于时间序列数据的异常检测，可以捕捉时间序列中的长期依赖关系。
- 生成对抗网络 (GAN)： 可以用于生成正常数据样本，然后通过判别器来检测异常。
- Transformer 架构： 通过自注意力机制，可以捕捉时序数据中的长程依赖关系，用于异常检测。

## 规则阈值法

根据业务规则或专家经验定义一系列规则，不符合规则的数据点被认为是异常值。