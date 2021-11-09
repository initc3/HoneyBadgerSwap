import matplotlib.pyplot as plt
import numpy as np
import re
import sys

from ratel.src.python.utils import replay

dir = sys.argv[1]

interval = 20
width = 0.25

def trunc_time(t):
    return float(t) // interval

def add(map, key, num=1):
    if key not in map.keys():
        map[key] = num
    else:
        map[key] += num

def plot(plt, map, offset, label):
    lists = sorted(map.items())
    x, y = zip(*lists)
    x, y = np.array(x), np.array(y)
    plt.bar(x + width * offset, y, width=width, label=label)

start = {}
start_mpc = {}
end = {}

file = f'{dir}/latency_0.csv'
with open(file, 'r') as f:
    lines = f.readlines()
    for line in lines:
        element = re.split('\t|\n', line)
        if element[0] == 'trade-1':
            time = trunc_time(element[4])
            add(start, time)
        elif element[0] == 'trade-6':
            time = trunc_time(element[4])
            add(start_mpc, time)
        elif element[0] == 'trade-12':
            time = trunc_time(element[4])
            add(end, time)

client = {}

file = f'{dir}/gas.csv'
with open(file, 'r') as f:
    lines = f.readlines()
    for line in lines:
        element = re.split('\t|\n', line)
        time = trunc_time(element[10])
        add(client, time, replay)

fig = plt.figure(figsize=(13, 4))

plot(plt, client, 0, 'client')
plot(plt, start, 1, 'start')
plot(plt, start_mpc, 2, 'start_mpc')
plot(plt, end, 3, 'end')

plt.xlabel(f"time(/{interval}secs)")
plt.ylabel("count")
plt.legend()
plt.savefig(f"{dir}/fig.pdf")
