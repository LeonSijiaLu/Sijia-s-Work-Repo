import matplotlib.pyplot as plt

#
# All these data are collected through real experiments
# You can find these data from .txt files in this directory
#

def draw_efficiency_vs_nodes_non_persistent():
    nodes_for_arrival_7 = [20, 40, 60, 80, 100]
    efficiency_for_arrival_7 = [0.978791228382956, 0.9548324273423789, 0.9285117850229627, 0.9006178137479518, 0.8716632849722776]

    nodes_for_arrival_10 = [20, 40, 60, 80, 100]
    efficiency_for_arrival_10 = [0.9709305224528058, 0.9389480298938454, 0.9073773841282919, 0.8762434640218827, 0.8482826339032393]

    nodes_for_arrival_20 = [20, 40, 60, 80, 100]
    efficiency_for_arrival_20 = [0.9537725704277185, 0.9196971453719669, 0.895999782170849, 0.8791822781517653, 0.8672548570794404]

    plt.plot(nodes_for_arrival_7, efficiency_for_arrival_7, color="red", marker="^", linewidth=1, label="Arrival Rate 7")
    plt.plot(nodes_for_arrival_10, efficiency_for_arrival_10, color="green", marker="x", linewidth=1, label="Arrival Rate 10")
    plt.plot(nodes_for_arrival_20, efficiency_for_arrival_20, color="orange", marker="s", linewidth=1, label="Arrival Rate 20")

    plt.xlabel("Number of Nodes")
    plt.ylabel("Efficiency")
    plt.legend(loc='best')
    plt.title("Efficiency (Y) vs The number of nodes (X), Non-Persistent")
    plt.show()

def draw_throughput_vs_nodes_non_persistent():
    nodes_for_arrival_7 = [20, 40, 60, 80, 100]
    throughput_for_arrival_7 = [205141.5, 401539.5, 587124.0, 761820.0, 925356.0]

    nodes_for_arrival_10 = [20, 40, 60, 80, 100]
    throughput_for_arrival_10 = [291520.5, 564804.0, 819364.5, 1059529.5, 1289379.0]

    nodes_for_arrival_20 = [20, 40, 60, 80, 100]
    throughput_for_arrival_20 = [574800.0, 1106535.0, 1628872.5, 2129463.0, 2625693.0]

    plt.plot(nodes_for_arrival_7, throughput_for_arrival_7, color="red", marker="^", linewidth=1, label="Arrival Rate 7")
    plt.plot(nodes_for_arrival_10, throughput_for_arrival_10, color="green", marker="x", linewidth=1, label="Arrival Rate 10")
    plt.plot(nodes_for_arrival_20, throughput_for_arrival_20, color="orange", marker="s", linewidth=1, label="Arrival Rate 20")

    plt.xlabel("Number of Nodes")
    plt.ylabel("Throughput")
    plt.legend(loc='best')
    plt.title("Throughput (Y) vs The number of nodes (X), Non-Persistent")
    plt.show()

def draw_efficiency_vs_nodes_persistent():
    nodes_for_arrival_7 = [20, 40, 60, 80, 100]
    efficiency_for_arrival_7 = [0.9744480365593984, 0.9289746709132414, 0.8467308923661534, 0.7038136542900171, 0.5119170952575429]

    nodes_for_arrival_10 = [20, 40, 60, 80, 100]
    efficiency_for_arrival_10 = [0.9590459067951462, 0.8665987963125272, 0.6661224831026359, 0.445407287272243, 0.3669437273921348]

    nodes_for_arrival_20 = [20, 40, 60, 80, 100]
    efficiency_for_arrival_20 = [0.8782121938110897, 0.5157340229389651, 0.4072542241624277, 0.3767778564694048, 0.3596140693744758]

    plt.plot(nodes_for_arrival_7, efficiency_for_arrival_7, color="red", marker="^", linewidth=1, label="Arrival Rate 7")
    plt.plot(nodes_for_arrival_10, efficiency_for_arrival_10, color="green", marker="x", linewidth=1, label="Arrival Rate 10")
    plt.plot(nodes_for_arrival_20, efficiency_for_arrival_20, color="orange", marker="s", linewidth=1, label="Arrival Rate 20")

    plt.xlabel("Number of Nodes")
    plt.ylabel("Efficiency")
    plt.legend(loc='best')
    plt.title("Efficiency (Y) vs The number of nodes (X), Persistent")
    plt.show()

def draw_throughput_vs_nodes_persistent():
    nodes_for_arrival_7 = [20, 40, 60, 80, 100]
    throughput_for_arrival_7 = [206620.5, 407766.0, 606976.5, 805482.0, 998385.0]

    nodes_for_arrival_10 = [20, 40, 60, 80, 100]
    throughput_for_arrival_10 = [293656.5, 578410.5, 860526.0, 1123800.0, 1328389.5]

    nodes_for_arrival_20 = [20, 40, 60, 80, 100]
    throughput_for_arrival_20 = [578445.0, 1130776.5, 1565034.0, 1927182.0, 2258875.5]

    plt.plot(nodes_for_arrival_7, throughput_for_arrival_7, color="red", marker="^", linewidth=1, label="Arrival Rate 7")
    plt.plot(nodes_for_arrival_10, throughput_for_arrival_10, color="green", marker="x", linewidth=1, label="Arrival Rate 10")
    plt.plot(nodes_for_arrival_20, throughput_for_arrival_20, color="orange", marker="s", linewidth=1, label="Arrival Rate 20")

    plt.xlabel("Number of Nodes")
    plt.ylabel("Throughput")
    plt.legend(loc='best')
    plt.title("Throughput (Y) vs The number of nodes (X), Persistent")
    plt.show()


draw_efficiency_vs_nodes_non_persistent()
draw_throughput_vs_nodes_non_persistent()
draw_efficiency_vs_nodes_persistent()
draw_throughput_vs_nodes_persistent()