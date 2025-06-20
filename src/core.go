package gda

func Run(root string) error {
	rootSize, detailsInfo, err := getDirSize(root)
	if err != nil {
		return err
	}
	buildResult(root, rootSize, detailsInfo)
	return nil
}
