# Goker

In its current state, Goker is a poker hand evaluator written in Go. The evaluator is an adaptation of [ashelly's](https://github.com/ashelly/ACE_eval) 
algorithm in C (without the goal of minimizing source code size).

The algorithm is quite fast, at the expense of source code legibility. It uses primitive data types and bitwise arithmetic to 
achieve good performance.

All the program does for now is generate all possible 7-card hands, evaluate 
them and print basic profiling info to the console.

Output from running the progam gives the following on my machine:

```
Number of hands evaluated: 133784560
Time elapsed: 6.330070 seconds
Evaluations per second: 21134767
```

That's roughly 21 million hands evaluated per second. Not bad for my purpose.


## Getting Started

Grab the source code, compile and run.  


### Prerequisites

Requires Go 1.9


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE) file for details


## Acknowledgments

* Hat's off to [ashelly](https://github.com/ashelly/ACE_eval) for their awesome algorithm and giving me an fun opportunity to practice my Go chops!
