import math
from random import random, randrange

def get_random_variable(mean):
    return (-mean) * math.log(1 - random())

def exponential_backoff(collisions):
    return randrange(0, 2**collisions - 1) * 512

