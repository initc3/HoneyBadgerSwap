import re
import sys

import numpy as np
import matplotlib.pyplot as plt

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
    dir = sys.argv[1]
    players = int(sys.argv[2])
    threshold = int(sys.argv[3])
    batch_num = int(sys.argv[4])

    x = np.zeros(batch_num)
    y = np.zeros(batch_num)

    for server_id in range(players):
        with open(f'{dir}/inputmask_generation_latency_{server_id}.csv', 'r') as f:
            for line in f.readlines():
                element = re.split('\t|\n', line)

                if element[0] == 'x':
                    x += eval(element[1])
                elif element[0] == 'y':
                    y += eval(element[1])
    x /= players
    y /= players

    # find line of best fit
    a, b = np.polyfit(x, y, 1)

    # add points to plot
    plt.scatter(x, y)

    # add line of best fit to plot
    plt.plot(x, a * x + b)

    plt.xlabel(f"Input mask number")
    plt.ylabel("Time cost(sec)")
    # ax = plt.gca()
    # ax.xaxis.set_major_locator(MultipleLocator(1))
    # ax.yaxis.set_major_locator(MultipleLocator(10))
    # plt.legend()
    plt.show()
    # plt.savefig(f"{dir}/inputmask.pdf")

    print(a, b)