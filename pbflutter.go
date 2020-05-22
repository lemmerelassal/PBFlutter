package PBFlutter

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	pb "github.com/lemmerelassal/PBFlutter/proto"
	"google.golang.org/protobuf/runtime/protoiface"
)

func packAny(message protoiface.MessageV1) *any.Any {
	msg, _ := ptypes.MarshalAny(message)
	return msg
}

func Text(text string, fontSize float32, colour *pb.PBColor) *any.Any {
	if colour == nil {
		colour = &pb.PBColor{
			Red: 0, Green: 0, Blue: 0, Opacity: 1.0,
		}
	}
	return packAny(&pb.PBText{Text: text, FontSize: fontSize, Colour: colour})
}

func Column(children []*any.Any) *any.Any {
	return packAny(&pb.PBColumn{
		Children: children,
	})
}

func SVG(picture []byte) *any.Any {
	return packAny(&pb.PBSVG{
		Picture: picture,
		Height:  285.0,
	})
}

func Align(child *any.Any, x, y float32) *any.Any {
	return packAny(&pb.PBAlign{
		Child:  child,
		AlignX: x,
		AlignY: y,
	})
}

func Children(children ...*any.Any) []*any.Any {
	res := []*any.Any{}
	for _, child := range children {
		res = append(res, child)
	}
	return res
}

func Size(width, height float32) *any.Any {
	return packAny(&pb.PBSize{
		Height: height,
		Width:  width,
	})
}

func PreferredSizeWidget(child *any.Any, size *pb.PBSize) *any.Any {
	return packAny(&pb.PBPreferredSizeWidget{
		Child:         child,
		PreferredSize: size,
	})
}

func SliverAppBar(leading *any.Any, backgroundColor *pb.PBColor, flexibleSpace *any.Any, expandedHeight, elevation float32, pinned, floating, forceElevated bool, bottom *pb.PreferredSizeWidget) { // many params

}

func IconButton(icon *pb.PBIcon, onPressed string) {

}

func FlexibleSpaceBar(collapseMode pb.PBCollapseMode, background *any.Any) {

}

func Icon(icon *pb.PBIconData, color *pb.PBColor) {

}

func Wrap(children []*any.Any) {

}

func Scaffold(body *any.Any) {

}
