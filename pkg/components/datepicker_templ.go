// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/axzilla/goilerplate/pkg/icons"
	"github.com/axzilla/goilerplate/pkg/utils"
	"time"
)

// DateFormat defines the available date formatting styles
type DateFormat string

const (
	// DateFormatISO represents ISO format (YYYY-MM-DD)
	DateFormatISO DateFormat = "iso"
	// DateFormatEU represents European format (DD.MM.YYYY)
	DateFormatEU DateFormat = "eu"
	// DateFormatUK represents UK format (DD/MM/YYYY)
	DateFormatUK DateFormat = "uk"
	// DateFormatUS represents US format (MM/DD/YYYY)
	DateFormatUS DateFormat = "us"
	// DateFormatLONG represents long format (Month DD, YYYY)
	DateFormatLONG DateFormat = "long"
)

// dateFormatMapping maps DateFormat to Go time format strings
var dateFormatMapping = map[DateFormat]string{
	DateFormatISO:  "2006-01-02",
	DateFormatEU:   "02.01.2006",
	DateFormatUK:   "02/01/2006",
	DateFormatUS:   "01/02/2006",
	DateFormatLONG: "January 2, 2006",
}

// DateLocale defines localization settings for the datepicker
type DateLocale struct {
	// MonthNames contains the localized names of months
	MonthNames []string
	// DayNames contains the localized abbreviated day names
	DayNames []string
}

// DateLocaleDefault provides English localization
var DateLocaleDefault = DateLocale{
	MonthNames: []string{"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December"},
	DayNames: []string{"Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"},
}

// Pre-defined locales for different languages
var (
	// DateLocaleSpanish provides Spanish localization
	DateLocaleSpanish = DateLocale{
		MonthNames: []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio",
			"Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"},
		DayNames: []string{"Lu", "Ma", "Mi", "Ju", "Vi", "Sa", "Do"},
	}

	// DateLocaleGerman provides German localization
	DateLocaleGerman = DateLocale{
		MonthNames: []string{"Januar", "Februar", "März", "April", "Mai", "Juni",
			"Juli", "August", "September", "Oktober", "November", "Dezember"},
		DayNames: []string{"Mo", "Di", "Mi", "Do", "Fr", "Sa", "So"},
	}

	// DateLocaleFrench provides French localization
	DateLocaleFrench = DateLocale{
		MonthNames: []string{"Janvier", "Février", "Mars", "Avril", "Mai", "Juin",
			"Juillet", "Août", "Septembre", "Octobre", "Novembre", "Décembre"},
		DayNames: []string{"Lu", "Ma", "Me", "Je", "Ve", "Sa", "Di"},
	}

	// DateLocaleItalian provides Italian localization
	DateLocaleItalian = DateLocale{
		MonthNames: []string{"Gennaio", "Febbraio", "Marzo", "Aprile", "Maggio", "Giugno",
			"Luglio", "Agosto", "Settembre", "Ottobre", "Novembre", "Dicembre"},
		DayNames: []string{"Lu", "Ma", "Me", "Gi", "Ve", "Sa", "Do"},
	}
)

// DatepickerConfig combines format and locale settings
type DatepickerConfig struct {
	Format DateFormat
	Locale DateLocale
}

// Pre-defined datepicker configurations
var (
	// DatePickerISO provides ISO format with default locale
	DatePickerISO = DatepickerConfig{
		Format: DateFormatISO,
		Locale: DateLocaleDefault,
	}

	// DatePickerEU provides European format with default locale
	DatePickerEU = DatepickerConfig{
		Format: DateFormatEU,
		Locale: DateLocaleDefault,
	}

	// DatePickerUK provides UK format with default locale
	DatePickerUK = DatepickerConfig{
		Format: DateFormatUK,
		Locale: DateLocaleDefault,
	}

	// DatePickerUS provides US format with default locale
	DatePickerUS = DatepickerConfig{
		Format: DateFormatUS,
		Locale: DateLocaleDefault,
	}

	// DatePickerUS provides US format with default locale
	DatePickerLONG = DatepickerConfig{
		Format: DateFormatLONG,
		Locale: DateLocaleDefault,
	}
)

// NewDatepickerConfig creates a new configuration with specified format and locale
func NewDatepickerConfig(format DateFormat, locale DateLocale) DatepickerConfig {
	return DatepickerConfig{
		Format: format,
		Locale: locale,
	}
}

// GetGoFormat returns the Go time format string for the current configuration
func (c DatepickerConfig) GetGoFormat() string {
	if format, ok := dateFormatMapping[c.Format]; ok {
		return format
	}
	return dateFormatMapping[DateFormatISO] // Default to ISO
}

