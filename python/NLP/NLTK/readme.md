# NLTK

NLTK，全称为Natural Language Toolkit，是一个功能强大、易于使用的自然语言处理（NLP）的Python库，是目前最受欢迎和广泛使用的NLP库之一。

NLTK包含了大量的自然语言处理工具和数据集，包括文本预处理、词性标注、句法分析、情感分析、机器翻译、语音识别等多个方面的工具和算法。这些工具和算法都是经过精心设计和实现的，并且支持多种自然语言处理任务的处理。

除此之外，NLTK还是一个开源项目，意味着用户可以自由地使用、修改和分发它的代码，而且有着强大的社区支持和贡献。

运行环境：

```python3
import nltk
nltk.download('stopwords')
nltk.download('wordnet')
nltk.download('averaged_perceptron_tagger')
```

output:

```
python3 hello.py
Sentences: ['This is an example sentence.', "It contains punctuation, numbers like 1, 2, and 3, and stop words such as 'the' and 'and'."]
Tokens: ['example', 'sentence', 'contains', 'punctuation', 'numbers', 'like', 'stop', 'words']
Lemmatized Tokens: ['example', 'sentence', 'contains', 'punctuation', 'number', 'like', 'stop', 'word']
POS Tags: [('example', 'NN'), ('sentence', 'NN'), ('contains', 'VBZ'), ('punctuation', 'NN'), ('number', 'NN'), ('like', 'IN'), ('stop', 'NN'), ('word', 'NN')]
```

## 词性标注返回值

- NN：名词（Noun）
- VB：动词原形（Verb）
- VBD：动词过去式（Verb, past tense）
- VBG：动词现在分词（Verb, gerund or present participle）
- VBN：动词过去分词（Verb, past participle）
- VBP：动词非第三人称单数现在时（Verb, non-3rd person singular present）
- VBZ：动词第三人称单数现在时（Verb, 3rd person singular present）
- JJ：形容词（Adjective）
- RB：副词（Adverb）
- IN：介词（Preposition）
- PRP：代词（Pronoun）
- CC：连词（Coordinating conjunction）
- DT：限定词（Determiner）
- CD：基数词（Cardinal number）
- EX：存在句中的there（Existential there）

