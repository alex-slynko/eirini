// Code generated by counterfeiter. DO NOT EDIT.
package bufio_fake

import (
	"bufio"
	"io"
	"sync"

	"code.cloudfoundry.org/goshims/bufioshim"
)

type FakeBufio struct {
	NewReaderSizeStub        func(rd io.Reader, size int) bufioshim.Reader
	newReaderSizeMutex       sync.RWMutex
	newReaderSizeArgsForCall []struct {
		rd   io.Reader
		size int
	}
	newReaderSizeReturns struct {
		result1 bufioshim.Reader
	}
	newReaderSizeReturnsOnCall map[int]struct {
		result1 bufioshim.Reader
	}
	NewReaderStub        func(rd io.Reader) bufioshim.Reader
	newReaderMutex       sync.RWMutex
	newReaderArgsForCall []struct {
		rd io.Reader
	}
	newReaderReturns struct {
		result1 bufioshim.Reader
	}
	newReaderReturnsOnCall map[int]struct {
		result1 bufioshim.Reader
	}
	NewWriterSizeStub        func(w io.Writer, size int) *bufio.Writer
	newWriterSizeMutex       sync.RWMutex
	newWriterSizeArgsForCall []struct {
		w    io.Writer
		size int
	}
	newWriterSizeReturns struct {
		result1 *bufio.Writer
	}
	newWriterSizeReturnsOnCall map[int]struct {
		result1 *bufio.Writer
	}
	NewWriterStub        func(w io.Writer) *bufio.Writer
	newWriterMutex       sync.RWMutex
	newWriterArgsForCall []struct {
		w io.Writer
	}
	newWriterReturns struct {
		result1 *bufio.Writer
	}
	newWriterReturnsOnCall map[int]struct {
		result1 *bufio.Writer
	}
	NewReadWriterStub        func(r *bufio.Reader, w *bufio.Writer) *bufio.ReadWriter
	newReadWriterMutex       sync.RWMutex
	newReadWriterArgsForCall []struct {
		r *bufio.Reader
		w *bufio.Writer
	}
	newReadWriterReturns struct {
		result1 *bufio.ReadWriter
	}
	newReadWriterReturnsOnCall map[int]struct {
		result1 *bufio.ReadWriter
	}
	NewScannerStub        func(r io.Reader) *bufio.Scanner
	newScannerMutex       sync.RWMutex
	newScannerArgsForCall []struct {
		r io.Reader
	}
	newScannerReturns struct {
		result1 *bufio.Scanner
	}
	newScannerReturnsOnCall map[int]struct {
		result1 *bufio.Scanner
	}
	ScanBytesStub        func(data []byte, atEOF bool) (advance int, token []byte, err error)
	scanBytesMutex       sync.RWMutex
	scanBytesArgsForCall []struct {
		data  []byte
		atEOF bool
	}
	scanBytesReturns struct {
		result1 int
		result2 []byte
		result3 error
	}
	scanBytesReturnsOnCall map[int]struct {
		result1 int
		result2 []byte
		result3 error
	}
	ScanRunesStub        func(data []byte, atEOF bool) (advance int, token []byte, err error)
	scanRunesMutex       sync.RWMutex
	scanRunesArgsForCall []struct {
		data  []byte
		atEOF bool
	}
	scanRunesReturns struct {
		result1 int
		result2 []byte
		result3 error
	}
	scanRunesReturnsOnCall map[int]struct {
		result1 int
		result2 []byte
		result3 error
	}
	ScanLinesStub        func(data []byte, atEOF bool) (advance int, token []byte, err error)
	scanLinesMutex       sync.RWMutex
	scanLinesArgsForCall []struct {
		data  []byte
		atEOF bool
	}
	scanLinesReturns struct {
		result1 int
		result2 []byte
		result3 error
	}
	scanLinesReturnsOnCall map[int]struct {
		result1 int
		result2 []byte
		result3 error
	}
	ScanWordsStub        func(data []byte, atEOF bool) (advance int, token []byte, err error)
	scanWordsMutex       sync.RWMutex
	scanWordsArgsForCall []struct {
		data  []byte
		atEOF bool
	}
	scanWordsReturns struct {
		result1 int
		result2 []byte
		result3 error
	}
	scanWordsReturnsOnCall map[int]struct {
		result1 int
		result2 []byte
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBufio) NewReaderSize(rd io.Reader, size int) bufioshim.Reader {
	fake.newReaderSizeMutex.Lock()
	ret, specificReturn := fake.newReaderSizeReturnsOnCall[len(fake.newReaderSizeArgsForCall)]
	fake.newReaderSizeArgsForCall = append(fake.newReaderSizeArgsForCall, struct {
		rd   io.Reader
		size int
	}{rd, size})
	fake.recordInvocation("NewReaderSize", []interface{}{rd, size})
	fake.newReaderSizeMutex.Unlock()
	if fake.NewReaderSizeStub != nil {
		return fake.NewReaderSizeStub(rd, size)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.newReaderSizeReturns.result1
}

func (fake *FakeBufio) NewReaderSizeCallCount() int {
	fake.newReaderSizeMutex.RLock()
	defer fake.newReaderSizeMutex.RUnlock()
	return len(fake.newReaderSizeArgsForCall)
}

func (fake *FakeBufio) NewReaderSizeArgsForCall(i int) (io.Reader, int) {
	fake.newReaderSizeMutex.RLock()
	defer fake.newReaderSizeMutex.RUnlock()
	return fake.newReaderSizeArgsForCall[i].rd, fake.newReaderSizeArgsForCall[i].size
}

func (fake *FakeBufio) NewReaderSizeReturns(result1 bufioshim.Reader) {
	fake.NewReaderSizeStub = nil
	fake.newReaderSizeReturns = struct {
		result1 bufioshim.Reader
	}{result1}
}

func (fake *FakeBufio) NewReaderSizeReturnsOnCall(i int, result1 bufioshim.Reader) {
	fake.NewReaderSizeStub = nil
	if fake.newReaderSizeReturnsOnCall == nil {
		fake.newReaderSizeReturnsOnCall = make(map[int]struct {
			result1 bufioshim.Reader
		})
	}
	fake.newReaderSizeReturnsOnCall[i] = struct {
		result1 bufioshim.Reader
	}{result1}
}

func (fake *FakeBufio) NewReader(rd io.Reader) bufioshim.Reader {
	fake.newReaderMutex.Lock()
	ret, specificReturn := fake.newReaderReturnsOnCall[len(fake.newReaderArgsForCall)]
	fake.newReaderArgsForCall = append(fake.newReaderArgsForCall, struct {
		rd io.Reader
	}{rd})
	fake.recordInvocation("NewReader", []interface{}{rd})
	fake.newReaderMutex.Unlock()
	if fake.NewReaderStub != nil {
		return fake.NewReaderStub(rd)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.newReaderReturns.result1
}

func (fake *FakeBufio) NewReaderCallCount() int {
	fake.newReaderMutex.RLock()
	defer fake.newReaderMutex.RUnlock()
	return len(fake.newReaderArgsForCall)
}

func (fake *FakeBufio) NewReaderArgsForCall(i int) io.Reader {
	fake.newReaderMutex.RLock()
	defer fake.newReaderMutex.RUnlock()
	return fake.newReaderArgsForCall[i].rd
}

func (fake *FakeBufio) NewReaderReturns(result1 bufioshim.Reader) {
	fake.NewReaderStub = nil
	fake.newReaderReturns = struct {
		result1 bufioshim.Reader
	}{result1}
}

func (fake *FakeBufio) NewReaderReturnsOnCall(i int, result1 bufioshim.Reader) {
	fake.NewReaderStub = nil
	if fake.newReaderReturnsOnCall == nil {
		fake.newReaderReturnsOnCall = make(map[int]struct {
			result1 bufioshim.Reader
		})
	}
	fake.newReaderReturnsOnCall[i] = struct {
		result1 bufioshim.Reader
	}{result1}
}

func (fake *FakeBufio) NewWriterSize(w io.Writer, size int) *bufio.Writer {
	fake.newWriterSizeMutex.Lock()
	ret, specificReturn := fake.newWriterSizeReturnsOnCall[len(fake.newWriterSizeArgsForCall)]
	fake.newWriterSizeArgsForCall = append(fake.newWriterSizeArgsForCall, struct {
		w    io.Writer
		size int
	}{w, size})
	fake.recordInvocation("NewWriterSize", []interface{}{w, size})
	fake.newWriterSizeMutex.Unlock()
	if fake.NewWriterSizeStub != nil {
		return fake.NewWriterSizeStub(w, size)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.newWriterSizeReturns.result1
}

func (fake *FakeBufio) NewWriterSizeCallCount() int {
	fake.newWriterSizeMutex.RLock()
	defer fake.newWriterSizeMutex.RUnlock()
	return len(fake.newWriterSizeArgsForCall)
}

func (fake *FakeBufio) NewWriterSizeArgsForCall(i int) (io.Writer, int) {
	fake.newWriterSizeMutex.RLock()
	defer fake.newWriterSizeMutex.RUnlock()
	return fake.newWriterSizeArgsForCall[i].w, fake.newWriterSizeArgsForCall[i].size
}

func (fake *FakeBufio) NewWriterSizeReturns(result1 *bufio.Writer) {
	fake.NewWriterSizeStub = nil
	fake.newWriterSizeReturns = struct {
		result1 *bufio.Writer
	}{result1}
}

func (fake *FakeBufio) NewWriterSizeReturnsOnCall(i int, result1 *bufio.Writer) {
	fake.NewWriterSizeStub = nil
	if fake.newWriterSizeReturnsOnCall == nil {
		fake.newWriterSizeReturnsOnCall = make(map[int]struct {
			result1 *bufio.Writer
		})
	}
	fake.newWriterSizeReturnsOnCall[i] = struct {
		result1 *bufio.Writer
	}{result1}
}

func (fake *FakeBufio) NewWriter(w io.Writer) *bufio.Writer {
	fake.newWriterMutex.Lock()
	ret, specificReturn := fake.newWriterReturnsOnCall[len(fake.newWriterArgsForCall)]
	fake.newWriterArgsForCall = append(fake.newWriterArgsForCall, struct {
		w io.Writer
	}{w})
	fake.recordInvocation("NewWriter", []interface{}{w})
	fake.newWriterMutex.Unlock()
	if fake.NewWriterStub != nil {
		return fake.NewWriterStub(w)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.newWriterReturns.result1
}

func (fake *FakeBufio) NewWriterCallCount() int {
	fake.newWriterMutex.RLock()
	defer fake.newWriterMutex.RUnlock()
	return len(fake.newWriterArgsForCall)
}

func (fake *FakeBufio) NewWriterArgsForCall(i int) io.Writer {
	fake.newWriterMutex.RLock()
	defer fake.newWriterMutex.RUnlock()
	return fake.newWriterArgsForCall[i].w
}

func (fake *FakeBufio) NewWriterReturns(result1 *bufio.Writer) {
	fake.NewWriterStub = nil
	fake.newWriterReturns = struct {
		result1 *bufio.Writer
	}{result1}
}

func (fake *FakeBufio) NewWriterReturnsOnCall(i int, result1 *bufio.Writer) {
	fake.NewWriterStub = nil
	if fake.newWriterReturnsOnCall == nil {
		fake.newWriterReturnsOnCall = make(map[int]struct {
			result1 *bufio.Writer
		})
	}
	fake.newWriterReturnsOnCall[i] = struct {
		result1 *bufio.Writer
	}{result1}
}

func (fake *FakeBufio) NewReadWriter(r *bufio.Reader, w *bufio.Writer) *bufio.ReadWriter {
	fake.newReadWriterMutex.Lock()
	ret, specificReturn := fake.newReadWriterReturnsOnCall[len(fake.newReadWriterArgsForCall)]
	fake.newReadWriterArgsForCall = append(fake.newReadWriterArgsForCall, struct {
		r *bufio.Reader
		w *bufio.Writer
	}{r, w})
	fake.recordInvocation("NewReadWriter", []interface{}{r, w})
	fake.newReadWriterMutex.Unlock()
	if fake.NewReadWriterStub != nil {
		return fake.NewReadWriterStub(r, w)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.newReadWriterReturns.result1
}

func (fake *FakeBufio) NewReadWriterCallCount() int {
	fake.newReadWriterMutex.RLock()
	defer fake.newReadWriterMutex.RUnlock()
	return len(fake.newReadWriterArgsForCall)
}

func (fake *FakeBufio) NewReadWriterArgsForCall(i int) (*bufio.Reader, *bufio.Writer) {
	fake.newReadWriterMutex.RLock()
	defer fake.newReadWriterMutex.RUnlock()
	return fake.newReadWriterArgsForCall[i].r, fake.newReadWriterArgsForCall[i].w
}

func (fake *FakeBufio) NewReadWriterReturns(result1 *bufio.ReadWriter) {
	fake.NewReadWriterStub = nil
	fake.newReadWriterReturns = struct {
		result1 *bufio.ReadWriter
	}{result1}
}

func (fake *FakeBufio) NewReadWriterReturnsOnCall(i int, result1 *bufio.ReadWriter) {
	fake.NewReadWriterStub = nil
	if fake.newReadWriterReturnsOnCall == nil {
		fake.newReadWriterReturnsOnCall = make(map[int]struct {
			result1 *bufio.ReadWriter
		})
	}
	fake.newReadWriterReturnsOnCall[i] = struct {
		result1 *bufio.ReadWriter
	}{result1}
}

func (fake *FakeBufio) NewScanner(r io.Reader) *bufio.Scanner {
	fake.newScannerMutex.Lock()
	ret, specificReturn := fake.newScannerReturnsOnCall[len(fake.newScannerArgsForCall)]
	fake.newScannerArgsForCall = append(fake.newScannerArgsForCall, struct {
		r io.Reader
	}{r})
	fake.recordInvocation("NewScanner", []interface{}{r})
	fake.newScannerMutex.Unlock()
	if fake.NewScannerStub != nil {
		return fake.NewScannerStub(r)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.newScannerReturns.result1
}

func (fake *FakeBufio) NewScannerCallCount() int {
	fake.newScannerMutex.RLock()
	defer fake.newScannerMutex.RUnlock()
	return len(fake.newScannerArgsForCall)
}

func (fake *FakeBufio) NewScannerArgsForCall(i int) io.Reader {
	fake.newScannerMutex.RLock()
	defer fake.newScannerMutex.RUnlock()
	return fake.newScannerArgsForCall[i].r
}

func (fake *FakeBufio) NewScannerReturns(result1 *bufio.Scanner) {
	fake.NewScannerStub = nil
	fake.newScannerReturns = struct {
		result1 *bufio.Scanner
	}{result1}
}

func (fake *FakeBufio) NewScannerReturnsOnCall(i int, result1 *bufio.Scanner) {
	fake.NewScannerStub = nil
	if fake.newScannerReturnsOnCall == nil {
		fake.newScannerReturnsOnCall = make(map[int]struct {
			result1 *bufio.Scanner
		})
	}
	fake.newScannerReturnsOnCall[i] = struct {
		result1 *bufio.Scanner
	}{result1}
}

func (fake *FakeBufio) ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error) {
	var dataCopy []byte
	if data != nil {
		dataCopy = make([]byte, len(data))
		copy(dataCopy, data)
	}
	fake.scanBytesMutex.Lock()
	ret, specificReturn := fake.scanBytesReturnsOnCall[len(fake.scanBytesArgsForCall)]
	fake.scanBytesArgsForCall = append(fake.scanBytesArgsForCall, struct {
		data  []byte
		atEOF bool
	}{dataCopy, atEOF})
	fake.recordInvocation("ScanBytes", []interface{}{dataCopy, atEOF})
	fake.scanBytesMutex.Unlock()
	if fake.ScanBytesStub != nil {
		return fake.ScanBytesStub(data, atEOF)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.scanBytesReturns.result1, fake.scanBytesReturns.result2, fake.scanBytesReturns.result3
}

func (fake *FakeBufio) ScanBytesCallCount() int {
	fake.scanBytesMutex.RLock()
	defer fake.scanBytesMutex.RUnlock()
	return len(fake.scanBytesArgsForCall)
}

func (fake *FakeBufio) ScanBytesArgsForCall(i int) ([]byte, bool) {
	fake.scanBytesMutex.RLock()
	defer fake.scanBytesMutex.RUnlock()
	return fake.scanBytesArgsForCall[i].data, fake.scanBytesArgsForCall[i].atEOF
}

func (fake *FakeBufio) ScanBytesReturns(result1 int, result2 []byte, result3 error) {
	fake.ScanBytesStub = nil
	fake.scanBytesReturns = struct {
		result1 int
		result2 []byte
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBufio) ScanBytesReturnsOnCall(i int, result1 int, result2 []byte, result3 error) {
	fake.ScanBytesStub = nil
	if fake.scanBytesReturnsOnCall == nil {
		fake.scanBytesReturnsOnCall = make(map[int]struct {
			result1 int
			result2 []byte
			result3 error
		})
	}
	fake.scanBytesReturnsOnCall[i] = struct {
		result1 int
		result2 []byte
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBufio) ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error) {
	var dataCopy []byte
	if data != nil {
		dataCopy = make([]byte, len(data))
		copy(dataCopy, data)
	}
	fake.scanRunesMutex.Lock()
	ret, specificReturn := fake.scanRunesReturnsOnCall[len(fake.scanRunesArgsForCall)]
	fake.scanRunesArgsForCall = append(fake.scanRunesArgsForCall, struct {
		data  []byte
		atEOF bool
	}{dataCopy, atEOF})
	fake.recordInvocation("ScanRunes", []interface{}{dataCopy, atEOF})
	fake.scanRunesMutex.Unlock()
	if fake.ScanRunesStub != nil {
		return fake.ScanRunesStub(data, atEOF)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.scanRunesReturns.result1, fake.scanRunesReturns.result2, fake.scanRunesReturns.result3
}

func (fake *FakeBufio) ScanRunesCallCount() int {
	fake.scanRunesMutex.RLock()
	defer fake.scanRunesMutex.RUnlock()
	return len(fake.scanRunesArgsForCall)
}

func (fake *FakeBufio) ScanRunesArgsForCall(i int) ([]byte, bool) {
	fake.scanRunesMutex.RLock()
	defer fake.scanRunesMutex.RUnlock()
	return fake.scanRunesArgsForCall[i].data, fake.scanRunesArgsForCall[i].atEOF
}

func (fake *FakeBufio) ScanRunesReturns(result1 int, result2 []byte, result3 error) {
	fake.ScanRunesStub = nil
	fake.scanRunesReturns = struct {
		result1 int
		result2 []byte
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBufio) ScanRunesReturnsOnCall(i int, result1 int, result2 []byte, result3 error) {
	fake.ScanRunesStub = nil
	if fake.scanRunesReturnsOnCall == nil {
		fake.scanRunesReturnsOnCall = make(map[int]struct {
			result1 int
			result2 []byte
			result3 error
		})
	}
	fake.scanRunesReturnsOnCall[i] = struct {
		result1 int
		result2 []byte
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBufio) ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	var dataCopy []byte
	if data != nil {
		dataCopy = make([]byte, len(data))
		copy(dataCopy, data)
	}
	fake.scanLinesMutex.Lock()
	ret, specificReturn := fake.scanLinesReturnsOnCall[len(fake.scanLinesArgsForCall)]
	fake.scanLinesArgsForCall = append(fake.scanLinesArgsForCall, struct {
		data  []byte
		atEOF bool
	}{dataCopy, atEOF})
	fake.recordInvocation("ScanLines", []interface{}{dataCopy, atEOF})
	fake.scanLinesMutex.Unlock()
	if fake.ScanLinesStub != nil {
		return fake.ScanLinesStub(data, atEOF)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.scanLinesReturns.result1, fake.scanLinesReturns.result2, fake.scanLinesReturns.result3
}

func (fake *FakeBufio) ScanLinesCallCount() int {
	fake.scanLinesMutex.RLock()
	defer fake.scanLinesMutex.RUnlock()
	return len(fake.scanLinesArgsForCall)
}

func (fake *FakeBufio) ScanLinesArgsForCall(i int) ([]byte, bool) {
	fake.scanLinesMutex.RLock()
	defer fake.scanLinesMutex.RUnlock()
	return fake.scanLinesArgsForCall[i].data, fake.scanLinesArgsForCall[i].atEOF
}

func (fake *FakeBufio) ScanLinesReturns(result1 int, result2 []byte, result3 error) {
	fake.ScanLinesStub = nil
	fake.scanLinesReturns = struct {
		result1 int
		result2 []byte
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBufio) ScanLinesReturnsOnCall(i int, result1 int, result2 []byte, result3 error) {
	fake.ScanLinesStub = nil
	if fake.scanLinesReturnsOnCall == nil {
		fake.scanLinesReturnsOnCall = make(map[int]struct {
			result1 int
			result2 []byte
			result3 error
		})
	}
	fake.scanLinesReturnsOnCall[i] = struct {
		result1 int
		result2 []byte
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBufio) ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
	var dataCopy []byte
	if data != nil {
		dataCopy = make([]byte, len(data))
		copy(dataCopy, data)
	}
	fake.scanWordsMutex.Lock()
	ret, specificReturn := fake.scanWordsReturnsOnCall[len(fake.scanWordsArgsForCall)]
	fake.scanWordsArgsForCall = append(fake.scanWordsArgsForCall, struct {
		data  []byte
		atEOF bool
	}{dataCopy, atEOF})
	fake.recordInvocation("ScanWords", []interface{}{dataCopy, atEOF})
	fake.scanWordsMutex.Unlock()
	if fake.ScanWordsStub != nil {
		return fake.ScanWordsStub(data, atEOF)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.scanWordsReturns.result1, fake.scanWordsReturns.result2, fake.scanWordsReturns.result3
}

func (fake *FakeBufio) ScanWordsCallCount() int {
	fake.scanWordsMutex.RLock()
	defer fake.scanWordsMutex.RUnlock()
	return len(fake.scanWordsArgsForCall)
}

func (fake *FakeBufio) ScanWordsArgsForCall(i int) ([]byte, bool) {
	fake.scanWordsMutex.RLock()
	defer fake.scanWordsMutex.RUnlock()
	return fake.scanWordsArgsForCall[i].data, fake.scanWordsArgsForCall[i].atEOF
}

func (fake *FakeBufio) ScanWordsReturns(result1 int, result2 []byte, result3 error) {
	fake.ScanWordsStub = nil
	fake.scanWordsReturns = struct {
		result1 int
		result2 []byte
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBufio) ScanWordsReturnsOnCall(i int, result1 int, result2 []byte, result3 error) {
	fake.ScanWordsStub = nil
	if fake.scanWordsReturnsOnCall == nil {
		fake.scanWordsReturnsOnCall = make(map[int]struct {
			result1 int
			result2 []byte
			result3 error
		})
	}
	fake.scanWordsReturnsOnCall[i] = struct {
		result1 int
		result2 []byte
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBufio) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newReaderSizeMutex.RLock()
	defer fake.newReaderSizeMutex.RUnlock()
	fake.newReaderMutex.RLock()
	defer fake.newReaderMutex.RUnlock()
	fake.newWriterSizeMutex.RLock()
	defer fake.newWriterSizeMutex.RUnlock()
	fake.newWriterMutex.RLock()
	defer fake.newWriterMutex.RUnlock()
	fake.newReadWriterMutex.RLock()
	defer fake.newReadWriterMutex.RUnlock()
	fake.newScannerMutex.RLock()
	defer fake.newScannerMutex.RUnlock()
	fake.scanBytesMutex.RLock()
	defer fake.scanBytesMutex.RUnlock()
	fake.scanRunesMutex.RLock()
	defer fake.scanRunesMutex.RUnlock()
	fake.scanLinesMutex.RLock()
	defer fake.scanLinesMutex.RUnlock()
	fake.scanWordsMutex.RLock()
	defer fake.scanWordsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBufio) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ bufioshim.Bufio = new(FakeBufio)
