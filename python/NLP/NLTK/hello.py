import nltk
from nltk.tokenize import word_tokenize, sent_tokenize
from nltk.corpus import stopwords
from nltk.stem import WordNetLemmatizer

# 定义待处理的文本
text = "This is an example sentence. It contains punctuation, numbers like 1, 2, and 3, and stop words such as 'the' and 'and'."

# 分句
sentences = sent_tokenize(text)

# 分词并去除停用词和标点符号
stop_words = stopwords.words("english")
tokens = []
for sentence in sentences:
    words = word_tokenize(sentence)
    words = [word.lower() for word in words if word.isalpha()]
    words = [word for word in words if word not in stop_words]
    tokens.extend(words)

# 词形还原
lemmatizer = WordNetLemmatizer()
lemmatized_tokens = [lemmatizer.lemmatize(token) for token in tokens]

# 词性标注
pos_tags = nltk.pos_tag(lemmatized_tokens)

print("Sentences:", sentences)
print("Tokens:", tokens)
print("Lemmatized Tokens:", lemmatized_tokens)
print("POS Tags:", pos_tags)
# Sentences: ['This is an example sentence.', "It contains punctuation, numbers like 1, 2, and 3, and stop words such as 'the' and 'and'."]
# Tokens: ['example', 'sentence', 'contains', 'punctuation', 'numbers', 'like', 'stop', 'words']
# Lemmatized Tokens: ['example', 'sentence', 'contains', 'punctuation', 'number', 'like', 'stop', 'word']
# POS Tags: [('example', 'NN'), ('sentence', 'NN'), ('contains', 'VBZ'), ('punctuation', 'NN'), ('number', 'NN'), ('like', 'IN'), ('stop', 'NN'), ('word', 'NN')]

