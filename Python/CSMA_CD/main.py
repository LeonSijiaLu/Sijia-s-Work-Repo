import non_persistent
import persistent

non_persistent_result = []
persistent_result = []

def run_non_persistent_csma_cd_simu(N, A, T):
    non_persistent_csma_cd = non_persistent.non_persistent_csma_cd(N, A, T, False)
    return non_persistent_csma_cd.start_non_persistent_csma_cd_simulation()

def show_non_persistent_performance():
    for A in [7, 10, 20]:
        for N in [20, 40, 60, 80, 100]:
            non_persistent_result.append(run_non_persistent_csma_cd_simu(N, A, 1000))
    print(non_persistent_result)

