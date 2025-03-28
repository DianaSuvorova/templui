// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.856
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/axzilla/templui/utils"
	"strconv"
)

type PopoverPosition string

const (
	PopoverTop         PopoverPosition = "top"
	PopoverTopStart    PopoverPosition = "top-start"
	PopoverTopEnd      PopoverPosition = "top-end"
	PopoverRight       PopoverPosition = "right"
	PopoverRightStart  PopoverPosition = "right-start"
	PopoverRightEnd    PopoverPosition = "right-end"
	PopoverBottom      PopoverPosition = "bottom"
	PopoverBottomStart PopoverPosition = "bottom-start"
	PopoverBottomEnd   PopoverPosition = "bottom-end"
	PopoverLeft        PopoverPosition = "left"
	PopoverLeftStart   PopoverPosition = "left-start"
	PopoverLeftEnd     PopoverPosition = "left-end"
)

type PopoverTriggerType string

const (
	PopoverTriggerTypeHover PopoverTriggerType = "hover"
	PopoverTriggerTypeClick PopoverTriggerType = "click"
)

type PopoverProps struct {
	Class string
}

type PopoverTriggerProps struct {
	ID          string
	TriggerType PopoverTriggerType
}

type PopoverContentProps struct {
	ID               string
	Class            string
	Attributes       templ.Attributes
	Position         PopoverPosition
	DisableClickAway bool
	DisableESC       bool
	ShowArrow        bool
}

