#!/usr/bin/env python

import argparse

import numpy as np
import pandas as pd
import seaborn as sns
import matplotlib.pyplot as plt

sns.set_style('whitegrid')
sns.set_context('notebook')


def draw_benchmark(path, vtype='line'):
    df = pd.read_csv(path)

    if vtype == 'both':
        _, axes = plt.subplots(ncols=2, figsize=(18,6), sharey=True)
        draw_line_benchmark(df, axes[0])
        draw_bar_benchmark(df, axes[1])

    else:
        _, ax = plt.subplots(figsize=(9,6))

        if vtype == 'line':
            draw_line_benchmark(df, ax)
        elif vtype == 'bar':
            draw_bar_benchmark(df, ax)
        else:
            raise ValueError("unknown viz type: '{}'".format(vtype))

    plt.tight_layout()
    plt.savefig("benchmark.png")


def draw_line_benchmark(df, ax):
    max_clients = df['clients'].max()

    for stype in df['server'].unique():
        sample = df[df['server'] == stype]
        means = sample.groupby('clients')['throughput'].mean()
        std = sample.groupby('clients')['throughput'].std()

        ax.plot(means, label=stype)
        ax.fill_between(np.arange(1, max_clients+1), means+std, means-std, alpha=0.25)

    ax.set_xlim(1, max_clients)
    ax.set_ylabel("throughput (requests/second)")
    ax.set_xlabel("concurrent clients")
    ax.set_title("Ipseity Benchmark")
    ax.legend(frameon=True)
    return ax


def draw_bar_benchmark(df, ax):
    g = sns.barplot('clients', 'throughput', hue='server', ax=ax, data=df)
    ax.set_ylabel("")
    ax.set_xlabel("concurrent clients")
    ax.set_title("Ipseity Benchmark")
    return ax


if __name__ == '__main__':
    parser = argparse.ArgumentParser(
        description="draw the benchmark visualization from a dataset",
    )

    parser.add_argument(
        '-t', '--type', choices=('bar', 'line', 'both'), default='line',
        help='specify the type of chart to produce'
    )
    parser.add_argument("data")

    args = parser.parse_args()
    draw_benchmark(args.data, args.type)
