import matplotlib.pyplot as plt
import numpy as np
import re


# Set the default text font size
plt.rc('font', size=15)
# Set the axes title font size
plt.rc('axes', titlesize=15)
# Set the axes labels font size
plt.rc('axes', labelsize=15)
# Set the font size for x tick labels
plt.rc('xtick', labelsize=12)
# Set the font size for y tick labels
plt.rc('ytick', labelsize=15)
# Set the legend font size
plt.rc('legend', fontsize=12)
# Set the font size of the figure title
plt.rc('figure', titlesize=20)


if __name__ == '__main__':
    x, y = [], []
    with open(f'ratel/benchmark/src/swap/pool_data/pool_rankings_rate.csv', 'r') as f:
        lines = f.readlines()
        for line in lines:
            element = re.split(',|\n', line)
            pool_name = element[0]
            rate = float(element[1])
            x.append(pool_name)
            y.append(rate)

    fig = plt.figure()
    ax1 = fig.add_subplot(111)

    num = 40
    xx = np.arange(0, num)

    ax1.bar(xx, y[:num], color='cornflowerblue')
    ax1.set_xlim(-1, num)
    ax1.set_xlabel('Pool name')
    ax1.set_ylabel('Swaps per hour')
    ax1.set_ylim(0, 350)
    ax1.set_xticks(xx)
    ax1.set_xticklabels(x[:num], rotation=90)
    plt.show()
