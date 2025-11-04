# Markov Chain Text Generator

## Description

The main aim of this project is to create a Markov chain algorithm that randomly generates sentences. Similar algorithms are used in smartphone keyboards - when you type a word, the keyboard suggests the next most probable word based on learned patterns.

## Installation

### Clone the Repository

```bash
git clone git@github.com:abu166/Markovify.git
cd Markovify
```

### Build

```bash
go build -o markovchain .
```

## Usage

### Basic Usage

Generate text by piping input from a text file:

```bash
cat text.txt | ./markovchain
```

**Example output:**

```
Chapter 1 In my younger and more stable, become for a job. He hadn't eat anything for a long, silent time. It was the sound of someone splashing after us over the confusion a long many-windowed room which overhung the terrace. Eluding Jordan's undergraduate who was well over sixty, and Maurice A. Flink and the great bursts of leaves growing on the air now. "How do you want? What do you like Europe?" she exclaimed surprisingly. "I just got here a minute. "Yes." He hesitated. "Was she killed?" "Yes." "I thought you didn't, if you'll pardon my--you see, I carry
```

### Count Generated Words

```bash
cat text.txt | ./markovchain | wc -w
```

**Output:**

```
100
```

### Error Handling

Running without input will produce an error:

```bash
./markovchain
```

**Output:**

```
Error: Empty stdin
```

### Advanced Options

#### Limit Word Count

Generate text with a maximum of 10 words:

```bash
cat text.txt | ./markovchain -w 10
```

**Output:**

```
Chapter 1 In ...
```

#### Custom Starting Prefix

Generate text starting with a specific phrase and custom prefix length:

```bash
cat text.txt | ./markovchain -w 10 -p "to something funny" -l 3
```

**Output:**

```
to something funny ...
```

## Options

View all available options:

```bash
./markovchain --help
```

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
```

## How It Works

A Markov chain is a stochastic model that predicts the next state based on the current state. In this text generator:

1. The algorithm analyzes input text to learn word patterns
2. It builds a chain of word associations based on frequency
3. Starting from a prefix, it probabilistically selects the next word
4. This process continues to generate coherent-looking text


## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.