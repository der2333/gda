package gda

func Run(root string) error {
	rootSize, detailsInfo, err := GetDirSize(root)
	if err != nil {
		return err
	}
	buildResult(rootSize, detailsInfo)
	return err
}
