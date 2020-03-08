from collections import deque

class Node:
    def __init__(self, event_array, index):
        self.eventQueue = deque(event_array)                   # packet stack of each node
        self.next_event_time = self.eventQueue[0].event_time   # next event of the packet stack, or the earliest event
        self.index = index                                     # the index of the node in LAN
        self.busy_count = 0                         
        self.collision_count = 0
        self.dropped_packets = 0
        self.successful_packets = 0
        self.total_packets = 0