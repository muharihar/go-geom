package wkb

import (
	"testing"

	"github.com/twpayne/gogeom2/geom"
	. "launchpad.net/gocheck"
)

func Test(t *testing.T) { TestingT(t) }

type WKBSuite struct{}

var _ = Suite(&WKBSuite{})

func test(c *C, g geom.T, xdr []byte, ndr []byte) {

	got, err := Read(xdr)
	c.Check(err, IsNil)
	c.Check(got, Not(IsNil))
	c.Check(got, DeepEquals, g)

	gotXDR, err := Write(g, XDR)
	c.Check(err, IsNil)
	c.Check(gotXDR, DeepEquals, xdr)

	got, err = Read(ndr)
	c.Check(err, IsNil)
	c.Check(got, Not(IsNil))
	c.Check(got, DeepEquals, g)

	gotNDR, err := Write(g, NDR)
	c.Check(err, IsNil)
	c.Check(gotNDR, DeepEquals, ndr)

}

func (s *WKBSuite) TestPoint(c *C) {
	g, err := geom.NewPoint(geom.XY, []float64{1, 2})
	c.Assert(err, IsNil)
	c.Assert(g, Not(IsNil))

	xdr := []byte("\x00\x00\x00\x00\x01?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00")
	ndr := []byte("\x01\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@")
	test(c, g, xdr, ndr)
}

func (s *WKBSuite) TestPointZ(c *C) {
	g, err := geom.NewPoint(geom.XYZ, []float64{1, 2, 3})
	c.Assert(err, IsNil)
	c.Assert(g, Not(IsNil))
	xdr := []byte("\x00\x00\x00\x03\xe9?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00")
	ndr := []byte("\x01\xe9\x03\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@")
	test(c, g, xdr, ndr)
}

func (s *WKBSuite) TestPointM(c *C) {
	g, err := geom.NewPoint(geom.XYM, []float64{1, 2, 3})
	c.Assert(err, IsNil)
	c.Assert(g, Not(IsNil))
	xdr := []byte("\x00\x00\x00\x07\xd1?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00")
	ndr := []byte("\x01\xd1\x07\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@")
	test(c, g, xdr, ndr)
}

func (s *WKBSuite) TestPointZM(c *C) {
	g, err := geom.NewPoint(geom.XYZM, []float64{1, 2, 3, 4})
	c.Assert(err, IsNil)
	c.Assert(g, Not(IsNil))
	xdr := []byte("\x00\x00\x00\x0b\xb9?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00")
	ndr := []byte("\x01\xb9\x0b\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@")
	test(c, g, xdr, ndr)
}

func (s *WKBSuite) TestLineString(c *C) {
	g, err := geom.NewLineString(geom.XY, [][]float64{{1, 2}, {3, 4}})
	c.Assert(err, IsNil)
	c.Assert(g, Not(IsNil))
	xdr := []byte("\x00\x00\x00\x00\x02\x00\x00\x00\x02?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00")
	ndr := []byte("\x01\x02\x00\x00\x00\x02\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@")
	test(c, g, xdr, ndr)
}

func (s *WKBSuite) TestLineStringZ(c *C) {
	g, err := geom.NewLineString(geom.XYZ, [][]float64{{1, 2, 3}, {4, 5, 6}})
	c.Assert(err, IsNil)
	c.Assert(g, Not(IsNil))
	xdr := []byte("\x00\x00\x00\x03\xea\x00\x00\x00\x02?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00")
	ndr := []byte("\x01\xea\x03\x00\x00\x02\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@")
	test(c, g, xdr, ndr)
}

func (s *WKBSuite) TestLineStringM(c *C) {
	g, err := geom.NewLineString(geom.XYM, [][]float64{{1, 2, 3}, {4, 5, 6}})
	c.Assert(err, IsNil)
	c.Assert(g, Not(IsNil))
	xdr := []byte("\x00\x00\x00\x07\xd2\x00\x00\x00\x02?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00")
	ndr := []byte("\x01\xd2\x07\x00\x00\x02\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@")
	test(c, g, xdr, ndr)
}

func (s *WKBSuite) TestLineStringZM(c *C) {
	g, err := geom.NewLineString(geom.XYZM, [][]float64{{1, 2, 3, 4}, {5, 6, 7, 8}})
	c.Assert(err, IsNil)
	c.Assert(g, Not(IsNil))
	xdr := []byte("\x00\x00\x00\x0b\xba\x00\x00\x00\x02?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00@\x1c\x00\x00\x00\x00\x00\x00@ \x00\x00\x00\x00\x00\x00")
	ndr := []byte("\x01\xba\x0b\x00\x00\x02\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@\x00\x00\x00\x00\x00\x00\x1c@\x00\x00\x00\x00\x00\x00 @")
	test(c, g, xdr, ndr)
}

func (s *WKBSuite) TestPolygon(c *C) {
	g, err := geom.NewPolygon(geom.XY, [][][]float64{{{1, 2}, {3, 4}, {5, 6}, {1, 2}}})
	c.Assert(err, IsNil)
	c.Assert(g, Not(IsNil))
	xdr := []byte("\x00\x00\x00\x00\x03\x00\x00\x00\x01\x00\x00\x00\x04?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00")
	ndr := []byte("\x01\x03\x00\x00\x00\x01\x00\x00\x00\x04\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@")
	test(c, g, xdr, ndr)
}

