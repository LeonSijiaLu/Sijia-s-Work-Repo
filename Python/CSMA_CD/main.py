import non_persistent
import persistent

def run_non_persistent_csma_cd_simu():
    non_persistent_csma_cd = non_persistent.non_persistent_csma_cd(60, 7, 1000, False)
    non_persistent_csma_cd.start_non_persistent_csma_cd_simulation()
    #non_persistent_csma_cd.print_lan_config()

#def run_persistent_csma_cd_simu():
 #   persistent_csma_cd = persistent.persistent_csma_cd(60,7,1000,False)

run_non_persistent_csma_cd_simu()