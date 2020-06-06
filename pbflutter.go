package PBFlutter

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	pb "github.com/lemmerelassal/PBFlutter/proto"
	"google.golang.org/protobuf/runtime/protoiface"
)

func PackAny(message protoiface.MessageV1) *any.Any {
	msg, _ := ptypes.MarshalAny(message)
	return msg
}

func ColorFromRGBO(red, green, blue int32, opacity float32) *any.Any {
	return PackAny(&pb.PBColor{
		Red:     red,
		Green:   green,
		Blue:    blue,
		Opacity: opacity,
	})
}

func Color(orgb uint32) *any.Any {

	var opacity = float32(orgb>>24) / 255.0

	return ColorFromRGBO(
		int32((orgb>>16)&0xFF),
		int32((orgb>>8)&0xFF),
		int32((orgb>>0)&0xFF),
		opacity,
	)
}

func Text(text string, textStyle *any.Any) *any.Any {
	return PackAny(&pb.PBText{Text: text, TextStyle: textStyle})
}

func AcidText(text string, textStyle *any.Any) *any.Any {
	return PackAny(&pb.PBAcidText{Text: text, TextStyle: textStyle})
}

func TextStyle(fontFamily string, color *any.Any, fontSize float32, shadows []*any.Any) *any.Any {
	return PackAny(&pb.PBTextStyle{FontFamily: fontFamily, Color: color, FontSize: fontSize, Shadows: shadows})
}

func Shadow(offset, color *any.Any) *any.Any {
	return PackAny(&pb.PBShadow{Offset: offset, Color: color})
}

func Offset(x, y float32) *any.Any {
	return PackAny(&pb.PBOffset{X: x, Y: y})
}

func Column(children ...*any.Any) *any.Any {
	return PackAny(&pb.PBColumn{
		Children: children,
	})
}

func SVG(picture []byte) *any.Any {
	return PackAny(&pb.PBSVG{
		Picture: picture,
		Height:  285.0,
	})
}

func Align(child *any.Any, x, y float32) *any.Any {
	return PackAny(&pb.PBAlign{
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
	return PackAny(&pb.PBSize{
		Height: height,
		Width:  width,
	})
}

func PreferredSizeWidget(child, size *any.Any) *any.Any {
	return PackAny(&pb.PBPreferredSizeWidget{
		Child:         child,
		PreferredSize: size,
	})
}

func SliverAppBar(leading, backgroundColor, flexibleSpace, bottom *any.Any, expandedHeight, elevation float32, pinned, floating, forceElevated bool) *any.Any {
	return PackAny(&pb.PBSliverAppBar{
		BackgroundColor: backgroundColor,
		Bottom:          bottom,
		Elevation:       elevation,
		ExpandedHeight:  expandedHeight,
		FlexibleSpace:   flexibleSpace,
		Floating:        floating,
		ForceElevated:   forceElevated,
		Leading:         leading,
		Pinned:          pinned,
	})
}

func IconButton(icon *any.Any, onPressed string) *any.Any {
	return PackAny(&pb.PBIconButton{
		Icon:      icon,
		OnPressed: onPressed,
	})
}

func FlexibleSpaceBar(collapseMode pb.PBCollapseMode, background *any.Any) *any.Any {
	return PackAny(&pb.PBFlexibleSpaceBar{
		Background:   background,
		CollapseMode: collapseMode,
	})
}

func Icon(icon *any.Any, color *any.Any) *any.Any {
	return PackAny(&pb.PBIcon{Color: color, Icon: icon})
}

func Wrap(children ...*any.Any) *any.Any {
	return PackAny(&pb.PBWrap{
		Children: children,
	})
}

func Scaffold(body *any.Any) *any.Any {
	return PackAny(&pb.PBScaffold{
		Body: body,
	})
}

func Positioned(left, top, right, bottom, width, height float32, child *any.Any) *any.Any {
	return PackAny(&pb.PBPositioned{
		Left:   left,
		Top:    top,
		Right:  right,
		Bottom: bottom,
		Width:  width,
		Height: height,
		Child:  child,
	})
}

func EdgeInsetsGeometry(bottom, end, left, right, start, top float32) *any.Any {
	return PackAny(&pb.PBEdgeInsetsGeometry{

		Bottom: bottom,
		End:    end,
		Left:   left,
		Right:  right,
		Start:  start,
		Top:    top,
	})
}

func Padding(padding, child *any.Any) *any.Any {
	return PackAny(&pb.PBPadding{
		Padding: padding,
		Child:   child,
	})
}

func Stack(children ...*any.Any) *any.Any {
	return PackAny(&pb.PBStack{
		Children: children,
	})
}

func RawMaterialButton(constraints, child, fillColor, shape *any.Any, elevation float32) *any.Any {
	return PackAny(&pb.PBRawMaterialButton{
		Constraints: constraints,
		Child:       child,
		Shape:       shape,
		Elevation:   elevation,
	})
}

func CircleBorder() *any.Any {
	return PackAny(&pb.PBCircleBorder{})
}

func Shimmer(baseColor, highlightColor, child *any.Any) *any.Any {
	return PackAny(&pb.PBShimmer{
		BaseColor:      baseColor,
		HighlightColor: highlightColor,
		Child:          child,
	})
}

func ColorTween(begin, end *any.Any) *any.Any {
	return PackAny(&pb.PBColorTween{Begin: begin, End: end})
}

func AnimatedBackground(colorTweens ...*any.Any) *any.Any {
	return PackAny(&pb.PBAnimatedBackground{
		ColorTweens: colorTweens,
	})
}
