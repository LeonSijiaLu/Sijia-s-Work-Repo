import matplotlib.pyplot as plt

#
# All these data are collected through real experiments
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

    plt.legend(loc='upper left', bbox_to_anchor=(0.0, 1.0))
    plt.title("Efficiency vs The number of nodes, Non-Persistent")
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

    plt.legend(loc='upper left', bbox_to_anchor=(0.0, 1.0))
    plt.title("Throughput vs The number of nodes, Non-Persistent")
    plt.show()


draw_efficiency_vs_nodes_non_persistent()
draw_throughput_vs_nodes_non_persistent()