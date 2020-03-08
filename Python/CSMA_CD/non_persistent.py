import lan_node
import arrival
import math
from random import random, randrange

class non_persistent_csma_cd:
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
        self.t_prop = self.D / self.S # Propagation Time
        self.t_trans = self.L / self.R # Transmission Time
        self.jamming_time = 48/self.R
        self.lan = []
        self.timer = 0
        
    def get_random_variable(self, mean):
        return (-mean) * math.log(1 - random())

    def exponential_backoff(self, collisions):
        return randrange(0, 2**collisions - 1) * 512 / self.R

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

    def populate_node_stack(self):  # populate packet stack of each node
        for i in range(self.N):
            node_time = 0
            node_events = []
            while node_time < self.T:
                random_interval_time = self.get_random_variable(1.0 / self.A)  
                node_time = node_time + random_interval_time
                node_events.append(arrival.ArrivalEvent(node_time)) # node_events is an array, elements are ArrivalEvent
            self.lan.append(lan_node.Node(node_events, i))

    def determine_next_sender(self): # find the next sender in LAN by simply comparing each node's earliest time
        sender_index = -1
        next_packet_time = self.T
        for i in self.lan:
            if next_packet_time > i.next_event_time and len(i.eventQueue) > 0:
                next_packet_time = i.next_event_time
                sender_index = i.index
        return sender_index

    def have_collision(self, node, new_time):   # call this function when having collisions
        node.collision_count += 1
        node.total_packets += 1
        node.busy_count = 0
        self.total_collisions += 1
        self.total_packets += 1
        if node.collision_count > self.K_max:   # If collision count exceeds K_max, drop the packet
            node.eventQueue.popleft()
            node.collision_count = 0
            node.dropped_packets += 1
            self.dropped_packets += 1
            if len(node.eventQueue)>0 and node.eventQueue[0].event_time > node.next_event_time:
                node.next_event_time = node.eventQueue[0].event_time
        else:                                   # exponential backoff
            backoff_time = self.exponential_backoff(node.collision_count)
            node.next_event_time = max(new_time + backoff_time, node.eventQueue[0].event_time)

    def have_busy_wait(self, node, new_time):
        node.busy_count += 1
        backoff_time = self.exponential_backoff(node.busy_count)
        if node.busy_count >= self.K_max:
            node.eventQueue.popleft()
            node.busy_count = 0
            node.dropped_packets += 1
            node.total_packets += 1
            self.dropped_packets += 1
            self.total_packets += 1
        else:
            node.next_event_time = max(new_time + backoff_time, node.eventQueue[0].event_time)

    def start_non_persistent_csma_cd_simulation(self):
        self.populate_node_stack()
        while self.timer < self.T:
            sender_index = self.determine_next_sender()
            sender_node = self.lan[sender_index]
            self.timer = sender_node.next_event_time

            for i in range(self.N):                                 # Check if possible collisions
                if i == sender_index or len(self.lan[i].eventQueue) == 0:
                    continue
                first_bit_arrival = abs(i-sender_index)*self.t_prop + self.timer
                last_bit_arrival = first_bit_arrival + self.t_trans
                if self.lan[i].next_event_time < first_bit_arrival: # detect collision
                    self.have_collision(self.lan[i], last_bit_arrival) # delay self.lan[i] packets

            for i in range(self.N):                                 # Check if possible busy waits
                if i == sender_index or len(self.lan[i].eventQueue) == 0:
                    continue
                first_bit_arrival = abs(i-sender_index)*self.t_prop + self.timer
                last_bit_arrival = first_bit_arrival + self.t_trans
                if first_bit_arrival <= self.lan[i].next_event_time <= last_bit_arrival: # busy wait time
                    self.have_busy_wait(self.lan[i], last_bit_arrival)

            sender_node.collision_count = 0
            sender_node.eventQueue.popleft()
            if len(sender_node.eventQueue) > 0 and sender_node.eventQueue[0].event_time > sender_node.next_event_time:
                sender_node.next_event_time = sender_node.eventQueue[0].event_time
            self.successful_packets += 1
            self.total_packets += 1
            sender_node.successful_packets += 1

        for n in self.lan:
            self.total_packets += len(n.eventQueue)

        print('N =', self.N)
        print('A =', self.A)
        print('Successfully transmitted:', self.successful_packets)
        print('Total packets:', self.total_packets)
        print('Efficiency:', self.successful_packets / self.total_packets)
        print('Throughput:', self.successful_packets * self.L / self.T)
        print('Total collisions:', self.total_collisions)
        print('Dropped packets:', self.dropped_packets)

        return [self.N, self.successful_packets / self.total_packets, self.A]