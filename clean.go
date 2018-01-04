package main

// GetPkgbuild gets the pkgbuild of the package 'pkg' trying the ABS first and then the AUR trying the ABS first and then the AUR.

// RemoveMakeDeps receives a make dependency list and removes those
// that are no longer necessary.
func removeMakeDeps(depS []string) (err error) {
	hanging := sliceHangingPackages(depS)

	if len(hanging) != 0 {
		if !continueTask("Confirm Removal?", "nN") {
			return nil
		}
		err = cleanRemove(hanging)
	}

	return
}

// RemovePackage removes package from VCS information
func removeVCSPackage(pkgs []string) {
	for _, pkgName := range pkgs {
		for i, e := range savedInfo {
			if e.Package == pkgName {
				savedInfo[i] = savedInfo[len(savedInfo)-1]
				savedInfo = savedInfo[:len(savedInfo)-1]
			}
		}
	}

	_ = saveVCSInfo()
}

// CleanDependencies removes all dangling dependencies in system
func cleanDependencies() error {
	hanging, err := hangingPackages()
	if err != nil {
		return err
	}

	if len(hanging) != 0 {
		if !continueTask("Confirm Removal?", "nN") {
			return nil
		}
		err = cleanRemove(hanging)
	}

	return err
}

// CleanRemove sends a full removal command to pacman with the pkgName slice
func cleanRemove(pkgNames []string) (err error) {
	if len(pkgNames) == 0 {
		return nil
	}
	
	arguments := makeArguments()
	arguments.addArg("R", "s", "n", "s", "noconfirm")
	arguments.addTarget(pkgNames...)

	err = passToPacman(arguments)
	return err
}