// DatepickerProps defines all available props for the Datepicker component
type DatepickerProps struct {
	// ID sets the input element's ID
	ID string
	// Name sets the input element's name
	Name string
	// Value sets the initial date
	Value time.Time
	// Config provides format and locale settings
	Config DatepickerConfig
	// Placeholder sets the input placeholder text
	Placeholder string
	// Disabled controls whether the datepicker can be interacted with
	Disabled bool
	// Class adds custom CSS classes
	Class string
	// Attributes allows additional HTML attributes
	Attributes templ.Attributes
}

// Datepicker renders a localized date picker component with configurable formatting.
// Features:
// - Multiple date formats (ISO, EU, US, UK, Long)
// - Localization support for months and weekdays
// - Popup calendar for date selection
// - Keyboard navigation support
// - Responsive positioning
// - Custom styling support
//
// Props:
// - ID: Input element identifier
// - Name: Form field name
// - Value: Initial date value
// - Config: Format and locale settings
// - Placeholder: Input placeholder text
// - Disabled: Interactivity state
// - Class: Additional CSS classes
// - Attributes: Additional HTML attributes
func Datepicker(props DatepickerProps) templ.Component {
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
		var templ_7745c5c3_Var2 = []any{utils.TwMerge("relative", props.Class)}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var2).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/datepicker.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.Value != (time.Time{}) {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" data-value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(props.Value.Format(props.Config.GetGoFormat()))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/datepicker.templ`, Line: 177, Col: 62}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" data-format=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(string(props.Config.Format))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/datepicker.templ`, Line: 179, Col: 43}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" data-monthnames=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(templ.JSONString(props.Config.Locale.MonthNames))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/datepicker.templ`, Line: 180, Col: 68}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" data-daynames=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(templ.JSONString(props.Config.Locale.DayNames))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/datepicker.templ`, Line: 181, Col: 64}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" x-data=\"{\n            open: false,\n            value: null,\n            format: $el.dataset.format,\n            currentMonth: 5,\n            currentYear: new Date().getFullYear(),\n            monthDays: [],\n            blankDays: [],\n            months: JSON.parse($el.dataset.monthnames) || [&#39;January&#39;, &#39;February&#39;, &#39;March&#39;, &#39;April&#39;, &#39;May&#39;, &#39;June&#39;, &#39;July&#39;, &#39;August&#39;, &#39;September&#39;, &#39;October&#39;, &#39;November&#39;, &#39;December&#39;],\n            days: JSON.parse($el.dataset.daynames) || [&#39;Mo&#39;, &#39;Tu&#39;, &#39;We&#39;, &#39;Th&#39;, &#39;Fr&#39;, &#39;Sa&#39;, &#39;Su&#39;],\n\n            position: &#39;bottom&#39;,\n\n            init() {\n                const initialDate = $el.dataset.value ? new Date(this.parseDate($el.dataset.value)) : new Date();\n                this.currentMonth = initialDate.getMonth();\n                this.currentYear = initialDate.getFullYear();\n                this.calculateDays();\n                // Format the initial value using the correct locale\n                if ($el.dataset.value) {\n                    this.value = this.formatDate(initialDate);\n                }\n            },\n\n            toggleDatePicker() {\n                this.open = !this.open;\n                if (this.open) {\n                    this.$nextTick(() =&gt; this.updatePosition());\n                }\n            },\n\n            updatePosition() {\n                const trigger = this.$refs.datePickerInput;\n                const popup = this.$refs.datePickerPopup;\n                const rect = trigger.getBoundingClientRect();\n                const popupRect = popup.getBoundingClientRect();\n                const viewportHeight = window.innerHeight || document.documentElement.clientHeight;\n                \n                if (rect.bottom + popupRect.height &gt; viewportHeight &amp;&amp; rect.top &gt; popupRect.height) {\n                    this.position = &#39;top&#39;;\n                } else {\n                    this.position = &#39;bottom&#39;;\n                }\n            },\n\n            calculateDays() {\n                const firstDay = new Date(this.currentYear, this.currentMonth, 1).getDay();\n                const daysInMonth = new Date(this.currentYear, this.currentMonth + 1, 0).getDate();\n                \n                this.blankDays = Array.from({ length: firstDay }, (_, i) =&gt; i);\n                this.monthDays = Array.from({ length: daysInMonth }, (_, i) =&gt; i + 1);\n            },\n\n\t\t\tparseDate(dateStr) {\n\t\t\t\tconst parts = dateStr.split(/[-/.]/);\n\t\t\t\tswitch(this.format) {\n\t\t\t\t\tcase &#39;eu&#39;:\n\t\t\t\t\t\treturn `${parts[2]}-${parts[1]}-${parts[0]}`;\n\t\t\t\t\tcase &#39;us&#39;:\n\t\t\t\t\t\treturn `${parts[2]}-${parts[0]}-${parts[1]}`;\n\t\t\t\t\tcase &#39;uk&#39;:\n\t\t\t\t\t\treturn `${parts[2]}-${parts[1]}-${parts[0]}`;\n\t\t\t\t\tcase &#39;long&#39;:\n\t\t\t\t\tcase &#39;iso&#39;:\n\t\t\t\t\tdefault:\n\t\t\t\t\t\treturn dateStr;  // Für ISO und long Format das Datum unverändert lassen\n\t\t\t\t}\n\t\t\t}, \n\n\t\t\tformatDate(date) {\n                const d = date.getDate().toString().padStart(2, &#39;0&#39;);\n                const m = (date.getMonth() + 1).toString().padStart(2, &#39;0&#39;);\n                const y = date.getFullYear();\n\n                switch(this.format) {\n                    case &#39;eu&#39;:\n                        return `${d}.${m}.${y}`;\n\t\t\t\t\tcase &#39;uk&#39;:\n\t\t\t\t\t\treturn `${d}/${m}/${y}`;\n                    case &#39;us&#39;:\n                        return `${m}/${d}/${y}`;\n                    case &#39;long&#39;:\n                        // Use the months array from the provided locale\n                        return `${this.months[date.getMonth()]} ${d}, ${y}`;\n\t\t\t\t\t// case &#39;iso&#39;\n                    default:\n                        return `${y}-${m}-${d}`;\n                }\n            },\n\n            isToday(day) {\n                const today = new Date();\n                const date = new Date(this.currentYear, this.currentMonth, day);\n                return date.toDateString() === today.toDateString();\n            },\n\n            isSelected(day) {\n                if (!this.value) return false;\n                const date = new Date(this.currentYear, this.currentMonth, day);\n                const selected = new Date(this.parseDate(this.value));\n                return date.toDateString() === selected.toDateString();\n            },\n\n            selectDate(day) {\n                const date = new Date(this.currentYear, this.currentMonth, day);\n                this.value = this.formatDate(date);\n                this.open = false;\n            }\n        }\" @resize.window=\"if (open) updatePosition()\"><div class=\"relative\"><input x-ref=\"datePickerInput\" type=\"text\" id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(props.ID)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/datepicker.templ`, Line: 297, Col: 17}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" name=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var9 string
		templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(props.Name)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/datepicker.templ`, Line: 298, Col: 21}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.Placeholder != "" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" placeholder=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var10 string
			templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(props.Placeholder)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/datepicker.templ`, Line: 300, Col: 36}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" placeholder=\"Select date\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if props.Disabled {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" disabled")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" :value=\"value\" x-modelable=\"value\" @click=\"toggleDatePicker()\" readonly class=\"peer flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, props.Attributes)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("> <button type=\"button\" @click=\"toggleDatePicker()\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.Disabled {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" disabled")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" class=\"peer-disabled:pointer-events-none peer-disabled:opacity-50 absolute top-0 right-0 px-3 py-2 cursor-pointer text-muted-foreground hover:text-foreground\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = icons.Calendar(icons.IconProps{}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button></div><div x-show=\"open\" x-ref=\"datePickerPopup\" @click.away=\"open = false\" x-transition.opacity class=\"absolute left-0 z-50 mt-1 w-64 rounded-lg border bg-popover p-4 shadow-md\" :class=\"{&#39;top-full mt-1&#39;: position === &#39;bottom&#39;,&#39;bottom-full mb-1&#39;: position === &#39;top&#39;}\"><div class=\"flex items-center justify-between mb-4\"><span x-text=\"months[currentMonth] + &#39; &#39; + currentYear\" class=\"text-sm font-medium\"></span><div class=\"flex gap-1\"><button type=\"button\" @click=\"currentMonth--; if(currentMonth &lt; 0) { currentMonth = 11; currentYear--; } calculateDays()\" class=\"inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 hover:bg-accent hover:text-accent-foreground h-7 w-7\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = icons.ChevronLeft(icons.IconProps{}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button> <button type=\"button\" @click=\"currentMonth++; if(currentMonth &gt; 11) { currentMonth = 0; currentYear++; } calculateDays()\" class=\"inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 hover:bg-accent hover:text-accent-foreground h-7 w-7\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = icons.ChevronRight(icons.IconProps{}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button></div></div><div class=\"grid grid-cols-7 gap-1 mb-2\"><template x-for=\"day in days\" :key=\"day\"><div class=\"text-center text-xs text-muted-foreground\" x-text=\"day\"></div></template></div><div class=\"grid grid-cols-7 gap-1\"><template x-for=\"blank in blankDays\" :key=\"&#39;blank&#39; + blank\"><div class=\"h-8 w-8\"></div></template><template x-for=\"day in monthDays\" :key=\"day\"><button type=\"button\" @click=\"selectDate(day)\" x-text=\"day\" :class=\"{\n                            &#39;bg-primary text-primary-foreground&#39;: isSelected(day),\n                            &#39;text-red-500&#39;: isToday(day) &amp;&amp; !isSelected(day),\n                            &#39;hover:bg-accent hover:text-accent-foreground&#39;: !isSelected(day)\n                        }\" class=\"inline-flex h-8 w-8 items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2\"></button></template></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
