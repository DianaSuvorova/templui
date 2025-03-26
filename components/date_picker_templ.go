// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.856
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/axzilla/templui/icons"
	"github.com/axzilla/templui/utils"
	"time"
)

type DateFormat string

const (
	DateFormatISO  DateFormat = "iso"
	DateFormatEU   DateFormat = "eu"
	DateFormatUK   DateFormat = "uk"
	DateFormatUS   DateFormat = "us"
	DateFormatLONG DateFormat = "long"
)

var dateFormatMapping = map[DateFormat]string{
	DateFormatISO:  "2006-01-02",
	DateFormatEU:   "02.01.2006",
	DateFormatUK:   "02/01/2006",
	DateFormatUS:   "01/02/2006",
	DateFormatLONG: "January 2, 2006",
}

type DateLocale struct {
	MonthNames []string
	DayNames   []string
}

var (
	DateLocaleDefault = DateLocale{
		MonthNames: []string{"January", "February", "March", "April", "May", "June",
			"July", "August", "September", "October", "November", "December"},
		DayNames: []string{"Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"},
	}

	DateLocaleSpanish = DateLocale{
		MonthNames: []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio",
			"Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"},
		DayNames: []string{"Lu", "Ma", "Mi", "Ju", "Vi", "Sa", "Do"},
	}

	DateLocaleGerman = DateLocale{
		MonthNames: []string{"Januar", "Februar", "März", "April", "Mai", "Juni",
			"Juli", "August", "September", "Oktober", "November", "Dezember"},
		DayNames: []string{"Mo", "Di", "Mi", "Do", "Fr", "Sa", "So"},
	}

	DateLocaleFrench = DateLocale{
		MonthNames: []string{"Janvier", "Février", "Mars", "Avril", "Mai", "Juin",
			"Juillet", "Août", "Septembre", "Octobre", "Novembre", "Décembre"},
		DayNames: []string{"Lu", "Ma", "Me", "Je", "Ve", "Sa", "Di"},
	}

	DateLocaleItalian = DateLocale{
		MonthNames: []string{"Gennaio", "Febbraio", "Marzo", "Aprile", "Maggio", "Giugno",
			"Luglio", "Agosto", "Settembre", "Ottobre", "Novembre", "Dicembre"},
		DayNames: []string{"Lu", "Ma", "Me", "Gi", "Ve", "Sa", "Do"},
	}

	DateLocaleJapanese = DateLocale{
		MonthNames: []string{"1月", "2月", "3月", "4月", "5月", "6月",
			"7月", "8月", "9月", "10月", "11月", "12月"},
		DayNames: []string{"日", "月", "火", "水", "木", "金", "土"},
	}
)

var (
	DatePickerISO = DatePickerConfig{
		Format: DateFormatISO,
		Locale: DateLocaleDefault,
	}

	DatePickerEU = DatePickerConfig{
		Format: DateFormatEU,
		Locale: DateLocaleDefault,
	}

	DatePickerUK = DatePickerConfig{
		Format: DateFormatUK,
		Locale: DateLocaleDefault,
	}

	DatePickerUS = DatePickerConfig{
		Format: DateFormatUS,
		Locale: DateLocaleDefault,
	}

	DatePickerLONG = DatePickerConfig{
		Format: DateFormatLONG,
		Locale: DateLocaleDefault,
	}
)

func NewDatePickerConfig(format DateFormat, locale DateLocale) DatePickerConfig {
	return DatePickerConfig{
		Format: format,
		Locale: locale,
	}
}

type DatePickerConfig struct {
	Format DateFormat
	Locale DateLocale
}

// All other DateFormat, DateLocale, etc. definitions remain unchanged
type DatePickerProps struct {
	ID          string
	Class       string
	Attributes  templ.Attributes
	Value       time.Time
	Config      DatePickerConfig
	Placeholder string
	Disabled    bool
	HasError    bool
	Name        string
}

