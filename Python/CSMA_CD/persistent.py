import lan_node
import utils
import arrival

class persistent_csma_cd:
    def __init__(self, N, A, T, persistent):
        self.N = N          # Number of nodes
        self.A = A          # Average number of packets per second
        self.T = T          # Simulation Time
        self.D = 10         # Distance between two neighbouring nodes, 10m
        self.K_max = 10     # max_wait
        self.dropped_packets = 0
        self.successful_packets = 0
        self.total_packets = 0
        self.total_collisions = 0
        self.R = 1e6        # Transmission speed, 1Mb
        self.L = 1500       # packet length is 1500 bits
        self.C = 3e8        # Speed of light
        self.S = (2/3) * self.C  # Propagation speed
        self.t_prop = self.D / self.S
        self.t_trans = self.L / self.R
        self.jamming_time = 48/self.R
        self.lan = []
        self.timer = 0

    def print_lan_config(self):
        print("N is ", self.N)
        print("A is ", self.A)
        print("T is ", self.T)
        print("D is ", self.D)
        print("K_max is ", self.K_max)
        print("R is ", self.R)
        print("L is ", self.L)
        print("C is ", self.C)
        print("S is ", self.S)
        print("t_prop is ", self.t_prop)
        print("t_trans is ", self.t_trans)
        print("jamming_time is ", self.jamming_time)

    def populate_node_stack(self):
        for i in range(self.N):
            node_time = 0
            node_events = []
            while node_time < self.T:
                random_interval_time = utils.get_random_variable(1.0 / self.A)
                node_time = node_time + random_interval_time
                node_events.append(arrival.ArrivalEvent(node_time)) # node_events is an array, elements are ArrivalEvent
            self.lan.append(lan_node.Node(node_events, i))

    def determine_next_sender(self):
        sender_index = -1
        next_packet_time = self.T
        for i in self.lan:
            if next_packet_time > i.next_event_time:
                next_packet_time = i.next_event_time
                sender_index = i.index
        return sender_index

    