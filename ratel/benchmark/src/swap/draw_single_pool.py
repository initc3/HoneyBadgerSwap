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
plt.rc('xtick', labelsize=15)
# Set the font size for y tick labels
plt.rc('ytick', labelsize=15)
# Set the legend font size
plt.rc('legend', fontsize=12)
# Set the font size of the figure title
plt.rc('figure', titlesize=20)


if __name__ == '__main__':
    times = {}
    with open(f'ratel/benchmark/src/swap/pool_data/traderjoev2_USDC.e_WAVAX.csv', 'r') as f:
        lines = f.readlines()
        for line in lines[1:]:
            element = re.split(',|\n', line)
            t = float(element[0]) // 3600
            if t not in times.keys():
                times[t] = 0
            times[t] += 1

    fig = plt.figure()
    ax1 = fig.add_subplot(111)

    lists = sorted(times.items())
    x, y = zip(*lists)
    ax1.bar(x, y, color='cornflowerblue')
    ax1.set_xlim(0, len(x))
    ax1.set_xlabel('Date')
    ax1.set_ylabel('Swaps per hour')
    # ax1.set_ylim(0, 350)
    x_labels = ['05/13', '05/14', '05/15', '05/16', '05/17', '05/18']
    ax1.set_xticks(x[10::24])
    ax1.set_xticklabels(x_labels)
    plt.show()
