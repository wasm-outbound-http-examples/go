// Based on https://github.com/realPy/learn-hogosuru/blob/08bf39e0f7e2139445b4a09a8263682d703f9034/iofetch/front/main.go
// and https://github.com/realPy/hogosuru-jwtdecoder/blob/fe2805340113e5cc030f27f5f62025b4f976f830/main.go
package main

import (
	"fmt"

	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/base/document"
	"github.com/realPy/hogosuru/base/event"
	"github.com/realPy/hogosuru/base/fetch"
	"github.com/realPy/hogosuru/base/htmlareaelement"
	"github.com/realPy/hogosuru/base/htmlbuttonelement"
	"github.com/realPy/hogosuru/base/htmldivelement"
	"github.com/realPy/hogosuru/base/htmllinkelement"
	"github.com/realPy/hogosuru/base/promise"
	"github.com/realPy/hogosuru/base/response"
)

func main() {
	hogosuru.Init()
	if doc, err := document.New(); hogosuru.AssertErr(err) {
		if head, err := doc.Head(); hogosuru.AssertErr(err) {
			if link, err := htmllinkelement.New(doc); hogosuru.AssertErr(err) {
				link.SetRel("stylesheet")
				link.SetHref("https://cdn.jsdelivr.net/npm/bulma@0.9.3/css/bulma.min.css")
				head.AppendChild(link.Node)
			}
		}

		if body := doc.Body_(); hogosuru.AssertErr(err) {
			if div, err := htmldivelement.New(doc); hogosuru.AssertErr(err) {
				if list, err := div.ClassList(); hogosuru.AssertErr(err) {
					list.Add("buttons")
				}
				if buttonprimary, err := htmlbuttonelement.New(doc); hogosuru.AssertErr(err) {
					buttonprimary.SetTextContent("Press this button to make a HTTP request")
					if list, err := buttonprimary.ClassList(); hogosuru.AssertErr(err) {
						list.Add("button")
						list.Add("is-primary")
					}

					buttonprimary.OnClick(func(e event.Event) {
						if ftch, err := fetch.New("https://httpbin.org/anything", nil); hogosuru.AssertErr(err) {
							ftch.Then(func(r response.Response) *promise.Promise {

								textpromise, _ := ftch.Then(func(r response.Response) *promise.Promise {
									if promise, err := r.Text(); hogosuru.AssertErr(err) {
										return &promise
									}
									return nil
								}, nil)

								textpromise.Then(func(i interface{}) *promise.Promise {
									fmt.Println(i.(string))
									if element, err := doc.GetElementById("outtxt"); hogosuru.AssertErr(err) {
										element.SetTextContent(i.(string))
									}
									return nil
								}, nil)

								return nil
							}, func(e error) {
								fmt.Printf("An error occured: %s\n", e.Error())
							})
						}
					})

					div.Append(buttonprimary.Element)
				}
				if outtext, err := htmlareaelement.New(doc); hogosuru.AssertErr(err) {
					outtext.SetID("outtxt")
					outtext.SetTextContent("[Result will be displayed here too]")
					if list, err := outtext.ClassList(); hogosuru.AssertErr(err) {
						list.Add("textarea")
						list.Add("is-primary")
					}

					div.Append(outtext.Element)
				}

				body.Append(div.Element)
			}
		}
	}

	ch := make(chan struct{})
	<-ch
}
