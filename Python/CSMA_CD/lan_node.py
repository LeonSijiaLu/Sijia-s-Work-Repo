from collections import deque

class Node:
    def __init__(self, event_array, index):
        self.eventQueue = deque(event_array) 
        self.next_event_time = self.eventQueue[0].event_time
        self.index = index
        self.busy_count = 0
        self.collision_count = 0
        self.dropped_packets = 0
        self.successful_packets = 0
        self.total_packets = 0