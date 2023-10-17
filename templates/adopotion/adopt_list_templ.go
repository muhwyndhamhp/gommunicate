// Code generated by templ@v0.2.364 DO NOT EDIT.

package adoption

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "fmt"

import (
	"github.com/muhwyndhamhp/gotes-mx/modules/pet"
	"github.com/muhwyndhamhp/gotes-mx/templates/base"
	"github.com/muhwyndhamhp/gotes-mx/utils/timehelper"
	"time"
)

func isLastItem(index, length int) bool {
	return index == length-1
}

func AdoptionList(pets []pet.Pet, petTypes []pet.PetTypeSet, nextURL string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		err = base.DocType().Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		err = base.Head().Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		err = base.Body(page(pets, petTypes, nextURL)).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func List(pets []pet.Pet, nextURL string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_2 := templ.GetChildren(ctx)
		if var_2 == nil {
			var_2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		for i, pet := range pets {
			err = item(pet, i, len(pets), nextURL).Render(ctx, templBuffer)
			if err != nil {
				return err
			}
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func item(petData pet.Pet, index, length int, nextURL string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_3 := templ.GetChildren(ctx)
		if var_3 == nil {
			var_3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div class=\"col-span-1 relative mb-16\"")
		if err != nil {
			return err
		}
		if index == length-1 {
			_, err = templBuffer.WriteString(" hx-get=\"")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(templ.EscapeString(nextURL))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("\" hx-swap=\"afterend\" hx-trigger=\"revealed\"")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		err = imageCarousel(petData.Images).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		err = petInfo(petData).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func petInfo(petData pet.Pet) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_4 := templ.GetChildren(ctx)
		if var_4 == nil {
			var_4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div class=\"rounded-2xl absolute -bottom-16 p-4 w-full bg-white shadow-xl border outline-slate-200 shadow-cyan-800/20 z-20\"><p class=\"text-2xl\">")
		if err != nil {
			return err
		}
		var var_5 string = petData.Name
		_, err = templBuffer.WriteString(templ.EscapeString(var_5))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" <span class=\"text-sm\">")
		if err != nil {
			return err
		}
		var_6 := `(`
		_, err = templBuffer.WriteString(var_6)
		if err != nil {
			return err
		}
		var var_7 string = petData.Nickname
		_, err = templBuffer.WriteString(templ.EscapeString(var_7))
		if err != nil {
			return err
		}
		var_8 := `)`
		_, err = templBuffer.WriteString(var_8)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</span></p><p class=\"text-xs text-gray-700 mt-1\">")
		if err != nil {
			return err
		}
		var var_9 string = getBirthday(petData.Birthday)
		_, err = templBuffer.WriteString(templ.EscapeString(var_9))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" ")
		if err != nil {
			return err
		}
		var_10 := `| `
		_, err = templBuffer.WriteString(var_10)
		if err != nil {
			return err
		}
		var var_11 string = petData.Race
		_, err = templBuffer.WriteString(templ.EscapeString(var_11))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p><p class=\"text-sm mt-2\">")
		if err != nil {
			return err
		}
		var var_12 string = petData.Address1
		_, err = templBuffer.WriteString(templ.EscapeString(var_12))
		if err != nil {
			return err
		}
		var_13 := `, `
		_, err = templBuffer.WriteString(var_13)
		if err != nil {
			return err
		}
		var var_14 string = petData.Address2
		_, err = templBuffer.WriteString(templ.EscapeString(var_14))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p></div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func getBirthday(bd *time.Time) string {
	return fmt.Sprintf("%v Years Old", timeHelper.RoundTime(time.Since(*bd).Seconds()/31207680))
}

func page(pets []pet.Pet, petTypes []pet.PetTypeSet, nextURL string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_15 := templ.GetChildren(ctx)
		if var_15 == nil {
			var_15 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<p>")
		if err != nil {
			return err
		}
		var_16 := `Followings are list of adoptable pets!`
		_, err = templBuffer.WriteString(var_16)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p><br>")
		if err != nil {
			return err
		}
		err = searchBar(petTypes).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("<br><div class=\"relative\">")
		if err != nil {
			return err
		}
		err = loading().Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("<div class=\"grid gap-6 grid-cols-1 mt-4 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-5\" id=\"pets-parent\">")
		if err != nil {
			return err
		}
		err = List(pets, nextURL).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div></div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
