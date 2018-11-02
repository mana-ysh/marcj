# marcj
Morphological Analysis Result Converter for Japanese

## Installation

`git clone https://github.com/mana-ysh/marcj.git && cd marcj && go install`

## Usage

1. prepare tokenized file
```
⟩⟩⟩ cat ./test/test.txt.sudachi
丹生	名詞,固有名詞,人名,姓,*,*	丹生
明里	名詞,固有名詞,人名,名,*,*	明里
は	助詞,係助詞,*,*,*,*	は
可愛い	形容詞,一般,*,*,形容詞,終止形-一般	可愛い
EOS
松田	名詞,固有名詞,人名,姓,*,*	松田
好	名詞,固有名詞,人名,名,*,*	好
花	名詞,普通名詞,一般,*,*,*	花
も	助詞,係助詞,*,*,*,*	も
かわいい	形容詞,一般,*,*,形容詞,終止形-一般	可愛い
EOS
```

2. fed into `marcj` command
```
⟩⟩⟩ cat ./test/test.txt.sudachi | marcj
丹生 明里 は 可愛い
松田 好 花 も かわいい
```

Output as normalized
```
⟩⟩⟩ cat test/test.txt.sudachi | marcj -n
丹生 明里 は 可愛い
松田 好 花 も 可愛い
```

Output as morpheme sequence
```
⟩⟩⟩ cat test/test.txt.sudachi| marcj -m
丹生,名詞,固有名詞,人名,姓,*,*,丹生 明里,名詞,固有名詞,人名,名,*,*,明里 は,助詞,係助詞,*,*,*,*,は 可愛い,形容詞,一般,*,*,形容詞,終止形-一般,可愛い
松田,名詞,固有名詞,人名,姓,*,*,松田 好,名詞,固有名詞,人名,名,*,*,好 花,名詞,普通名詞,一般,*,*,*,花 も,助詞,係助詞,*,*,*,*,も かわいい,形容詞,一般,*,*,形容詞,終止形-一般,可愛い
```
