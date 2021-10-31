package example

/* 应该wrap errors
   遇到问题向上抛，既让开发者知道是什么问题，同时便于判定处理，干净明了 
*/
func WriteDao(w io.Write, buf []byte) error {
	_, err := w.Write(buf)
	return errors.Wrap(err, "write filed")
}

