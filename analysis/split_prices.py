import matplotlib
import os

matplotlib.use('agg')
from matplotlib import pyplot as plt
from matplotlib.dates import date2num
from datetime import datetime


def split_by_token(data):
    btc_lines = list()
    eth_lines = list()
    ltc_lines = list()
    for line in data:
        if line[0].isdigit():
            if "Bitcoin" in line:
                btc_lines.append(line)
            if "Ethereum" in line:
                eth_lines.append(line)
            if "Litecoin" in line:
                ltc_lines.append(line)
    return btc_lines, eth_lines, ltc_lines


def create_time_tuple(data):
    tograph = list()
    for line in data:
        lines = line.split(", ")
        datetime_object = datetime.strptime(lines[1], '%d:%M:%S.%f%p')
        tograph.append((datetime_object, lines[3]))
    return tograph


def output_tuple_graph(tograph, name):
    x = [date2num(date) for (date, value) in tograph]
    y = [value for (date, value) in tograph]
    fig = plt.figure(figsize=(20,20))
    graph = fig.add_subplot(111)
    # Plot the data as a red line with round markers
    graph.plot(x, y, '-')
    # Set the xtick locations to correspond to just the dates you entered.
    graph.set_xticks(x)
    # Set the xtick labels to correspond to just the dates you entered.
    graph.set_xticklabels([date.strftime("%M") for (date, value) in tograph])
    plt.savefig(name)


if __name__ == "__main__":
    data = open("28-Nov-17.txt", "r").read().split("\n")
    btc, eth, ltc = split_by_token(data)
    gbtc = create_time_tuple(btc)
    geth = create_time_tuple(eth)
    gltc = create_time_tuple(ltc)
    os.makedirs("output")
    output_tuple_graph(gbtc, "output/btc")
    output_tuple_graph(geth, "output/eth")
    output_tuple_graph(gltc, "output/ltc")
