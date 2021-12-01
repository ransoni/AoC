#!/usr/bin/env python3


def main():
    with open('data.txt') as f:
        comb = []
        sum = 0
        for line in f:
            line = line.rstrip()
            if not line:
                sum += len(set([c for l in comb for c in l]))
                comb = []
                continue
            comb.append(line)
        sum += len(set([c for l in comb for c in l]))
    print(sum)

if __name__ == '__main__':
    main()