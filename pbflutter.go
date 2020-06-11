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

func ListView(children []*any.Any, padding *any.Any) *any.Any {
	return PackAny(&pb.PBListView{
		Children: children,
		Padding:  padding,
	})
}

func Karte(uuid string,
	preview, title *any.Any,
	gradientColor ...*any.Any) *any.Any {
	return PackAny(&pb.PBKarte{
		Uuid:          uuid,
		Title:         title,
		Preview:       preview,
		GradientColor: gradientColor,
	})

}

func ScreenWidth(multiplier, offset float32) *any.Any {
	return PackAny(&pb.PBScreenWidth{Offset: offset, Multiplier: multiplier})
}

func ScreenHeight(multiplier, offset float32) *any.Any {
	return PackAny(&pb.PBScreenHeight{Offset: offset, Multiplier: multiplier})
}

func RadiusElliptical(x, y float32) *any.Any {
	return PackAny(&pb.PBRadiusElliptical{X: x, Y: y})
}

func RadiusCircular(x float32) *any.Any {
	return PackAny(&pb.PBRadiusElliptical{X: x, Y: x})
}

func BoxDecoration(color, image, border, borderRadius, gradient, backgroundBlendMode, shape *any.Any, boxShadow []*any.Any) *any.Any {
	if shape == nil {
		shape = PackAny(&pb.PBBoxShape{BoxShape: pb.EnumBoxShape_rectangle})
	}
	return PackAny(&pb.PBBoxDecoration{
		Color: color, Image: image, Border: border, BorderRadius: borderRadius, Gradient: gradient, BackgroundBlendMode: backgroundBlendMode, Shape: shape, BoxShadow: boxShadow,
	})
}

func BlendMode(blendMode pb.EnumBlendMode) *any.Any {
	return PackAny(&pb.PBBlendMode{BlendMode: blendMode})
}

func TransformTranslate(offset, child *any.Any) *any.Any {
	return PackAny(&pb.PBTransformTranslate{Offset: offset, Child: child})
}

func TransformRotate(
	angle float32,
	origin,
	alignment,
	child *any.Any,
) *any.Any {
	return PackAny(&pb.PBTransformRotate{
		Alignment: alignment, Angle: angle, Child: child, Origin: origin,
	})
}

func TransformScale(
	scale float32,
	origin,
	alignment,
	child *any.Any,
) *any.Any {
	return PackAny(&pb.PBTransformScale{
		Alignment: alignment, Child: child, Origin: origin, Scale: scale,
	})

}

func LinearGradient(
	begin,
	end,
	tileMode,
	transform *any.Any,
	colors []*any.Any,
	stops []float32,
) *any.Any {

	if tileMode == nil {
		tileMode = TileMode(pb.EnumTileMode_clamp)
	}

	return PackAny(&pb.PBLinearGradient{
		Begin:     begin,
		End:       end,
		TileMode:  tileMode,
		Transform: transform,
		Colors:    colors,
		Stops:     stops,
	})
}

func Gradient(colors []*any.Any, stops []float32, transform *any.Any) *any.Any {
	return PackAny(&pb.PBGradient{
		Colors: colors, Stops: stops, Transform: transform,
	})
}

func DecorationImage(
	image,
	//        this.onError,
	colorFilter,
	fit,
	alignment,
	centerSlice,
	repeat *any.Any,
	//matchTextDirection
) *any.Any {

	if alignment == nil {
		alignment = PackAny(&pb.PBAlignment{X: 0.0, Y: 0.0})
	}

	if repeat == nil {
		repeat = PackAny(&pb.PBImageRepeat{Repeat: pb.EnumImageRepeat_noRepeat})
	}
	return PackAny(&pb.PBDecorationImage{
		Image: image, ColorFilter: colorFilter, Fit: fit, Alignment: alignment, CenterSlice: centerSlice, Repeat: repeat,
	})

}

func BoxFit(
	boxFit pb.EnumBoxFit) *any.Any {
	return PackAny(&pb.PBBoxFit{
		BoxFit: boxFit,
	})

}

func NetworkImage(url string, scale float32, headers map[string]string) *any.Any {
	return PackAny(&pb.PBNetworkImage{Url: url, Scale: scale, Headers: headers})
}

func Container(width, height, color, decoration, child *any.Any) *any.Any {
	return PackAny(&pb.PBContainer{Color: color, Width: width, Height: height, Child: child, Decoration: decoration})
}

func Alignment(x, y float32) *any.Any {
	return PackAny(&pb.PBAlignment{X: x, Y: y})
}

func TileMode(
	tileMode pb.EnumTileMode) *any.Any {
	return PackAny(&pb.PBTileMode{
		TileMode: tileMode,
	})

}

func BorderRadiusAll(radius *any.Any) *any.Any {
	return PackAny(&pb.PBBorderRadiusAll{Radius: radius})
}

func BorderRadiusOnly(topLeft, topRight, bottomLeft, bottomRight *any.Any) *any.Any {
	if topLeft == nil {
		topLeft = RadiusCircular(0.0)
	}

	if topRight == nil {
		topRight = RadiusCircular(0.0)
	}

	if bottomLeft == nil {
		bottomLeft = RadiusCircular(0.0)
	}

	if bottomRight == nil {
		bottomRight = RadiusCircular(0.0)
	}
	return PackAny(&pb.PBBorderRadiusOnly{TopLeft: topLeft, TopRight: topRight, BottomLeft: bottomLeft, BottomRight: bottomRight})
}

func Float(value float32) *any.Any {
	return PackAny(&pb.PBFloat{Value: value})
}

func Int(value int32) *any.Any {
	return PackAny(&pb.PBInt{Value: value})
}

func String(value string) *any.Any {
	return PackAny(&pb.PBString{Value: value})
}

func BorderAll(width, color, style *any.Any) *any.Any {
	return PackAny(&pb.PBBorderAll{Width: width, Color: color, Style: style})
}

func BorderStyle(borderStyle pb.EnumBorderStyle) *any.Any {
	/*	if borderStyle == nil {
		borderStyle = pb.EnumBorderStyle_BORDER_STYLE_NONE
	}*/
	return PackAny(&pb.PBBorderStyle{
		BorderStyle: borderStyle,
	})
}

func BoxShadow(color, offset, blurRadius, spreadRadius *any.Any) *any.Any {
	return PackAny(&pb.PBBoxShadow{Color: color, Offset: offset, BlurRadius: blurRadius, SpreadRadius: spreadRadius})
}
