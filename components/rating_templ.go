// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"context"
	"github.com/axzilla/templui/icons"
	"github.com/axzilla/templui/utils"
	"strconv"
)

type RatingStyle string

const (
	RatingStyleStar  RatingStyle = "star"
	RatingStyleHeart RatingStyle = "heart"
	RatingStyleEmoji RatingStyle = "emoji"
)

type RatingProps struct {
	ID          string
	Class       string
	Attributes  templ.Attributes
	Value       float64
	ReadOnly    bool
	Precision   float64
	Name        string
	OnlyInteger bool
}

type RatingGroupProps struct {
	ID         string
	Class      string
	Attributes templ.Attributes
}

type RatingItemProps struct {
	ID         string
	Class      string
	Attributes templ.Attributes
	Value      int
	Style      RatingStyle
}

type RatingLabelProps struct {
	ID         string
	Class      string
	Attributes templ.Attributes
}

type RatingValueProps struct {
	ID         string
	Class      string
	Attributes templ.Attributes
}

func Rating(props RatingProps) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		props.setDefaults()
		ctx = context.WithValue(ctx, "readOnly", props.ReadOnly)
		var templ_7745c5c3_Var2 = []any{utils.TwMerge(
			"flex flex-col items-start gap-1",
			props.Class,
		)}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div x-data=\"rating\" data-value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.FormatFloat(props.Value, 'f', -1, 64))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/rating.templ`, Line: 60, Col: 60}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "\" data-precision=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.FormatFloat(props.Precision, 'f', -1, 64))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/rating.templ`, Line: 61, Col: 68}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "\" data-readonly=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.FormatBool(props.ReadOnly))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/rating.templ`, Line: 62, Col: 52}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "\" data-name=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(props.Name)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/rating.templ`, Line: 63, Col: 24}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "\" data-onlyinteger=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.FormatBool(props.OnlyInteger))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/rating.templ`, Line: 64, Col: 58}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var2).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/rating.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, props.Attributes)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, ">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "<input type=\"hidden\" name=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var9 string
		templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(props.Name)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/rating.templ`, Line: 74, Col: 20}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "\" x-bind:value=\"value\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.Name == "" {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, " disabled")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, " x-ref=\"input\"></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func RatingGroup(props RatingGroupProps) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var10 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var10 == nil {
			templ_7745c5c3_Var10 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var templ_7745c5c3_Var11 = []any{utils.TwMerge("flex items-center gap-1", props.Class)}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var11...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, "<div class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var12 string
		templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var11).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/rating.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 14, "\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, props.Attributes)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 15, ">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var10.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 16, "</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func RatingItem(props RatingItemProps) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var13 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var13 == nil {
			templ_7745c5c3_Var13 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		props.setDefaults()
		readOnly, _ := ctx.Value("readOnly").(bool)
		var templ_7745c5c3_Var14 = []any{
			utils.TwMerge(
				"relative",
				getColorClass(props.Style),
				"transition-opacity",
				"cursor-pointer",
				utils.If(readOnly, "cursor-default"),
				props.Class,
			),
		}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var14...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 17, "<div class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var15 string
		templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var14).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/rating.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 18, "\" data-rating-value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var16 string
		templ_7745c5c3_Var16, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(props.Value))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/rating.templ`, Line: 105, Col: 47}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var16))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 19, "\" @click=\"setValue\" @mouseover=\"hover\" @mouseleave=\"resetPreview\"><div class=\"opacity-30\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = getRatingIcon(props.Style, false, float64(props.Value)).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 20, "</div><div class=\"absolute inset-0 overflow-hidden\" x-bind:style=\"getItemStyle\" data-index=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var17 string
		templ_7745c5c3_Var17, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(props.Value))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/rating.templ`, Line: 116, Col: 41}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var17))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 21, "\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = getRatingIcon(props.Style, true, float64(props.Value)).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 22, "</div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func getColorClass(style RatingStyle) string {
	switch style {
	case RatingStyleHeart:
		return "text-destructive"
	case RatingStyleEmoji:
		return "text-yellow-500"
	default:
		return "text-yellow-400"
	}
}

func getRatingIcon(style RatingStyle, filled bool, value float64) templ.Component {
	if style == RatingStyleEmoji {
		if filled {
			switch {
			case value <= 1:
				return icons.Angry(icons.IconProps{})
			case value <= 2:
				return icons.Frown(icons.IconProps{})
			case value <= 3:
				return icons.Meh(icons.IconProps{})
			case value <= 4:
				return icons.Smile(icons.IconProps{})
			default:
				return icons.Laugh(icons.IconProps{})
			}
		}
		return icons.Meh(icons.IconProps{})
	}
	if filled {
		switch style {
		case RatingStyleHeart:
			return icons.Heart(icons.IconProps{Fill: "currentColor"})
		default:
			return icons.Star(icons.IconProps{Fill: "currentColor"})
		}
	} else {
		switch style {
		case RatingStyleHeart:
			return icons.Heart(icons.IconProps{})
		default:
			return icons.Star(icons.IconProps{})
		}
	}
}

func (p *RatingItemProps) setDefaults() {
	if p.Style == "" {
		p.Style = RatingStyleStar
	}
}

func (p *RatingProps) setDefaults() {
	if p.Precision <= 0 {
		p.Precision = 1.0
	}
}

func RatingScript() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var18 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var18 == nil {
			templ_7745c5c3_Var18 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		handle := templ.NewOnceHandle()
		templ_7745c5c3_Var19 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 23, "<script defer nonce=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var20 string
			templ_7745c5c3_Var20, templ_7745c5c3_Err = templ.JoinStringErrs(templ.GetNonce(ctx))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/rating.templ`, Line: 184, Col: 43}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var20))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 24, "\">\n\t\t\tdocument.addEventListener('alpine:init', () => {\n\t\t\t\tAlpine.data('rating', () => ({\n\t\t\t\t\tvalue: 0,\n\t\t\t\t\tmaxValue: 5, // Default value, will be dynamically updated\n\t\t\t\t\tprecision: 1,\n\t\t\t\t\treadonly: false,\n\t\t\t\t\tname: '',\n\t\t\t\t\tonlyInteger: false,\n\t\t\t\t\tpreviewValue: 0,\n\t\t\t\t\t\n\t\t\t\t\tinit() {\n\t\t\t\t\t\tthis.value = parseFloat(this.$el.dataset.value) || 0;\n\t\t\t\t\t\tthis.precision = parseFloat(this.$el.dataset.precision) || 1;\n\t\t\t\t\t\tthis.readonly = this.$el.dataset.readonly === 'true';\n\t\t\t\t\t\tthis.name = this.$el.dataset.name || '';\n\t\t\t\t\t\tthis.onlyInteger = this.$el.dataset.onlyinteger === 'true';\n\t\t\t\t\t\t\n\t\t\t\t\t\t// Dynamically calculate maxValue based on the items\n\t\t\t\t\t\tthis.calculateMaxValue();\n\t\t\t\t\t\t\n\t\t\t\t\t\t// Ensure value is within valid range\n\t\t\t\t\t\tthis.value = Math.min(this.maxValue, this.value);\n\t\t\t\t\t\tthis.value = Math.round(this.value / this.precision) * this.precision;\n\t\t\t\t\t\t\n\t\t\t\t\t\t// Initialize the form element for proper form integration\n\t\t\t\t\t\tif (this.$refs.input) {\n\t\t\t\t\t\t\tthis.$refs.input.value = this.value;\n\t\t\t\t\t\t}\n\t\t\t\t\t\t\n\t\t\t\t\t\t// Check if we're in a form context\n\t\t\t\t\t\tthis.form = this.$el.closest('form');\n\t\t\t\t\t\tif (this.form && this.name) {\n\t\t\t\t\t\t\t// Add validation support\n\t\t\t\t\t\t\tthis.form.addEventListener('submit', () => {\n\t\t\t\t\t\t\t\tthis.validate();\n\t\t\t\t\t\t\t});\n\t\t\t\t\t\t}\n\t\t\t\t\t\t\n\t\t\t\t\t\t// Setup mutation observer to react to DOM changes\n\t\t\t\t\t\tthis.observeDOMChanges();\n\t\t\t\t\t},\n\t\t\t\t\t\n\t\t\t\t\tcalculateMaxValue() {\n\t\t\t\t\t\t// Find all rating items and determine the highest value\n\t\t\t\t\t\tconst items = this.$el.querySelectorAll('[data-rating-value]');\n\t\t\t\t\t\tlet highestValue = 0;\n\t\t\t\t\t\t\n\t\t\t\t\t\titems.forEach(item => {\n\t\t\t\t\t\t\tconst value = parseInt(item.dataset.ratingValue, 10);\n\t\t\t\t\t\t\tif (value > highestValue) {\n\t\t\t\t\t\t\t\thighestValue = value;\n\t\t\t\t\t\t\t}\n\t\t\t\t\t\t});\n\t\t\t\t\t\t\n\t\t\t\t\t\t// Set minimum maxValue of 1\n\t\t\t\t\t\tthis.maxValue = Math.max(1, highestValue);\n\t\t\t\t\t},\n\t\t\t\t\t\n\t\t\t\t\tobserveDOMChanges() {\n\t\t\t\t\t\t// Use MutationObserver to react to DOM changes\n\t\t\t\t\t\tconst observer = new MutationObserver(() => {\n\t\t\t\t\t\t\tthis.calculateMaxValue();\n\t\t\t\t\t\t});\n\t\t\t\t\t\t\n\t\t\t\t\t\t// Watch for changes in the child node list\n\t\t\t\t\t\tobserver.observe(this.$el, { childList: true, subtree: true });\n\t\t\t\t\t},\n\t\t\t\t\t\n\t\t\t\t\tvalidate() {\n\t\t\t\t\t\t// Basic validation - can be extended as needed\n\t\t\t\t\t\tconst isValid = this.value > 0;\n\t\t\t\t\t\t// Trigger custom event for form validation\n\t\t\t\t\t\tthis.$dispatch('rating-validate', { \n\t\t\t\t\t\t\tname: this.name, \n\t\t\t\t\t\t\tvalue: this.value,\n\t\t\t\t\t\t\tvalid: isValid\n\t\t\t\t\t\t});\n\t\t\t\t\t\treturn isValid;\n\t\t\t\t\t},\n\t\t\t\t\t\n\t\t\t\t\tsetValue() {\n\t\t\t\t\t\tif (this.readonly) return;\n\t\t\t\t\t\tconst item = this.$event.target.closest('[data-rating-value]');\n\t\t\t\t\t\tif (!item) return;\n\t\t\t\t\t\tconst newValue = parseInt(item.dataset.ratingValue);\n\t\t\t\t\t\tif (this.onlyInteger) {\n\t\t\t\t\t\t\tthis.value = Math.round(newValue);\n\t\t\t\t\t\t} else {\n\t\t\t\t\t\t\tthis.value = Math.round(newValue / this.precision) * this.precision;\n\t\t\t\t\t\t}\n\t\t\t\t\t\tthis.value = Math.max(0, Math.min(this.maxValue, this.value));\n\t\t\t\t\t\t\n\t\t\t\t\t\t// Update the hidden input value\n\t\t\t\t\t\tif (this.$refs.input) {\n\t\t\t\t\t\t\tthis.$refs.input.value = this.value;\n\t\t\t\t\t\t}\n\t\t\t\t\t\t\n\t\t\t\t\t\t// Trigger change events for form integration\n\t\t\t\t\t\tthis.$dispatch('rating-change', { \n\t\t\t\t\t\t\tname: this.name, \n\t\t\t\t\t\t\tvalue: this.value,\n\t\t\t\t\t\t\tmaxValue: this.maxValue\n\t\t\t\t\t\t});\n\t\t\t\t\t\t\n\t\t\t\t\t\t// Trigger input event for better form integration\n\t\t\t\t\t\tif (this.$refs.input) {\n\t\t\t\t\t\t\tthis.$refs.input.dispatchEvent(new Event('input', { bubbles: true }));\n\t\t\t\t\t\t\tthis.$refs.input.dispatchEvent(new Event('change', { bubbles: true }));\n\t\t\t\t\t\t}\n\t\t\t\t\t},\n\t\t\t\t\t\n\t\t\t\t\tgetFormattedValue() {\n\t\t\t\t\t\treturn Math.round(this.value * 100) / 100;\n\t\t\t\t\t},\n\t\t\t\t\t\n\t\t\t\t\tgetItemStyle() {\n\t\t\t\t\t\tconst index = parseInt(this.$el.dataset.index || '0');\n\t\t\t\t\t\tconst filled = index <= Math.floor(this.value);\n\t\t\t\t\t\tconst partial = !filled && (index - 1 < this.value && this.value < index);\n\t\t\t\t\t\tconst percentage = partial ? (this.value - Math.floor(this.value)) * 100 : 0;\n\t\t\t\t\t\treturn {\n\t\t\t\t\t\t\twidth: filled ? '100%' : (partial ? percentage + '%' : '0%')\n\t\t\t\t\t\t};\n\t\t\t\t\t},\n\t\t\t\t\t\n\t\t\t\t\thover() {\n\t\t\t\t\t\tif (this.readonly) return;\n\t\t\t\t\t\tconst item = this.$event.target.closest('[data-rating-value]');\n\t\t\t\t\t\tif (!item) return;\n\t\t\t\t\t\tthis.previewValue = parseInt(item.dataset.ratingValue);\n\t\t\t\t\t},\n\t\t\t\t\t\n\t\t\t\t\tresetPreview() {\n\t\t\t\t\t\tif (this.readonly) return;\n\t\t\t\t\t\tthis.previewValue = 0;\n\t\t\t\t\t}\n\t\t\t\t}));\n\t\t\t});\n\t\t</script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = handle.Once().Render(templ.WithChildren(ctx, templ_7745c5c3_Var19), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
