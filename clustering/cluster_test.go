package clustering

import (
	"testing"
)

func compSlices(slice1 []int, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i, elem := range slice1 {
		if elem != slice2[i] {
			return false
		}
	}

	return true
}

func compSlicesOfSlices(sliceOfSlices1 [][]int, sliceOfSlices2 [][]int) bool {
	if len(sliceOfSlices1) != len(sliceOfSlices2) {
		return false
	}

	for i, slice1 := range sliceOfSlices1 {
		if !compSlices(slice1, sliceOfSlices2[i]) {
			return false
		}
	}

	return true
}

func TestCluster(t *testing.T) {
	data := []float64{0.4025641025641026, 0.727891156462585, 0.6793349168646081, 0.767156862745098, 0.6204986149584488, 0.48083623693379796, 0.7092511013215859, 0.6519721577726219, 0.2466216216216216, 0.30514705882352944, 0.7048832271762209, 0.6525612472160356, 0.16554054054054057, 0.3614864864864865, 0.08108108108108103, 0.8601516427969671, 0.8729096989966555, 0.8768558951965065, 0.9084444444444444, 0.8737113402061856, 0.8618925831202047, 0.8710382513661202, 0.8628318584070797, 0.8661971830985915, 0.8899755501222494, 0.8556581986143187, 0.847900113507378, 0.7370689655172413, 0.8624078624078624, 0.8754578754578755, 0.8973042362002568, 0.9131614654002713, 0.8991282689912827, 0.8886168910648715, 0.41097208854667955, 0.8052064631956912, 0.8584729981378026, 0.8704902867715079, 0.8780251694094869, 0.9077380952380952, 0.8766603415559773, 0.8646616541353384, 0.12223291626564003, 0.7501915708812261, 0.32894736842105265, 0.8348778433024431, 0.8027801911381407, 0.8374558303886925, 0.8757819481680071, 0.8328981723237598, 0.8161318300086731, 0.6547543075941289, 0.7099567099567099, 0.7675953079178885, 0.6941331575477917, 0.8340192043895748, 0.8382453735435229, 0.8714581893572909, 0.9001396648044693, 0.8604810996563574, 0.8495212038303693, 0.5843549328663165, 0.6827284105131415, 0.7370417193426043, 0.629695885509839, 0.296321998612075, 0.8210272873194222, 0.7900826446280992, 0.8277591973244147, 0.8610169491525423, 0.8167495854063018, 0.7998345740281224, 0.630461922596754, 0.6977067407922168, 0.7573839662447257, 0.6681701030927836, 0.07445887445887445, 0.296916890080429, 0.8360258481421647, 0.8092868988391376, 0.8414839797639123, 0.88, 0.8361064891846922, 0.8200663349917081, 0.6377358490566039, 0.7018284106891701, 0.7587437544610992, 0.6759740259740259, 0.051464063886424105, 0.2951153324287653, 0.024242424242424288}
	nPts := 14

	expected := [][]int{{12, 6, 7, 9, 10, 11, 13}, {1, 0, 2, 3, 4, 5}, {8}}

	res, err := Cluster(data, nPts)
	if err != nil {
		t.Fatalf("Clustering test failed!")
	}

	if compSlicesOfSlices(res, expected) {
		t.Log("Clustering test was successfully executed!")
	} else {
		t.Fatalf("Clustering test failed!")
	}
}