import matplotlib.pyplot as plt
import matplotlib.colors as mcolors

from ratel.benchmark.src import trade_throughput, trade_latency

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
    pool_num_list = [1, 2, 4, 8, 16]
    rep = 20
    players = 3

    throughput_list = []
    for pool_num in pool_num_list:
        throughput = 0
        for server_id in range(players):
            dir = f'ratel/benchmark/data/{players}_{pool_num}_{pool_num}_{rep}'
            _, _, _, _, _, mean = trade_throughput.scan(dir, server_id)
            throughput += mean
        throughput /= players
        throughput_list.append(throughput)
    print(throughput_list)

    latency_list = []
    for pool_num in pool_num_list:
        latency = 0
        for server_id in range(players):
            dir = f'ratel/benchmark/data/{players}_{pool_num}_{pool_num}_{rep}'
            _, mean, _ = trade_latency.scan(dir, server_id)
            latency += mean
        latency /= players
        latency_list.append(latency)
    print(latency_list)

    colors = list(mcolors.TABLEAU_COLORS.keys())

    fig = plt.figure()
    ax1 = fig.add_subplot(111)
    ax2 = ax1.twinx()

    ax1.plot(pool_num_list, latency_list, mcolors.TABLEAU_COLORS[colors[0]], label=f'latency')
    ax2.plot(pool_num_list, throughput_list, mcolors.TABLEAU_COLORS[colors[1]], label=f'throughput')

    ax1.set_xlabel('Trading pool number')
    ax1.set_ylabel('Average Trade Latency(s)')
    ax2.set_ylabel('Average Trade Throughput(/min)')
    ax1.set_ylim(0, 8)
    ax2.set_ylim(0, 180)
    ax1.legend(loc='upper left')
    ax2.legend(loc='upper right')
    plt.show()
    # save_file = f'{fig_dir}/scale_chain_num_{max_chain_num}_val_num_{sys.argv[1]}.pdf'
    # plt.savefig(save_file)