func popoverPortalContainer() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div id=\"popover-portal-container\" class=\"fixed inset-0 z-[9999] pointer-events-none\"></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func PopoverTrigger(props ...PopoverTriggerProps) templ.Component {
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
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var p PopoverTriggerProps
		if len(props) > 0 {
			p = props[0]
		}
		if p.TriggerType == "" {
			p.TriggerType = PopoverTriggerTypeClick
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<span data-popover-trigger data-popover-id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(p.ID)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/popover.templ`, Line: 68, Col: 24}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "\" data-popover-type=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(string(p.TriggerType))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/popover.templ`, Line: 69, Col: 43}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var2.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "</span>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func Popover(props ...PopoverProps) templ.Component {
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
		templ_7745c5c3_Var5 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var5 == nil {
			templ_7745c5c3_Var5 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var p PopoverProps
		if len(props) > 0 {
			p = props[0]
		}
		var templ_7745c5c3_Var6 = []any{utils.TwMerge("relative inline-block", p.Class)}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var6...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<div class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var6).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/popover.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var5.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = popoverPortalContainer().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func PopoverContent(props ...PopoverContentProps) templ.Component {
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
		templ_7745c5c3_Var8 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var8 == nil {
			templ_7745c5c3_Var8 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var p PopoverContentProps
		if len(props) > 0 {
			p = props[0]
		}
		if p.Position == "" {
			p.Position = PopoverBottom
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "<template data-popover-content-template data-popover-id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var9 string
		templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(p.ID)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/popover.templ`, Line: 96, Col: 24}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "\" data-popover-position=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var10 string
		templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(string(p.Position))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/popover.templ`, Line: 97, Col: 44}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, "\" data-popover-disable-clickaway=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var11 string
		templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.FormatBool(p.DisableClickAway))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/popover.templ`, Line: 98, Col: 73}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, "\" data-popover-disable-esc=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var12 string
		templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.FormatBool(p.DisableESC))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/popover.templ`, Line: 99, Col: 61}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, "\" data-popover-show-arrow=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var13 string
		templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.FormatBool(p.ShowArrow))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/popover.templ`, Line: 100, Col: 59}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 14, "\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var14 = []any{utils.TwMerge(
			"bg-background rounded-lg border text-sm shadow-lg pointer-events-auto absolute z-[9999]",
			p.Class,
		)}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var14...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 15, "<div class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var15 string
		templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var14).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/popover.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 16, "\"><div class=\"w-full overflow-hidden\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var8.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 17, "</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if p.ShowArrow {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 18, "<!-- Wir erzeugen alle Pfeile mit eindeutigen Data-Attributen zur einfacheren Identifikation --> <!-- Top-Pfeile --> <div data-arrow=\"top\" class=\"absolute h-2.5 w-2.5 rotate-45 bg-background top-[-5px] left-1/2 -translate-x-1/2 border-t border-l hidden\"></div><div data-arrow=\"top-start\" class=\"absolute h-2.5 w-2.5 rotate-45 bg-background top-[-5px] left-4 border-t border-l hidden\"></div><div data-arrow=\"top-end\" class=\"absolute h-2.5 w-2.5 rotate-45 bg-background top-[-5px] right-4 border-t border-l hidden\"></div><!-- Bottom-Pfeile --> <div data-arrow=\"bottom\" class=\"absolute h-2.5 w-2.5 rotate-45 bg-background bottom-[-5px] left-1/2 -translate-x-1/2 border-b border-r hidden\"></div><div data-arrow=\"bottom-start\" class=\"absolute h-2.5 w-2.5 rotate-45 bg-background bottom-[-5px] left-4 border-b border-r hidden\"></div><div data-arrow=\"bottom-end\" class=\"absolute h-2.5 w-2.5 rotate-45 bg-background bottom-[-5px] right-4 border-b border-r hidden\"></div><!-- Left-Pfeile --> <div data-arrow=\"left\" class=\"absolute h-2.5 w-2.5 rotate-45 bg-background left-[-5px] top-1/2 -translate-y-1/2 border-b border-l hidden\"></div><div data-arrow=\"left-start\" class=\"absolute h-2.5 w-2.5 rotate-45 bg-background left-[-5px] top-2 border-b border-l hidden\"></div><div data-arrow=\"left-end\" class=\"absolute h-2.5 w-2.5 rotate-45 bg-background left-[-5px] bottom-2 border-b border-l hidden\"></div><!-- Right-Pfeile --> <div data-arrow=\"right\" class=\"absolute h-2.5 w-2.5 rotate-45 bg-background right-[-5px] top-1/2 -translate-y-1/2 border-t border-r hidden\"></div><div data-arrow=\"right-start\" class=\"absolute h-2.5 w-2.5 rotate-45 bg-background right-[-5px] top-2 border-t border-r hidden\"></div><div data-arrow=\"right-end\" class=\"absolute h-2.5 w-2.5 rotate-45 bg-background right-[-5px] bottom-2 border-t border-r hidden\"></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 19, "</div></template>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func PopoverScript() templ.Component {
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
		templ_7745c5c3_Var16 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var16 == nil {
			templ_7745c5c3_Var16 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 20, "<script defer nonce=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var17 string
		templ_7745c5c3_Var17, templ_7745c5c3_Err = templ.JoinStringErrs(templ.GetNonce(ctx))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/popover.templ`, Line: 135, Col: 42}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var17))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 21, "\">\n        document.addEventListener('DOMContentLoaded', () => {\n            const portalContainer = document.getElementById('popover-portal-container');\n            const triggers = document.querySelectorAll('[data-popover-trigger]');\n            const templates = document.querySelectorAll('[data-popover-content-template]');\n            \n            // Aktive Popovers\n            const activePopovers = new Map();\n            \n            // Funktion zum Aktualisieren des Pfeils basierend auf der Position\n            const updateArrow = (popoverElement, position) => {\n                if (popoverElement.dataset.popoverShowArrow !== 'true') return;\n                \n                // Alle Pfeile zunächst verstecken\n                popoverElement.querySelectorAll('[data-arrow]').forEach(arrow => {\n                    arrow.classList.add('hidden');\n                });\n                \n                // Pfeilrichtung ist immer gegenteilig zur Position\n                // z.B. wenn Popover unten ist, muss der Pfeil nach oben zeigen\n                let arrowPosition;\n                \n                if (position.startsWith('top')) {\n                    // Wenn Popover oben ist, muss Pfeil nach unten zeigen\n                    arrowPosition = position.replace('top', 'bottom');\n                } else if (position.startsWith('bottom')) {\n                    // Wenn Popover unten ist, muss Pfeil nach oben zeigen\n                    arrowPosition = position.replace('bottom', 'top');\n                } else if (position.startsWith('left')) {\n                    // Wenn Popover links ist, muss Pfeil nach rechts zeigen\n                    arrowPosition = position.replace('left', 'right');\n                } else if (position.startsWith('right')) {\n                    // Wenn Popover rechts ist, muss Pfeil nach links zeigen\n                    arrowPosition = position.replace('right', 'left');\n                } else {\n                    // Fallback\n                    arrowPosition = position;\n                }\n                \n                // Den korrekten Pfeil basierend auf der Richtung anzeigen\n                const arrow = popoverElement.querySelector(`[data-arrow=\"${arrowPosition}\"]`);\n                if (arrow) {\n                    arrow.classList.remove('hidden');\n                }\n            };\n            \n            // Positionierungsfunktion\n            const positionPopover = (trigger, popoverElement) => {\n                // Finde das tatsächliche Element, auf das sich der Popover bezieht\n                let triggerElement = trigger;\n                let largestArea = 0;\n                \n                // Prüfe alle direkten Kinder des Triggers\n                const children = trigger.children;\n                for (let i = 0; i < children.length; i++) {\n                    const child = children[i];\n                    const rect = child.getBoundingClientRect();\n                    const area = rect.width * rect.height;\n                    \n                    if (area > largestArea) {\n                        largestArea = area;\n                        triggerElement = child;\n                    }\n                }\n                \n                const triggerRect = triggerElement.getBoundingClientRect();\n                const contentRect = popoverElement.getBoundingClientRect();\n                const margin = popoverElement.dataset.popoverShowArrow === 'true' ? 8 : 4;\n                const scrollY = window.scrollY || window.pageYOffset;\n                const scrollX = window.scrollX || window.pageXOffset;\n\n                // Position aus dem Datensatz abrufen\n                const requestedPosition = popoverElement.dataset.popoverPosition || 'bottom';\n                // Wir speichern die finale Position, die je nach Viewport-Raum angepasst werden kann\n                let finalPosition = requestedPosition;\n\n                // Viewport-Dimensionen\n                const viewportWidth = window.innerWidth;\n                const viewportHeight = window.innerHeight;\n\n                // Get element heights and widths\n                const triggerHeight = triggerRect.height;\n                const contentHeight = contentRect.height;\n                const contentWidth = contentRect.width;\n                \n                // Definiere Ankerpunkte\n                const triggerTop = triggerRect.top + scrollY;\n                const triggerBottom = triggerRect.bottom + scrollY;\n                const triggerLeft = triggerRect.left + scrollX;\n                const triggerRight = triggerRect.right + scrollX;\n                \n                // Berechne verfügbaren Platz in jede Richtung\n                const spaceAbove = triggerRect.top;\n                const spaceBelow = viewportHeight - triggerRect.bottom;\n                const spaceLeft = triggerRect.left;\n                const spaceRight = viewportWidth - triggerRect.right;\n\n                // Intelligente Positionsanpassung\n                // Wir prüfen die gegensätzliche Position, wenn nicht genug Platz vorhanden ist\n                if (finalPosition.startsWith('top') && spaceAbove < contentHeight + margin) {\n                    // Wenn oben nicht genug Platz ist, unten anzeigen\n                    finalPosition = finalPosition.replace('top', 'bottom');\n                } else if (finalPosition.startsWith('bottom') && spaceBelow < contentHeight + margin) {\n                    // Wenn unten nicht genug Platz ist, oben anzeigen\n                    finalPosition = finalPosition.replace('bottom', 'top');\n                } else if (finalPosition.startsWith('left') && spaceLeft < contentWidth + margin) {\n                    // Wenn links nicht genug Platz ist, rechts anzeigen\n                    finalPosition = finalPosition.replace('left', 'right');\n                } else if (finalPosition.startsWith('right') && spaceRight < contentWidth + margin) {\n                    // Wenn rechts nicht genug Platz ist, links anzeigen\n                    finalPosition = finalPosition.replace('right', 'left');\n                }\n\n                // Speichere die aktuelle Position im Element für CSS-Anpassungen (z.B. Pfeil-Position)\n                popoverElement.dataset.popoverCurrentPosition = finalPosition;\n                \n                // Den richtigen Pfeil anzeigen\n                updateArrow(popoverElement, finalPosition);\n                \n                let top, left;\n                \n                // Positionierungslogik mit der finalen Position\n                switch (finalPosition) {\n                    case 'top':\n                        top = triggerTop - contentHeight - margin;\n                        left = triggerLeft + (triggerRect.width / 2) - (contentRect.width / 2);\n                        break;\n                    case 'top-start':\n                        top = triggerTop - contentHeight - margin;\n                        left = triggerLeft;\n                        break;\n                    case 'top-end':\n                        top = triggerTop - contentHeight - margin;\n                        left = triggerRight - contentRect.width;\n                        break;\n                    case 'right':\n                        top = triggerTop + (triggerHeight / 2) - (contentHeight / 2);\n                        left = triggerRight + margin;\n                        break;\n                    case 'right-start':\n                        top = triggerTop;\n                        left = triggerRight + margin;\n                        break;\n                    case 'right-end':\n                        top = triggerBottom - contentHeight;\n                        left = triggerRight + margin;\n                        break;\n                    case 'bottom':\n                        top = triggerBottom + margin;\n                        left = triggerLeft + (triggerRect.width / 2) - (contentRect.width / 2);\n                        break;\n                    case 'bottom-start':\n                        top = triggerBottom + margin;\n                        left = triggerLeft;\n                        break;\n                    case 'bottom-end':\n                        top = triggerBottom + margin;\n                        left = triggerRight - contentRect.width;\n                        break;\n                    case 'left':\n                        top = triggerTop + (triggerHeight / 2) - (contentHeight / 2);\n                        left = triggerLeft - contentRect.width - margin;\n                        break;\n                    case 'left-start':\n                        top = triggerTop;\n                        left = triggerLeft - contentRect.width - margin;\n                        break;\n                    case 'left-end':\n                        top = triggerBottom - contentHeight;\n                        left = triggerLeft - contentRect.width - margin;\n                        break;\n                    default:\n                        top = triggerBottom + margin;\n                        left = triggerLeft + (triggerRect.width / 2) - (contentRect.width / 2);\n                }\n\n                // Horizontale Begrenzung - sorgt dafür, dass der Popover nicht aus dem Viewport ragt\n                if (left < 10) {\n                    left = 10; // Minimaler Abstand vom linken Rand\n                } else if (left + contentWidth > viewportWidth - 10) {\n                    left = viewportWidth - contentWidth - 10; // Minimaler Abstand vom rechten Rand\n                }\n\n                // Vertikale Begrenzung - Optional, kann in manchen Fällen problematisch sein\n                if (top < 10) {\n                    top = 10; // Minimaler Abstand vom oberen Rand\n                } else if (top + contentHeight > viewportHeight - 10) {\n                    top = viewportHeight - contentHeight - 10; // Minimaler Abstand vom unteren Rand\n                }\n\n                popoverElement.style.top = `${top}px`;\n                popoverElement.style.left = `${left}px`;\n            };\n            \n            // Event-Handler einrichten\n            function setupTrigger(trigger) {\n                const popoverId = trigger.dataset.popoverId;\n                const template = document.querySelector(`[data-popover-content-template][data-popover-id=\"${popoverId}\"]`);\n                \n                if (!template) return;\n                \n                const triggerType = trigger.dataset.popoverType;\n                \n                // Click-Handler\n                if (triggerType === 'click') {\n                    trigger.addEventListener('click', () => {\n                        // Wenn Popover bereits aktiv ist, entfernen\n                        if (activePopovers.has(popoverId)) {\n                            const popover = activePopovers.get(popoverId);\n                            popover.remove();\n                            activePopovers.delete(popoverId);\n                            return;\n                        }\n                        \n                        // Sonst neues Popover erstellen\n                        const content = template.content.cloneNode(true).firstElementChild;\n                        \n                        // Attribute vom Template übertragen\n                        Object.keys(template.dataset).forEach(key => {\n                            if (key.startsWith('popover')) {\n                                content.dataset[key] = template.dataset[key];\n                            }\n                        });\n                        \n                        // Initiale Position setzen\n                        content.dataset.popoverCurrentPosition = content.dataset.popoverPosition;\n                        \n                        // Zum Portal-Container hinzufügen\n                        portalContainer.appendChild(content);\n                        \n                        // Positionieren\n                        positionPopover(trigger, content);\n                        \n                        // In aktive Popovers einfügen\n                        activePopovers.set(popoverId, content);\n                        \n                        // Clickaway-Handler hinzufügen\n                        if (content.dataset.popoverDisableClickaway !== 'true') {\n                            const clickHandler = (e) => {\n                                if (!trigger.contains(e.target) && !content.contains(e.target)) {\n                                    content.remove();\n                                    activePopovers.delete(popoverId);\n                                    document.removeEventListener('click', clickHandler);\n                                }\n                            };\n                            \n                            // Verzögerung, damit der aktuelle Klick nicht sofort wieder schließt\n                            setTimeout(() => {\n                                document.addEventListener('click', clickHandler);\n                            }, 0);\n                        }\n                        \n                        // ESC-Handler hinzufügen\n                        if (content.dataset.popoverDisableEsc !== 'true') {\n                            const keyHandler = (e) => {\n                                if (e.key === 'Escape') {\n                                    content.remove();\n                                    activePopovers.delete(popoverId);\n                                    document.removeEventListener('keydown', keyHandler);\n                                }\n                            };\n                            document.addEventListener('keydown', keyHandler);\n                        }\n                    });\n                } else if (triggerType === 'hover') {\n                    // Hover-Handlers\n                    let hoverTimeout;\n                    let leaveTimeout;\n                    \n                    trigger.addEventListener('mouseenter', () => {\n                        clearTimeout(leaveTimeout);\n                        \n                        // Wenn Popover bereits aktiv, nicht neu erstellen\n                        if (activePopovers.has(popoverId)) return;\n                        \n                        // Neues Popover erstellen\n                        const content = template.content.cloneNode(true).firstElementChild;\n                        \n                        // Attribute übertragen\n                        Object.keys(template.dataset).forEach(key => {\n                            if (key.startsWith('popover')) {\n                                content.dataset[key] = template.dataset[key];\n                            }\n                        });\n                        \n                        // Zum Portal-Container hinzufügen\n                        portalContainer.appendChild(content);\n                        \n                        // Positionieren\n                        positionPopover(trigger, content);\n                        \n                        // In aktive Popovers einfügen\n                        activePopovers.set(popoverId, content);\n                        \n                        // Hover-Handler für das Content-Element\n                        content.addEventListener('mouseenter', () => {\n                            clearTimeout(leaveTimeout);\n                        });\n                        \n                        content.addEventListener('mouseleave', () => {\n                            leaveTimeout = setTimeout(() => {\n                                if (activePopovers.has(popoverId)) {\n                                    const popover = activePopovers.get(popoverId);\n                                    popover.remove();\n                                    activePopovers.delete(popoverId);\n                                }\n                            }, 100);\n                        });\n                    });\n                    \n                    trigger.addEventListener('mouseleave', (e) => {\n                        // Prüfen, ob wir zum Content hovern\n                        const related = e.relatedTarget;\n                        const content = activePopovers.get(popoverId);\n                        \n                        // Wenn wir direkt zum Content hovern, nicht schließen\n                        if (content && content.contains(related)) {\n                            return;\n                        }\n                        \n                        leaveTimeout = setTimeout(() => {\n                            if (activePopovers.has(popoverId)) {\n                                const popover = activePopovers.get(popoverId);\n                                popover.remove();\n                                activePopovers.delete(popoverId);\n                            }\n                        }, 100);\n                    });\n                }\n            }\n            \n            // Für jeden Trigger die Handler einrichten\n            triggers.forEach(setupTrigger);\n            \n            // Scrolling-Handler für alle Popovers\n            window.addEventListener('scroll', () => {\n                activePopovers.forEach((content, popoverId) => {\n                    const trigger = document.querySelector(`[data-popover-trigger][data-popover-id=\"${popoverId}\"]`);\n                    if (trigger) {\n                        positionPopover(trigger, content);\n                    }\n                });\n            }, { passive: true });\n            \n            // Resize-Handler\n            window.addEventListener('resize', () => {\n                activePopovers.forEach((content, popoverId) => {\n                    const trigger = document.querySelector(`[data-popover-trigger][data-popover-id=\"${popoverId}\"]`);\n                    if (trigger) {\n                        positionPopover(trigger, content);\n                    }\n                });\n            });\n            \n            // Finde alle scrollbaren Elternelemente und füge Scroll-Handler hinzu\n            function setupScrollHandlers() {\n                const scrollableElements = new Set();\n                \n                // Für jeden Trigger die scrollbaren Eltern finden\n                triggers.forEach(trigger => {\n                    let element = trigger.parentElement;\n                    \n                    while (element) {\n                        const style = window.getComputedStyle(element);\n                        const overflow = style.overflow + style.overflowY + style.overflowX;\n                        \n                        if (overflow.includes('scroll') || overflow.includes('auto') || \n                            element.scrollHeight > element.clientHeight) {\n                            scrollableElements.add(element);\n                        }\n                        \n                        element = element.parentElement;\n                    }\n                });\n                \n                // Scroll-Handler für jedes scrollbare Element\n                scrollableElements.forEach(element => {\n                    element.addEventListener('scroll', () => {\n                        activePopovers.forEach((content, popoverId) => {\n                            const trigger = document.querySelector(`[data-popover-trigger][data-popover-id=\"${popoverId}\"]`);\n                            if (trigger) {\n                                positionPopover(trigger, content);\n                            }\n                        });\n                    }, { passive: true });\n                });\n            }\n            \n            setupScrollHandlers();\n        });\n    </script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
