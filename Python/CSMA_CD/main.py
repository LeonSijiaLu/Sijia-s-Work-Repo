import non_persistent
import persistent

non_persistent_result = []
persistent_result = []

def run_non_persistent_csma_cd_simu(N, A, T):
    non_persistent_csma_cd = non_persistent.non_persistent_csma_cd(N, A, T, False)
    return non_persistent_csma_cd.start_non_persistent_csma_cd_simulation()

def show_non_persistent_performance():  # call non_persistent csma/cd
    for A in [7, 10, 20]:               # loops throught arrival rate
        for N in [20, 40, 60, 80, 100]: # loops through number of nodes
            non_persistent_result.append(run_non_persistent_csma_cd_simu(N, A, 1000))
    print(non_persistent_result)

def run_persistent_csma_cd_simu(N, A, T):
    persistent_csma_cd = persistent.persistent_csma_cd(N, A, T, False)
    return persistent_csma_cd.start_persistent_csma_cd_simulation()

def show_persistent_performance():      # call persistent csma/cd
    for A in [7, 10, 20]:               # loops throught arrival rate
        for N in [20, 40, 60, 80, 100]: # loops through number of nodes
            persistent_result.append(run_persistent_csma_cd_simu(N, A, 1000))
    print(persistent_result)

show_non_persistent_performance()
show_persistent_performance()
    