func (s *WKBSuite) TestPolygonZ(c *C) {
	g, err := geom.NewPolygon(geom.XYZ, [][][]float64{{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {1, 2, 3}}})
	c.Assert(err, IsNil)
	c.Assert(g, Not(IsNil))
	xdr := []byte("\x00\x00\x00\x03\xeb\x00\x00\x00\x01\x00\x00\x00\x04?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00@\x1c\x00\x00\x00\x00\x00\x00@ \x00\x00\x00\x00\x00\x00@\"\x00\x00\x00\x00\x00\x00?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00")
	ndr := []byte("\x01\xeb\x03\x00\x00\x01\x00\x00\x00\x04\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@\x00\x00\x00\x00\x00\x00\x1c@\x00\x00\x00\x00\x00\x00 @\x00\x00\x00\x00\x00\x00\"@\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@")
	test(c, g, xdr, ndr)
}

func (s *WKBSuite) TestPolygonM(c *C) {
	g, err := geom.NewPolygon(geom.XYM, [][][]float64{{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {1, 2, 3}}})
	c.Assert(err, IsNil)
	c.Assert(g, Not(IsNil))
	xdr := []byte("\x00\x00\x00\x07\xd3\x00\x00\x00\x01\x00\x00\x00\x04?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00@\x1c\x00\x00\x00\x00\x00\x00@ \x00\x00\x00\x00\x00\x00@\"\x00\x00\x00\x00\x00\x00?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00")
	ndr := []byte("\x01\xd3\x07\x00\x00\x01\x00\x00\x00\x04\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@\x00\x00\x00\x00\x00\x00\x1c@\x00\x00\x00\x00\x00\x00 @\x00\x00\x00\x00\x00\x00\"@\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@")
	test(c, g, xdr, ndr)
}

func (s *WKBSuite) TestPolygonZM(c *C) {
	g, err := geom.NewPolygon(geom.XYZM, [][][]float64{{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {1, 2, 3, 4}}})
	c.Assert(err, IsNil)
	c.Assert(g, Not(IsNil))
	xdr := []byte("\x00\x00\x00\x0b\xbb\x00\x00\x00\x01\x00\x00\x00\x04?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00@\x1c\x00\x00\x00\x00\x00\x00@ \x00\x00\x00\x00\x00\x00@\"\x00\x00\x00\x00\x00\x00@$\x00\x00\x00\x00\x00\x00@&\x00\x00\x00\x00\x00\x00@(\x00\x00\x00\x00\x00\x00?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00")
	ndr := []byte("\x01\xbb\x0b\x00\x00\x01\x00\x00\x00\x04\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@\x00\x00\x00\x00\x00\x00\x1c@\x00\x00\x00\x00\x00\x00 @\x00\x00\x00\x00\x00\x00\"@\x00\x00\x00\x00\x00\x00$@\x00\x00\x00\x00\x00\x00&@\x00\x00\x00\x00\x00\x00(@\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@")
	test(c, g, xdr, ndr)
}

func (s *WKBSuite) TestMultiPoint(c *C) {
	g, err := geom.NewMultiPoint(geom.XY, [][]float64{{1, 2}, {3, 4}})
	c.Assert(err, IsNil)
	c.Assert(g, Not(IsNil))
	xdr := []byte("\x00\x00\x00\x00\x04\x00\x00\x00\x02\x00\x00\x00\x00\x01?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00")
	ndr := []byte("\x01\x04\x00\x00\x00\x02\x00\x00\x00\x01\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x01\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@")
	test(c, g, xdr, ndr)
}

func (s *WKBSuite) TestMultiPointZ(c *C) {
	g, err := geom.NewMultiPoint(geom.XYZ, [][]float64{{1, 2, 3}, {4, 5, 6}})
	c.Assert(err, IsNil)
	c.Assert(g, Not(IsNil))
	xdr := []byte("\x00\x00\x00\x03\xec\x00\x00\x00\x02\x00\x00\x00\x03\xe9?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00\x00\x00\x00\x03\xe9@\x10\x00\x00\x00\x00\x00\x00@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00")
	ndr := []byte("\x01\xec\x03\x00\x00\x02\x00\x00\x00\x01\xe9\x03\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x01\xe9\x03\x00\x00\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@")
	test(c, g, xdr, ndr)
}

func (s *WKBSuite) TestMultiPointM(c *C) {
	g, err := geom.NewMultiPoint(geom.XYM, [][]float64{{1, 2, 3}, {4, 5, 6}})
	xdr := []byte("\x00\x00\x00\x07\xd4\x00\x00\x00\x02\x00\x00\x00\x07\xd1?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00\x00\x00\x00\x07\xd1@\x10\x00\x00\x00\x00\x00\x00@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00")
	ndr := []byte("\x01\xd4\x07\x00\x00\x02\x00\x00\x00\x01\xd1\x07\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x01\xd1\x07\x00\x00\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@")
	c.Assert(err, IsNil)
	c.Assert(g, Not(IsNil))
	test(c, g, xdr, ndr)
}

func (s *WKBSuite) TestMultiPointZM(c *C) {
	g, err := geom.NewMultiPoint(geom.XYZM, [][]float64{{1, 2, 3, 4}, {5, 6, 7, 8}})
	c.Assert(err, IsNil)
	c.Assert(g, Not(IsNil))
	xdr := []byte("\x00\x00\x00\x0b\xbc\x00\x00\x00\x02\x00\x00\x00\x0b\xb9?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0b\xb9@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00@\x1c\x00\x00\x00\x00\x00\x00@ \x00\x00\x00\x00\x00\x00")
	ndr := []byte("\x01\xbc\x0b\x00\x00\x02\x00\x00\x00\x01\xb9\x0b\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@\x01\xb9\x0b\x00\x00\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@\x00\x00\x00\x00\x00\x00\x1c@\x00\x00\x00\x00\x00\x00 @")
	test(c, g, xdr, ndr)
}
