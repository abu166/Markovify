#MARKOV-CHAIN

### Description
Main aim of this project is to create markov-chain algorithm with random creating sentences

Similar algorithms are used in your phones. For example, when you type a word in the
keyboard it suggests you the next the most probable word


### Clone

```bash
git clone git@github.com:abu166/Markovify.git
```

### Build

```bash
go build -o markovchain .
```

### Examples

```bash
cat text.txt | ./markovchain
```

Chapter 1 In my younger and more stable, become for a job. He hadn't eat anything for a long, silent time. It was the sound of someone splashing after us over the confusion a long many-windowed room which overhung the terrace. Eluding Jordan's undergraduate who was well over sixty, and Maurice A. Flink and the great bursts of leaves growing on the air now. "How do you want? What do you like Europe?" she exclaimed surprisingly. "I just got here a minute. "Yes." He hesitated. "Was she killed?" "Yes." "I thought you didn't, if you'll pardon my--you see, I carry$


```bash
cat text.txt | ./markovchain | wc -w
```

   100

```bash
./markovchain
```

Error: Empty stdin

```bash
cat text.txt | ./markovchain -w 10 
```

Chapter 1 In ...

```bash
cat text.txt | ./markovchain -w 10 -p "to something funny" -l 3
```

to something funny ...

```bash
./markovchain --help
```

Markov Chain text generator.

Usage:
  markovchain [-w <N>] [-p <S>] [-l <N>]
  markovchain --help

Options:
  --help  Show this screen.
  -w N    Number of maximum words
  -p S    Starting prefix
  -l N    Prefix length