func DatePicker(props ...DatePickerProps) templ.Component {
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
		var p DatePickerProps
		if len(props) > 0 {
			p = props[0]
		}
		if p.ID == "" {
			p.ID = utils.RandomID()
		}
		if p.Placeholder == "" {
			p.Placeholder = "Select a date"
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<!-- Container for the entire DatePicker -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 = []any{utils.TwMerge("relative", p.Class)}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<div id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(p.ID)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/date_picker.templ`, Line: 135, Col: 11}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var2).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/date_picker.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if p.Value != (time.Time{}) {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, " data-value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(p.Value.Format(p.Config.getGoFormat()))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/date_picker.templ`, Line: 138, Col: 54}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, " data-format=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(string(p.Config.Format))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/date_picker.templ`, Line: 140, Col: 39}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "\" data-monthnames=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(templ.JSONString(p.Config.Locale.MonthNames))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/date_picker.templ`, Line: 141, Col: 64}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "\" data-daynames=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(templ.JSONString(p.Config.Locale.DayNames))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/date_picker.templ`, Line: 142, Col: 60}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "\" x-data=\"datePicker\" @resize.window=\"updatePosition\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, p.Attributes)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, "><!-- Hidden input for form value --><input type=\"hidden\" id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var9 string
		templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(p.ID + "-hidden")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/date_picker.templ`, Line: 150, Col: 24}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, "\" name=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var10 string
		templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(p.Name)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/date_picker.templ`, Line: 151, Col: 16}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, "\" x-ref=\"hiddenInput\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if p.Disabled {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 14, " disabled")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 15, "><div class=\"relative\"><!-- Button instead of input -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var11 = []any{utils.TwMerge(
			"w-full h-10 px-3 py-2 rounded-md border bg-background text-sm flex items-center justify-between",
			"focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2",
			"disabled:cursor-not-allowed disabled:opacity-50",
			utils.IfElse(p.HasError, "border-destructive ring-destructive", "border-input"),
		)}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var11...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 16, "<button type=\"button\" id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var12 string
		templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(p.ID + "-trigger")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/date_picker.templ`, Line: 159, Col: 26}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 17, "\" @click=\"toggleDatePicker\" x-ref=\"triggerButton\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if p.Disabled {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 18, " disabled")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 19, " class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var13 string
		templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var11).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/date_picker.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 20, "\"><!-- Display for the selected date --><span x-ref=\"displayText\" class=\"text-left grow\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if p.Value != (time.Time{}) {
			var templ_7745c5c3_Var14 string
			templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(p.Value.Format(p.Config.getGoFormat()))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/date_picker.templ`, Line: 176, Col: 46}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 21, "<span class=\"text-muted-foreground\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var15 string
			templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(p.Placeholder)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/date_picker.templ`, Line: 178, Col: 57}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 22, "</span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 23, "</span><!-- Calendar icon --><span class=\"text-muted-foreground flex items-center\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = icons.Calendar().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 24, "</span></button></div><!-- DatePicker popup with DOM-generated content -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var16 = []any{
			utils.TwMerge(
				"absolute left-0 z-50 w-64 p-4",
				"rounded-lg border bg-popover shadow-md",
			),
		}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var16...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 25, "<div x-show=\"open\" x-ref=\"datePickerPopup\" @click.away=\"closeDatePicker\" x-transition.opacity class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var17 string
		templ_7745c5c3_Var17, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var16).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/date_picker.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var17))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 26, "\" style=\"display: none;\"><div class=\"flex items-center justify-between mb-4\"><span x-ref=\"monthDisplay\" class=\"text-sm font-medium\"></span><div class=\"flex gap-1\"><button type=\"button\" @click=\"prevMonth\" class=\"inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-hidden focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 hover:bg-accent hover:text-accent-foreground h-7 w-7\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = icons.ChevronLeft().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 27, "</button> <button type=\"button\" @click=\"nextMonth\" class=\"inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-hidden focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 hover:bg-accent hover:text-accent-foreground h-7 w-7\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = icons.ChevronRight().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 28, "</button></div></div><!-- Weekday container --><div x-ref=\"weekdaysContainer\" class=\"grid grid-cols-7 gap-1 mb-2\"></div><!-- Calendar days container --><div x-ref=\"daysContainer\" class=\"grid grid-cols-7 gap-1\"></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func (c DatePickerConfig) getGoFormat() string {
	if format, ok := dateFormatMapping[c.Format]; ok {
		return format
	}
	return dateFormatMapping[DateFormatISO]
}

func DatePickerScript() templ.Component {
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 29, "<script defer nonce=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var20 string
			templ_7745c5c3_Var20, templ_7745c5c3_Err = templ.JoinStringErrs(templ.GetNonce(ctx))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/date_picker.templ`, Line: 238, Col: 43}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var20))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 30, "\">\n\t\t\tdocument.addEventListener('alpine:init', () => {\n\t\t\t\tAlpine.data('datePicker', () => ({\n\t\t\t\t\topen: false,\n\t\t\t\t\tvalue: null,\n\t\t\t\t\tformat: null,\n\t\t\t\t\tcurrentMonth: new Date().getMonth(),\n\t\t\t\t\tcurrentYear: new Date().getFullYear(),\n\t\t\t\t\tmonthNames: [],\n\t\t\t\t\tdayNames: [],\n\t\t\t\t\tposition: 'bottom',\n\n\t\t\t\t\tinit() {\n\t\t\t\t\t\ttry {\n\t\t\t\t\t\t    // Parse month names\n\t\t\t\t\t\t    this.monthNames = JSON.parse(this.$el.dataset.monthnames || '[]');\n\t\t\t\t\t\t    if (!this.monthNames || !this.monthNames.length) {\n\t\t\t\t\t\t        this.monthNames = ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'];\n\t\t\t\t\t\t    }\n\t\t\t\t\t\t    \n\t\t\t\t\t\t    // Parse day names\n\t\t\t\t\t\t    this.dayNames = JSON.parse(this.$el.dataset.daynames || '[]');\n\t\t\t\t\t\t    if (!this.dayNames || !this.dayNames.length) {\n\t\t\t\t\t\t        this.dayNames = ['Mo', 'Tu', 'We', 'Th', 'Fr', 'Sa', 'Su'];\n\t\t\t\t\t\t    }\n\t\t\t\t\t\t    \n                            // Set format\n\t\t\t\t\t\t\tthis.format = this.$el.dataset.format;\n\t\t\t\t\t\t\t\n\t\t\t\t\t\t\t// Process initial value\n\t\t\t\t\t\t\tconst initialValue = this.$el.dataset.value;\n\t\t\t\t\t\t\tif (initialValue) {\n\t\t\t\t\t\t\t\tconst initialDate = new Date(this.parseDate(initialValue));\n\t\t\t\t\t\t\t\tthis.currentMonth = initialDate.getMonth();\n\t\t\t\t\t\t\t\tthis.currentYear = initialDate.getFullYear();\n\t\t\t\t\t\t\t\tthis.value = this.formatDate(initialDate);\n\t\t\t\t\t\t\t\t\n\t\t\t\t\t\t\t\t// Update hidden input\n\t\t\t\t\t\t\t\tthis.$refs.hiddenInput.value = this.value;\n\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\n\t\t\t\t\t\t\t// Initialize month display\n\t\t\t\t\t\t\tthis.updateMonthDisplay();\n\t\t\t\t\t\t\t\n\t\t\t\t\t\t\t// Set weekdays\n\t\t\t\t\t\t\tthis.renderWeekdayLabels();\n\t\t\t\t\t\t} catch(e) {\n\t\t\t\t\t\t    console.error('DatePicker initialization error:', e);\n\t\t\t\t\t\t}\n\t\t\t\t\t},\n\n\t\t\t\t\ttoggleDatePicker() {\n\t\t\t\t\t\tif (this.$refs.triggerButton.disabled) {\n\t\t\t\t\t\t\treturn;\n\t\t\t\t\t\t}\n\t\t\t\t\t\t\n\t\t\t\t\t\tthis.open = !this.open;\n\t\t\t\t\t\t\n\t\t\t\t\t\tif (this.open) {\n\t\t\t\t\t\t\t// Adjust position and fill calendar\n\t\t\t\t\t\t\tthis.$nextTick(() => {\n\t\t\t\t\t\t\t\tthis.updatePosition();\n\t\t\t\t\t\t\t\tthis.renderCalendarDays();\n\t\t\t\t\t\t\t});\n\t\t\t\t\t\t}\n\t\t\t\t\t},\n\t\t\t\t\t\n\t\t\t\t\tupdateMonthDisplay() {\n\t\t\t\t\t    if (this.$refs.monthDisplay) {\n\t\t\t\t\t        this.$refs.monthDisplay.textContent = this.monthNames[this.currentMonth] + ' ' + this.currentYear;\n\t\t\t\t\t    }\n\t\t\t\t\t},\n\t\t\t\t\t\n\t\t\t\t\trenderWeekdayLabels() {\n\t\t\t\t\t    const container = this.$refs.weekdaysContainer;\n\t\t\t\t\t    if (!container) return;\n\t\t\t\t\t    \n\t\t\t\t\t    // Clear container\n\t\t\t\t\t    container.innerHTML = '';\n\t\t\t\t\t    \n\t\t\t\t\t    // Add weekdays\n\t\t\t\t\t    for (let i = 0; i < this.dayNames.length; i++) {\n\t\t\t\t\t        const dayEl = document.createElement('div');\n\t\t\t\t\t        dayEl.className = 'text-center text-xs text-muted-foreground';\n\t\t\t\t\t        dayEl.textContent = this.dayNames[i];\n\t\t\t\t\t        container.appendChild(dayEl);\n\t\t\t\t\t    }\n\t\t\t\t\t},\n\t\t\t\t\t\n\t\t\t\t\trenderCalendarDays() {\n\t\t\t\t\t    const container = this.$refs.daysContainer;\n\t\t\t\t\t    if (!container) return;\n\t\t\t\t\t    \n\t\t\t\t\t    // Clear container\n\t\t\t\t\t    container.innerHTML = '';\n\t\t\t\t\t    \n\t\t\t\t\t    // Calculate the first day of the month and number of days\n\t\t\t\t\t    let firstDay = new Date(this.currentYear, this.currentMonth, 1).getDay();\n\t\t\t\t\t    firstDay = firstDay === 0 ? 6 : firstDay - 1; // Convert Sunday (0) to position 6\n\t\t\t\t\t    \n\t\t\t\t\t    const daysInMonth = new Date(this.currentYear, this.currentMonth + 1, 0).getDate();\n\t\t\t\t\t    \n\t\t\t\t\t    // Empty days at the beginning\n\t\t\t\t\t    for (let i = 0; i < firstDay; i++) {\n\t\t\t\t\t        const blankDay = document.createElement('div');\n\t\t\t\t\t        blankDay.className = 'h-8 w-8';\n\t\t\t\t\t        container.appendChild(blankDay);\n\t\t\t\t\t    }\n\t\t\t\t\t    \n\t\t\t\t\t    // Days of the month\n\t\t\t\t\t    for (let day = 1; day <= daysInMonth; day++) {\n\t\t\t\t\t        const dayBtn = document.createElement('button');\n\t\t\t\t\t        dayBtn.type = 'button';\n\t\t\t\t\t        dayBtn.className = 'inline-flex h-8 w-8 items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-hidden focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2';\n\t\t\t\t\t        dayBtn.textContent = day;\n\t\t\t\t\t        dayBtn.dataset.day = day;\n\t\t\t\t\t        \n\t\t\t\t\t        // Click handler\n\t\t\t\t\t        dayBtn.addEventListener('click', (e) => this.selectDate(e));\n\t\t\t\t\t        \n\t\t\t\t\t        // Determine if today or selected\n\t\t\t\t\t        this.styleDayButton(dayBtn, day);\n\t\t\t\t\t        \n\t\t\t\t\t        // Add to container\n\t\t\t\t\t        container.appendChild(dayBtn);\n\t\t\t\t\t    }\n\t\t\t\t\t},\n\t\t\t\t\t\n\t\t\t\t\tstyleDayButton(button, day) {\n\t\t\t\t\t    // Mark today\n\t\t\t\t\t    const today = new Date();\n\t\t\t\t\t    const isToday = today.getDate() === day && \n\t\t\t\t\t                   today.getMonth() === this.currentMonth && \n\t\t\t\t\t                   today.getFullYear() === this.currentYear;\n\t\t\t\t\t                   \n\t\t\t\t\t    // Mark selected date\n\t\t\t\t\t    const isSelected = this.isSelectedDate(day);\n\t\t\t\t\t    \n\t\t\t\t\t    if (isSelected) {\n\t\t\t\t\t        button.classList.add('bg-primary', 'text-primary-foreground');\n\t\t\t\t\t    } else if (isToday) {\n\t\t\t\t\t        button.classList.add('text-red-500');\n\t\t\t\t\t    } else {\n\t\t\t\t\t        button.classList.add('hover:bg-accent', 'hover:text-accent-foreground');\n\t\t\t\t\t    }\n\t\t\t\t\t},\n\n\t\t\t\t\tcloseDatePicker() {\n\t\t\t\t\t\tthis.open = false;\n\t\t\t\t\t},\n\n\t\t\t\t\tupdatePosition() {\n    \t\t\t\t\tconst trigger = this.$refs.triggerButton;\n\t\t\t\t\t\tconst popup = this.$refs.datePickerPopup;\n\t\t\t\t\t\t\n\t\t\t\t\t\tif (!trigger || !popup) return;\n\t\t\t\t\t\t\n\t\t\t\t\t\tconst rect = trigger.getBoundingClientRect();\n\t\t\t\t\t\tconst popupRect = popup.getBoundingClientRect();\n\t\t\t\t\t\tconst viewportHeight = window.innerHeight;\n\t\t\t\t\t\t\n\t\t\t\t\t\tif (rect.bottom + popupRect.height > viewportHeight && rect.top > popupRect.height) {\n\t\t\t\t\t\t\tpopup.classList.add('bottom-full', 'mb-1');\n\t\t\t\t\t\t\tpopup.classList.remove('top-full', 'mt-1');\n\t\t\t\t\t\t} else {\n\t\t\t\t\t\t\tpopup.classList.add('top-full', 'mt-1');\n\t\t\t\t\t\t\tpopup.classList.remove('bottom-full', 'mb-1');\n\t\t\t\t\t\t}\n\t\t\t\t\t},\n\n\t\t\t\t\tprevMonth() {\n\t\t\t\t\t\tthis.currentMonth--;\n\t\t\t\t\t\tif (this.currentMonth < 0) {\n\t\t\t\t\t\t\tthis.currentMonth = 11;\n\t\t\t\t\t\t\tthis.currentYear--;\n\t\t\t\t\t\t}\n\t\t\t\t\t\tthis.updateMonthDisplay();\n\t\t\t\t\t\tthis.renderCalendarDays();\n\t\t\t\t\t},\n\n\t\t\t\t\tnextMonth() {\n\t\t\t\t\t\tthis.currentMonth++;\n\t\t\t\t\t\tif (this.currentMonth > 11) {\n\t\t\t\t\t\t\tthis.currentMonth = 0;\n\t\t\t\t\t\t\tthis.currentYear++;\n\t\t\t\t\t\t}\n\t\t\t\t\t\tthis.updateMonthDisplay();\n\t\t\t\t\t\tthis.renderCalendarDays();\n\t\t\t\t\t},\n\n\t\t\t\t\tparseDate(dateStr) {\n\t\t\t\t\t\tconst parts = dateStr.split(/[-/.]/);\n\t\t\t\t\t\tswitch(this.format) {\n\t\t\t\t\t\t\tcase 'eu':\n\t\t\t\t\t\t\t\treturn `${parts[2]}-${parts[1]}-${parts[0]}`;\n\t\t\t\t\t\t\tcase 'us':\n\t\t\t\t\t\t\t\treturn `${parts[2]}-${parts[0]}-${parts[1]}`;\n\t\t\t\t\t\t\tcase 'uk':\n\t\t\t\t\t\t\t\treturn `${parts[2]}-${parts[1]}-${parts[0]}`;\n\t\t\t\t\t\t\tcase 'long':\n\t\t\t\t\t\t\tcase 'iso':\n\t\t\t\t\t\t\tdefault:\n\t\t\t\t\t\t\t\treturn dateStr;\n\t\t\t\t\t\t}\n\t\t\t\t\t},\n\n\t\t\t\t\tformatDate(date) {\n\t\t\t\t\t\tconst d = date.getDate().toString().padStart(2, '0');\n\t\t\t\t\t\tconst m = (date.getMonth() + 1).toString().padStart(2, '0');\n\t\t\t\t\t\tconst y = date.getFullYear();\n\n\t\t\t\t\t\tswitch(this.format) {\n\t\t\t\t\t\t\tcase 'eu':\n\t\t\t\t\t\t\t\treturn `${d}.${m}.${y}`;\n\t\t\t\t\t\t\tcase 'uk':\n\t\t\t\t\t\t\t\treturn `${d}/${m}/${y}`;\n\t\t\t\t\t\t\tcase 'us':\n\t\t\t\t\t\t\t\treturn `${m}/${d}/${y}`;\n\t\t\t\t\t\t\tcase 'long':\n\t\t\t\t\t\t\t\treturn `${this.monthNames[date.getMonth()]} ${d}, ${y}`;\n\t\t\t\t\t\t\tcase 'iso':\n\t\t\t\t\t\t\tdefault:\n\t\t\t\t\t\t\t\treturn `${y}-${m}-${d}`;\n\t\t\t\t\t\t}\n\t\t\t\t\t},\n\n\t\t\t\t\tisSelectedDate(day) {\n\t\t\t\t\t\tif (!this.value) return false;\n\t\t\t\t\t\t\n\t\t\t\t\t\ttry {\n\t\t\t\t\t\t\tconst date = new Date(this.currentYear, this.currentMonth, day);\n\t\t\t\t\t\t\tconst selected = new Date(this.parseDate(this.value));\n\t\t\t\t\t\t\treturn date.toDateString() === selected.toDateString();\n\t\t\t\t\t\t} catch(e) {\n\t\t\t\t\t\t\treturn false;\n\t\t\t\t\t\t}\n\t\t\t\t\t},\n\n\t\t\t\t\tselectDate(e) {\n\t\t\t\t\t\tconst day = e.target.dataset.day;\n\t\t\t\t\t\tif (!day) return;\n\t\t\t\t\t\t\n\t\t\t\t\t\tconst date = new Date(this.currentYear, this.currentMonth, parseInt(day));\n\t\t\t\t\t\tthis.value = this.formatDate(date);\n\t\t\t\t\t\t\n\t\t\t\t\t\t// Update hidden input\n\t\t\t\t\t\tthis.$refs.hiddenInput.value = this.value;\n\t\t\t\t\t\tthis.$refs.hiddenInput.dispatchEvent(new Event('change', { bubbles: true }));\n\t\t\t\t\t\t\n\t\t\t\t\t\t// Update display text\n\t\t\t\t\t\tthis.$refs.displayText.innerHTML = '';\n\t\t\t\t\t\tthis.$refs.displayText.textContent = this.value;\n\t\t\t\t\t\tthis.$refs.displayText.classList.remove('text-muted-foreground');\n\t\t\t\t\t\t\n\t\t\t\t\t\t// Close DatePicker\n\t\t\t\t\t\tthis.open = false;\n\t\t\t\t\t}\n\t\t\t\t}));\n\t\t\t});\n\t\t</script>")
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
